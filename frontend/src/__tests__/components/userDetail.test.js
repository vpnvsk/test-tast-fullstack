import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import UserDetail from '../../components/userDetail';

const mockUser = {
    id: 1,
    name: 'John',
    surname: 'Doe',
    age: 25,
};

const mockHandleUpdate = jest.fn();
const mockHandleDelete = jest.fn();

describe('UserDetail', () => {
    it('renders user details correctly', () => {
        render(
            <UserDetail
                user={mockUser}
                handleUpdate={mockHandleUpdate}
                handleDelete={mockHandleDelete}
                isFormVisible={false}
            />
        );

        expect(screen.getByText(`Name: ${mockUser.name}`)).toBeInTheDocument();
        expect(screen.getByText(`Surname: ${mockUser.surname}`)).toBeInTheDocument();
        expect(screen.getByText(`Age: ${mockUser.age}`)).toBeInTheDocument();
    });

    it('calls handleUpdate when "Update User" button is clicked', () => {
        render(
            <UserDetail
                user={mockUser}
                handleUpdate={mockHandleUpdate}
                handleDelete={mockHandleDelete}
                isFormVisible={false}
            />
        );

        fireEvent.click(screen.getByText('Update User'));

        expect(mockHandleUpdate).toHaveBeenCalledWith(mockUser.id);
    });

    it('calls handleDelete when "Delete User" button is clicked', () => {
        render(
            <UserDetail
                user={mockUser}
                handleUpdate={mockHandleUpdate}
                handleDelete={mockHandleDelete}
                isFormVisible={false}
            />
        );

        fireEvent.click(screen.getByText('Delete User'));

        expect(mockHandleDelete).toHaveBeenCalledWith(mockUser.id);
    });

    it('does not render "Update User" button when isFormVisible is true', () => {
        render(
            <UserDetail
                user={mockUser}
                handleUpdate={mockHandleUpdate}
                handleDelete={mockHandleDelete}
                isFormVisible={true}
            />
        );

        expect(screen.queryByText('Update User')).toBeNull();
    });
});
