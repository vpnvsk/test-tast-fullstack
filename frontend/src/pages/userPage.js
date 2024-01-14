import React, { useState, useEffect, useCallback } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { useParams } from 'react-router-dom';
import UserDetail from '../components/userDetail';
import '../styles/userPage.css';
import UpdateUserForm from '../components/updateUserForm';


const UserPage = () => {
  const { userId } = useParams();
  const [user, setUser] = useState(null);
  const [isFormVisible, setFormVisible] = useState(false);

  const navigate = useNavigate();
  const navigateToHomePage = () => {
        navigate(`/`);
      };

  const fetchUserDetails = useCallback(async () => {
    try {
      const response = await axios.get(`http://localhost:8000/api/user/${userId}`);
      setUser(response.data);
    } catch (error) {
      console.error('Error fetching user details:', error);
    }
  }, [userId]);

  useEffect(() => {
    fetchUserDetails();
  }, [fetchUserDetails]);

  const handleUpdate = () => {
    setFormVisible(!isFormVisible);
  };
  const handleFormSubmit = async (updatedUserData) => {
    try {
      const { name, surname, age } = updatedUserData;
  
      await axios.put(`http://localhost:8000/api/user/${userId}`, { name, surname, age });
  
      fetchUserDetails();
  
      setFormVisible(false);
    } catch (error) {
      console.error('Error updating user:', error);
    }
  };

  const handleDelete = async (id) => {
    try {
      await axios.delete(`http://localhost:8000/api/user/${id}`);
      navigateToHomePage();
    } catch (error) {
      console.error('Error deleting user:', error);
    }
  };


  return (
    <div className="user-page-container">
      {user ? (
        <div className="user-details">
            <UserDetail user={user} handleUpdate={handleUpdate} handleDelete={handleDelete} isFormVisible={isFormVisible} />
          {isFormVisible && (
            <div className="user-update-form">
              <UpdateUserForm user={user} onSubmit={handleFormSubmit} onCancel={handleUpdate} />
            </div>
          )}
        </div>
      ) : (
        <p>Loading...</p>
      )}
      <div className="user-actions">
      <button onClick={navigateToHomePage} className="button-secondary">
          Back to Home
        </button>
</div>
    </div>
  );
};


export default UserPage;
