import { useDispatch, useSelector } from "react-redux";
import { closeModal } from "../../redux/modal";
import "./Modal.css";
import { PublicationForm } from "../PublicationForm/PublicationForm";
import {CommentForm} from "../CommentForm/CommentForm";
import {PublicationDeleteForm} from "../DeleteForm/PublicationDeleteForm";

export function Modal() {
    const dispatch = useDispatch();
    const modalData = useSelector((state) => state.modal.modalData);
    const modalType = useSelector((state) => state.modal.modalType);

    if (!modalData || !modalType) {
        return null;
    }

    let modal = null;

    switch (modalType) {
        case "editPublication": {
            modal = <PublicationForm {...modalData} />;
            break;
        }
        case "addComment": {
            modal = <CommentForm {...modalData} />;
            break;
        }
        case "deletePublication": {
            modal = <PublicationDeleteForm {...modalData}/>;
            break;
        }
        default: {
            modal = null;
        }
    }

    if (!modal) {
        return null;
    }

    return (
        <div className="modalContainer" onClick={() => dispatch(closeModal())}>
            <div className="modal" onClick={(event) => event.stopPropagation()}>
                {modal}
            </div>
        </div>
    );
}
