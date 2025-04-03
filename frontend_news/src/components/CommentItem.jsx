
import React from "react";

const CommentItem = ({ comment }) => {
  return (
    <div className="border p-3 rounded-lg shadow-sm mb-2">
      <h4 className="font-semibold text-blue-600">@{comment.by}</h4>
      <p dangerouslySetInnerHTML={{ __html: comment.text }}></p>
    </div>
  );
};

export default CommentItem;
