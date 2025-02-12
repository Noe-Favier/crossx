import { User } from "@/models/user";

export const getPostsMock = async () => {
    return [
        {
            id: 1,
            created_at: "2025-01-01T10:00:00Z",
            updated_at: "2025-01-01T10:30:00Z",
            content: "Ceci est un post d'exemple numéro 1.",
            media_url: "https://example.com/media1.jpg",
            user: {
                id: 1,
                bio: 'hey',
                email: 'mail@mail.mail',
                profile_picture_url: 'https://avatars.githubusercontent.com/u/91917411?v=4',
                username: "John Doe",
            },
            user_id: 1,
        },
        {
            id: 2,
            created_at: "2025-01-02T10:00:00Z",
            updated_at: "2025-01-02T10:30:00Z",
            content: "Ceci est un post d'exemple numéro 2.",
            media_url: "https://example.com/media2.jpg",
            user: {
                id: 2,
                username: "Jane Smith",
                bio: 'hey',
                email: 'mail@mail.mail',
                profile_picture_url: 'https://avatars.githubusercontent.com/u/91917411?v=4',
            },
            user_id: 2,
        },
        {
            id: 3,
            created_at: "2025-01-03T10:00:00Z",
            updated_at: "2025-01-03T10:30:00Z",
            content: "Ceci est un post d'exemple numéro 3.",
            media_url: null,
            user: {
                id: 3,
                username: "Alice Johnson",
                bio: 'hey',
                email: 'mail@mail.mail',
                profile_picture_url: 'https://avatars.githubusercontent.com/u/91917411?v=4',
            },
            user_id: 3,
        },
        {
            id: 4,
            created_at: "2025-01-04T10:00:00Z",
            updated_at: "2025-01-04T10:30:00Z",
            content: "Ceci est un post d'exemple numéro 4.",
            media_url: "https://example.com/media4.jpg",
            user: {
                id: 4,
                username: "Bob Lee",
                bio: 'hey',
                email: 'mail@mail.mail',
                profile_picture_url: 'https://avatars.githubusercontent.com/u/91917411?v=4',
            },
            user_id: 4,
        },
        {
            id: 5,
            created_at: "2025-01-05T10:00:00Z",
            updated_at: "2025-01-05T10:30:00Z",
            content: "Ceci est un post d'exemple numéro 5.",
            media_url: "https://example.com/media5.jpg",
            user: {
                id: 5,
                username: "Charlie Brown",
                bio: 'hey',
                email: 'mail@mail.mail',
                profile_picture_url: 'https://avatars.githubusercontent.com/u/91917411?v=4',
            },
            user_id: 5,
        },
        {
            id: 6,
            created_at: "2025-01-06T10:00:00Z",
            updated_at: "2025-01-06T10:30:00Z",
            content: "Ceci est un post d'exemple numéro 6.",
            media_url: null,
            user: {
                id: 6,
                username: "David Wilson",
                bio: 'hey',
                email: 'mail@mail.mail',
                profile_picture_url: 'https://avatars.githubusercontent.com/u/91917411?v=4',
            },
            user_id: 6,
        },
        {
            id: 7,
            created_at: "2025-01-07T10:00:00Z",
            updated_at: "2025-01-07T10:30:00Z",
            content: "Ceci est un post d'exemple numéro 7.",
            media_url: "https://example.com/media7.jpg",
            user: {
                id: 7,
                username: "Emma Scott",
                bio: 'hey',
                email: 'mail@mail.mail',
                profile_picture_url: 'https://avatars.githubusercontent.com/u/91917411?v=4',
            },
            user_id: 7,
        },
        {
            id: 8,
            created_at: "2025-01-08T10:00:00Z",
            updated_at: "2025-01-08T10:30:00Z",
            content: "Ceci est un post d'exemple numéro 8.",
            media_url: null,
            user: {
                id: 8,
                username: "Frank Harris",
                bio: 'hey',
                email: 'mail@mail.mail',
                profile_picture_url: 'https://avatars.githubusercontent.com/u/91917411?v=4',
            },
            user_id: 8,
        },
        {
            id: 9,
            created_at: "2025-01-09T10:00:00Z",
            updated_at: "2025-01-09T10:30:00Z",
            content: "Ceci est un post d'exemple numéro 9.",
            media_url: "https://example.com/media9.jpg",
            user: {
                id: 9,
                username: "Grace Lee",
                bio: 'hey',
                email: 'mail@mail.mail',
                profile_picture_url: 'https://avatars.githubusercontent.com/u/91917411?v=4',
            },
            user_id: 9,
        },
        {
            id: 10,
            created_at: "2025-01-10T10:00:00Z",
            updated_at: "2025-01-10T10:30:00Z",
            content: "Ceci est un post d'exemple numéro 10.",
            media_url: null,
            user: {
                id: 10,
                username: "Henry Clark",
                bio: 'hey',
                email: 'mail@mail.mail',
                profile_picture_url: 'https://avatars.githubusercontent.com/u/91917411?v=4',
            },
            user_id: 10,
        },
        {
            id: 11,
            created_at: "2025-01-11T10:00:00Z",
            updated_at: "2025-01-11T10:30:00Z",
            content: "Ceci est un post d'exemple numéro 11.",
            media_url: "https://example.com/media11.jpg",
            user: {
                id: 11,
                username: "Isla Adams",
                bio: 'hey',
                email: 'mail@mail.mail',
                profile_picture_url: 'https://avatars.githubusercontent.com/u/91917411?v=4',
            },
            user_id: 11,
        },
        {
            id: 12,
            created_at: "2025-01-12T10:00:00Z",
            updated_at: "2025-01-12T10:30:00Z",
            content: "Ceci est un post d'exemple numéro 12.",
            media_url: null,
            user: {
                id: 12,
                username: "Jack Walker",
                bio: 'hey',
                email: 'mail@mail.mail',
                profile_picture_url: 'https://avatars.githubusercontent.com/u/91917411?v=4',
            },
            user_id: 12,
        },
        {
            id: 13,
            created_at: "2025-01-13T10:00:00Z",
            updated_at: "2025-01-13T10:30:00Z",
            content: "Ceci est un post d'exemple numéro 13.",
            media_url: "https://example.com/media13.jpg",
            user: {
                id: 13,
                username: "Kimberly Young",
                bio: 'hey',
                email: 'mail@mail.mail',
                profile_picture_url: 'https://avatars.githubusercontent.com/u/91917411?v=4',
            },
            user_id: 13,
        },
        {
            id: 14,
            created_at: "2025-01-14T10:00:00Z",
            updated_at: "2025-01-14T10:30:00Z",
            content: "Ceci est un post d'exemple numéro 14.",
            media_url: "https://example.com/media14.jpg",
            user: {
                id: 14,
                username: "Liam Martin",
                bio: 'hey',
                email: 'mail@mail.mail',
                profile_picture_url: 'https://avatars.githubusercontent.com/u/91917411?v=4',
            },
            user_id: 14,
        },
        {
            id: 15,
            created_at: "2025-01-15T10:00:00Z",
            updated_at: "2025-01-15T10:30:00Z",
            content: "Ceci est un post d'exemple numéro 15.",
            media_url: null,
            user: {
                id: 15,
                username: "Mia Thompson",
                bio: 'hey',
                email: 'mail@mail.mail',
                profile_picture_url: 'https://avatars.githubusercontent.com/u/91917411?v=4',
            },
            user_id: 15,
        },
    ];
};

export const getMeMock = async (): Promise<{ user: User, token: string }> => {
    return {
        user: {
            id: 1,
            bio: 'Salut la team',
            email: 'noe.favier@mail.com',
            username: "Noé Favier",
            profile_picture_url: 'https://avatars.githubusercontent.com/u/91917411?v=4',
        },
        token: 'abcdef',
    };
}
