import { Flex, Icon, WingBlank } from '@ant-design/react-native';
import { NavigationContainer, NavigationIndependentTree } from '@react-navigation/native';
import { useFonts } from 'expo-font';
import { Stack, router } from 'expo-router';
import * as SplashScreen from 'expo-splash-screen';
import { Text, TouchableOpacity, View } from 'react-native';
import { useNavigationState } from '@react-navigation/native';

SplashScreen.preventAutoHideAsync();

export default function RootLayout() {
  SplashScreen.hideAsync();

  const [loaded] = useFonts({
    SpaceMono: require('../assets/fonts/SpaceMono-Regular.ttf'),
    Macondo: require('../assets/fonts/Macondo-Regular.ttf'),
    antoutline: require('@ant-design/icons-react-native/fonts/antoutline.ttf'),
  });

  if (!loaded) {
    return null;
  }


  function HomeButton() {
    const state = useNavigationState(state => state);
    const currentRoute = state?.routes[state.index]?.name;
    console.log('(main)', currentRoute);
    if (currentRoute !== '(main)') {
      return (
        <TouchableOpacity
          style={{
            position: 'absolute',
            bottom: 20,
            left: 20,
            backgroundColor: '#fff',
            padding: 10,
            borderRadius: 5,
            elevation: 5,
          }}
          onPress={() => router.push('/')}
        >
          <Text style={{ fontSize: 16, color: '#000' }}>Home</Text>
        </TouchableOpacity>
      );
    }
  }

  return (
    <View style={{ flex: 1 }}>
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
        <WingBlank style={{ marginRight: 0 }}>
          <TouchableOpacity style={{ backgroundColor: 'none' }} onPress={() => router.push('/(main)/profile')}>
            <Icon name='plus-circle' size='lg' color='black' style={{ marginRight: 5 }} />
          </TouchableOpacity>
        </WingBlank>
      </Flex>

      <HomeButton />

      {/* router outlet */}
      <View style={{ padding: 5, flex: 1 }}>
        <View style={{ borderRadius: 12, flex: 1, overflow: 'hidden' }}>
          <Stack screenOptions={{ headerShown: false }}>
            <Stack.Screen name="(main)" />
          </Stack>
        </View>
      </View>
    </View>
  );
}
