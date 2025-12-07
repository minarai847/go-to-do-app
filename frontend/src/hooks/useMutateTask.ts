import axios from 'axios'
import { useQueryClient, useMutation } from '@tanstack/react-query'
import { Task } from '../types'
import useStore from '../store'
import { useError } from './useError'

export const useMutateTask = () => {
    const queryClient = useQueryClient()
    const { switchErrorHandling } = useError()
    const resetEditedTask = useStore((state) => state.resetEditedTask)
}
