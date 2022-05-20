import Button from "./button"
import TextArea from "./textArea"

const TextAndButton = () => {
    return (
        <div className="textAndButton">
            <div className="promptAndTitle">
                <h2>Journal entry</h2>
                <hr></hr>
                 <p className="promptText"> What were you grateful for today? Start small, and work you way up.</p>
            </div>
            <Button />
        </div>
    )
}

export default TextAndButton