import { Post } from "./post";
import { User } from "./user";

export interface Comment {
    id: number;
    created_at: string;
    updated_at: string;
    content: string;
    post: Post;
    user: User;
    post_id: number;
    user_id: number;
}