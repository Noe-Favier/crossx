import React, { createContext, useContext, useEffect, useState } from 'react';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { User } from '@/models/user';
import api, { apiGetMe, apiLogin } from '@/services/api';

interface AuthContextType {
    user: { user: User | null, token: string } | null;
    login: (email: string, password: string) => Promise<void>;
    logout: () => Promise<void>;
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

    const login = async (email: string, password: string) => {
        const res = await apiLogin({ email, password });
        const token = res.token;
        await AsyncStorage.setItem('token', token);
        api.defaults.headers.common['Authorization'] = `Bearer ${token}`;
        setUser(res);
    };

    const logout = async () => {
        await AsyncStorage.removeItem('token');
        delete api.defaults.headers.common['Authorization'];
        setUser(null);
    };

    const getMe = async () => {
        return user;
    }

    return <AuthContext.Provider value={{ user, login, logout }}>{children}</AuthContext.Provider>;
};