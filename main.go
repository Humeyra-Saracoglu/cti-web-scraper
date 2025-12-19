package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func fetchHTML(targetURL, outputDir string) error {

	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return fmt.Errorf("request oluşturulamadı: %v", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP isteği başarısız: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP hata kodu: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("HTML okunamadı: %v", err)
	}

	filename := filepath.Join(outputDir, "site_data.html")
	err = os.WriteFile(filename, body, 0644)
	if err != nil {
		return fmt.Errorf("dosyaya yazılamadı: %v", err)
	}

	fmt.Printf("[+] HTML kaydedildi (%s) - %d bytes\n", filename, len(body))
	return nil
}

func saveTextVersion(htmlContent, outputDir string) error {
	re := regexp.MustCompile(`<[^>]*>`)
	textContent := re.ReplaceAllString(htmlContent, " ")

	spaceRe := regexp.MustCompile(`\s+`)
	textContent = spaceRe.ReplaceAllString(textContent, " ")
	textContent = strings.TrimSpace(textContent)

	filename := filepath.Join(outputDir, "output.txt")
	err := os.WriteFile(filename, []byte(textContent), 0644)

	if err != nil {
		return fmt.Errorf("text dosyası yazılamadı: %v", err)
	}

	fmt.Printf("[+] Text dosyası kaydedildi (%s) - %d karakter\n", filename, len(textContent))
	return nil
}

func extractURLs(htmlContent, outputDir string) error {
	fmt.Println("[*] URL'ler çıkarılıyor...")

	re := regexp.MustCompile(`href\s*=\s*["']([^"']+)["']`)
	matches := re.FindAllStringSubmatch(htmlContent, -1)

	filename := filepath.Join(outputDir, "urls.txt")
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("URL dosyası oluşturulamadı: %v", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "# Web Scraper - URL Listesi\n")
	fmt.Fprintf(file, "# Helin Hümeyra SARAÇOĞLU\n")
	fmt.Fprintf(file, "# Tarih: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Fprintf(file, "# Toplam URL: %d\n\n", len(matches))

	validURLs := 0
	for _, match := range matches {
		if len(match) > 1 {
			urlStr := strings.TrimSpace(match[1])
			if urlStr != "" && urlStr != "#" && !strings.HasPrefix(urlStr, "javascript:") {
				fmt.Fprintf(file, "%d. %s\n", validURLs+1, urlStr)
				validURLs++
			}
		}
	}

	fmt.Printf("[+] URL listesi kaydedildi (%s) - %d URL bulundu\n", filename, validURLs)
	return nil
}

func takeScreenshot(targetURL, outputDir string) error {
	fmt.Println("[*] Ekran görüntüsü alınıyor...")

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var screenshot []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(targetURL),
		chromedp.Sleep(3*time.Second),
		chromedp.FullScreenshot(&screenshot, 90),
	)
	if err != nil {
		return fmt.Errorf("screenshot alınamadı: %v", err)
	}

	filename := filepath.Join(outputDir, "screenshot.png")
	err = os.WriteFile(filename, screenshot, 0644)
	if err != nil {
		return fmt.Errorf("screenshot dosyaya yazılamadı: %v", err)
	}

	fmt.Printf("[+] Screenshot kaydedildi (%s) - %d bytes\n", filename, len(screenshot))
	return nil
}

func createSafeFilename(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "unknown_site"
	}

	domain := u.Host
	domain = strings.ReplaceAll(domain, ".", "_")
	domain = strings.ReplaceAll(domain, ":", "_")
	domain = strings.ReplaceAll(domain, "/", "_")

	return domain
}

