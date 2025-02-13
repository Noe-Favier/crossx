import { View, Text, StyleProp, TextStyle, ViewStyle } from 'react-native';
import React, { useEffect, useState } from 'react';
import { Button, Icon, Input } from '@ant-design/react-native';
import { router, useLocalSearchParams } from 'expo-router';
import api, { apiGetMe, apiGetPost, apiLikePost, apiPostComment, apiUnlikePost } from '@/services/api';
import { Post } from '@/models/post';
import { Comment } from '@/models/comment';
import { useAuth } from '@/context/AuthContext';
import { FlatList, GestureHandlerRootView } from 'react-native-gesture-handler';
import { ListRenderItemInfo } from 'react-native';

export default function PostScreen() {
    const { id } = useLocalSearchParams();
    const [post, setPost] = useState<Post | null>(null);
    const [liked, setLiked] = useState(false);
    const [commentInput, setCommentInput] = useState<string>('');
    const user = useAuth().userState?.user;


    const renderComment = ({ item: comment }: ListRenderItemInfo<Comment>) => {
        return (
            <View>
                <Text style={{ fontWeight: 'bold' }}>{comment.user?.username}</Text>
                <Text>{comment.content}</Text>
            </View>
        );
    }

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
                    <Text>{new Date(post?.created_at ?? 0).toLocaleString()}</Text>
                    <Text>{post.user?.username || 'Unknown'}</Text>
                </View>
                <Text style={styles.content}>{post.content}</Text>
            </View>
            {(
                <>
                    <GestureHandlerRootView style={{ flex: 1, marginTop: 20, borderTopColor: 'black', borderTopWidth: 1 }}>
                        <FlatList
                            data={post.comments}
                            renderItem={renderComment}
                            scrollEnabled={true}
                            keyExtractor={(item) => item!.id!.toString()}
                            ListEmptyComponent={<Text style={{ fontStyle: 'italic' }}>No comments</Text>}
                        />
                    </GestureHandlerRootView>
                    <View style={{ flexDirection: 'row', justifyContent: 'space-between', alignItems: 'center', borderLeftColor: 'gray', borderLeftWidth: 1, height: 'auto' }}>
                        <Input
                            style={{ flexBasis: '70%' }}
                            placeholder="Add comment..."
                            value={commentInput}
                            onChangeText={(text) => setCommentInput(text)} />
                        <Button
                            onPress={() => { apiPostComment(post.id!, commentInput).then((cm: Comment) => { setCommentInput(''); post.comments.push(cm); }) }}
                            style={{ flexBasis: '20%' }}
                            type="primary"
                            disabled={!commentInput}>
                            <Icon name="send" />
                        </Button>
                    </View>
                </>
            )
            }
            <View>
                <Button
                    onPress={() => liked ? apiUnlikePost(post.id!).then(() => setLiked(false)) : apiLikePost(post.id!).then(() => setLiked(true))}
                    style={[styles.backButton, { borderColor: liked ? 'pink' : 'gray' }]}>
                    {liked ? 'Liked 💖' : 'Like'}
                </Button>
                <Button onPress={() => { router.navigate('/(main)'); }} type="primary" style={styles.backButton}>
                    Back
                </Button>
            </View>
        </View >
    );
}

const styles = {
    container: {
        flex: 1,
        padding: 20,
        justifyContent: 'space-between' as 'space-between',
    } as StyleProp<ViewStyle>,
    loadingText: {
        fontSize: 16,
        fontWeight: 'bold',
        textAlign: 'center',
    } as StyleProp<TextStyle>,
    title: {
        fontSize: 20,
        fontWeight: 'bold',
        marginBottom: 10,
    } as StyleProp<TextStyle>,
    metaContainer: {
        flexDirection: 'row',
        justifyContent: 'space-between',
    } as StyleProp<ViewStyle>,
    content: {
        marginVertical: 10,
        textAlign: 'justify',
    } as StyleProp<TextStyle>,
    backButton: {
        marginTop: 20,
    } as StyleProp<ViewStyle>,
};
