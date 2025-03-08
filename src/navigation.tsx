import { useState } from "react";
import { Link, useLocation } from "react-router-dom";
import { Menu, X } from "lucide-react";

export function NavigationBar() {
    const [isOpen, setIsOpen] = useState(false);
    const location = useLocation();
    const isActive = (path: string) => location.pathname.startsWith(path);

    return (
        <nav className="bg-sky-600 text-white shadow-md p-4 fixed top-0 left-0 w-full z-50 flex justify-between items-center">
            <div className="container mx-auto flex justify-between items-center">
                <div className="flex items-center space-x-4">
                    <Link to="/" className="text-xl text-white font-bold flex items-center space-x-2">
                        <span className="font-semibold">Tanuki</span>
                    </Link>
                    <span className="text-gray-300">|</span>
                    <Link to="/" className="text-gray-300 hover:text-white">Registry</Link>
                </div>

                <div className="hidden md:flex space-x-6 items-center">
                    <input
                        type="text"
                        placeholder="Search"
                        className="px-4 py-1 rounded-md text-black w-64 focus:outline-none"
                    />
                    <Link to="/providers" className="hover:text-gray-300">Providers</Link>
                    <Link to="/modules" className="hover:text-gray-300">Modules</Link>
                    <div className="w-8 h-8 rounded-full bg-white"></div>
                </div>

                <button
                    className="md:hidden p-2 focus:outline-none"
                    onClick={() => setIsOpen(!isOpen)}
                >
                    {isOpen ? <X className="w-6 h-6" /> : <Menu className="w-6 h-6" />}
                </button>
            </div>
            {isOpen && (
                <div className="md:hidden flex flex-col space-y-4 mt-4 p-4 bg-[#4b2ea8] rounded-lg">
                    <input
                        type="text"
                        placeholder="Search"
                        className="px-4 py-1 rounded-md text-black w-full focus:outline-none"
                    />
                    <Link to="/providers" className="hover:text-gray-300">Providers</Link>
                    <Link to="/modules" className="hover:text-gray-300">Modules</Link>
                </div>
            )}
        </nav>
    );
}
