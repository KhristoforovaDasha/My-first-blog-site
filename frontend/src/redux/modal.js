import { createSlice } from "@reduxjs/toolkit";

const initialState = {
    modalData: null,
    modalType: null,
};

const modalSlice = createSlice({
    name: "modal",
    initialState,
    reducers: {
        openModal: (state, action) => {
            state.modalType = action.payload.modalType;
            state.modalData = action.payload.modalData;
        },
        closeModal: (state) => {
            state.modalData = null;
            state.modalType = null;
        },
    },
});

export const modalReducer = modalSlice.reducer;
export const { openModal, closeModal } = modalSlice.actions;
