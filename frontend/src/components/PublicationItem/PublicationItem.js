import {Link} from "react-router-dom";
import {useDispatch} from "react-redux";
import "./PublicationItem.css";
import { openModal, closeModal } from "../../redux/modal";
import FavoriteIcon from '@mui/icons-material/Favorite';
import ChatBubbleOutlineIcon from '@mui/icons-material/ChatBubbleOutline';
import {updatePublication} from "../../redux/publication";
import {useState} from "react";


export function PublicationItem(props) {
    const { post } = props;
    const dispatch = useDispatch();
    const {
        id,
        title,
        post_text,
        imageUrl,
        likesCount,
    } = post;

    const [likes, setLikes] = useState(likesCount)
    const [title_post, setTitle] = useState(title)
    const [text, setText] = useState(post_text)

    console.log(post)

    const handleSuccess = async (editedTitle, editedDescription) => {
        setTitle(editedTitle)
        setText(editedDescription)
        dispatch(closeModal());
        await handleEdit({
            id,
            title: editedTitle,
            post_text: editedDescription,
            imageUrl,
            likesCount,
        });
    };
    const handleEdit = async (post) => {
        console.log(post)
        const response = await fetch(`http://localhost:3001/posts/${id}/update`, {
            method: "POST",
            body: JSON.stringify(post),
            headers: {
                "Content-Type": "application/json",
            },
        });
        const updatedPublication = await response.json();
        console.log("updatedPublication", updatedPublication)
        setLikes(likes + 1)
        dispatch(updatePublication(updatedPublication["data"]));
    };

    const handleDelete = async () => {
        await fetch(`http://localhost:3001/posts/${id}/delete`, {
            method: "GET",
        });
        window.location.href = "/";
    };

    const handleDeleteSuccess = async () => {
        dispatch(closeModal());
        await handleDelete();
    };

    const editModal = () => {
        dispatch(
            openModal({
                modalType: "editPublication",
                modalData: {
                    formTitle: "Редактировать",
                    defaultTitle: title_post,
                    defaultText: text,
                    onSuccess: handleSuccess,
                },
            })
        )
    }

    const deleteModal = () => {
        dispatch(
            openModal({
                modalType: "deletePublication",
                modalData: {
                    formTitle: "Вы точно хотите удалить эту публикацию?",
                    onDeleteSuccess: handleDeleteSuccess,
                },
            })
        )
    }

    return (
        <div className="post">
            <p>{title_post}</p>
            <Link to={`posts/${id}`}>
                <img src={imageUrl} alt="" />
            </Link>
            <b>{text}</b>
            <div className="reactions">
                <div className="like">
                    {likes}
                    <FavoriteIcon onClick={() =>
                        handleEdit({ id, title, post_text, imageUrl, likesCount: likesCount + 1 })
                    }/>
                </div>
                <div className="comment">
                    <ChatBubbleOutlineIcon />
                </div>
                <div className="redact">
                    <button onClick={() =>
                        deleteModal()
                    }>Удалить</button>
                    <button onClick={() =>
                        editModal()
                    }
                    >Редактировать</button>
                </div>
            </div>
        </div>
    );
}
