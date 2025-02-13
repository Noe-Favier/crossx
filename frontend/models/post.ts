import { User } from "./user";
import { Comment } from "./comment";

export interface Post {
    id?: number;
    created_at?: string;
    updated_at?: string;
    title?: string;
    content?: string;
    media_url?: string | null;
    user?: User;
    user_id?: number;

    views: User[];
    likes: User[];
    comments: Comment[];
}