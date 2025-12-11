import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { CsrfToken } from '../types';
import useStore from '../store';

export const useError = () => {

    const navigate = useNavigate();
    const resetEditedTask = useStore((state) => state.resetEditedTask);
    const getCsrfToken = async () => {
        try {
            const { data } = await axios.get<CsrfToken>(
                `${process.env.REACT_APP_API_URL}/csrf`,
                { withCredentials: true }
            );
            axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token;
            return data.csrf_token;
        } catch (error) {
            console.error('Failed to get CSRF token:', error);
            return null;
        }
    }
    const switchErrorHandling = (msg: string | any) => {
        // オブジェクトの場合は文字列に変換
        let errorMessage: string;
        if (typeof msg === 'string') {
            errorMessage = msg;
        } else if (msg && typeof msg === 'object') {
            // エラーオブジェクトからメッセージを抽出
            if (msg.error) {
                errorMessage = typeof msg.error === 'string' ? msg.error : JSON.stringify(msg.error);
            } else if (msg.message) {
                errorMessage = msg.message;
            } else {
                errorMessage = JSON.stringify(msg);
            }
        } else {
            errorMessage = String(msg);
        }

        switch (errorMessage) {
            case 'invalid csrf token':
                getCsrfToken();
                alert('CSRF token is invalid, please try again');
                break;
            case 'invalid or expired jwt':
                alert(`access token expired, please login`);
                resetEditedTask();
                navigate('/');
                break;
            case 'missin or malformed jwt':
                alert(`access token is not valid, please login`);
                navigate('/');
                break;
            case 'duplicated key not allowed':
                alert(`email already exist, please use another one`);
                break;
            case 'crypto/bcrypt: bcrypt match failed':
                alert(`password is incorrect, please try again`);
                break;
            case 'record not found':
                alert(`user not found, please signup`);
                navigate('/');
                break;
            default:
                alert(errorMessage);
        }
    }
    return { switchErrorHandling, getCsrfToken }
}
