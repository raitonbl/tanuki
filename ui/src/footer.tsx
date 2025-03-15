export default function Footer() {
    return (
        <footer className="w-full bg-white text-gray-600 text-sm py-4 border-t border-gray-300 flex flex-col items-center">
            <div className="flex flex-wrap justify-center gap-6">
                <a href="#" className="hover:underline">INTRO</a>
                <a href="#" className="hover:underline">DOCS</a>
                <a href="#" className="hover:underline">TERMS</a>
            </div>
            <div className="mt-2 flex items-center">
                <img src="/tanuki-logo.svg" alt="Tanuki Logo" className="h-4 mr-2"/>
                <span>© Tanuki 2025</span>
            </div>
        </footer>
    );
}
