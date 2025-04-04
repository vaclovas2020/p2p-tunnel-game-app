import { useState } from "react"
import Welcome from "./pages/Welcome";

function App() {

    const [gameStarted, setGameStarted] = useState(false);

    return (
        <div className="w-full h-[100dvh]">
            {!gameStarted && (<Welcome />)}
        </div>
    )
}

export default App
