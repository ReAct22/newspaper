import React from "react";
import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <nav>
      <Link to="/">Home</Link>
      <Link to="/category/programming">Programming</Link>
      <Link to="/category/cybersecurity">Cybersecurity</Link>
      <Link to="/category/ai">AI</Link>
      <Link to="/comment">Comment</Link>
    </nav>
  );
};

export default Navbar;
