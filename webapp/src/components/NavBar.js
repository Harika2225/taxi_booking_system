import { useState, useEffect } from 'react';
import { Navbar, Nav, Container } from 'react-bootstrap';
import navIcon2 from '../assets/img/nav-icon2.svg';
import { BiChevronDown } from 'react-icons/bi';
import { useAuth } from 'react-oidc-context';
import { BrowserRouter as Router } from 'react-router-dom';
import { FaUser, FaTaxi } from "react-icons/fa";
import { MdPayments } from "react-icons/md";

export const NavBar = () => {
  const [activeLink, setActiveLink] = useState('home');
  const [scrolled, setScrolled] = useState(false);
  const [isSubMenuOpen, setIsSubMenuOpen] = useState(false);
  const [isSubMenuOpenForApp, setIsSubMenuOpenForApp] = useState(false);
  const auth = useAuth();

  useEffect(() => {
    const onScroll = () => {
      if (window.scrollY > 50) {
        setScrolled(true);
      } else {
        setScrolled(false);
      }
    };

    window.addEventListener('scroll', onScroll);

    return () => window.removeEventListener('scroll', onScroll);
  }, []);

  const toggleMenu = () => {
    console.log('Toggle menu swagger clicked');
    setIsSubMenuOpen(!isSubMenuOpen);
  };
  const toggleAppMenu = () => {
    console.log('Toggle menu app clicked');
    setIsSubMenuOpenForApp(!isSubMenuOpenForApp);
  };

  const onUpdateActiveLink = value => {
    console.log('Update active link clicked');
    setActiveLink(value);
  };
  if (auth.isLoading) {
    return <></>;
  }

  let isloggedIn;
  if (auth.isAuthenticated) {
    window.history.replaceState({}, document.title, window.location.pathname);
    isloggedIn = auth.isAuthenticated;
  }

  return (
    <Router>
      <Navbar expand="md" className={scrolled ? 'scrolled' : ''}>
        <Container>
          <Navbar.Brand href="/"></Navbar.Brand>
          <Navbar.Toggle aria-controls="basic-navbar-nav">
            <span className="navbar-toggler-icon"></span>
          </Navbar.Toggle>
          <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="ms-auto">
              <Nav.Link
                href="/"
                className={activeLink === 'home' ? 'active navbar-link' : 'navbar-link'}
                onClick={() => onUpdateActiveLink('home')}
              >
                Home
              </Nav.Link>
              <Nav.Link
                className={activeLink === 'app' ? 'active navbar-link' : 'navbar-link'}
                onClick={() => {
                  toggleAppMenu();
                  onUpdateActiveLink('app');
                }}
              >
                Taxi Booking <BiChevronDown />
              </Nav.Link>

              <Nav.Link
                href="/docs"
                className={activeLink === 'projects' ? 'active navbar-link' : 'navbar-link'}
                onClick={() => onUpdateActiveLink('projects')}
              >
                Docs
              </Nav.Link>
            </Nav>
            <div className={isSubMenuOpenForApp ? 'sub-menu-wrap open-menu' : 'sub-menu-wrap'} id="subMenu">
              <div class="sub-menu">
                <a href="/bookings/bookingmanagementmodule" class="sub-menu-link">
                <span><FaTaxi /></span>
                  <h5>BookingManagement</h5>
                </a>
                <hr></hr>
                <span></span>
                <a href="/customers/customermanagementmodule" class="sub-menu-link">
                  <span><FaUser /></span>
                  <h5>CustomerManagement</h5>
                  {/* <span>&gt;</span> */}
                </a>
                <hr></hr>
                <span></span>
                <a href="/drivers/drivermanagementmodule" class="sub-menu-link">
                <span><FaUser /></span>
                  <h5>DriverManagement</h5>
                  {/* <span>&gt;</span> */}
                </a>
                <hr></hr>
                <span></span>
                <a href="/payments/paymentmanagementmodule" class="sub-menu-link">
                  <span><MdPayments /></span>
                  <h5>PaymentManagement</h5>
                  {/* <span>&gt;</span> */}
                </a>
              </div>
            </div>

            <span className="navbar-text">
              <div className="social-icon">
                <a onClick={() => window.open(process.env.REACT_APP_WEDAA_GITHUB, '_blank')}>
                  <img src={navIcon2} alt="" />
                </a>
              </div>
              {isloggedIn ? (
                <a target="_blank" rel="noopener noreferrer">
                  <button
                    className="vvd"
                    onClick={() =>
                      auth.signoutRedirect({
                        post_logout_redirect_uri: process.env.REACT_APP_PROJECT_URL,
                      })
                    }
                  >
                    <span>Sign Out</span>
                  </button>
                </a>
              ) : (
                <a target="_blank" rel="noopener noreferrer">
                  <button className="vvd" onClick={() => auth.signinRedirect()}>
                    <span>Sign In</span>
                  </button>
                </a>
              )}
            </span>
          </Navbar.Collapse>
        </Container>
      </Navbar>
    </Router>
  );
};
