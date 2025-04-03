import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { getNewsByCategory } from "../api/news";
import NewsCard from "../components/NewsCard";

const Category = () => {
  const { category } = useParams();
  const [news, setNews] = useState([]);

  useEffect(() => {
    const fetchNews = async () => {
      const data = await getNewsByCategory(category);
      setNews(data);
    };
    fetchNews();
  }, [category]);

  return (
    <div>
      <h2>News in {category}</h2>
      {news.length === 0 ? (
        <p>Loading...</p>
      ) : (
        news.map((item) => <NewsCard key={item.id} title={item.title} url={item.url} />)
      )}
    </div>
  );
};

export default Category;
