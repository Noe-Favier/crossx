import { AuthProvider, useAuth } from '@/context/AuthContext';
import { Flex, Icon, WingBlank } from '@ant-design/react-native';
import { useFonts } from 'expo-font';
import { Stack, router } from 'expo-router';
import * as SplashScreen from 'expo-splash-screen';
import React, { useEffect } from 'react';
import { Image, Pressable, Text, View } from 'react-native';
import { SafeAreaProvider, SafeAreaView } from 'react-native-safe-area-context';
import LoginScreen from './(main)/login';

SplashScreen.preventAutoHideAsync();


export default function RootLayout() {
  const [loaded] = useFonts({
    SpaceMono: require('../assets/fonts/SpaceMono-Regular.ttf'),
    Macondo: require('../assets/fonts/Macondo-Regular.ttf'),
    antoutline: require('@ant-design/icons-react-native/fonts/antoutline.ttf'),
  });

  useEffect(() => {
    if (loaded) {
      SplashScreen.hideAsync();
    }
  }, [loaded]);

  if (!loaded) {
    return null;
  }

  return (
    <AuthProvider>
      <AuthenticatedStack />
    </AuthProvider>
  );
}

export function AuthenticatedStack() {
  const userState = useAuth().userState;
  const tabBarItems = [
    {
      title: 'Home',
      icon: 'home',
      url: '/' as const,
    },
    {
      title: 'Profile',
      icon: 'user',
      url: '/profile' as const,
    },
    {
      title: 'Settings',
      icon: 'setting',
      url: '/settings' as const,
    }
  ]

  if (!userState) {
    return (
      <LoginScreen />
    );
  }

  return (
    <SafeAreaProvider>
      <SafeAreaView style={{ flex: 1 }}>

        {/* header */}
        <Flex style={{ backgroundColor: '#f5f5f9' }} justify='between'>
          <Flex align='center' justify='start'>
            <WingBlank style={{ paddingLeft: 1, paddingRight: 0, marginRight: 0, marginLeft: 10 }}>
              <Icon name='heat-map' size='lg' color='#000' />
            </WingBlank>
            <WingBlank>
              <Text style={{ fontSize: 17, paddingTop: 4, fontFamily: 'Macondo' }}>
                Donjons et Confessions
              </Text>
            </WingBlank>
          </Flex>
          <Pressable
            onPress={() => router.push('/profile')}
            style={{ marginRight: 0 }}>
            <Image source={{ uri: userState.user?.profile_picture_url ?? '' }} style={{ width: 30, height: 30, borderRadius: 20, marginRight: 5 }} />
          </Pressable>
        </Flex>

        {/* router outlet */}
        <View style={{ padding: 5, flex: 1 }}>
          <View style={{ borderRadius: 12, flex: 1, overflow: 'hidden' }}>
            <Stack screenOptions={{ headerShown: false }}>
              <Stack.Screen name="(main)" />
            </Stack>
          </View>
        </View>

        {/* tab bar */}
        <Flex justify='around' align='center' style={{ backgroundColor: '#f5f5f9', height: '8%' }}>
          {tabBarItems.map((item) => (
            <Pressable
              key={item.url}
              onPress={() => router.push(`/(main)${item.url}` as any)}
              android_ripple={{ color: 'lightgray' }}
              style={{ flexBasis: '33%', alignItems: 'center', justifyContent: 'center', height: 50 }}>
              <Icon name={item.icon as any} size='lg' color='#000' />
              <Text>{item.title}</Text>
            </Pressable>
          ))}
        </Flex>
      </SafeAreaView >
    </SafeAreaProvider >
  );
}