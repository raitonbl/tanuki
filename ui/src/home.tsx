import { Link } from "react-router-dom";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faGlobe, faFileZipper } from '@fortawesome/free-solid-svg-icons';
import { useState } from "react";

export default function Home() {
    // @ts-ignore
    const [providersCount, setProvidersCount] = useState(0);
    // @ts-ignore
    const [modulesCount, setModulesCount] = useState(0);

    return (
        <div className="flex flex-col items-center justify-center h-screen bg-white text-gray-900">
            <h1 className="text-5xl font-bold font-serif text-red-600">Tanuki Registry</h1>
            <p className="text-lg text-center mt-4 max-w-2xl leading-relaxed font-bold">
                Discover hosted OpenTofu providers that are available for different resource types,
                or find modules for quickly deploying common infrastructure configurations.
            </p>
            <div className="mt-8 flex justify-center gap-6">
                <Link to="/providers" className="flex items-center gap-2 outline-2 outline-offset-2 outline-red-600 bg-white text-red-600 hover:text-white hover:bg-red-600 px-6 py-3 rounded-lg font-semibold shadow-lg text-lg transition">
                    <FontAwesomeIcon icon={faGlobe} /> Browse Providers
                </Link>
                <Link to="/modules" className="flex items-center gap-2 outline-2 outline-offset-2 outline-red-600 bg-white text-red-600 hover:text-white hover:bg-red-600 px-6 py-3 rounded-lg font-semibold shadow-lg text-lg transition">
                    <FontAwesomeIcon icon={faFileZipper} /> Browse Modules
                </Link>
            </div>
            <p className="mt-6 text-sm font-bold text-red-600">
                {providersCount} Providers, {modulesCount} Modules
            </p>
        </div>
    );
}
