import React, { useState, useEffect } from 'react';
import Modal from './Modal';
import { MdDeleteOutline, MdEdit } from 'react-icons/md';
import { useAuth } from 'react-oidc-context';

const BookingsList = ({ bookingsApp }) => {
  const [bookings, setBookings] = useState([]);
  const [openModal, setOpenModal] = useState(false);
  const [editingBooking, setEditingBooking] = useState(null);
  const auth = useAuth();

  useEffect(() => {
    // Fetch bookings from the backend API
    const getBookings = async () => {
      try {
        let envString = 'REACT_APP_MICROSERVICE_' + bookingsApp.toUpperCase();
        const response = await fetch(process.env[envString] + `/api/getBooking`, {
          method: 'GET',
          headers: {
            Authorization: `Bearer ${auth.user.access_token}`,
            'Content-Type': 'application/json',
          },
        }); // Replace with your actual API endpoint
        const data = await response.json();
        console.log(data,"pppp")
        if (data != null) setBookings(data); // Assuming the API response is an array of bookings
      } catch (error) {
        console.error('Error fetching bookings:', error);
      }
    };
    getBookings();
  }, []);

  const handleSubmit = async Data => {
    try {
      // Make API call to post the collected data
      let envString = 'REACT_APP_MICROSERVICE_' + bookingsApp.toUpperCase();
      const response = await fetch(process.env[envString] +`/api/createBooking`, {
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
        setBookings((prevBookings) => [...prevBookings, responseData]);
      } else {
        const errorText = await response.text();
        console.error('Error submitting data:', errorText);
      }
    } catch (error) {
      console.error('Error:', error.message);
    }
  };

  const handleUpdateBooking = async (id, data) => {
    try {
      let envString = 'REACT_APP_MICROSERVICE_' + bookingsApp.toUpperCase();
      const response = await fetch(`${process.env[envString]}/api/updateBookingById/${id}`, {
        method: 'PUT',
        headers: {
          Authorization: `Bearer ${auth.user.access_token}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });
      console.log(response,"response")
      if (response.ok) {
        const updatedBooking = await response.json();
        // Assuming the update was successful, update the local state with the modified booking
        setBookings((prevBookings) =>
          prevBookings.map((booking) => (booking.id === id ? updatedBooking : booking))
        );
        // Close the modal
        setOpenModal(false);
      } else {
        const errorText = await response.text();
        console.error('Error updating booking:', errorText);
      }
    } catch (error) {
      console.error('Error updating booking:', error);
    }
  };
  

  const handleEditBooking = (booking) => {
    // Set the booking being edited in the state and open the modal
    setEditingBooking(booking);
    setOpenModal(true);
  };

  const handleCloseModal = () => {
    // Clear the editingBooking state when closing the modal
    setEditingBooking(null);
    setOpenModal(false);
  };
  
  const handleDeleteBooking = async id => {
    console.log(id,"psdapsdasp")
    // Add logic to delete note using the backend API
    try {
      let envString = 'REACT_APP_MICROSERVICE_' + bookingsApp.toUpperCase();
      await fetch(`${process.env[envString]}/api/deleteBookingById?id=${id}`, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${auth.user.access_token}`,
          'Content-Type': 'application/json',
        },
      });
      // Assuming the delete was successful, update the local state
      setBookings(prevBookings => prevBookings.filter(booking => booking.id !== id));
    } catch (error) {
      console.error('Error deleting note:', error);
    }
  };

  return (
    <div className="container ping">
      <button className="ping-button" onClick={() => setOpenModal(true)} style={{ alignItems: 'revert-layer' }}>
        Add Booking
      </button>
      <Modal isOpen={openModal} onClose={handleCloseModal} 
      // onSubmit={handleSubmit} 
       onSubmit={(data) => {
          // If editingBooking is set, call handleUpdateBooking; otherwise, call handleSubmit
          editingBooking ? handleUpdateBooking(editingBooking.id, data) : handleSubmit(data);
        }}
        initialValues={editingBooking} // Pass the initialValues to the Modal for pre-filling
      />
      <table style={{ borderCollapse: 'collapse', width: '100%' }}>
        <thead>
          <tr style={{ color: 'black', backgroundColor: '#f2f2f2' }}>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Sno</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Pickup Address</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Destination</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Date</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Time</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Status</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Amount</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}></th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}></th>
          </tr>
        </thead>
        <tbody>
          {bookings.map((booking, index) => (
            <tr
              key={index}
              style={{
                color: 'black',
                backgroundColor: index % 2 === 0 ? '#f9f9f9' : 'white',
              }}
            >
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{index + 1}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{booking.pickupaddress}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{booking.destination}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{booking.date}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{booking.time}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{booking.status}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{booking.amount}</td>
              <td
                style={{
                  padding: '10px',
                  border: '1px solid #dddddd',
                  textAlign: 'center',
                }}
              >
                <MdEdit style={{ fontSize: 20, color: 'red', cursor: 'pointer' }} onClick={() => handleEditBooking(booking)}
                 />{' '}
              </td>
              <td
                style={{
                  padding: '10px',
                  border: '1px solid #dddddd',
                  textAlign: 'center',
                }}
              >
                <MdDeleteOutline style={{ fontSize: 20, color: 'red', cursor: 'pointer' }} onClick={() => handleDeleteBooking(booking.id)} />{' '}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default BookingsList;
