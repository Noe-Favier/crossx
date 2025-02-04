import { Link, Stack } from 'expo-router';
import { StyleSheet } from 'react-native';

import React from 'react';
import { View } from '@ant-design/react-native';

export default function NotFoundScreen() {
  return (
    <>
      <Stack.Screen options={{ title: 'Oops!' }} />
      <View style={styles.container}>
        <Link href="/" style={{ color: 'blue' }}>Go to home screen</Link>
      </View>
    </>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    padding: 20,
  },
  link: {
    marginTop: 15,
    paddingVertical: 15,
  },
});
