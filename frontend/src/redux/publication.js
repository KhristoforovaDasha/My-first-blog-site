import { createSlice } from "@reduxjs/toolkit";


const initialState = {
    publicationsList: [],
    pubs: {}
};

const publicationSlice = createSlice({
    name: "publication",
    initialState,
    reducers: {
        setPublications: (state, action) => {
            state.publicationsList = action.payload.publications;
            state.pubs = new Map()
            for (let id in state.publicationsList) {
                const post = action.payload.publications[id]
                state.pubs.set(post["id"], post)
            }
            console.log("action", action.payload.publications)
            console.log("action", state.pubs)
        },
        updatePublication: (state, action) => {
            for (let pub in state.publicationsList) {
                if (action.payload.id === pub["id"]) {
                    pub = action.payload;
                    console.log("pub")
                    console.log(pub)
                }
            }
            state.pubs[action.payload.id] = action.payload;
            //console.log("update")
            //console.log(action.payload)
        },
    },
});

export const publicationReducer = publicationSlice.reducer;
export const { setPublications, updatePublication} =
    publicationSlice.actions;

