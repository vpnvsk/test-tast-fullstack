import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import UserList from '../components/userList';
import '../styles/homePage.css';

const HomePage = () => {
  const [users, setUsers] = useState([]);
  const [formData, setFormData] = useState({
    name: '',
    surname: '',
    age: 0,
  });
  const [errors, setErrors] = useState({});
  const navigate = useNavigate();

  useEffect(() => {
    fetchUsers()
  }, []);

  const fetchUsers = async () => {
    try {
      const response = await axios.get('http://localhost:8000/api/user/');
      setUsers(response.data);
    } catch (error) {
      console.error('Error fetching users:', error);
    }
  };

  const handleInputChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
    setErrors({ ...errors, [e.target.name]: '' });
  };

  const validateForm = () => {
    const newErrors = {};

    if (!formData.name.trim()) {
      newErrors.name = 'Name is required';
    }

    if (!formData.surname.trim()) {
      newErrors.surname = 'Surname is required';
    }

    if (formData.age <= 0) {
      newErrors.age = 'Age must be a positive number';
    }

    setErrors(newErrors);

    return Object.keys(newErrors).length === 0;
  };

  const handleCreateUser = async (e) => {
    e.preventDefault();

    if (validateForm()) {
      const numericAge = parseInt(formData.age, 10);

      try {
        await axios.post('http://localhost:8000/api/user/', { ...formData, age: numericAge });
        setFormData({
          name: '',
          surname: '',
          age: 0,
        });
        fetchUsers();
      } catch (error) {
        console.error('Error creating user:', error);
      }
    }
  };

  const handleDelete = async (id) => {
    try {
      await axios.delete(`http://localhost:8000/api/user/${id}`);
      fetchUsers();
    } catch (error) {
      console.error('Error deleting user:', error);
    }
  };

  const navigateToUserPage = (userId) => {
    navigate(`/user/${userId}`);
  };

  return (
    <div className="homepage-container">
      <h1>CRUD Application</h1>
      <form className="user-form" onSubmit={handleCreateUser}>
        <label>Name:</label>
        <input type="text" name="name" value={formData.name} onChange={handleInputChange} />
        <div className="error-message">{errors.name}</div>

        <label>Surname:</label>
        <input type="text" name="surname" value={formData.surname} onChange={handleInputChange} />
        <div className="error-message">{errors.surname}</div>

        <label>Age:</label>
        <input type="number" name="age" value={formData.age} onChange={handleInputChange} />
        <div className="error-message">{errors.age}</div>

        <button type="submit">Create User</button>
      </form>
      <div className="user-list-container">
        <h2>User List</h2>
        <UserList users={users} handleDelete={handleDelete} navigateToUserPage={navigateToUserPage} />
      </div>
    </div>
  );
};

export default HomePage;

