import axios from 'axios';
import { getMeMock, getPostsMock } from './mock';
import { Post } from '@/models/post';
import { User } from '@/models/user';

const api = axios.create({
    baseURL: 'https://localhost:8080/api',
    timeout: 5000,
    headers: { 'Content-Type': 'application/json' },
});

export const apiGetPosts = async (): Promise<Post[]> => {
    //return await api.get('/posts');
    return getPostsMock();
};

export const apiLogin = async (data: { email: string, password: string }): Promise<{ user: User, token: string }> => {
    //return await api.post('/login', data);
    return getMeMock();
}

export const apiGetMe = async (): Promise<User> => {
    //return await api.get('/me');
    return (await getMeMock()).user;
}

export default api;