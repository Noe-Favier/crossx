import { StyleSheet, View, Text, Image, Pressable } from 'react-native';
import { Flex } from '@ant-design/react-native';
import { useAuth } from '@/context/AuthContext';
import React from 'react';
import { User } from '@/models/user';

export default function TabTwoScreen() {
  const auth = useAuth();
  const userState = auth.userState?.user;
  const logout = auth.logout;

  //TODO: enhance
  const unknownUserImg = 'https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png';

  return (
    <View style={{ flex: 1, padding: 20 }}>
      <Flex style={{ flex: 1 }} align='center' justify='center' direction='column'>
        <View style={{ marginBottom: 20 }}>
          <Image
            source={{ uri: ((userState?.profile_picture_url?.length ?? 0) > 0) ? userState?.profile_picture_url! : unknownUserImg }}
            style={{ width: 100, height: 100, borderRadius: 50 }}
          />
        </View>
        <Text style={{ fontSize: 20 }}>
          Profile of <Text style={{ fontWeight: 'bold' }}>{userState?.username}</Text>
        </Text>
        <View style={{ minWidth: 300, marginTop: 20, height: '10%' }}>
          <Flex
            justify={(userState?.bio?.length ?? 0) <= 20 ? 'between' : 'start'}
            direction={(userState?.bio?.length ?? 0) <= 15 ? 'row' : 'column'}
          >
            <Text>Bio</Text>
            <Text>{
              userState?.bio != '' ? userState?.bio : <Text style={{ fontStyle: 'italic' }}>Non défini</Text>
            }</Text>
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
