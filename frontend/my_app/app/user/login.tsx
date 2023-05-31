import { useState } from "react";
import { Button, TextInput, StyleSheet, Alert } from "react-native";
import { Text, View } from "../../components/Themed";
import { Link } from '@react-navigation/native';
import type { Props } from '../../constants/NavigationType';
import { Icon } from "../../components/FontAwesomeIcon";
import { postData } from '../../components/Api';
import { storage } from "../../components/Storage";

export default function LoginScreen({ route, navigation }: Props<'login'>) {
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
          textContentType='username'
          placeholder="用户名"
          onChangeText={username => onChangeUsername(username)}
          defaultValue={username}
          maxLength={255}
        />

        <TextInput 
          style={styles.input_text}
          secureTextEntry={true}
          placeholder="密码"
          onChangeText={password => onChangePassword(password)}
          defaultValue={password}
          maxLength={32}
        />
      </View>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <View style={styles.botton}>
        <Button
          title="登录"
          onPress={() => HandleLogin(route.params.setisLogin, username, password)}
        />

        <Text 
          style={styles.link_text} 
          onPress={() => {
            navigation.navigate('register', {
              setisLogin : route.params.setisLogin,
            });
          }}>
          没账号？注册！
        </Text>
      </View>
    </View>
  );
}

function failToast(msg: string) {
  Alert.alert(msg);
}

function HandleLogin(setisLogin: Function, username: string, password: string) {
  if(username.length < 1) {
    failToast("用户名不能为空！");
    return;
  }
  if(password.length < 1) {
    failToast("密码不能为空！");
    return;
  }
  postData("/user/login", {"username" : username, "password" : password}).then(
    ret => {
      if(ret.code != 0) {
        throw new Error(ret.err_msg);
      } else {
        storage.save("token", ret.data.token.slice(0));
        setisLogin(true);
      }
    }
  ).catch(
    err => {
      failToast(err+" 登录失败");
    }
  )
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
    width: '70%',
  },
  input_text: {
    alignItems: 'center',
    justifyContent: 'center',
    textAlign: "left",
    paddingVertical: 5,
    borderBottomWidth: 1,
    fontSize: 15,
    // borderColor: 'gray', borderWidth: 1
  },
  link_text: {
    alignItems: 'center',
    justifyContent: 'center',
    textAlign: "left",
    paddingVertical: 5,
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
    width: '0%',
  },
});