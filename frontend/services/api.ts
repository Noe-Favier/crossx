import axios from 'axios';
import { getMeMock, getPostsMock } from './mock';
import { Post } from '@/models/post';
import { User } from '@/models/user';
import { useAuth } from '@/context/AuthContext';
import { Stack, router, usePathname } from 'expo-router';

const api = axios.create({
    baseURL: 'http://noais.fr:8080/api/v1',
    timeout: 5000,
    headers: { 'Content-Type': 'application/json' },
});

api.interceptors.request.use(async (config) => {
    //on request ...
    return config;
});

api.interceptors.response.use(
    response => {
        console.log(`${response.config.method?.toUpperCase()} ${response.config.url} => ${response.status} ${JSON.stringify(response.data)}`);
        return response.data;
    },
    error => {
        if (error.response.status === 401) {
            router.replace('/login');
            return Promise.reject(error);
        }

        console.error('api error', error);
        return Promise.reject(error);
    }
);

export const apiGetPosts = async (): Promise<Post[]> => {
    return api.get('/post');
};

export const apiLogin = async (data: { username: string, password: string }): Promise<{ user: User, token: string }> => {
    return api.post('/public/login', data);
}

export const apiGetMe = async (): Promise<User> => {
    return api.get('/public/me');
}

export const apiSignup = async (data: { email: string, password: string, username: string, image: string }): Promise<{ user: User, token: string }> => {
    return api.post('/public/signup', data);
}

export default api;