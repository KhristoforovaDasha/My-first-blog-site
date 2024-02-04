import { useParams } from "react-router-dom";
import { PublicationItem } from "../PublicationItem/PublicationItem";
import "./Publication.css"
import {useSelector} from "react-redux";


export function Publication() {
    const params = useParams();
    const post = useSelector((state) => state.publication.pubs.get(Number(params.id)) || null)
    return (
        <div className="publication-page">
            <PublicationItem post={post} />
        </div>
    );
}
