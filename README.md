Project Documentation: Newspaper Assignment

Introduction
This project consists of backend_news built with Golang and frontend_news using ReactJS. The backend serves as a gateway between the news database and the frontend, while the frontend is responsible for displaying news and comments.

Project Structure
project_root/
│── backend_news/ # Backend with Golang
│ ├── config/ # Database & environment configuration
│ ├── routes/ # API routing
│ ├── service/ # Business logic (fetching data, caching, etc.)
│ ├── main.go # Backend entry point
│── frontend_news/ # Frontend with ReactJS
│ ├── src/
│ │ ├── components/ # UI components (CommentList, NewsItem, etc.)
│ │ ├── pages/ # Main & news detail pages
│ │ ├── App.jsx # Main React component
│ │ ├── main.jsx # React entry point
│── README.md # Project documentation

Backend (Golang) - backend_news
Backend Installation
Ensure Golang is installed, then run:
cd backend_news
go mod init backend_news
go get -u github.com/gin-gonic/gin
go get -u github.com/go-sql-driver/mysql
go run main.go

API Endpoints
Method  Endpoint Description

GET     /news     Fetches the list of news articles

GET     /category/:category  Fetches news by category

GET     /comment?story_id=:id  Fetches comments for a specific news item

Database Connection (MySQL)
Edit config/database.go for MySQL configuration:
package config
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
    dsn := "user:password@tcp(localhost:3306)/news_db"
    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        panic(err)
    }
}

Security Middleware
Add middleware in main.go for security:

r.Use(func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Next()
})

Sql Query
##
CREATE DATABASE news_cache;

USE news_cache;

CREATE TABLE cached_news (
    id INT PRIMARY KEY,
    title VARCHAR(255),
    url TEXT,
    category VARCHAR(50),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);



Frontend (ReactJS) - frontend_news
Frontend Installation
cd frontend_news
npm create vite@latest . --template react
npm install
npm install axios react-router-dom
npm run dev

Frontend Folder Structure
frontend_news/
│── src/
│   ├── components/
│   │   ├── NewsCard.jsx   # Component to display news data from API
│   │   ├── CommentItem.jsx   # Component to display comments
│   │   ├── Navbar.jsx      # Component for navigation bar
│   ├── pages/
│   │   ├── Home.jsx    # Home page
│   │   ├── Category.jsx # Page for category-based news
│   │   ├── Comment.jsx # Page for displaying comments
│   ├── App.jsx               # Main React component
│   ├── main.jsx              # Root file for React

Fetching Data from Backend
import axios from "axios";
useEffect(() => {
    axios.get("http://localhost:8080/news")
        .then(response => setNews(response.data))
        .catch(error => console.error(error));
}, []);

Displaying Comments in CommentList.jsx
const fetchComments = async () => {
    try {
        const response = await axios.get(`http://localhost:8080/comment?story_id=${storyId}`);
        setComments(response.data);
    } catch (error) {
        console.error("Error fetching comments:", error);
    }
};
