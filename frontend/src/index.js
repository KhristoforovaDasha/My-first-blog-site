import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import { Provider } from "react-redux";
import { store } from "./store";
import {enableMapSet} from "immer";

const root = ReactDOM.createRoot(document.getElementById('root'));
enableMapSet()

root.render(
    <Provider store={store}>
        <App />
    </Provider>
);
