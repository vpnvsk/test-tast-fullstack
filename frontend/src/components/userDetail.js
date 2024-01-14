import React from 'react';

const UserDetail = ({ user, handleUpdate, handleDelete, isFormVisible }) => (
    <div>
      <h2>User Details</h2>
      <p>Name: {user.name}</p>
      <p>Surname: {user.surname}</p>
      <p>Age: {user.age}</p>
      {!isFormVisible && (
        <button onClick={() => handleUpdate(user.id)}>Update User</button>
      )}
      <span className="button-space"></span>
      <button onClick={() => handleDelete(user.id)}>Delete User</button>
    </div>
  );
  
  export default UserDetail;
  
