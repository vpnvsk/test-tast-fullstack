import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import UserItem from '../../components/userItem';

const mockUser = {
    id: 1,
    name: 'John',
    surname: 'Doe',
    age: 25,
};

const mockHandleDelete = jest.fn();
const mockNavigateToUserPage = jest.fn();

describe('UserItem', () => {
    it('renders user details correctly', () => {
        render(
            <UserItem
                user={mockUser}
                handleDelete={mockHandleDelete}
                navigateToUserPage={mockNavigateToUserPage}
            />
        );

        expect(screen.getByText(`${mockUser.name} ${mockUser.surname}`)).toBeInTheDocument();
        expect(screen.getByText(`Age: ${mockUser.age}`)).toBeInTheDocument();
    });

    it('calls handleDelete when "Delete" button is clicked', () => {
        render(
            <UserItem
                user={mockUser}
                handleDelete={mockHandleDelete}
                navigateToUserPage={mockNavigateToUserPage}
            />
        );

        fireEvent.click(screen.getByText('Delete'));

        expect(mockHandleDelete).toHaveBeenCalledWith(mockUser.id);
    });

    it('calls navigateToUserPage when "View Details" button is clicked', () => {
        render(
            <UserItem
                user={mockUser}
                handleDelete={mockHandleDelete}
                navigateToUserPage={mockNavigateToUserPage}
            />
        );

        fireEvent.click(screen.getByText('View Details'));

        expect(mockNavigateToUserPage).toHaveBeenCalledWith(mockUser.id);
    });
});
