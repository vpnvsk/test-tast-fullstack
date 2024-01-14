import React from 'react';

const UserItem = ({ user, handleDelete, navigateToUserPage }) => {
  return (
    <div className="user-item">
      <h3>{user.name} {user.surname}</h3>
      <p>Age: {user.age}</p>
      <button onClick={() => navigateToUserPage(user.id)}>View Details</button>
      <button onClick={() => handleDelete(user.id)}>Delete</button>
    </div>
  );
};

export default UserItem;
