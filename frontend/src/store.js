import { configureStore } from "@reduxjs/toolkit";
import { publicationReducer } from "./redux/publication";
import {modalReducer} from "./redux/modal";

export const store = configureStore({
    reducer: {
        modal: modalReducer,
        publication: publicationReducer,
    },
    devTools: true,
});
