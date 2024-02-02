import React, { useState } from 'react';
import { useAuth } from 'react-oidc-context';

const PingDropdown = ({ onDropdownChange }) => {
  const [selectedValue, setSelectedValue] = useState('');

  const handleDropdownChange = event => {
    const newValue = event.target.value;
    setSelectedValue(newValue);
    onDropdownChange(newValue);
  };

  return (
    <div className="dropdownContainer">
      <select id="pingDropdown" value={selectedValue} onChange={handleDropdownChange} className="pingDropdown">
        <option value="" className="dropdownOption">
          Select an option
        </option>
      </select>
    </div>
  );
};

const Ping = () => {
  const [selectedService, setSelectedService] = useState('');
  const [responseData, setResponseData] = useState(null);
  const auth = useAuth();

  const handleDropdownChange = value => {
    setSelectedService(value);
  };

  const handlePing = () => {
    if (selectedService) {
      // Make the API call using fetch
      let envString = 'REACT_APP_MICROSERVICE_' + selectedService.toUpperCase();
      fetch(process.env[envString] + `/api/services/${selectedService}`, {
        method: 'GET',
        headers: {
          Authorization: `Bearer ${auth.user.access_token}`,
          'Content-Type': 'application/json',
        },
      })
        .then(response => response.json())
        .then(data => {
          // Set the response data in state
          setResponseData(data.server);
        })
        .catch(error => {
          // Handle errors
          console.error('API Error:', error);
        });
    } else {
      console.error('Please select a service before pinging.');
      setResponseData('');
    }
  };

  return (
    <div className="container ping">
      <h2 style={{ color: 'black' }}> Ping your service</h2>
      <div className="select-service">
        <h4 style={{ color: 'black' }}> Select Your service to be pinged</h4>
        <PingDropdown onDropdownChange={handleDropdownChange} />
        <button className="ping-button" onClick={handlePing}>
          Ping
        </button>
        <div>
          <label style={{ color: 'black' }}>Response body</label>
          <div className="response-container">
            <h6>{responseData ? '{\n\t"server": "' + responseData + '"\n}' : 'No response yet'}</h6>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Ping;
