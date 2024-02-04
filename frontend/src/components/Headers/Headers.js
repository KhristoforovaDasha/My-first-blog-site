import "./Headers.css";
import { Link } from "react-router-dom";
import SearchOutlinedIcon from "@mui/icons-material/SearchOutlined";
import PersonOutlinedIcon from "@mui/icons-material/PersonOutlined";

export function Header() {

    return (
        <header className="header">
            <Link to="/create">
                <button>Создать публикацию</button>
            </Link>
            <div className="center">
                <div className="search">
                    <input type="text" placeholder="Search..."/>
                    <SearchOutlinedIcon/>
                </div>
                <p className="blog-name">Я создал Даше Канал</p>
            </div>
            <div className="right">
                <PersonOutlinedIcon/>
                <Link to="/registration">
                    <button>Регистрация/Вход</button>
                </Link>
            </div>
        </header>
    );
}

