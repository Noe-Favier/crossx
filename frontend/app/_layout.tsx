import { useState } from 'react';
import { Button, Flex, Icon, WingBlank } from '@ant-design/react-native';
import { Stack } from 'expo-router';
import { View, Text } from 'react-native';
import * as SplashScreen from 'expo-splash-screen';
import { useFonts } from 'expo-font';

SplashScreen.preventAutoHideAsync();

export default function RootLayout() {
  SplashScreen.hideAsync();

  const [darkMode, setDarkMode] = useState(false);
  const [loaded] = useFonts({
    SpaceMono: require('../assets/fonts/SpaceMono-Regular.ttf'),
    Macondo: require('../assets/fonts/Macondo-Regular.ttf'),
    antoutline: require('@ant-design/icons-react-native/fonts/antoutline.ttf'),
  });

  if (!loaded) {
    return null;
  }

  function toggleDarkMode() {
    setDarkMode(!darkMode);
  }

  return (
    <View>
      <Flex style={{ backgroundColor: '#f5f5f9' }} justify='between'>
        <Flex align='center' justify='start'>
          <WingBlank style={{ paddingLeft: 2, paddingRight: 2, marginRight: 0, marginLeft: 15 }}>
            <Icon name='heat-map' size='lg' color='#000' />
          </WingBlank>
          <WingBlank>
            <Text style={{ fontSize: 24, fontFamily: 'Macondo' }}>
              Donjons et Confessions
            </Text>
          </WingBlank>
        </Flex>
        <WingBlank style={{ marginRight: 0 }}>
          <Button style={{ backgroundColor: 'none' }} onPress={() => toggleDarkMode()}>
            <Icon name={darkMode ? 'eye-invisible' : 'bulb'} size='md' />
          </Button>
        </WingBlank>
      </Flex>

      {/* router outlet */}
      <Stack screenOptions={{ headerShown: false }}>
        <Stack.Screen name='(main)' />
      </Stack>
    </View >
  );
}