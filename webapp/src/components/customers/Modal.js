import React, { useEffect, useState } from 'react';

const CustomerModal = ({ isOpen, onClose, onSubmit, initialValues }) => {
  const initialData = {
    firstName: '',
    lastName: '',
    email: '',
    phone: '',
    address: '',
  };
  const [formData, setFormData] = useState(initialData);

  useEffect(() => {
    // Update form data when initialValues change
    setFormData(initialValues || {});
  }, [initialValues]);

  const handleInputChange = (fieldName, value) => {
    setFormData(prevData => ({
      ...prevData,
      [fieldName]: value,
    }));
  };

  const handleSubmit = async () => {
    onSubmit(formData);
    setFormData(initialData);
    onClose();
  };

  if (!isOpen) return null;

  return (
    <div
      style={{
        position: 'fixed',
        top: 0,
        left: 0,
        width: '100%',
        height: '100%',
        background: 'rgba(0, 0, 0, 0.5)',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
      }}
    >
      <div
        style={{
          position: 'relative',
          background: 'white',
          width: 500,
          height: 'auto',
          margin: 'auto',
          padding: '4%',
          border: '2px solid #000',
          borderRadius: '10px',
          boxShadow: '2px solid black',
        }}
      >
        <button
          onClick={onClose}
          style={{
            position: 'absolute',
            top: 5,
            right: 4,
            background: 'none',
            border: 'none',
            cursor: 'pointer',
            padding: 0,
            fontSize: '20px',
          }}
        >
          <span
            style={{
              display: 'inline-block',
              width: '30px',
              height: '30px',
              borderRadius: '50%',
              background: '#eee',
              textAlign: 'center',
              lineHeight: '30px',
            }}
          >
            &times;
          </span>
        </button>

        <div
          style={{
            display: 'flex',
            flexDirection: 'column',
            marginBottom: '10px',
          }}
        >
          <div style={{ marginBottom: '10px', display: 'flex' }}>
            <label
              style={{
                color: 'black',
                paddingRight: '20px',
                width: '120px',
                flex: '0 0 120px',
                boxSizing: 'border-box',
                marginTop: '1px',
              }}
            >
              First Name:
            </label>
            <input
              type="text"
              value={formData.firstName}
              onChange={e => handleInputChange('firstName', e.target.value)}
              style={{
                paddingLeft: '5px',
                flex: 1,
                height: '30px',
                borderRadius: '5px',
                border: '1px solid #ccc',
              }}
            />
          </div>
          <div style={{ marginBottom: '10px', display: 'flex' }}>
          <label
              style={{
                color: 'black',
                paddingRight: '20px',
                width: '120px',
                flex: '0 0 120px',
                boxSizing: 'border-box',
                marginTop: '1px',
              }}
            >
              Last Name:
            </label>
            <input
              type="text"
              value={formData.lastName}
              onChange={e => handleInputChange('lastName', e.target.value)}
              style={{
                paddingLeft: '5px',
                flex: 1,
                height: '30px',
                borderRadius: '5px',
                border: '1px solid #ccc',
              }}
            />
          </div>
          <div style={{ marginBottom: '10px', display: 'flex' }}>
          <label
              style={{
                color: 'black',
                paddingRight: '20px',
                width: '120px',
                flex: '0 0 120px',
                boxSizing: 'border-box',
                marginTop: '1px',
              }}
            >
              Email:
            </label>
            <input
              type="text"
              value={formData.email}
              onChange={e => handleInputChange('email', e.target.value)}
              style={{
                paddingLeft: '5px',
                flex: 1,
                height: '30px',
                borderRadius: '5px',
                border: '1px solid #ccc',
              }}
            />
          </div>
          <div style={{ marginBottom: '10px', display: 'flex' }}>
          <label
              style={{
                color: 'black',
                paddingRight: '20px',
                width: '120px',
                flex: '0 0 120px',
                boxSizing: 'border-box',
                marginTop: '1px',
              }}
            >
              Phone Number:
            </label>
            <input
              type="text"
              value={formData.phone}
              onChange={e => handleInputChange('phone', e.target.value)}
              style={{
                paddingLeft: '5px',
                flex: 1,
                height: '30px',
                borderRadius: '5px',
                border: '1px solid #ccc',
              }}
            />
          </div>
          <div style={{ marginBottom: '10px', display: 'flex' }}>
          <label
              style={{
                color: 'black',
                paddingRight: '20px',
                width: '120px',
                flex: '0 0 120px',
                boxSizing: 'border-box',
                marginTop: '1px',
              }}
            >
              Address:
            </label>
            <input
              type="text"
              value={formData.address}
              onChange={e => handleInputChange('address', e.target.value)}
              style={{
                paddingLeft: '5px',
                flex: 1,
                height: '30px',
                borderRadius: '5px',
                border: '1px solid #ccc',
              }}
            />
          </div>
        </div>

        <div style={{ textAlign: 'center' }}>
          <button
            onClick={handleSubmit}
            style={{
              backgroundColor: '#4CAF50',
              color: 'white',
              padding: '10px',
              border: 'none',
              borderRadius: '5px',
              cursor: 'pointer',
            }}
          >
            Submit
          </button>
        </div>
      </div>
    </div>
  );
};

export default CustomerModal;
