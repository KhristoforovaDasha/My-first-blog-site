import {useEffect} from "react";
import {PublicationItem} from "../PublicationItem/PublicationItem";
import {ApiService} from "../../services/ApiService";
import "./Main.css"
import {setPublications} from "../../redux/publication";
import {useDispatch, useSelector} from "react-redux";

export function Main() {
    const dispatch = useDispatch();
    const publicationsList = useSelector(
        (state) => state.publication.publicationsList
    );

    useEffect(() => {
        (async () => {
            const posts = await ApiService("posts");
            console.log("posts", posts)

            dispatch(
                setPublications({
                    publications: posts["data"],
                })
            );
        })();
    }, []);

    console.log("main", publicationsList)

    return (
        <div className="list">
            {publicationsList.map((item) => (
                <PublicationItem key={item} post={item} />
            ))}
        </div>
    );
}

