import { PublicationForm } from "../PublicationForm/PublicationForm";
import bato from "../PublicationItem/images/bato.jpg";
import garden from "../PublicationItem/images/garden.jpg";
import insomnia from "../PublicationItem/images/insomnia.jpg";
import sheregesh from "../PublicationItem/images/sheregesh.jpg";

const images = [bato, garden, insomnia, sheregesh]
function randomInteger(min, max) {
    let rand = min + Math.random() * (max + 1 - min);
    return Math.floor(rand);
}

export function CreatePublication() {
    const handleCreate = async (title, post_text) => {
        const imageUrl = images[randomInteger(0, 3)]
        const values = {
            title,
            post_text,
            imageUrl,
        };
        console.log(title, post_text)

        await fetch("http://localhost:3001/create", {
            method: "POST",
            body: JSON.stringify(values),
            headers: {
                "Content-Type": "application/json",
            },
        });
        window.location.href = "/";
    };

    return (
        <PublicationForm onSuccess={handleCreate} formTitle="Создать публикацию" />
    );
}
