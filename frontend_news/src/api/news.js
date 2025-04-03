import axios from "axios";

const API_URL = "http://localhost:8080";

export const getNews = async () => {
    try {
        const response = await axios.get(`${API_URL}/news`);
        return response.data;
    } catch (error) {
        console.error("Error fetching news:", error);
        return [];
    }
}

export const getNewsByCategory = async (category) => {
    try {
        const response = await axios.get(`${API_URL}/category/${category}`);
        return response.data;
    } catch (error) {
        console.error("Error fetching category news:", error);
        return [];
    }
}

export const getComment = async () => {
    try {
        const response = await axios.get(`${API_URL}/comment?story_id=2921983`);
        return response.data;
    } catch(error){
        console.error("Error fetching Comment: ", error);
        return [];

    }
}