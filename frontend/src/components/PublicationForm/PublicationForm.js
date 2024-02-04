import { useState } from "react";
import "./PublicationForm.css";
export function PublicationForm(props) {
    const { onSuccess, formTitle, defaultTitle, defaultText } = props;
    const [title, setTitle] = useState(defaultTitle);
    const [post_text, setText] = useState(defaultText);
    console.log("default", defaultTitle, defaultText)

    return (
        <main className="main-publication">
            <h1>{formTitle}</h1>
            <div className="publication-create">
                <form>
                    <div className="form-item-title">
                        <label id="title">Название</label>
                        <input name="title" value={title}
                               onChange={(event) => setTitle(event.target.value)} />
                    </div>
                    <div className="form-item-descr">
                        <label id="post_text">Описание</label>
                        <textarea
                            name="post_text"
                            value={post_text}
                            onChange={(event) => setText(event.target.value)}
                        />
                    </div>
                    <div>
                        <button
                            type="submit"
                            id="create-button"
                            onClick={(event) => {
                                event.preventDefault();
                                console.log("create-button", JSON.stringify({title, post_text}))
                                onSuccess(title, post_text);
                            }}
                        >
                            {formTitle}
                        </button>
                    </div>
                </form>
            </div>
        </main>
    );
}
