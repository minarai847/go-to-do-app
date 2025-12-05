import axios from 'axios';
import { useNavigate } from '@tanstack/react-query';
import { CsrfToken } from '../types';
import useStore from '../store';

export const useError = () => {

    const navigate = useNavigate();
    const resetEditedTask = useStore((state) => state.resetEditedTask);
    const getCsrfToken = async () => {
        const { data } = await axios.get<CsrfToken>(`${process.env.REACT_APP_API_URL}/csrf`);
        axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token;
    }
    const switchErrorHandling = (msg: string) => {
        switch (msg) {
            case 'invalid csrf token';
                getCsrfToken();
                alert('CSRF token is invalid, please try again');
                break;
            default:
                alert(msg)
        }
    }
}
