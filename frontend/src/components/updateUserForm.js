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
        <form data-testid="update-user-form" onSubmit={handleSubmit}>
            <label htmlFor="name">Name:</label>
            <input
                type="text"
                id="name"
                name="name"
                value={updatedUser.name}
                onChange={handleInputChange}
                data-testid="name"
            />
            <div className="error-message" data-testid="name-error">{errors.name}</div>

            <label htmlFor="surname">Surname:</label>
            <input
                type="text"
                id="surname"
                name="surname"
                value={updatedUser.surname}
                onChange={handleInputChange}
                data-testid="surname"
            />
            <div className="error-message" data-testid="surname-error">{errors.surname}</div>

            <label htmlFor="age">Age:</label>
            <input
                type="number"
                id="age"
                name="age"
                value={updatedUser.age}
                onChange={handleInputChange}
                data-testid="age"
            />
            <div className="error-message" data-testid="age-error">{errors.age}</div>

            <div className="button-group">
                <button type="submit" data-testid="submit-button">
                    Update User
                </button>
                <span className="button-space"></span>
                <button type="button" onClick={onCancel} data-testid="cancel-button">
                    Cancel
                </button>
            </div>
        </form>
    );
};

export default UpdateUserForm;
