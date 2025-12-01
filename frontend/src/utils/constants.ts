// 定数定義

export const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

export const ROUTES = {
    HOME: '/',
    LOGIN: '/login',
    SIGNUP: '/signup',
    TASKS: '/tasks',
} as const;

