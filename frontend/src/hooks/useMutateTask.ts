import axios from 'axios';
import { useQueryClient, useMutation } from '@tanstack/react-query';
import { Task, TaskResponse } from '../types';
import useStore from '../store';
import { useError } from './useError';

export const useMutateTask = () => {
    const queryClient = useQueryClient();
    const { switchErrorHandling } = useError();
    const resetEditedTask = useStore((state) => state.resetEditedTask);

    const createTaskMutation = useMutation({
        mutationFn: async (task: { title: string }) => {
            const response = await axios.post<TaskResponse>(
                `${process.env.REACT_APP_API_URL}/tasks`,
                task,
                { withCredentials: true }
            );
            return response.data;
        },
        onSuccess: () => {
            resetEditedTask();
            queryClient.invalidateQueries({ queryKey: ['tasks'] });
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

    const updateTaskMutation = useMutation({
        mutationFn: async ({ id, title }: { id: number; title: string }) => {
            const response = await axios.put<TaskResponse>(
                `${process.env.REACT_APP_API_URL}/tasks/${id}`,
                { title },
                { withCredentials: true }
            );
            return response.data;
        },
        onSuccess: () => {
            resetEditedTask();
            queryClient.invalidateQueries({ queryKey: ['tasks'] });
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

    const deleteTaskMutation = useMutation({
        mutationFn: async (id: number) => {
            await axios.delete(
                `${process.env.REACT_APP_API_URL}/tasks/${id}`,
                { withCredentials: true }
            );
            return id;
        },
        onSuccess: () => {
            resetEditedTask();
            queryClient.invalidateQueries({ queryKey: ['tasks'] });
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
        createTaskMutation,
        updateTaskMutation,
        deleteTaskMutation,
    };
};
