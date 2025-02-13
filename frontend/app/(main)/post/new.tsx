import { View, Text } from 'react-native';
import React from 'react';
import { Flex, Input, TextareaItem } from '@ant-design/react-native';
import { Button } from '@ant-design/react-native';
import api, { apiPostNewPost } from '@/services/api';
import { Post } from '@/models/post';
import { router } from 'expo-router';

export default function NewPostScreen() {
    const [title, setTitle] = React.useState('');
    const [content, setContent] = React.useState('');

    return (
        <View style={{ padding: 20 }}>
            <Text style={{ fontSize: 20, paddingBottom: 20 }}>New post</Text>
            <Input
                value={title}
                onChangeText={setTitle}
                placeholder='Title'
                style={{ marginBottom: 10, borderBottomColor: '#ddd', borderBottomWidth: 1 }}
            />

            <TextareaItem rows={10}
                value={content}
                onChangeText={setContent}
                placeholder='Content'
                style={{ marginBottom: 10 }}
            />

            <Button
                onPress={() => { apiPostNewPost({ title, content }).then(() => { setTitle(''); setContent(''); router.replace('/') }) }}
                type='primary'
                style={{ marginTop: 10 }}>
                Submit
            </Button>
        </View>
    );
}