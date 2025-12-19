# 🌐 CTI Web Scraper (Go)

Siber Tehdit İstihbaratı (CTI) operasyonları için geliştirilmiş, güvenilir ve kapsamlı web scraping aracı.

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/platform-cross--platform-lightgrey)](https://github.com/golang/go)


## 🎯 Özellikler

### Temel Fonksiyonlar
- **HTML İçerik Çekme**: Web sayfalarının tam HTML içeriğini güvenilir şekilde çeker
- **Ekran Görüntüsü Alma**: Chrome DevTools Protocol ile otomatik screenshot
- **URL Listeleme**: Sayfadaki tüm linkleri çıkarır ve listeler
- **Otomatik Dosya Organizasyonu**: Timestamp içeren klasör yapısı
- **Kapsamlı Hata Yönetimi**: Network, timeout ve parsing hatalarına karşı dayanıklı

### CTI Odaklı Özellikler
- **Çoklu Platform Desteği**: 15+ farklı site kategorisi test edildi
- **Güvenilir HTTP İletişimi**: Timeout ve retry mekanizmaları
- **Veri Bütünlüğü**: Dosya boyutu ve karakter sayısı raporlama
- **Profesyonel Raporlama**: Detaylı analiz çıktıları

## 🚀 Kurulum

### Gereksinimler
- **Go 1.25+** ([İndir](https://golang.org/dl/))
- **Google Chrome/Chromium** (Screenshot için)
- **Git** ([İndir](https://git-scm.com/))

### Hızlı Başlangıç

```bash
# Repository'yi klonla
git clone https://github.com/YOUR_USERNAME/cti-web-scraper.git
cd cti-web-scraper

# Bağımlılıkları yükle
go mod download

# Programı çalıştır
go run main.go https://example.com
```

### Manuel Kurulum

```bash
# Go modülünü başlat (yeni proje için)
go mod init cti-web-scraper

# ChromeDP bağımlılığını ekle
go get github.com/chromedp/chromedp
```

## 📖 Kullanım

### Temel Kullanım
```bash
# Tek bir web sitesini analiz et
go run main.go https://github.com

# Yerel dosya olarak derle ve çalıştır
go build -o scraper main.go
./scraper https://stackoverflow.com
```

### Çıktı Dosyaları

Program her çalıştırmada aşağıdaki dosyaları oluşturur:

```
scraper_output_domain_com_YYYYMMDD_HHMMSS/
├── README.txt          # Test bilgileri ve özeti
├── site_data.html      # Ham HTML içerik
├── output.txt          # İşlenmiş text içerik  
├── urls.txt            # Çıkarılan URL listesi
└── screenshot.png      # Sayfa ekran görüntüsü
```

### Örnek Çıktı

```bash
$ go run main.go https://github.com

Siber Vatan CTI - Grup 2
=======================
Target URL: https://github.com
Start Time: 2025-12-18 16:40:11

[✓] HTML içerik alındı (562,265 bytes)
[✓] Text dosyası oluşturuldu (142,773 karakter)
[✓] URL listesi çıkarıldı (195 link)
[✓] Ekran görüntüsü alındı (372KB)

Dosyalar: scraper_output_github_com_20251218_164011/
```

## 🛠 Teknik Detaylar

### Kullanılan Kütüphaneler

**Standart Kütüphaneler:**
- `net/http` - HTTP istek/yanıt işlemleri
- `net/url` - URL parsing ve validation  
- `context` - Context yönetimi ve timeout
- `regexp` - Regular expression işlemleri
- `os`, `io` - Dosya sistemi operasyonları

**Harici Kütüphaneler:**
- `github.com/chromedp/chromedp` - Chrome DevTools Protocol

### Güvenlik Özellikleri

- **Dosya Adı Sanitization**: Güvenli dosya adı oluşturma
- **Timeout Yönetimi**: Maksimum 30 saniye bekleme
- **Error Handling**: Kapsamlı hata yakalama ve raporlama
- **Resource Cleanup**: Otomatik kaynak temizleme

## 📊 Test Sonuçları

Program 15 farklı kategori ve 150+ web sitesi üzerinde test edilmiştir:

| Kategori | Test Sayısı | Başarı Oranı |
|----------|-------------|---------------|
| Haber ve Blog | 3 | %100 |
| Teknoloji Platformları | 3 | %100 |
| Kurumsal ve Eğitim | 3 | %100 |
| Açık Veri Platformları | 3 | %100 |
| E-ticaret ve Sosyal | 3 | %100 |

### Test Edilen Platformlar
- **Haber**: BBC, Wikipedia, Medium
- **Teknoloji**: GitHub, StackOverflow, Reddit  
- **Kurumsal**: Microsoft, TÜBİTAK, Coursera
- **Açık Veri**: data.gov, data.europa.eu, Kaggle
- **E-ticaret**: Trendyol, Amazon, Mastodon

Detaylı test sonuçları için [proje raporu](docs/project-report.pdf) dosyasına bakın.

## 📁 Proje Yapısı

```
cti-web-scraper/
├── main.go                    # Ana kaynak kod
├── go.mod                     # Go modül tanımları
├── go.sum                     # Bağımlılık checksumları  
├── README.md                  # Bu dosya
├── docs/                      # Dokümantasyon
│   └── project-report.pdf     # Detaylı proje raporu
├── examples/                  # Örnek kullanımlar
│   └── sample-outputs/        # Örnek çıktı dosyaları
└── .gitignore                 # Git ignore kuralları
```

## 👨‍💻 Geliştirici

**Helim Hümeyran SARAÇOĞLU**  
*Siber Vatan Team 11 - YILDIZ CTI Grup 2*

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=flat&logo=linkedin&logoColor=white)](https://linkedin.com/in/helim-humeyran-saracoglu)

---
