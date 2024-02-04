import {useState} from "react";
import "./Registration.css"
export function RegistrationAndSignIn() {
    const [login, setLogin] = useState();
    const [password, setPassword] = useState();
    const handleUser = async (login, password, action) => {
        const values = {
            login,
            password,
        };
        console.log(login, password)
        const response = await fetch(`http://localhost:3001/user/${action}`, {
            method: "POST",
            body: JSON.stringify(values),
            headers: {
                "Content-Type": "application/json",
            },
        });
        console.log(response.status)
        if (response.status !== 200 && response.status !== 201) {
            alert("Что-то пошло не так, возможно вы ввели что-то неправильно или длина пароля < 8 символов " +
                "или такое имя пользователя уже занято")
        } else {
            alert("Ура, вы залогинились")
            window.location.href = "/";
        }
    };

    return (
        <main className="autorization">
            <h1>Авторизуйтесь</h1>
            <div className="auth">
                <form>
                    <div className="login">
                        <label id="login">Логин</label>
                        <input name="login" value={login}
                               onChange={(event) => setLogin(event.target.value)} />
                    </div>
                    <div className="password">
                        <label id="password">Пароль</label>
                        <input
                            name="password"
                            value={password}
                            onChange={(event) => setPassword(event.target.value)}
                        />
                    </div>
                    <div>
                        <button
                            type="submit"
                            id="create-button"
                            onClick={(event) => {
                                event.preventDefault();
                                handleUser(login, password, "login");
                            }}
                        >
                            Войти
                        </button>
                        <button
                            type="submit"
                            id="create-button"
                            onClick={(event) => {
                                event.preventDefault();
                                handleUser(login, password, "registration");
                            }}
                        >
                            Зарегистрироваться
                        </button>
                    </div>
                </form>
            </div>
        </main>
    );
}