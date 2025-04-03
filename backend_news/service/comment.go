package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Comment struct {
	ID     int    `json:"id"`
	By     string `json:"by"`
	Text   string `json:"text"`
	Parent int    `json:"parent"`
	Time   int64  `json:"time"`
	Type   string `json:"type"`
}

func GetComments(storyID int) ([]Comment, error) {
	itemURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", storyID)
	resp, err := http.Get(itemURL)
	if err != nil {
		log.Println("Error fetching story details:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var storyData struct {
		Kids []int `json:"kids"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&storyData); err != nil {
		log.Println("Error decoding story response:", err)
		return nil, err
	}

	var comments []Comment
	for _, commentID := range storyData.Kids {
		commentURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", commentID)
		commentResp, err := http.Get(commentURL)
		if err != nil {
			log.Println("Error fetching comment:", err)
			continue
		}
		defer commentResp.Body.Close()

		var comment Comment
		if err := json.NewDecoder(commentResp.Body).Decode(&comment); err != nil {
			log.Println("Error decoding comment response:", err)
			continue
		}

		comments = append(comments, comment)
	}
	return comments, nil
}
