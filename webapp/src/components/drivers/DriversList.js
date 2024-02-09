import React, { useState, useEffect } from 'react';
import Modal from './Modal'
import { MdDeleteOutline, MdEdit } from 'react-icons/md';
import { useAuth } from 'react-oidc-context';

const DriversList = ({ driversApp }) => {
  const [drivers, setDrivers] = useState([]);
  const [editingDriver, setEditingDriver] = useState(null);
  const [openModal, setOpenModal] = useState(false);
  const auth = useAuth();

  useEffect(() => {
    const getDrivers = async () => {
      try {
        let envString = 'REACT_APP_MICROSERVICE_' + driversApp.toUpperCase();
        const response = await fetch(process.env[envString] + `/api/getDriver`, {
            method: 'GET',
            headers: {
                Authorization: `Bearer ${auth.user.access_token}`,
                'Content-Type': 'application/json',
            },
        });
        const data = await response.json();
        console.log(data, "pppp");
        if (data != null) setDrivers(data);
        } catch (error) {
            console.error('Error fetching drivers:', error);
        }
    };
    getDrivers();
  }, []);

  const handleSubmit = async Data => {
    try {
      let envString = 'REACT_APP_MICROSERVICE_' + driversApp.toUpperCase();
      const response = await fetch(process.env[envString] + `/api/createDriver`, {
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
        console.log(responseData,"rd")
        setDrivers((prevDrivers) => [...prevDrivers, responseData]);
      } else {
        const errorText = await response.text();
        console.error('Error submitting data:', errorText);
      }
    } catch (error) {
      console.error('Error:', error.message);
    }
  };

  const handleUpdateDriver = async (id, data) => {
    try {
        let envString = 'REACT_APP_MICROSERVICE_' + driversApp.toUpperCase();
        const response = await fetch(`${process.env[envString]}/api/updateDriverById/${id}`, {
            method: 'PUT',
            headers: {
                Authorization: `Bearer ${auth.user.access_token}`,
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data), // Use the formData state
        });
        console.log(response, "response");
        if (response.ok) {
            const updatedDriver = await response.json();
            // Assuming the update was successful, update the local state with the modified driver
            setDrivers((prevDrivers) =>
                prevDrivers.map((driver) => (driver.id === id ? updatedDriver : driver))
            );
            // Close the modal
            setOpenModal(false);
        } else {
            const errorText = await response.text();
            console.error('Error updating driver:', errorText);
        }
    } catch (error) {
        console.error('Error updating driver:', error);
    }
};


  

  const handleEditCustomer = (driver) => {
    // Set the driver being edited in the state and open the modal
    setEditingDriver(driver);
    setOpenModal(true);
  };

  const handleCloseModal = () => {
    // Clear the editingDriver state when closing the modal
    setEditingDriver(null);
    setOpenModal(false);
  };
  

  const handleDeleteDriver = async (id) => {
    try {
      let envString = 'REACT_APP_MICROSERVICE_' + driversApp.toUpperCase();
      await fetch(`${process.env[envString]}/api/deleteDriverById?id=${id}`, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${auth.user.access_token}`,
          'Content-Type': 'application/json',
        },
      });
      setDrivers(prevDrivers => prevDrivers.filter(driver => driver.id !== id));
    } catch (error) {
      console.error('Error deleting driver:', error);
    }
  };

  return (
    <div className="container ping">
      <button className="ping-button" onClick={() => setOpenModal(true)} style={{ alignItems: 'revert-layer' }}>
        Add Driver
      </button>
      <Modal isOpen={openModal} onClose={handleCloseModal} 
      onSubmit={(data) => {
        // If editingDriver is set, call handleUpdateDriver; otherwise, call handleSubmit
        editingDriver ? handleUpdateDriver(editingDriver.id, data) : handleSubmit(data);
      }}
      initialValues={editingDriver} // Pass the initialValues to the Modal for pre-filling
    />

      <table style={{ borderCollapse: 'collapse', width: '100%' }}>
        <thead>
          <tr style={{ color: 'black', backgroundColor: '#f2f2f2' }}>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Sno</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>First Name</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Last Name</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Phone Number</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>License Number</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}></th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}></th>
          </tr>
        </thead>
        <tbody>
          {drivers.map((driver, index) => (
            <tr
              key={index}
              style={{
                color: 'black',
                backgroundColor: index % 2 === 0 ? '#f9f9f9' : 'white',
              }}
            >
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{index + 1}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{driver.firstname}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{driver.lastname}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{driver.phone}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{driver.license}</td>
              <td
                style={{
                  padding: '10px',
                  border: '1px solid #dddddd',
                  textAlign: 'center',
                }}
              >
                <MdEdit style={{ fontSize: 20, color: 'red', cursor: 'pointer' }} onClick={() => handleEditCustomer(driver)} />
              </td>
              <td
                style={{
                  padding: '10px',
                  border: '1px solid #dddddd',
                  textAlign: 'center',
                }}
              >
                <MdDeleteOutline style={{ fontSize: 20, color: 'red', cursor: 'pointer' }} onClick={() => handleDeleteDriver(driver.id)} />
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default DriversList;
