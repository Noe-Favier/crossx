import React, { useState } from 'react';
import { View, TextInput } from 'react-native';
import { useAuth } from '@/context/AuthContext';
import { Button } from '@ant-design/react-native';

export default function LoginScreen() {
    const { login } = useAuth();
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    return (
        <View style={{ flex: 1, paddingHorizontal: 20, justifyContent: 'center' }}>
            <TextInput style={inputStyle} placeholder="Email" value={email} onChangeText={setEmail} />
            <TextInput style={inputStyle} placeholder="Password" secureTextEntry value={password} onChangeText={setPassword} />
            <Button type='ghost' onPress={() => login(email, password)}>
                Login
            </Button>
        </View>
    );
}

const inputStyle = {
    marginBottom: 20,
    padding: 10,
    borderBottomWidth: 1,
    borderBottomColor: 'black',
};
