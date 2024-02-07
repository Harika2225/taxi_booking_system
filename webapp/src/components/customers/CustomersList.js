import React, { useState, useEffect } from 'react';
import Modal from './Modal';
import { MdDeleteOutline, MdEdit } from 'react-icons/md';
import { useAuth } from 'react-oidc-context';

const CustomersList = ({ customersApp }) => {
  const [customers, setCustomers] = useState([]);
  const [showAddCustomerPopup, setShowAddCustomerPopup] = useState(false);
  const [newCustomer, setNewCustomer] = useState({ firstName: '', lastName: '', email: '', phone: '', address:'' });
  const [openModal, setOpenModal] = useState(false);
  const auth = useAuth();

  useEffect(() => {
    // Fetch customers from the backend API
    const getCustomers = async () => {
      try {
        let envString = 'REACT_APP_MICROSERVICE_' + customersApp.toUpperCase();
        const response = await fetch(process.env[envString] + `/api/getCustomer`, {
          method: 'GET',
          headers: {
            Authorization: `Bearer ${auth.user.access_token}`,
            'Content-Type': 'application/json',
          },
        }); // Replace with your actual API endpoint
        const data = await response.json();
        console.log(data,"pppp")
        if (data != null) setCustomers(data); // Assuming the API response is an array of customers
      } catch (error) {
        console.error('Error fetching customers:', error);
      }
    };
    getCustomers();
  }, []);

  const handleSubmit = async Data => {
    try {
      // Make API call to post the collected data
      let envString = 'REACT_APP_MICROSERVICE_' + customersApp.toUpperCase();
      const response = await fetch(process.env[envString] +`/api/createCustomer`, {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${auth.user.access_token}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(Data),
      });
      console.log(response, Data)
      if (response.ok) {
        console.log(response,"response")
        const responseData = await response.json(); // Assuming the response is in JSON format
        // Use a callback function with setCustomers to ensure you have the latest state
        setCustomers(prevCustomers => [
          ...prevCustomers,
          {
            id: responseData.id,
            firstName: responseData.firstName,
            lastName: responseData.lastName,
            email: responseData.email,
            phone: responseData.phone,
            address: responseData.address
          },
        ]);
        console.log(customers
          ,"prevsdlfj")
      } else {
        // Handle error, maybe show an error message
        const errorText = await response.text();
        console.error('Error submitting data:', errorText);
      }
    } catch (error) {
      console.error('Error:', error.message);
    }
  };

  // const handleInputChange = e => {
  //   const { name, value } = e.target;
  //   setNewCustomer({ ...newCustomer, [name]: value });
  // };

  // const handlePopupClose = () => {
  //   setShowAddCustomerPopup(false);
  //   // setNewCustomer({ firstName: '', lastName: '', email: '', phone: '', address:'' });
  // };

  // const handleAddNote = Data => {
  //   // Add logic to send newCustomer data to the backend via API
  //   console.log('New Note:', Data);
  // };

  const handleUpdateCustomer = async id => {
    console.log(id,"akjfsjskdj")
    try {
      let envString = 'REACT_APP_MICROSERVICE_' + customersApp.toUpperCase();
      await fetch(`${process.env[envString]}/api/updateCustomerById?id=${id}`, {
        method: 'PATCH',
        headers: {
          Authorization: `Bearer ${auth.user.access_token}`,
          'Content-Type': 'application/json',
        },
      });
      // Assuming the delete was successful, update the local state
      setCustomers(prevCustomers => prevCustomers.filter(customer => customer.id !== id));
    } catch (error) {
      console.error('Error deleting note:', error);
    }
  };

  const handleDeleteCustomer = async id => {
    console.log(id,"psdapsdasp")
    // Add logic to delete note using the backend API
    try {
      let envString = 'REACT_APP_MICROSERVICE_' + customersApp.toUpperCase();
      await fetch(`${process.env[envString]}/api/deleteCustomerById?id=${id}`, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${auth.user.access_token}`,
          'Content-Type': 'application/json',
        },
      });
      // Assuming the delete was successful, update the local state
      setCustomers(prevCustomers => prevCustomers.filter(customer => customer.id !== id));
    } catch (error) {
      console.error('Error deleting note:', error);
    }
  };

  return (
    <div className="container ping">
      <button className="ping-button" onClick={() => setOpenModal(true)} style={{ alignItems: 'revert-layer' }}>
        Add Customer
      </button>
      <Modal isOpen={openModal} onClose={() => setOpenModal(false)} onSubmit={handleSubmit} />
      {/* {showAddCustomerPopup && (
        <div className="popup">
          <div className="popup-content">
            <span className="close" onClick={handlePopupClose}>
              &times;
            </span>
            <div>sjfksldjk</div>
            <div style={{ padding: '20px', textAlign: 'center', color: 'black' }}>
              <label style={{ display: 'block', marginBottom: '10px' }}>
                Firstttttttt Name:
                <input
                  type="text"
                  name="firstName"
                  value={newCustomer.firstName}
                  onChange={handleInputChange}
                  style={{
                    width: '100%',
                    padding: '8px',
                    boxSizing: 'border-box',
                  }}
                />
              </label>
              <label style={{ display: 'block', marginBottom: '10px' }}>
                Last Name:
                <input
                  type="text"
                  name="lastName"
                  value={newCustomer.lastName}
                  onChange={handleInputChange}
                  style={{
                    width: '100%',
                    padding: '8px',
                    boxSizing: 'border-box',
                  }}
                />
              </label>
              <label style={{ display: 'block', marginBottom: '10px' }}>
                Email:
                <input
                  type="text"
                  name="email"
                  value={newCustomer.email}
                  onChange={handleInputChange}
                  style={{
                    width: '100%',
                    padding: '8px',
                    boxSizing: 'border-box',
                  }}
                />
              </label>
              <label style={{ display: 'block', marginBottom: '10px' }}>
                Phone Number:
                <input
                  type="text"
                  name="phone"
                  value={newCustomer.phone}
                  onChange={handleInputChange}
                  style={{
                    width: '100%',
                    padding: '8px',
                    boxSizing: 'border-box',
                  }}
                />
              </label>
              <label style={{ display: 'block', marginBottom: '10px' }}>
                Address:
                <input
                  type="text"
                  name="address"
                  value={newCustomer.address}
                  onChange={handleInputChange}
                  style={{
                    width: '100%',
                    padding: '8px',
                    boxSizing: 'border-box',
                  }}
                />
              </label>
              <button
                style={{
                  backgroundColor: '#4CAF50',
                  color: 'white',
                  padding: '10px 20px',
                  borderRadius: '5px',
                  cursor: 'pointer',
                }}
              >
                Submit
              </button>
            </div>
          </div>
        </div>
      )} */}

      <table style={{ borderCollapse: 'collapse', width: '100%' }}>
        <thead>
          <tr style={{ color: 'black', backgroundColor: '#f2f2f2' }}>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Sno</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>First Name</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Last Name</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Email</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Phone</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Address</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}></th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}></th>
          </tr>
        </thead>
        <tbody>
          {customers.map((customer, index) => (
            <tr
              key={index}
              style={{
                color: 'black',
                backgroundColor: index % 2 === 0 ? '#f9f9f9' : 'white',
              }}
            >
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{index + 1}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{customer.firstName}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{customer.lastName}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{customer.email}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{customer.phone}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{customer.address}</td>
              <td
                style={{
                  padding: '10px',
                  border: '1px solid #dddddd',
                  textAlign: 'center',
                }}
              >
                <MdEdit style={{ fontSize: 20, color: 'red', cursor: 'pointer' }} onClick={() => handleUpdateCustomer(customer.id)}
                 />{' '}
              </td>
              <td
                style={{
                  padding: '10px',
                  border: '1px solid #dddddd',
                  textAlign: 'center',
                }}
              >
                <MdDeleteOutline style={{ fontSize: 20, color: 'red', cursor: 'pointer' }} onClick={() => handleDeleteCustomer(customer.id)} />{' '}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default CustomersList;
