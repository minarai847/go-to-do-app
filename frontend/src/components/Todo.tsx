import {
    ArrowRightOnRectangleIcon,
} from '@heroicons/react/24/outline'
import { useMutateAuth } from '../hooks/useMatateAuth'

export const Todo = () => {
    const { logoutMutation } = useMutateAuth()
    const logout = async () => {
        await logoutMutation.mutateAsync()
    }
    return (
        <div>
            <ArrowRightOnRectangleIcon
                onClick={logout}
                className="w-6 h-6 my-6 text-blue-500 cursor-pointer"
            />
        </div>
    )
}
