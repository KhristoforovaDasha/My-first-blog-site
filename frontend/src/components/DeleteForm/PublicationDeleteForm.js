export function PublicationDeleteForm(props) {
    const {onDeleteSuccess, formTitle} = props;
    return (
        <main className="main-publication" style={{display: "grid", gap:"100px"}}>
            <h2>{formTitle}</h2>
            <button style={{width: "200px", justifySelf:"center"}} onClick={(event) => {
                event.preventDefault();
                onDeleteSuccess();
            }}>Да</button>
        </main>
    )

}
