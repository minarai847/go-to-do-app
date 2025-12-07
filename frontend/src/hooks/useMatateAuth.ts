import { useNavigate } from 'react-router-dom';
import { useMutation } from '@tanstack/react-query';
import axios from 'axios';
import useStore from '../store';
import { Credentials } from '../types';
import { useError } from './useError';

export const useMutateAuth = () => {
    const navigate = useNavigate();
    const resetEditedTask = useStore((state) => state.resetEditedTask);
    const { switchErrorHandling } = useError();

    const loginMutation = useMutation({
        mutationFn: async (user: Credentials) => {
            const response = await axios.post(
                `${process.env.REACT_APP_API_URL}/login`,
                user
            );
            return response.data;
        },
        onSuccess: () => {
            navigate('/todo');
        },
        onError: (err: any) => {
            if (err.response?.data?.message) {
                switchErrorHandling(err.response.data.message);
            } else if (err.response?.data) {
                switchErrorHandling(err.response.data);
            } else {
                switchErrorHandling(err.message || 'An error occurred');
            }
        },
    });

    const registerMutation = useMutation({
        mutationFn: async (user: Credentials) => {
            const response = await axios.post(
                `${process.env.REACT_APP_API_URL}/signup`,
                user
            );
            return response.data;
        },
        onSuccess: () => {
            navigate('/todo');
        },
        onError: (err: any) => {
            if (err.response?.data?.message) {
                switchErrorHandling(err.response.data.message);
            } else if (err.response?.data) {
                switchErrorHandling(err.response.data);
            } else {
                switchErrorHandling(err.message || 'An error occurred');
            }
        },
    });

    const logoutMutation = useMutation({
        mutationFn: async () => {
            const response = await axios.post(
                `${process.env.REACT_APP_API_URL}/logout`
            );
            return response.data;
        },
        onSuccess: () => {
            resetEditedTask();
            navigate('/');
        },
        onError: (err: any) => {
            if (err.response?.data?.message) {
                switchErrorHandling(err.response.data.message);
            } else if (err.response?.data) {
                switchErrorHandling(err.response.data);
            } else {
                switchErrorHandling(err.message || 'An error occurred');
            }
        },
    });

    return {
        loginMutation,
        registerMutation,
        logoutMutation,
    };
};
