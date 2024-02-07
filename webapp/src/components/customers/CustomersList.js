import React, { useState, useEffect } from 'react';
import Modal from './Modal';
import { MdDeleteOutline, MdEdit } from 'react-icons/md';
import { useAuth } from 'react-oidc-context';

const CustomersList = ({ customersApp }) => {
  const [customers, setCustomers] = useState([]);
  const [showAddCustomerPopup, setShowAddCustomerPopup] = useState(false);
  const [newCustomer, setNewCustomer] = useState({ firstName: '', lastName: '', email: '', phone: '', address:'' });
  const [openModal, setOpenModal] = useState(false);
  const [editingCustomer, setEditingCustomer] = useState(null);
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
        setCustomers((prevCustomers) => [...prevCustomers, responseData]);
      } else {
        const errorText = await response.text();
        console.error('Error submitting data:', errorText);
      }
    } catch (error) {
      console.error('Error:', error.message);
    }
  };

  const handleUpdateCustomer = async (id, data) => {
    try {
      let envString = 'REACT_APP_MICROSERVICE_' + customersApp.toUpperCase();
      const response = await fetch(`${process.env[envString]}/api/updateCustomerById/${id}`, {
        method: 'PUT',
        headers: {
          Authorization: `Bearer ${auth.user.access_token}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });
      console.log(response,"response")
      if (response.ok) {
        const updatedCustomer = await response.json();
        // Assuming the update was successful, update the local state with the modified customer
        setCustomers((prevCustomers) =>
          prevCustomers.map((customer) => (customer.id === id ? updatedCustomer : customer))
        );
        // Close the modal
        setOpenModal(false);
      } else {
        const errorText = await response.text();
        console.error('Error updating customer:', errorText);
      }
    } catch (error) {
      console.error('Error updating customer:', error);
    }
  };
  

  const handleEditCustomer = (customer) => {
    // Set the customer being edited in the state and open the modal
    setEditingCustomer(customer);
    setOpenModal(true);
  };

  const handleCloseModal = () => {
    // Clear the editingCustomer state when closing the modal
    setEditingCustomer(null);
    setOpenModal(false);
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
      <Modal isOpen={openModal} onClose={handleCloseModal} 
      // onSubmit={handleSubmit} 
       onSubmit={(data) => {
          // If editingCustomer is set, call handleUpdateCustomer; otherwise, call handleSubmit
          editingCustomer ? handleUpdateCustomer(editingCustomer.id, data) : handleSubmit(data);
        }}
        initialValues={editingCustomer} // Pass the initialValues to the Modal for pre-filling
      />
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
                <MdEdit style={{ fontSize: 20, color: 'red', cursor: 'pointer' }} onClick={() => handleEditCustomer(customer)}
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
