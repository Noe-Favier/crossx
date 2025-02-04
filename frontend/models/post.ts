import { User } from "./user";

export interface Post {
    id: number;
    created_at: string;
    updated_at: string;
    content: string;
    media_url: string | null;
    user: User;
    user_id: number;
}