import React, { useState, useEffect } from 'react';
import Modal from './Modal'
import { MdDeleteOutline, MdEdit } from 'react-icons/md';
import { useAuth } from 'react-oidc-context';

const PaymentsList = ({ paymentsApp }) => {
  const [payments, setPayments] = useState([]);
  const [editingPayment, setEditingPayment] = useState(null);
  const [openModal, setOpenModal] = useState(false);
  const auth = useAuth();

  useEffect(() => {
    const getPayments = async () => {
      try {
        let envString = 'REACT_APP_MICROSERVICE_' + paymentsApp.toUpperCase();
        const response = await fetch(process.env[envString] + `/api/getPayment`, {
            method: 'GET',
            headers: {
                Authorization: `Bearer ${auth.user.access_token}`,
                'Content-Type': 'application/json',
            },
        });
        const data = await response.json();
        console.log(data, "pppp");
        if (data != null) setPayments(data);
        } catch (error) {
            console.error('Error fetching payments:', error);
        }
    };
    getPayments();
  }, []);

  const handleSubmit = async Data => {
    try {
      let envString = 'REACT_APP_MICROSERVICE_' + paymentsApp.toUpperCase();
      const response = await fetch(process.env[envString] + `/api/createPayment`, {
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
        setPayments((prevPayments) => [...prevPayments, responseData]);
      } else {
        const errorText = await response.text();
        console.error('Error submitting data:', errorText);
      }
    } catch (error) {
      console.error('Error:', error.message);
    }
  };

  const handleUpdatePayment = async (id, data) => {
    try {
        let envString = 'REACT_APP_MICROSERVICE_' + paymentsApp.toUpperCase();
        const response = await fetch(`${process.env[envString]}/api/updatePaymentById/${id}`, {
            method: 'PUT',
            headers: {
                Authorization: `Bearer ${auth.user.access_token}`,
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data), // Use the formData state
        });
        console.log(response, "response");
        if (response.ok) {
            const updatedPayment = await response.json();
            // Assuming the update was successful, update the local state with the modified payment
            setPayments((prevPayments) =>
                prevPayments.map((payment) => (payment.id === id ? updatedPayment : payment))
            );
            // Close the modal
            setOpenModal(false);
        } else {
            const errorText = await response.text();
            console.error('Error updating payment:', errorText);
        }
    } catch (error) {
        console.error('Error updating payment:', error);
    }
};


  

  const handleEditPayment = (payment) => {
    // Set the payment being edited in the state and open the modal
    setEditingPayment(payment);
    setOpenModal(true);
  };

  const handleCloseModal = () => {
    // Clear the editingPayment state when closing the modal
    setEditingPayment(null);
    setOpenModal(false);
  };
  

  const handleDeletePayment = async (id) => {
    console.log(id,"id")
    let envString = 'REACT_APP_MICROSERVICE_' + paymentsApp.toUpperCase();
console.log(envString,"env")
    try {
      let envString = 'REACT_APP_MICROSERVICE_' + paymentsApp.toUpperCase();
      await fetch(`${process.env[envString]}/api/deletePaymentById?id=${id}`, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${auth.user.access_token}`,
          'Content-Type': 'application/json',
        },
      });
      setPayments(prevPayments => prevPayments.filter(payment => payment.id !== id));
    } catch (error) {
      console.error('Error deleting payment:', error);
    }
  };

  return (
    <div className="container ping">
      <button className="ping-button" onClick={() => setOpenModal(true)} style={{ alignItems: 'revert-layer' }}>
        Add Payment
      </button>
      <Modal isOpen={openModal} onClose={handleCloseModal} 
      onSubmit={(data) => {
        // If editingPayment is set, call handleUpdatePayment; otherwise, call handleSubmit
        editingPayment ? handleUpdatePayment(editingPayment.id, data) : handleSubmit(data);
      }}
      initialValues={editingPayment} // Pass the initialValues to the Modal for pre-filling
    />

      <table style={{ borderCollapse: 'collapse', width: '100%' }}>
        <thead>
          <tr style={{ color: 'black', backgroundColor: '#f2f2f2' }}>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Sno</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Amount</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Payment Date</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Payment_status</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Payment_method_id</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}></th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}></th>
          </tr>
        </thead>
        <tbody>
          {payments.map((payment, index) => (
            <tr
              key={index}
              style={{
                color: 'black',
                backgroundColor: index % 2 === 0 ? '#f9f9f9' : 'white',
              }}
            >
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{index + 1}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{payment.amount}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{payment.payment_date}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{payment.payment_status}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{payment.payment_method_id}</td>
              <td
                style={{
                  padding: '10px',
                  border: '1px solid #dddddd',
                  textAlign: 'center',
                }}
              >
                <MdEdit style={{ fontSize: 20, color: 'red', cursor: 'pointer' }} onClick={() => handleEditPayment(payment)} />
              </td>
              <td
                style={{
                  padding: '10px',
                  border: '1px solid #dddddd',
                  textAlign: 'center',
                }}
              >
                <MdDeleteOutline style={{ fontSize: 20, color: 'red', cursor: 'pointer' }} onClick={() => handleDeletePayment(payment.id)} />
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default PaymentsList;
