import './App.css';
import {Header} from "./components/Headers/Headers";
import {BrowserRouter, Routes, Route} from "react-router-dom";
import {Main} from "./components/Main/Main";
import {CreatePublication} from "./components/CreatePublication/CreatePublication";
import {Publication} from "./components/Publication/Publication";
import {Modal} from "./components/Modal/Modal";
import {RegistrationAndSignIn} from "./components/RegistrationAndSignIn/Registration";

function App() {

    return (
        <BrowserRouter>
            <Header />
            <main className="main">
                <Routes>
                    <Route path="/" element={<Main />} />
                    <Route path="/create" element={<CreatePublication />} />
                    <Route path="/registration" element={<RegistrationAndSignIn />} />
                    <Route path="/posts/:id" element={<Publication />} />
                </Routes>
            </main>
            <Modal />
        </BrowserRouter>
    );
}

export default App;
