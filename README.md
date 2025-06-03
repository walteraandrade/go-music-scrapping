# Go Web Scraper with Colly

A simple and efficient web scraper built with Go and the **Colly** framework to discover album recommendations based on music genres. Perfect for music enthusiasts looking to explore new albums within their favorite genres.

## Features

- 🎵 **Genre-based Discovery**: Search for albums by any music genre
- 📊 **Detailed Information**: Extracts album names, artists, URLs, and release years
- 🛡️ **Error Handling**: Simple, but eficient error handling for HTTP requests
- 🚀 **Fast & Lightweight**: Built with Go's performance and Colly's efficiency

## Quick Start

### Prerequisites

- Go 1.19 or higher
- Internet connection

### Installation

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd <project-folder>
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the scraper:**
   ```bash
   go run main.go
   ```

## Configuration

### Changing the Target Genre

The scraper currently targets a hardcoded genre in `main.go`. To explore different genres:

1. Open `main.go` in your editor
2. Locate this line:
   ```go
   baseUrl := "https://www.allmusic.com/search/all/afrobeat"
   ```
3. Replace `"afrobeat"` with your desired genre:
   ```go
   baseUrl := "https://www.allmusic.com/search/all/jazz"        // For jazz
   baseUrl := "https://www.allmusic.com/search/all/rock"        // For rock
   baseUrl := "https://www.allmusic.com/search/all/electronic"  // For electronic
   ```

## How It Works

The scraper targets **AllMusic.com** to gather comprehensive album data. For each discovered album, it extracts:

- 🎵 Album name and direct URL
- 👤 Artist name and profile URL  
- 📅 Release year
- 🔗 Additional metadata from individual album pages

All scraped data is displayed in the console with structured formatting for easy reading.

## Roadmap

### Upcoming Features

- [ ] **🔄 Infinite Scrolling Support**  
  Handle dynamic content loading for more comprehensive results

- [ ] **💻 CLI Genre Input**  
  Accept genre as command-line argument: `go run main.go --genre="jazz"`

- [ ] **🌐 Multiple Data Sources**  
  Expand to scrape from Discogs, Spotify, and other music databases

- [ ] **🖥️ Web Interface**  
  Full-stack application with **Svelte** frontend for interactive browsing

- [ ] **💾 Data Export**  
  Save results to JSON, CSV, or database formats

- [ ] **⚡ Concurrent Scraping**  
  Parallel processing for faster data collection

## Contributing

Contributions are welcome! Please feel free to submit issues, feature requests, or pull requests.

## License

This project is open source. Please check the LICENSE file for details.

---

**Note**: This scraper is designed for educational purposes. Please respect the terms of service of the websites you scrape and implement appropriate rate limiting.
