import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { Task } from '../types'
import { useError } from './useError'

export const useQueryTasks = () => {
    const { switchErrorHandling } = useError();
    const getTasks = async () => {
        try {
            const { data } = await axios.get<Task[]>(
                `${process.env.REACT_APP_API_URL}/tasks`,
                { withCredentials: true }
            );
            return data;
        } catch (err: any) {
            if (err.response?.data?.message) {
                switchErrorHandling(err.response.data.message);
            } else if (err.response?.data) {
                switchErrorHandling(err.response.data);
            } else {
                switchErrorHandling(err.message || 'An error occurred');
            }
            throw err;
        }
    };
    return useQuery<Task[], Error>({
        queryKey: ['tasks'],
        queryFn: getTasks,
        staleTime: Infinity,
    });
};