func createProjectInfo(targetURL, outputDir string) error {
	filename := filepath.Join(outputDir, "README.txt")

	content := fmt.Sprintf(`# Web Scraper Projesi - Çıktılar

	Geliştiren: Helin Hümeyra SARAÇOĞLU
	Takım: Siber Vatan Yıldız CTI - Grup 2
	Tarih: %s
	Hedef URL: %s

	Dosyalar:
	- site_data.html    # Ham HTML içerik
	- output.txt        # Temizlenmiş text içerik  
	- urls.txt          # Bulunan URL'lerin listesi
	- screenshot.png    # Ekran görüntüsü
	- README.txt        # Bu dosya

	--- YILDIZ CTI Web Scraping Projesi ----
	`, time.Now().Format("2006-01-02 15:04:05"), targetURL)

	return os.WriteFile(filename, []byte(content), 0644)
}

func main() {
	fmt.Println(" Go Web Scraper v2.0")
	fmt.Println(" Geliştiren: Helin Hümeyra SARAÇOĞLU")
	fmt.Println(" Siber Vatan Yıldız CTI - Grup 2")
	fmt.Println(strings.Repeat("=", 50))

	if len(os.Args) < 2 {
		fmt.Println("❌ Kullanım:")
		fmt.Println("   go run main.go <URL>")
		fmt.Println()
		fmt.Println(" Örnek:")
		fmt.Println("   go run main.go https://example.com")
		return
	}

	targetURL := os.Args[1]
	fmt.Printf(" *-*-> Hedef URL: %s\n", targetURL)
	fmt.Println()

	timestamp := time.Now().Format("20060102_150405")
	safeName := createSafeFilename(targetURL)
	outputDir := fmt.Sprintf("scraper_output_%s_%s", safeName, timestamp)

	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		fmt.Printf("[-] Klasör oluşturulamadı: %v\n", err)
		return
	}

	fmt.Printf("---> Çıktı klasörü: %s\n", outputDir)
	fmt.Println()

	// 1. HTML içeriği çek
	fmt.Println("[*] HTML içeriği çekiliyor...")
	if err := fetchHTML(targetURL, outputDir); err != nil {
		fmt.Printf("[-] HTML Hatası: %v\n", err)
		return
	}

	// 2. HTML dosyasını oku
	htmlFile := filepath.Join(outputDir, "site_data.html")
	htmlFiles, err := os.ReadFile(htmlFile)
	if err != nil {
		fmt.Printf("[-] HTML dosyası okunamadı: %v\n", err)
		return
	}
	htmlContent := string(htmlFiles)

	// 3. Text versiyonu oluştur
	fmt.Println("[*] Text dosyası oluşturuluyor...")
	if err := saveTextVersion(htmlContent, outputDir); err != nil {
		fmt.Printf("[-] Text Hatası: %v\n", err)
	}

	// 4. URL'leri çıkar
	if err := extractURLs(htmlContent, outputDir); err != nil {
		fmt.Printf("[-] URL Hatası: %v\n", err)
	}

	// 5. Screenshot al
	fmt.Println("[*] Ekran görüntüsü deneniyor...")
	if err := takeScreenshot(targetURL, outputDir); err != nil {
		fmt.Printf("[-] Screenshot Hatası: %v\n", err)
		fmt.Println(" !! Chrome/Chromium yüklü değil - Screenshot atlanıyor. !!")

		mockContent := fmt.Sprintf("# Screenshot Mock Dosyası\n# %s için ekran görüntüsü alınamadı\n# Chrome/Chromium kurulumu gerekli\n# Tarih: %s",
			targetURL, time.Now().Format("2006-01-02 15:04:05"))
		mockFile := filepath.Join(outputDir, "screenshot_mock.txt")
		os.WriteFile(mockFile, []byte(mockContent), 0644)
		fmt.Printf("[+] Mock screenshot dosyası oluşturuldu: %s\n", mockFile)
	}

	createProjectInfo(targetURL, outputDir)

	fmt.Println()
	fmt.Println("✅ Tüm işlemler tamamlandı!")
	fmt.Printf(" Sonuçlar klasörü: %s\n", outputDir)
	fmt.Println(" Oluşturulan dosyalar:")

	files, _ := os.ReadDir(outputDir)
	for _, file := range files {
		fmt.Printf("    %s\n", file.Name())
	}

	fmt.Println()
	fmt.Printf("💡 Klasörü açmak için: open %s\n", outputDir)
}
