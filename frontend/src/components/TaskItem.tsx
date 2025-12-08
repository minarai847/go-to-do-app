import { FC, memo } from 'react';
import { PencilIcon, TrashIcon } from '@heroicons/react/24/outline';
import useStore from '../store';
import { Task } from '../types';
import { useMutateTask } from '../hooks/useMutateTask';

type TaskItemProps = Omit<Task, 'created_at' | 'updated_at' | 'user_id'>;

const TaskItemMemo: FC<TaskItemProps> = ({ id, title }) => {
    const { updateEditedTask } = useStore();
    const { deleteTaskMutation } = useMutateTask();

    return (
        <li className="my-3">
            <span className="font-bold">{title}</span>
            <div className="float-right ml-20 flex">
                <PencilIcon
                    className="h-5 w-5 mx-1 text-blue-500 cursor-pointer"
                    onClick={() => updateEditedTask({ id, title })}
                />
                <TrashIcon
                    className="h-5 w-5 text-blue-500 cursor-pointer"
                    onClick={() => deleteTaskMutation.mutate(id)}
                />
            </div>
        </li>
    );
};

export const TaskItem = memo(TaskItemMemo);
