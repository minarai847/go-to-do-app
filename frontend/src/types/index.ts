// ユーザー関連の型定義
export interface User {
    id: number;
    email: string;
}

export interface UserResponse {
    id: number;
    email: string;
}

// タスク関連の型定義
export interface Task {
    id: number;
    title: string;
    user_id: number;
    created_at: string;
    updated_at: string;
}

export interface TaskResponse {
    id: number;
    title: string;
    created_at: string;
    updated_at: string;
}

// APIレスポンスの型定義
export interface ApiError {
    error: string | Record<string, string>;
    message?: string;
}

