// 認証関連のカスタムフック

import { useState, useEffect } from 'react';
import { userAPI } from '../services/api';
import { UserResponse } from '../types';

export const useAuth = () => {
    const [user, setUser] = useState<UserResponse | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        // 認証状態の確認（必要に応じて実装）
        setLoading(false);
    }, []);

    const login = async (email: string, password: string) => {
        try {
            await userAPI.login(email, password);
            // ログイン成功後の処理
            return true;
        } catch (error) {
            console.error('Login failed:', error);
            return false;
        }
    };

    const logout = async () => {
        try {
            await userAPI.logout();
            setUser(null);
        } catch (error) {
            console.error('Logout failed:', error);
        }
    };

    return {
        user,
        loading,
        login,
        logout,
    };
};

