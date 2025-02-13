import { View, Text } from 'react-native';
import React, { useEffect, useState } from 'react';
import { Button } from '@ant-design/react-native';
import { router, useLocalSearchParams } from 'expo-router';
import { apiGetMe, apiGetPost, apiLikePost, apiUnlikePost } from '@/services/api';
import { Post } from '@/models/post';
import { useAuth } from '@/context/AuthContext';

export default function PostScreen() {
    const { id } = useLocalSearchParams();
    const [post, setPost] = useState<Post | null>(null);
    const [liked, setLiked] = useState(false);

    const user = useAuth().userState?.user;

    useEffect(() => {
        if (!id) return;
        console.log('PostScreen', id);

        apiGetPost(Number(id))
            .then((res) => {
                setPost(res);
                setLiked(res.likes.some((usr) => usr.id === user?.id));
            })
            .catch((err) => console.error('Failed to fetch post:', err));
    }, [id]);

    if (!post) {
        return (
            <View style={styles.container}>
                <Text style={styles.loadingText}>Loading...</Text>
            </View>
        );
    }

    return (
        <View style={styles.container}>
            <View>
                <Text style={styles.title}>{post.title}</Text>
                <View style={styles.metaContainer}>
                    <Text>{new Date(post.created_at).toLocaleString()}</Text>
                    <Text>{post.user?.username || 'Unknown'}</Text>
                </View>
                <Text style={styles.content}>{post.content}</Text>
            </View>
            <View>
                <Button
                    onPress={() => liked ? apiUnlikePost(post.id!).then(() => setLiked(false)) : apiLikePost(post.id!).then(() => setLiked(true))}
                    type="primary"
                    style={{ ...styles.backButton, backgroundColor: liked ? 'red' : 'gray', borderColor: 'transparent' }}>
                    {liked ? 'unlike' : 'Like'}
                </Button>
                <Button onPress={() => { router.navigate('/(main)'); }} type="primary" style={styles.backButton}>
                    Back
                </Button>
            </View>
        </View>
    );
}

const styles = {
    container: {
        flex: 1,
        padding: 20,
        justifyContent: 'space-between',
    },
    loadingText: {
        fontSize: 16,
        fontWeight: 'bold',
        textAlign: 'center',
    },
    title: {
        fontSize: 20,
        fontWeight: 'bold',
        marginBottom: 10,
    },
    metaContainer: {
        flexDirection: 'row',
        justifyContent: 'space-between',
    },
    content: {
        marginVertical: 10,
        textAlign: 'justify',
    },
    backButton: {
        marginTop: 20,
    },
};
