import { ArrowRightOnRectangleIcon, ShieldCheckIcon } from '@heroicons/react/24/outline';
import { useQueryClient } from '@tanstack/react-query';
import { FormEvent } from 'react';
import { useMutateAuth } from '../hooks/useMatateAuth';
import { useQueryTasks } from '../hooks/useQueryTasks';
import { TaskItem } from './TaskItem';
import useStore from '../store';
import { useMutateTask } from '../hooks/useMutateTask';

export const Todo = () => {
    const queryClient = useQueryClient();
    const { editedTask } = useStore();
    const updateTask = useStore((state) => state.updateEditedTask);
    const { data, isLoading } = useQueryTasks();
    const { createTaskMutation, updateTaskMutation } = useMutateTask();
    const { logoutMutation } = useMutateAuth();

    const submitTaskHandler = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (editedTask.id === 0) {
            createTaskMutation.mutate({ title: editedTask.title });
        } else {
            updateTaskMutation.mutate({ id: editedTask.id, title: editedTask.title });
        }
    };

    const logout = async () => {
        await logoutMutation.mutateAsync();
        queryClient.removeQueries({ queryKey: ['tasks'] });
    };

    return (
        <div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
            <div className="flex items-center my-3">
                <ShieldCheckIcon className="h-8 w-8 mr-2 text-blue-500" />
                <span className="text-center text-3xl font-extrabold">
                    Task Manager
                </span>
            </div>
            <ArrowRightOnRectangleIcon
                onClick={logout}
                className="w-6 h-6 my-6 text-blue-500 cursor-pointer"
            />
            <form onSubmit={submitTaskHandler}>
                <input
                    className="mb-3 px-3 py-2 border border-gray-300"
                    placeholder="title ?"
                    type="text"
                    onChange={(e) => updateTask({ ...editedTask, title: e.target.value })}
                    value={editedTask.title || ''}
                />
                <button
                    className="disabled:opacity-40 disabled:cursor-not-allowed px-3 py-2 text-sm text-white bg-blue-600 hover:bg-blue-700 rounded"
                    disabled={!editedTask.title}
                    type="submit"
                >
                    {editedTask.id === 0 ? 'Create' : 'Update'}
                </button>
            </form>
            {isLoading ? (
                <p>Loading...</p>
            ) : (
                <ul className="my-5">
                    {data?.map((task) => (
                        <TaskItem key={task.id} id={task.id} title={task.title} />
                    ))}
                </ul>
            )}
        </div>
    );
};
