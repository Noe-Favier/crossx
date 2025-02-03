import { StyleSheet, View, Text, Image, Pressable } from 'react-native';
import { Flex } from '@ant-design/react-native';
import { useAuth } from '@/context/AuthContext';
import React from 'react';

export default function TabTwoScreen() {
  const auth = useAuth();
  const userState = auth.userState;
  const logout = auth.logout;

  return (
    <Flex direction="column" align="center" justify="center">
      <View>
        <Flex style={{ flex: 1 }} align='center' justify='center' direction='column'>
          <View style={{ marginBottom: 20 }}>
            <Image
              source={{ uri: userState?.user?.profile_picture_url || '' }}
              style={{ width: 100, height: 100, borderRadius: 50 }}
            />
          </View>
          <Text style={{ fontSize: 20 }}>
            Profile of <Text style={{ fontWeight: 'bold' }}>{userState?.user?.username}</Text>
          </Text>
          <View style={{ width: '100%', marginTop: 20, height: '10%' }}>
            <Flex justify='between'>
              <Text>Email</Text>
              <Text>{userState?.user?.email}</Text>
            </Flex>
            <Flex
              justify={userState?.user?.bio?.length ?? 0 <= 20 ? 'between' : 'start'}
              direction={userState?.user?.bio?.length ?? 0 <= 15 ? 'row' : 'column'}
            >
              <Text>Bio</Text>
              <Text>{userState?.user?.bio}</Text>
            </Flex>
          </View>
        </Flex>
        <View style={{ width: '100%' }}>
          <Pressable
            android_ripple={{ color: 'red' }}
            style={({ pressed }) => [
              styles.logoutButton,
              { backgroundColor: pressed ? '#FF6347' : '#FF4500' },
            ]}
            onPress={() => { logout(); }}>
            <Text style={styles.logoutText}>Logout</Text>
          </Pressable>
        </View>
      </View>
    </Flex>
  );
}

const styles = StyleSheet.create({
  logoutButton: {
    marginBottom: 20,
    marginHorizontal: 'auto',
    paddingVertical: 5,
    paddingHorizontal: 30,
    borderRadius: 5,
    alignItems: 'center',
    justifyContent: 'center',
    elevation: 2,

  },
  logoutText: {
    color: 'white',
    fontSize: 18,
    fontWeight: 'bold',
  },
});
