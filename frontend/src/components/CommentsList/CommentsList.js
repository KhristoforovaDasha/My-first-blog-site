import { useSelector } from "react-redux";
import { CommentsItem } from "../CommentsItem/CommentsItem";

export const CommentsList = (props) => {
    const commentsList = useSelector(
        (state) => (state.publication.pubs[props.id] || {}).comments || []
    );
    console.log(commentsList)

    return (
        <div>
            {commentsList.map((comment) => (
                <CommentsItem key={comment} id={comment} />
            ))}
        </div>
    );
};
