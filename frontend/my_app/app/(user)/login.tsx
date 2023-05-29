import { useState } from "react";
import { Button, TextInput, StyleSheet } from "react-native";
import { Text, View } from "../../components/Themed";
import { Link } from '@react-navigation/native';
import FontAwesome from '@expo/vector-icons/FontAwesome';

function Icon(props: {
  name: React.ComponentProps<typeof FontAwesome>['name'];
  color: string;
}) {
  return <FontAwesome size={28} style={{ marginBottom: -3 }} {...props} />;
}

export default function LoginScreen() {
  const [username, onChangeUsername] = useState('');
  const [password, onChangePassword] = useState('');
  
  return (
    <View style={styles.container}>
      <View style={styles.title}>
        <Text style={styles.welcome_text}>
          <Icon name="comments" color={'#2f95dc'} />  幽语
        </Text>
      </View>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <View style={styles.input}>
        <TextInput 
          style={styles.input_text}
          placeholder="用户名"
          onChangeText={username => onChangeUsername(username)}
          defaultValue={username}
        />

        <TextInput 
          style={styles.input_text}
          placeholder="密码"
          onChangeText={password => onChangePassword(password)}
          defaultValue={password}
        />
      </View>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <View style={styles.botton}>
        <Button
          title="登录"
          onPress={() => alert("login now!")}
        />

        <Link style={styles.link_text} to={{screen: '(tabs)'}}>
          没账号？注册！
        </Link>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
  },
  title: {
    alignItems: 'center',
    justifyContent: 'center',
  },
  welcome_text: {
    fontSize: 20,
    fontWeight: 'bold',
  },
  input: {
    width: '80%',
  },
  input_text: {
    alignItems: 'center',
    justifyContent: 'center',
    textAlign: "left",
    paddingVertical: 5,
    borderBottomWidth: 2,
    fontSize: 15,
  },
  link_text: {
    alignItems: 'center',
    justifyContent: 'center',
    textAlign: "left",
    paddingVertical: 5,
    borderBottomWidth: 2,
    fontSize: 12,
    textDecorationLine: 'none',
    color: 'gray'
  },
  botton: {
    width: '50%',
  },
  separator: {
    marginVertical: 30,
    height: 1,
    width: '80%',
  },
});