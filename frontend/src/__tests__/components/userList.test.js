import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import UserList from '../../components/userList';

const mockUsers = [
    { id: 1, name: 'John', surname: 'Doe', age: 25 },
    { id: 2, name: 'Jane', surname: 'Smith', age: 30 },
    // Add more mock user data as needed
];

const mockHandleDelete = jest.fn();
const mockNavigateToUserPage = jest.fn();

describe('UserList', () => {
    it('renders a list of users', () => {
        render(
            <UserList
                users={mockUsers}
                handleDelete={mockHandleDelete}
                navigateToUserPage={mockNavigateToUserPage}
            />
        );

        // Check if each user's details are rendered
        mockUsers.forEach((user) => {
            expect(screen.getByText(`${user.name} ${user.surname}`)).toBeInTheDocument();
            expect(screen.getByText(`Age: ${user.age}`)).toBeInTheDocument();
        });
    });

    it('calls navigateToUserPage when "View Details" button is clicked for each user', () => {
        render(
            <UserList
                users={mockUsers}
                handleDelete={mockHandleDelete}
                navigateToUserPage={mockNavigateToUserPage}
            />
        );

        // Click the "View Details" button for each user
        mockUsers.forEach((user) => {
            const viewDetailsButtons = screen.getAllByText('View Details', { selector: 'button' });
            const userViewDetailsButton = viewDetailsButtons.find((button) => button.closest('div').textContent.includes(`${user.name} ${user.surname}`));

            fireEvent.click(userViewDetailsButton);
            expect(mockNavigateToUserPage).toHaveBeenCalledWith(user.id);
        });
    });
});
