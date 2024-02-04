import {useSelector} from "react-redux";

export const CommentsItem = (props) => {
    console.log(props.id)
    const comment = useSelector((state) => state.publication.comments[props.id]);
    if (!comment) {
        return null;
    }
    console.log(comment)

    return (
        <div className="comment">
            <div>{comment.text}</div>
        </div>
    );

};
