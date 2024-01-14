import React, { useState } from 'react';
import '../styles/updateUserForm.css'; 

const UpdateUserForm = ({ user, onSubmit, onCancel }) => {
  const [updatedUser, setUpdatedUser] = useState({
    name: user.name || '',
    surname: user.surname || '',
    age: user.age || 0,
  });
  const [errors, setErrors] = useState({});

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    const updatedValue = name === 'age' ? parseInt(value, 10) : value;
    setUpdatedUser({ ...updatedUser, [name]: updatedValue });
    setErrors({ ...errors, [name]: '' });
  };

  const validateForm = () => {
    const newErrors = {};

    if (!updatedUser.name.trim()) {
      newErrors.name = 'Name is required';
    }

    if (!updatedUser.surname.trim()) {
      newErrors.surname = 'Surname is required';
    }

    if (updatedUser.age <= 0) {
      newErrors.age = 'Age must be a positive number';
    }

    setErrors(newErrors);

    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    if (validateForm()) {
      onSubmit(updatedUser);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>Name:</label>
      <input
        type="text"
        name="name"
        value={updatedUser.name}
        onChange={handleInputChange}
      />
      <div className="error-message">{errors.name}</div>

      <label>Surname:</label>
      <input
        type="text"
        name="surname"
        value={updatedUser.surname}
        onChange={handleInputChange}
      />
      <div className="error-message">{errors.surname}</div>

      <label>Age:</label>
      <input
        type="number"
        name="age"
        value={updatedUser.age}
        onChange={handleInputChange}
      />
      <div className="error-message">{errors.age}</div>

      <div className="button-group">
        <button type="submit">Update User</button>
        <span className="button-space"></span>
        <button type="button" onClick={onCancel}>
          Cancel
        </button>
      </div>
    </form>
  );
};

export default UpdateUserForm;
