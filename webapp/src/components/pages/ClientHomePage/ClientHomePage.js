import React from "react";
import "./ClientHomePage.css";
import taxi from "../../../assets/img/taxi.png";
import taxi1 from "../../../assets/img/taxi1.jpg";
import taxi2 from "../../../assets/img/taxi2.jpg";
import taxi3 from "../../../assets/img/taxi3.jpg";

import { ArrowRightCircle } from "react-bootstrap-icons";

function ClientHomePage() {
  return (
    <div className="container">
      <div className="block">
        <div className="left">
          <div className="justify-content-center">
            <h1> Go anywhere with a tap</h1>
            <br />
            <p>Request a ride, hop in, and go.</p>
            <div className="location-input">
              <input type="text" placeholder="Enter pickup location" />
            </div>
            <div className="location-output">
              <input type="text" placeholder="Enter destination"  />
            </div>
            {/* <div className="date">
              <input type="text" placeholder="Select date"  />
            </div>
            <div className="time">
              <input type="text" placeholder="Select time"  />
            </div> */}
            <br/>
            <button className="book">Book</button>
          </div>
        </div>
        <div className="right">
          <img src={taxi} />
        </div>
      </div>
      <br />
      <div className="block">
        <div className="right">
          <img src={taxi1} />
        </div>
        <div className="left">
          <div className="justify-content-center">
            <h1>Why we???</h1>
            <br />
            <span>
              Our goal is to bring riders and drivers together on one platform.
              <br />
              <br />
              We make traveling easier in this digital age by allowing you to
              move from one location to another with a single tap.
            </span>
            <button className="get-started-button">
              Get Started
              <ArrowRightCircle className="icon" size={25} />
            </button>
          </div>
        </div>
      </div>
      <br />
      <div className="block">
        <div className="left">
          <span>
            Experience peace of mind with our taxi booking service, offering
            real-time tracking capabilities that allow you to monitor your
            driver's whereabouts and access comprehensive driver details.
            <br />
            <br />
            Whether it's knowing their name, vehicle information, or estimated
            time of arrival, stay informed every step of the way for a
            hassle-free ride.
          </span>
        </div>
        <div className="right">
          <img src={taxi2} />
        </div>
      </div>
      <br />
      <div className="block">
        <div className="right">
          <img src={taxi3} />
        </div>
        <div className="left">
          <span>
            Sit back and unwind as you embark on your journey with us. Our
            commitment to safety is unparalleled, with all our drivers
            undergoing rigorous verification procedures.
            <br />
            <br />
            Rest assured knowing that every driver behind the wheel is
            thoroughly vetted, ensuring your security and comfort throughout
            your ride.
            <br />
            <br />
            Travel with confidence and reach your destination with ease, knowing
            you're in capable hands.
          </span>
        </div>
      </div>
    </div>
  );
}

export default ClientHomePage;
