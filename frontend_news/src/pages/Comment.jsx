import React, {useEffect, useState} from "react";
import { getComment } from "../api/news";
import CommentItem from "../components/CommentItem";

const Comment = () => {
    const [comment, setComments] = useState([]);

    useEffect(() => {
        const fetchComments = async () => {
            const data = await getComment();
            setComments(data)
        };

        fetchComments();
    }, []);

    return (
        <div className="mt-5">
            <h3 className="text-lg font-bold">Comment</h3>
            <div>
                {comment.length > 0 ? (
                    comment.map((comment) => <CommentItem key={comment.id} comment={comment} />) 
                ) : (
                    <p>Comment Not Exists</p>
                )}
            </div>
        </div>
    )
}

export default Comment