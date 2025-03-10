import './App.css'
import {NavigationBar} from "./navigation.tsx";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./home.tsx";
import Footer from "./footer.tsx";

export default function App() {
    return (
        <Router>
            <NavigationBar />
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/providers" element={<h1>Providers</h1>} />
                <Route path="/providers/:name" element={<h1>Provider Details</h1>} />
                <Route path="/providers/:name/content" element={<h1>Provider Content</h1>} />
                <Route path="/modules" element={<h1>Modules</h1>} />
                <Route path="/modules/:name" element={<h1>Module Details</h1>} />
                <Route path="/modules/:name/content" element={<h1>Module Content</h1>} />
            </Routes>
            <Footer />
        </Router>
    );
}

