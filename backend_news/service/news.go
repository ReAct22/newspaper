package service

import (
	"backend_news/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Struct untuk menampung data berita Hacker News
type NewsItem struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

var categories = map[string][]string{
	"programming":   {"Go", "Python", "JavaScript", "Rust", "Kotlin"},
	"tech-startups": {"Startup", "VC", "Fundraising", "Entrepreneurship"},
	"cybersecurity": {"Security", "Hacking", "Cyberattack", "Malware"},
	"open-source":   {"Open Source", "GitHub", "Linux", "BSD"},
	"science-ai":    {"AI", "ML", "Science", "Deep Learning"},
	"industry-news": {"Tech News", "Apple", "Google", "Microsoft"},
	"personal-dev":  {"Productivity", "Self-improvement", "Motivation"},
	"tech-culture":  {"Remote Work", "Work Culture", "Tech Community"},
}

// Fungsi untuk mendapatkan berita dari Hacker News API
func GetTopStories() ([]NewsItem, error) {
	// Ambil ID berita dari Hacker News API
	topStoriesURL := "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"
	resp, err := http.Get(topStoriesURL)
	if err != nil {
		return nil, fmt.Errorf("Error fetching Hacker News top stories: %v", err)
	}
	defer resp.Body.Close()

	var topStories []int
	if err := json.NewDecoder(resp.Body).Decode(&topStories); err != nil {
		return nil, fmt.Errorf("Error decoding response: %v", err)
	}

	// Ambil beberapa berita berdasarkan ID
	var newsItems []NewsItem
	for _, id := range topStories[:5] { // Ambil 5 berita pertama
		// Ambil detail dari tiap berita
		itemURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", id)
		itemResp, err := http.Get(itemURL)
		if err != nil {
			return nil, fmt.Errorf("Error fetching item: %v", err)
		}
		defer itemResp.Body.Close()

		var newsItem NewsItem
		if err := json.NewDecoder(itemResp.Body).Decode(&newsItem); err != nil {
			return nil, fmt.Errorf("Error decoding item response: %v", err)
		}

		// Tambahkan berita ke slice
		newsItems = append(newsItems, newsItem)
	}

	return newsItems, nil
}

func getCachedNews(category string) ([]NewsItem, error) {
	rows, err := config.DB.Query("SELECT id, title, url FROM cached_news WHERE category = ? AND updated_at > NOW() - INTERVAL 10 MINUTE", category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var news []NewsItem
	for rows.Next() {
		var item NewsItem
		if err := rows.Scan(&item.ID, &item.Title, &item.URL); err != nil {
			return nil, err
		}
		news = append(news, item)
	}
	return news, nil
}

func cacheNews(news []NewsItem, category string) {
	for _, item := range news {
		_, err := config.DB.Exec("INSERT INTO cached_news (id, title, url, category) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE title=?, url=?, updated_at=NOW()",
			item.ID, item.Title, item.URL)
		if err != nil {
			log.Println("Error caching news:", err)
		}
	}
}

// Fungsi untuk mendapatkan berita berdasarkan kategori
func GetNewsByCategory(category string) ([]NewsItem, error) {
	// Cek cache dulu
	cachedNews, err := getCachedNews(category)
	if err == nil && len(cachedNews) > 0 {
		fmt.Println("Using cached data for category:", category)
		return cachedNews, nil
	}

	// Jika tidak ada cache, ambil berita baru
	allNews, err := GetTopStories()
	if err != nil {
		return nil, err
	}

	keywords, exists := categories[category]
	if !exists {
		return nil, fmt.Errorf("Kategori tidak ditemukan")
	}

	// Filter berita sesuai kategori
	var filteredNews []NewsItem
	for _, item := range allNews {
		for _, keyword := range keywords {
			if strings.Contains(strings.ToLower(item.Title), strings.ToLower(keyword)) {
				filteredNews = append(filteredNews, item)
				break
			}
		}
	}

	// Simpan ke cache MySQL
	cacheNews(filteredNews, category)

	return filteredNews, nil
}
