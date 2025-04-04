import React from "react";
import { Link } from "react-router-dom";
import { Home, Code, Shield, BrainCircuit, MessageSquare } from "lucide-react";

const Navbar = () => {
  return (
    <nav className="bg-blue-600 text-white p-4 shadow-lg">
      <div className="container mx-auto flex justify-between items-center">
        <Link to="/" className="text-xl font-bold flex items-center gap-2">
          <Home className="w-6 h-6" /> Home
        </Link>
        <div className="flex gap-6">
          <Link to="/category/programming" className="flex items-center gap-2 hover:text-gray-200">
            <Code className="w-5 h-5" /> Programming
          </Link>
          <Link to="/category/cybersecurity" className="flex items-center gap-2 hover:text-gray-200">
            <Shield className="w-5 h-5" /> Cybersecurity
          </Link>
          <Link to="/category/ai" className="flex items-center gap-2 hover:text-gray-200">
            <BrainCircuit className="w-5 h-5" /> AI
          </Link>
          <Link to="/comment" className="flex items-center gap-2 hover:text-gray-200">
            <MessageSquare className="w-5 h-5" /> Comment
          </Link>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
