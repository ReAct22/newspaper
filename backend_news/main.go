package main

import (
	"backend_news/config"
	"backend_news/service"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Ganti dengan URL frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Message": "Selamat Datang di pemograman golang"})
	})

	r.GET("/news", func(c *gin.Context) {
		newsItems, err := service.GetTopStories()
		if err != nil {
			log.Printf("Error fetching news: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, newsItems)
	})

	r.GET("/category/:category", func(c *gin.Context) {
		category := c.Param("category")

		newsItems, err := service.GetNewsByCategory(category)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, newsItems)
	})

	r.GET("/comment", func(c *gin.Context) {
		storyIDStr := c.Query("story_id")
		if storyIDStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "story_id parameter is required"})
			return
		}

		storyID, err := strconv.Atoi(storyIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid story_id"})
			return
		}

		comments, err := service.GetComments(storyID)
		if err != nil {
			log.Println("Error getting comments:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
			return
		}

		c.JSON(http.StatusOK, comments)
	})

	r.Run(":8080")
}
