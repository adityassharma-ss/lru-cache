import React, { useState } from 'react';
import axios from 'axios';
import './App.css';

function App() {
  const [key, setKey] = useState('');
  const [value, setValue] = useState('');
  const [duration, setDuration] = useState('');
  const [dialogMessage, setDialogMessage] = useState('');
  const [error, setError] = useState('');

  const handleSet = async () => {
    if (!key.trim() || !value.trim() || !duration.trim()) {
      setError('Key, value, and duration must not be empty.');
      return;
    }

    try {
      const response = await axios.post('http://localhost:8080/set', {
        key: key.trim(),
        value: value.trim(),
        duration: parseInt(duration),
      });
      setDialogMessage(`Set operation successful for key: ${key}`);
      console.log(response.data);
    } catch (error) {
      setError(`Error setting data: ${error.message}`);
      console.error(error);
    }
  };

  const handleGet = async () => {
    if (!key.trim()) {
      setError('Key must not be empty.');
      return;
    }

    try {
      const response = await axios.get(`http://localhost:8080/get?key=${key.trim()}`);
      if (response.data) {
        setDialogMessage(`Data retrieved successfully for key: ${key}`);
        console.log(response.data);
      } else {
        setDialogMessage(`Key not found or expired: ${key}`);
      }
    } catch (error) {
      setError(`Error getting data: ${error.message}`);
      console.error(error);
    }
  };

  const handleDelete = async () => {
    if (!key.trim()) {
      setError('Key must not be empty.');
      return;
    }

    try {
      const response = await axios.delete(`http://localhost:8080/delete?key=${key.trim()}`);
      setDialogMessage(`Delete operation successful for key: ${key}`);
      console.log(response.data);
    } catch (error) {
      setError(`Error deleting data: ${error.message}`);
      console.error(error);
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>Cache Management Application</h1>
        <div className="operation-section">
          <label htmlFor="key">Key:</label>
          <input type="text" id="key" value={key} onChange={(e) => setKey(e.target.value)} />
          <label htmlFor="value">Value:</label>
          <input type="text" id="value" value={value} onChange={(e) => setValue(e.target.value)} />
          <label htmlFor="duration">Duration:</label>
          <input type="text" id="duration" value={duration} onChange={(e) => setDuration(e.target.value)} />
          <button onClick={handleSet}>Set</button>
          <button onClick={handleGet}>Get</button>
          <button onClick={handleDelete}>Delete</button>
        </div>
        {error && (
          <div className="error-box">
            <p>{error}</p>
            <button onClick={() => setError('')}>Close</button>
          </div>
        )}
        {dialogMessage && (
          <div className="dialog-box">
            <p>{dialogMessage}</p>
            <button onClick={() => setDialogMessage('')}>Close</button>
          </div>
        )}
      </header>
    </div>
  );
}

export default App;
