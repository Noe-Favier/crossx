import React, { useState } from 'react';
import { View, TextInput, Image, Pressable, KeyboardAvoidingView, Platform, ScrollView } from 'react-native';
import * as ImagePicker from 'expo-image-picker';
import { useAuth } from '@/context/AuthContext';
import { Button } from '@ant-design/react-native';

export default function SignupScreen() {
    const { login, signup } = useAuth();
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [email, setEmail] = useState('');
    const [image, setImage] = useState<string | null>(null);
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

    <Image 
    source={require('@/assets/images/logo_donjons_et_confessions_clean.png')}
    style={{ width: 60, height: 60 }} // Ajustez ces dimensions selon vos besoins
    />

    const proceedSignup = async () => {
        const fd = new FormData();
        fd.append('username', username);
        fd.append('password', password);
        fd.append('email', email);

        if (image) {
            // Assurez-vous que c'est bien un Blob ou un File
            console.log('Image type:', typeof image, typeof image === 'string');
            fd.append('profileImage', image);
        }

        // Log le contenu du FormData avant envoi
        fd.forEach((value, key) => {
            console.log(`${key}:`, value);
        });

        try {
            const response = await signup(fd);
            console.log('Signup response:', response);
        } catch (error) {
            console.error('Signup error:', error);
        }
    };

    const commonFormElt = (
        <>
            <TextInput style={inputStyle} placeholder="Username" value={username} onChangeText={setUsername} />
            <TextInput style={inputStyle} placeholder="Password" secureTextEntry value={password} onChangeText={setPassword} />
        </>
    );

    return (
        <KeyboardAvoidingView
            style={{ flex: 1 }}
            behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
        >
            <ScrollView contentContainerStyle={{ flexGrow: 1, justifyContent: 'center', paddingHorizontal: 20 }}>
                {isSigningUp ? (
                    <>
                        {commonFormElt}
                        <TextInput style={inputStyle} placeholder="E-mail" value={email} onChangeText={setEmail} />
                        <View style={{ alignItems: 'center' }}>
                            <Pressable onPress={pickImage}>
                                <Image
                                    source={image ? { uri: image } : require('@/assets/images/no-profile.webp')}
                                    style={{ width: 200, height: 200, marginTop: 20, marginBottom: 10 }} />
                            </Pressable>
                        </View>
                        <Button type='ghost' onPress={() => proceedSignup()}>
                            Sign up
                        </Button>
                        <View style={{ height: 1, backgroundColor: 'gray', opacity: .20, marginVertical: 20 }} />
                        <Button type='ghost' style={{ height: 'auto', borderColor: 'transparent', paddingVertical: 2, marginTop: 20 }} onPress={() => setIsSigningUp(false)}>
                            Login ?
                        </Button>
                    </>
                ) : (
                    <>
                        {commonFormElt}
                        <Button type='ghost' onPress={() => login(username, password)}>
                            Login
                        </Button>
                        <View style={{ height: 1, backgroundColor: 'gray', opacity: .20, marginVertical: 20 }} />
                        <Button type='ghost' style={{ height: 'auto', borderColor: 'transparent', paddingVertical: 2, marginTop: 20 }} onPress={() => setIsSigningUp(true)}>
                            Sign up ?
                        </Button>
                    </>
                )}
            </ScrollView>
        </KeyboardAvoidingView>
    );
}

const inputStyle = {
    marginBottom: 20,
    padding: 10,
    borderBottomWidth: 1,
    borderBottomColor: 'black',
};
