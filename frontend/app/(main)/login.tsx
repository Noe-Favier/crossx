import React, { useState } from 'react';
import { View, TextInput, Image, Pressable } from 'react-native';
import * as ImagePicker from 'expo-image-picker';
import { useAuth } from '@/context/AuthContext';
import { Button } from '@ant-design/react-native';

export default function SignupScreen() {
    const { login, signup } = useAuth();
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [email, setEmail] = useState('');
    const [image, setImage] = useState<string | null>(null);

    //
    const [isSigningUp, setIsSigningUp] = useState(false);

    const pickImage = async () => {
        const { status } = await ImagePicker.requestMediaLibraryPermissionsAsync();
        if (status !== 'granted') {
            alert('Permission refusée !');
            return;
        }

        const result = await ImagePicker.launchImageLibraryAsync({
            mediaTypes: 'images',
            allowsEditing: true,
            aspect: [4, 3],
            quality: 1,
        });

        if (!result.canceled) {
            setImage(result.assets[0].uri);
        }
    };

    //\\

    const commonFormElt = (
        <>
            <TextInput style={inputStyle} placeholder="Username" value={username} onChangeText={setUsername} />
            <TextInput style={inputStyle} placeholder="Password" secureTextEntry value={password} onChangeText={setPassword} />
        </>
    );

    if (isSigningUp) {
        return (
            <View style={{ flex: 1, paddingHorizontal: 20, justifyContent: 'center' }}>
                {commonFormElt}
                <TextInput style={inputStyle} placeholder="E-mail" secureTextEntry value={email} onChangeText={setEmail} />
                <View style={{ alignItems: 'center' }}>
                    <Pressable onPress={pickImage}>
                        <Image
                            source={image ? { uri: image } : require('@/assets/images/no-profile.webp')}
                            style={{ width: 200, height: 200, marginTop: 20, marginBottom: 10, }} />
                    </Pressable>
                </View>

                <Button type='ghost' onPress={() => signup(username, password, email, image)}>
                    Sign up
                </Button>
                <View style={{ height: 1, backgroundColor: 'gray', opacity: .20, marginVertical: 20 }} />
                <Button type='ghost' style={{ height: 'auto', borderColor: 'transparent', paddingVertical: 2, marginTop: 20 }} onPress={() => setIsSigningUp(false)}>
                    Login ?
                </Button>
            </View>
        );
    } else {
        return (<View style={{ flex: 1, paddingHorizontal: 20, justifyContent: 'center' }}>
            {commonFormElt}
            <Button type='ghost' onPress={() => login(username, password)}>
                Login
            </Button>
            <View style={{ height: 1, backgroundColor: 'gray', opacity: .20, marginVertical: 20 }} />
            <Button type='ghost' style={{ height: 'auto', borderColor: 'transparent', paddingVertical: 2, marginTop: 20 }} onPress={() => setIsSigningUp(true)}>
                Sign up ?
            </Button>
        </View>)
    }
}

const inputStyle = {
    marginBottom: 20,
    padding: 10,
    borderBottomWidth: 1,
    borderBottomColor: 'black',
};