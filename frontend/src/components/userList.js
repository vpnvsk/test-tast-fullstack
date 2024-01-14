import React from 'react';
import UserItem from './userItem';
import '../styles/userList.css'

const UserList = ({ users, handleDelete, navigateToUserPage }) => {
    return (
        <div className="user-list">
          {users.map((user) => (
            <UserItem key={user.id} user={user} handleDelete={handleDelete} navigateToUserPage={navigateToUserPage} />
          ))}
        </div>
      );
    };

export default UserList;


