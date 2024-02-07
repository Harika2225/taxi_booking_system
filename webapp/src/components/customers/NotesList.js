import React, { useState, useEffect } from 'react';
import Modal from './Modal';
import { MdDeleteOutline } from 'react-icons/md';
import { useAuth } from 'react-oidc-context';

const NotesList = ({ notesApp }) => {
  const [notes, setNotes] = useState([]);
  const [showAddNotePopup, setShowAddNotePopup] = useState(false);
  const [newNote, setNewNote] = useState({ subject: '', summary: '' });
  const [openModal, setOpenModal] = useState(false);
  const auth = useAuth();

  useEffect(() => {
    // Fetch notes from the backend API
    const fetchNotes = async () => {
      try {
        let envString = 'REACT_APP_MICROSERVICE_' + notesApp.toUpperCase();
        const response = await fetch(process.env[envString] + `/api/notes`, {
          method: 'GET',
          headers: {
            Authorization: `Bearer ${auth.user.access_token}`,
            'Content-Type': 'application/json',
          },
        }); // Replace with your actual API endpoint
        const data = await response.json();
        if (data != null) setNotes(data); // Assuming the API response is an array of notes
      } catch (error) {
        console.error('Error fetching notes:', error);
      }
    };
    fetchNotes();
  }, []);

  const handleSubmit = async Data => {
    try {
      // Make API call to post the collected data
      let envString = 'REACT_APP_MICROSERVICE_' + notesApp.toUpperCase();
      const response = await fetch(process.env[envString] + `/api/notes`, {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${auth.user.access_token}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(Data),
      });
      // Check if the request was successful (status code 2xx)
      if (response.ok) {
        const responseData = await response.json(); // Assuming the response is in JSON format
        // Use a callback function with setNotes to ensure you have the latest state
        setNotes(prevNotes => [
          ...prevNotes,
          {
            id: responseData.id,
            subject: responseData.subject,
            description: responseData.description,
          },
        ]);
      } else {
        // Handle error, maybe show an error message
        console.error('Error submitting data:', response.statusText);
      }
    } catch (error) {
      console.error('Error:', error.message);
    }
  };

  const handleInputChange = e => {
    const { name, value } = e.target;
    setNewNote({ ...newNote, [name]: value });
  };

  const handlePopupClose = () => {
    setShowAddNotePopup(false);
    setNewNote({ subject: '', summary: '' });
  };

  const handleAddNote = Data => {
    // Add logic to send newNote data to the backend via API
    console.log('New Note:', Data);
  };

  const handleDeleteNote = async id => {
    // Add logic to delete note using the backend API
    try {
      let envString = 'REACT_APP_MICROSERVICE_' + notesApp.toUpperCase();
      await fetch(process.env[envString] + `/api/notes?id=${id}`, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${auth.user.access_token}`,
          'Content-Type': 'application/json',
        },
      });
      // Assuming the delete was successful, update the local state
      setNotes(prevNotes => prevNotes.filter(note => note.id !== id));
    } catch (error) {
      console.error('Error deleting note:', error);
    }
  };

  return (
    <div className="container ping">
      <button className="ping-button" onClick={() => setOpenModal(true)} style={{ alignItems: 'revert-layer' }}>
        Add note
      </button>
      <Modal isOpen={openModal} onClose={() => setOpenModal(false)} onSubmit={handleSubmit} />
      {showAddNotePopup && (
        <div className="popup">
          <div className="popup-content">
            <span className="close" onClick={handlePopupClose}>
              &times;
            </span>
            <div style={{ padding: '20px', textAlign: 'center', color: 'black' }}>
              <label style={{ display: 'block', marginBottom: '10px' }}>
                Subject:
                <input
                  type="text"
                  name="subject"
                  value={newNote.subject}
                  onChange={handleInputChange}
                  style={{
                    width: '100%',
                    padding: '8px',
                    boxSizing: 'border-box',
                  }}
                />
              </label>
              <label style={{ display: 'block', marginBottom: '10px' }}>
                Summary:
                <textarea
                  name="summary"
                  value={newNote.summary}
                  onChange={handleInputChange}
                  style={{
                    width: '100%',
                    padding: '8px',
                    boxSizing: 'border-box',
                    minHeight: '80px',
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
      )}

      <table style={{ borderCollapse: 'collapse', width: '100%' }}>
        <thead>
          <tr style={{ color: 'black', backgroundColor: '#f2f2f2' }}>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Sno</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Subject</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}>Summary</th>
            <th style={{ padding: '10px', border: '1px solid #dddddd' }}></th>
          </tr>
        </thead>
        <tbody>
          {notes.map((note, index) => (
            <tr
              key={index}
              style={{
                color: 'black',
                backgroundColor: index % 2 === 0 ? '#f9f9f9' : 'white',
              }}
            >
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{index + 1}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{note.subject}</td>
              <td style={{ padding: '10px', border: '1px solid #dddddd' }}>{note.description}</td>
              <td
                style={{
                  padding: '10px',
                  border: '1px solid #dddddd',
                  textAlign: 'center',
                }}
              >
                <MdDeleteOutline style={{ fontSize: 20, color: 'red', cursor: 'pointer' }} onClick={() => handleDeleteNote(note.id)} />{' '}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default NotesList;
