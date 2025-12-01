// API呼び出し用のサービス

const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

// 共通のfetch関数
async function fetchAPI<T>(
    endpoint: string,
    options: RequestInit = {}
): Promise<T> {
    const response = await fetch(`${API_BASE_URL}${endpoint}`, {
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options.headers,
        },
        credentials: 'include', // Cookieを送信
    });

    if (!response.ok) {
        const error = await response.json();
        throw new Error(error.error || error.message || 'An error occurred');
    }

    return response.json();
}

// ユーザー関連のAPI
export const userAPI = {
    signup: async (email: string, password: string) => {
        return fetchAPI('/signup', {
            method: 'POST',
            body: JSON.stringify({ email, password }),
        });
    },

    login: async (email: string, password: string) => {
        return fetchAPI('/login', {
            method: 'POST',
            body: JSON.stringify({ email, password }),
        });
    },

    logout: async () => {
        return fetchAPI('/logout', {
            method: 'POST',
        });
    },
};

// タスク関連のAPI
export const taskAPI = {
    getAll: async () => {
        return fetchAPI('/tasks');
    },

    getById: async (taskId: number) => {
        return fetchAPI(`/tasks/${taskId}`);
    },

    create: async (title: string) => {
        return fetchAPI('/tasks', {
            method: 'POST',
            body: JSON.stringify({ title }),
        });
    },

    update: async (taskId: number, title: string) => {
        return fetchAPI(`/tasks/${taskId}`, {
            method: 'PUT',
            body: JSON.stringify({ title }),
        });
    },

    delete: async (taskId: number) => {
        return fetchAPI(`/tasks/${taskId}`, {
            method: 'DELETE',
        });
    },
};

