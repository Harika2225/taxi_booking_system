import './Navigation.css';
import { Navbar } from 'react-bootstrap';
import { BrowserRouter as Router } from 'react-router-dom';
import { RiAccountCircleFill } from "react-icons/ri";
import { FaHome } from "react-icons/fa";

export const Navigation = () => {
 
  return (
    <Router>
    <Navbar className='navBar'>
      <h3 style={{color:'white'}}>Taxi</h3>
      <Navbar.Collapse className="justify-content-end">
        <Navbar.Text>
          <FaHome className='iconStyle' />
          <RiAccountCircleFill className='iconStyle'/>
        </Navbar.Text>
      </Navbar.Collapse>
    </Navbar>
  </Router>
  );
};
