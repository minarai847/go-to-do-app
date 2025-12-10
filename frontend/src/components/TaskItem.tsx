import { FC, memo } from 'react';
import { PencilIcon, TrashIcon } from '@heroicons/react/24/outline';
import useStore from '../store';
import { Task } from '../types';

type TaskItemProps = Omit<Task, 'created_at' | 'updated_at' | 'user_id'> & {
    onDelete: (id: number) => void;
};

const TaskItemMemo: FC<TaskItemProps> = ({ id, title, onDelete }) => {
    const { updateEditedTask } = useStore();

    // 無効なデータの場合は何も表示しない
    if (!id || !title || title.trim() === '') {
        return null;
    }

    return (
        <li className="my-3 flex items-center justify-between w-full">
            <span className="font-bold flex-1">{title}</span>
            <div className="flex items-center gap-2 flex-shrink-0">
                <PencilIcon
                    className="h-5 w-5 text-blue-500 cursor-pointer hover:text-blue-700"
                    onClick={() => updateEditedTask({ id, title })}
                />
                <TrashIcon
                    className="h-5 w-5 text-blue-500 cursor-pointer hover:text-blue-700"
                    onClick={() => onDelete(id)}
                />
            </div>
        </li>
    );
};

export const TaskItem = memo(TaskItemMemo, (prevProps, nextProps) => {
    return prevProps.id === nextProps.id &&
        prevProps.title === nextProps.title &&
        prevProps.onDelete === nextProps.onDelete;
});
