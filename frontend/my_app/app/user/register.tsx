import { useState } from "react";
import { Button, TextInput, StyleSheet, Alert } from "react-native";
import { Text, View } from "../../components/Themed";
import { Link } from '@react-navigation/native';
import type { Props } from '../../constants/NavigationType';
import { Icon } from "../../components/FontAwesomeIcon";
import { postData } from '../../components/Api';
import { storage } from "../../components/Storage";
import g from '../globaldata';

export default function RegisterScreen({ route, navigation }: Props<'register'>) {
  const [username, onChangeUsername] = useState('');
  const [password, onChangePassword] = useState('');
  const [repassword, onChangeRepassword] = useState('');
  
  return (
    <View style={styles.container}>
      <View style={styles.title}>
        <Text style={styles.welcome_text}>
          <Icon name="comments" color={'#2f95dc'} />  幽语 注册
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
          textContentType='password'
          secureTextEntry={true}
          placeholder="密码"
          onChangeText={password => onChangePassword(password)}
          defaultValue={password}
          maxLength={32}
        />
        <TextInput 
          style={styles.input_text}
          textContentType='password'
          secureTextEntry={true}
          placeholder="再输入一遍密码"
          onChangeText={repassword => onChangeRepassword(repassword)}
          defaultValue={repassword}
          maxLength={32}
        />
      </View>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <View style={styles.botton}>
        <Button
          title="注册"
          onPress={() => HandleRegister(route.params.setisLogin, username, password, repassword)}
        />
      </View>

      <Text 
        style={styles.link_text} 
        onPress={() => {
        navigation.navigate('login', {
            setisLogin : route.params.setisLogin,
        });
        }}
      >
        已有账号，登录！
      </Text>
    </View>
  );
}

function failToast(msg: string) {
  Alert.alert(msg);
}

function HandleRegister(setisLogin: Function, username: string, password: string, repassword: string) {
  if(password != repassword) {
    failToast("两次输入的密码不相同！");
    return;
  }
  if(username.length < 1) {
    failToast("用户名不能为空！");
    return;
  }
  if(password.length < 1) {
    failToast("密码不能为空！");
    return;
  }
  postData("/user/register", {"username" : username, "password" : password}).then(
    ret => {
      if(ret.code != 0) {
        throw new Error(ret.err_msg);
      } else {
        Promise.all([
          storage.save({key:"token", data:ret.data.token.slice(0)}),
          storage.save({key:"username", data:ret.data.username.slice(0)}),
          storage.save({key:"userid", data:ret.data.id}),
        ]).then(() => {
          // 将数据存储到全局变量中
          g.token = ret.data.token.slice(0);
          g.username = ret.data.username.slice(0);
          g.userid = ret.data.id;

          setisLogin(true);
        }).catch((err) => {
          console.log(err);
        });
      }
    }
  ).catch(
    err => {
      failToast(err+" 注册失败");
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