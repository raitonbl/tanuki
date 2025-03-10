import { useState } from "react";
import { Link, useLocation } from "react-router-dom";
import { Menu, X, User } from "lucide-react";

export function NavigationBar() {
    const [isOpen, setIsOpen] = useState(false);
    const [userMenuOpen, setUserMenuOpen] = useState(false);
    const location = useLocation();
    const isActive = (path: string) => location.pathname.startsWith(path);

    return (
        <nav className="bg-red-600 text-white shadow-md p-4 fixed top-0 left-0 w-full z-50">
            <div className="container mx-auto grid grid-cols-3 items-center gap-4">
                {/* Left Side - Logo */}
                <div className="flex items-center space-x-2">
                    <Link to="/" className="text-xl font-bold">
                        Tanuki
                    </Link>
                    <span className="text-gray-300">|</span>
                    <Link to="/" className=" font-bold text-white hover:text-white">Registry</Link>
                </div>

                {/* Center - Search Bar */}
                <div className="hidden md:block">
                    <input
                        type="text"
                        placeholder="Search"
                        className="bg-white px-4 py-2 rounded-md text-black w-full focus:outline-none"
                    />
                </div>

                {/* Right Side - Navigation Links */}
                <div className="hidden md:flex items-center justify-end space-x-6">
                    <Link to="/providers" className={`hover:text-gray-300 ${isActive('/providers') ? 'font-semibold' : ''}`}>
                        Providers
                    </Link>
                    <Link to="/modules" className={`hover:text-gray-300 ${isActive('/modules') ? 'font-semibold' : ''}`}>
                        Modules
                    </Link>

                    {/* User Icon with Dropdown */}
                    <div className="relative">
                        <button
                            className="flex items-center space-x-2"
                            onClick={() => setUserMenuOpen(!userMenuOpen)}
                        >
                            <User className="w-6 h-6" />
                        </button>
                        {userMenuOpen && (
                            <div className="absolute right-0 mt-2 w-48 bg-white text-black rounded-lg shadow-lg p-2">
                                <Link to="/" className="block px-4 py-2 hover:bg-gray-100">Logout</Link>
                            </div>
                        )}
                    </div>
                </div>

                {/* Mobile Menu Toggle */}
                <button
                    className="md:hidden flex justify-end"
                    onClick={() => setIsOpen(!isOpen)}
                >
                    {isOpen ? <X className="w-6 h-6" /> : <Menu className="w-6 h-6" />}
                </button>
            </div>

            {/* Mobile Dropdown Menu */}
            {isOpen && (
                <div className="md:hidden flex flex-col space-y-4 mt-4 p-4 bg-red-600 rounded-lg">
                    <input
                        type="text"
                        placeholder="Search"
                        className="bg-white px-4 py-2 rounded-md text-black w-full focus:outline-none"
                    />
                    <Link to="/providers" className="hover:text-gray-300">Providers</Link>
                    <Link to="/modules" className="hover:text-gray-300">Modules</Link>
                </div>
            )}
        </nav>
    );
}
