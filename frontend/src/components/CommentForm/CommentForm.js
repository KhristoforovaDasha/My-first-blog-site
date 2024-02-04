import { useState } from "react";
export function CommentForm(props) {
    const {onSuccess} = props;
    const [comment, setComment] = useState();
    return (
        <main className="main-publication">
            <h1>Оставьте комментарий</h1>
            <form>
                <textarea name="comment" value={comment}
                          onChange={(event) => setComment(event.target.value)}/>
                <div>
                    <button
                        type="submit"
                        id="create-button"
                        onClick={(event) => {
                            event.preventDefault();
                            onSuccess(comment);
                        }}
                    >
                        Добавить
                    </button>
                </div>
            </form>
        </main>
    )
}