import React, { useState, useEffect } from 'react';
import { View } from '@ant-design/react-native';
import PostList from '@/components/PostList';
import { Post } from '@/models/post';
import { apiGetPosts } from '@/services/api';

export default function HomeScreen() {
  const [posts, setPosts] = useState<Post[]>([]);
  useEffect(() => { apiGetPosts().then((res) => setPosts(res)); }, []);

  return (
    <View style={{ width: '100%' }}>
      <PostList posts={posts} />
    </View>
  );
}
