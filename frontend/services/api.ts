import axios, { AxiosInstance } from 'axios';
import { getMeMock, getPostsMock } from './mock';
import { Post } from '@/models/post';
import { User } from '@/models/user';
import { useAuth } from '@/context/AuthContext';
import { Stack, router, usePathname } from 'expo-router';
import { Comment } from '@/models/comment';

const api = axios.create({
    baseURL: 'http://noais.fr:8080/api/v1',
    timeout: 5000,
    headers: { 'Content-Type': 'application/json' },
    // Ajouter cette configuration pour gérer correctement FormData
    transformRequest: [
        function (data, headers) {
            if (data instanceof FormData) {
                headers['Content-Type'] = 'multipart/form-data';
                return data;
            }
            if (typeof data === 'object') {
                headers['Content-Type'] = 'application/json';
                return JSON.stringify(data);
            }
            return data;
        }
    ],
});

api.interceptors.request.use(async (config) => {
    //on request ...
    return config;
});

api.interceptors.response.use(
    response => {
        console.log('api response', response.data);
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

export const apiSignup = async (data: FormData): Promise<{ user: User, token: string }> => {
    // Convertir FormData en objet pour voir le contenu
    const formDataObj: any = {};
    data.forEach((value, key) => {
        formDataObj[key] = value;
    });
    console.log('FormData content:', formDataObj);

    return api.post('/public/signup', data, {
        headers: {
            'Content-Type': 'multipart/form-data',
        },
        // Désactiver la transformation par défaut pour FormData
        transformRequest: [(data) => data],
    });
}

export const apiPostNewPost = async (data: Post): Promise<Post> => {
    return api.post('/post', data);
}

export const apiGetPost = async (id: number): Promise<Post> => {
    console.log('apiGetPost', id);
    return api.get(`/post/${id}`);
}

export const apiLikePost = async (id: number): Promise<void> => {
    return api.post(`/post/${id}/like`);
}

export const apiUnlikePost = async (id: number): Promise<void> => {
    return api.post(`/post/${id}/unlike`);
}

export const apiPostComment = async (postId: number, content: string): Promise<void> => {
    const comment: Comment = { content, post_id: postId };
    return api.post(`/comment`, comment);
}

export default api;