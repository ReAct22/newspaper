import React, {useEffect, useState} from "react";
import { getNews } from "../api/news";
import NewsCard from "../components/NewsCard";

const Home = () => {
    const [news, setNews] = useState([]);

    useEffect(() => {
        const fetchNews = async () => {
            const data = await getNews();
            setNews(data);
        };
        fetchNews();
    }, []);

    return (
        <div>
            <h2>Latest news</h2>
            {news.length === 0 ? (
                <p>Loading .....</p>
            ) : (
                news.map((item) => <NewsCard key={item.id} title={item.title} url={item.url} />)
            )}
        </div>
    );
}

export default Home