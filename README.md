# Dokumentasi Proyek Backend & Frontend News

## 📌 Pendahuluan
Proyek ini terdiri dari **backend_news** menggunakan **Golang** dan **frontend_news** menggunakan **ReactJS**. Backend berfungsi sebagai gateway antara **Hacker News API** dan frontend, sedangkan frontend bertugas untuk menampilkan berita dan komentar dari Hacker News.

---

## 📂 Struktur Proyek
```
project_root/
│── backend_news/        # Backend dengan Golang
│   ├── config/          # Konfigurasi database & environment
│   ├── routes/          # Routing API
│   ├── service/         # Logika bisnis (fetch data, caching, dll)
│   ├── main.go          # Entry point backend
│── frontend_news/       # Frontend dengan ReactJS
│   ├── src/
│   │   ├── components/  # Komponen UI (CommentList, NewsItem, dll)
│   │   ├── pages/       # Halaman utama & detail berita
│   │   ├── App.jsx      # Komponen utama React
│   │   ├── main.jsx     # Entry point React
│── README.md            # Dokumentasi proyek
```

---

## 🚀 Backend (Golang) - `backend_news`

### **1️⃣ Instalasi Backend**
Pastikan **Golang** sudah terinstal, lalu jalankan:
```sh
cd backend_news
go mod init backend_news
go get -u github.com/gin-gonic/gin
go get -u github.com/go-sql-driver/mysql
go run main.go
```

### **2️⃣ API Endpoint**
| Method | Endpoint         | Deskripsi                                  |
|--------|----------------|--------------------------------------------|
| GET    | `/news`        | Mengambil daftar berita dari Hacker News  |
| GET    | `/categori/:categori` | Menampilkan berita berdasarkan kategori |
| GET    | `/comment?story_id=:id` | Mengambil komentar dari berita tertentu |

### **3️⃣ Koneksi Database (MySQL)**
Edit `config/database.go` untuk konfigurasi MySQL:
```go
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
```

### **4️⃣ Middleware Keamanan**
Tambahkan middleware di `main.go` untuk keamanan:
```go
r.Use(func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Next()
})
```

---

## 🎨 Frontend (ReactJS) - `frontend_news`

### **1️⃣ Instalasi Frontend**
Jalankan perintah berikut untuk memulai proyek React:
```sh
cd frontend_news
npm create vite@latest . --template react
npm install
npm install axios react-router-dom
npm run dev
```

### **2️⃣ Struktur Folder Frontend**
```
frontend_news/
│── src/
│   ├── components/
│   │   ├── NewsCard.jsx   # component Show Data news by API
│   │   ├── CommentItem.jsx   # Komponen detail commentar
|   |   ├── Navbar.jsx      #Component for navigation bar
│   ├── pages/
│   │   ├── Home.jsx    # Home page
|   |   ├── Category.jsx # page for show data by category
|   |   ├── Comment.jsx # Page for show data comment
│   ├── App.jsx               # Entry point React
│   ├── main.jsx              # Root file React
```

### **3️⃣ Fetch Data dari Backend**
#### 📌 Mengambil Daftar Berita
```jsx
import axios from "axios";
useEffect(() => {
    axios.get("http://localhost:8080/news")
        .then(response => setNews(response.data))
        .catch(error => console.error(error));
}, []);
```

#### 📌 Menampilkan Komentar di `CommentList.jsx`
```jsx
const fetchComments = async () => {
    try {
        const response = await axios.get(`http://localhost:8080/comment?story_id=${storyId}`);
        setComments(response.data);
    } catch (error) {
        console.error("Error fetching comments:", error);
    }
};
```

---

## 🎯 Kesimpulan
Dengan proyek ini, kita berhasil membangun aplikasi **news aggregator** yang menampilkan berita dan komentar dari **Hacker News** dengan **Golang sebagai backend** dan **ReactJS sebagai frontend**.

💡 **Selamat Coding! 🚀**

