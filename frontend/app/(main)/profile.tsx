import { StyleSheet } from 'react-native';


import { Button, Flex, Toast, WingBlank } from '@ant-design/react-native';

export default function TabTwoScreen() {
  return (
    <Flex direction="column" align="center" justify="center">
      <WingBlank>
        <Button
          onPress={() => {
            Toast.info('This is a toast message.');
          }}
        >
          Show Toast
        </Button>
      </WingBlank>
    </Flex>
  );
}

const styles = StyleSheet.create({
  headerImage: {
    color: '#808080',
    bottom: -90,
    left: -35,
    position: 'absolute',
  },
  titleContainer: {
    flexDirection: 'row',
    gap: 8,
  },
});
