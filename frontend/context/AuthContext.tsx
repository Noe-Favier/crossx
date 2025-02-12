import React, { createContext, useContext, useEffect, useState } from 'react';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { User } from '@/models/user';
import api, { apiGetMe, apiLogin } from '@/services/api';

interface AuthContextType {
    userState: { user: User | null, token: string } | null;
    login: (email: string, password: string) => Promise<void>;
    logout: () => Promise<void>;
    getMe: () => Promise<User>;
}

const AuthContext = createContext<AuthContextType | null>(null);
export const useAuth = () => useContext(AuthContext)!;
export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
    const [user, setUser] = useState<{ user: User, token: string } | null>(null);

    useEffect(() => {
        const loadUser = async () => {
            const token = await AsyncStorage.getItem('token');
            if (token) {
                api.defaults.headers.common['Authorization'] = `Bearer ${token}`;
                const user = await apiGetMe();
                setUser({ user: user, token: token });
            }
        };
        loadUser();
    }, []);

    const login = async (username: string, password: string) => {
        try {
            const res = await apiLogin({ username, password });
            const token = res.token;
            await AsyncStorage.setItem('token', token);
            api.defaults.headers.common['Authorization'] = `Bearer ${token}`;
            setUser({ user: res.user, token });
        } catch (error) {
            console.error('Login failed', error);
        }
    };


    const logout = async () => {
        await AsyncStorage.removeItem('token');
        delete api.defaults.headers.common['Authorization'];
        setUser(null);
    };

    const getMe = async () => {
        return apiGetMe();
    }

    return <AuthContext.Provider value={{ userState: user, login, logout, getMe }}>{children}</AuthContext.Provider>;
};