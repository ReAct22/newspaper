import React from "react";

const NewsCard = ({title, url}) => {
    return (
        <div className="new-card">
            <h3>{title}</h3>
            <a href={url} target="_blank" rel="noopener noreferrer">read more</a>
        </div>
    )
}

export default NewsCard;