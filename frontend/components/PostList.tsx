import { Post } from '@/app/models/post';
import React from 'react';
import { FlatList, View, Text, Image, StyleSheet } from 'react-native';

interface PostListProps {
    posts: Post[];
}

const PostList: React.FC<PostListProps> = ({ posts }) => {
    const renderPost = ({ item }: { item: Post }) => {
        return (
            <View style={styles.postContainer}>
                <Text style={styles.userName}>{item.user.username}</Text>
                <Text style={styles.date}>
                    {new Date(item.created_at).toLocaleString()}
                </Text>
                <Text style={styles.content}>{item.content}</Text>

            </View>
        );
    };

    return (
        <FlatList
            data={posts}
            renderItem={renderPost}
            keyExtractor={(item) => item.id.toString()}
            ListEmptyComponent={
                <View style={{ alignItems: 'center', marginTop: 20 }}>
                    <Text>Aucun post à afficher</Text>
                </View>
            }
        />
    );
};

const styles = StyleSheet.create({
    postContainer: {
        marginBottom: 20,
        padding: 10,
        borderBottomWidth: 1,
        borderBottomColor: '#ddd',
    },
    userName: {
        fontWeight: 'bold',
        fontSize: 16,
    },
    date: {
        fontSize: 12,
        color: '#777',
    },
    content: {
        marginVertical: 10,
        fontSize: 14,
    },
    media: {
        width: '100%',
        height: 200,
        resizeMode: 'contain',
    },
});

export default PostList;