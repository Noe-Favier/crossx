export interface User {
    id: number;
    created_at: string;
    updated_at: string;
    bio: string;
    email: string;
    username: string;
    password_hash: string;
    profile_picture_url: string | null;
}