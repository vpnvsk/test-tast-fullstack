import React from 'react';
import { render, fireEvent, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import UpdateUserForm from '../../components/updateUserForm';

const mockSubmit = jest.fn();
const mockCancel = jest.fn();

describe('UpdateUserForm', () => {
    const mockUser = {
        name: 'John',
        surname: 'Doe',
        age: 25,
    };

    it('submits the form with valid data', async () => {
        const { getByTestId } = render(
            <UpdateUserForm user={mockUser} onSubmit={mockSubmit} onCancel={mockCancel} />
        );

        fireEvent.change(getByTestId('name'), { target: { value: 'Jane' } });
        fireEvent.change(getByTestId('surname'), { target: { value: 'Smith' } });
        fireEvent.change(getByTestId('age'), { target: { value: '30' } });

        fireEvent.click(getByTestId('submit-button'));

        await waitFor(() => {
            expect(mockSubmit).toHaveBeenCalledWith({
                name: 'Jane',
                surname: 'Smith',
                age: 30,
            });
        });
    });

    it('calls onCancel function when cancel button is clicked', () => {
        const { getByTestId } = render(
            <UpdateUserForm user={mockUser} onSubmit={mockSubmit} onCancel={mockCancel} />
        );

        fireEvent.click(getByTestId('cancel-button'));

        expect(mockCancel).toHaveBeenCalled();
    });
});
