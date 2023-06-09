import { useState } from "react";
import { Button, TextInput, StyleSheet, Alert, ImageBackground } from "react-native";
import { Text, View } from "../../components/Themed";
import { Link } from '@react-navigation/native';
import type { Props } from '../../constants/NavigationType';
import { Icon } from "../../components/FontAwesomeIcon";
import { postData } from '../../components/Api';
import { storage } from "../../components/Storage";
import g from '../globaldata';
import FontAwesome from "@expo/vector-icons/FontAwesome";

export default function RegisterScreen({ route, navigation }: Props<'register'>) {
  const [username, onChangeUsername] = useState('');
  const [password, onChangePassword] = useState('');
  const [repassword, onChangeRepassword] = useState('');
  
  return (
    <ImageBackground source={require("../../assets/images/b2.gif") } style={styles.container}>
      <View style={styles.container}>
        <View style={styles.title}>
          <FontAwesome
            name="comments-o"
            size={60}
            color={"white"}
            style={{ marginRight: 15 }}
          />
          <Text style={styles.welcome_text}>
              注册
          </Text>
        </View>
        <View style={styles.logincard}>
          <View style={styles.input}>
            <View style={styles.inputbox}>
              <TextInput 
                style={styles.input_text}
                textContentType='username'
                placeholder="用户名"
                onChangeText={username => onChangeUsername(username)}
                defaultValue={username}
                maxLength={255}
              />
            </View>
            
            <View style={styles.inputbox}>
              <TextInput 
                style={styles.input_text}
                secureTextEntry={true}
                placeholder="密码"
                onChangeText={password => onChangePassword(password)}
                defaultValue={password}
                maxLength={32}
              />
            </View>

            <View style={styles.inputbox}>
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
            
          </View>
        </View>
          <Text 
            style={styles.link_text} 
            onPress={() => {
              navigation.navigate('login', {
                setisLogin : route.params.setisLogin,
              });
            }}>
            已有账号，登录！
          </Text>

          <FontAwesome
            name="chevron-circle-right"
            size={90}
            color={"#6699FF"}
            style={{ marginRight: 15, top:200, }}
            onPress={() => HandleRegister(route.params.setisLogin, username, password, repassword)}
          />
      </View>
    </ImageBackground>
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
    backgroundColor:'transparent',
  },
  inputbox: {
    padding:15,
    backgroundColor:"#F6F6F6",
    shadowOffset: {width: 0, height: 3}, // 阴影偏移量
    shadowRadius: 6, // 阴影模糊半径
    shadowOpacity: 0.2, // 阴影不透明度
    shadowColor: '#5c5c5c', // 设置阴影色
    borderRadius: 40,
    marginVertical:15,
    top:-170,
    height:70,
    alignItems:'center',
  },
  logincard: {
    top:250,
    borderWidth: 0,
    position:'absolute',
    width:'100%',
    alignItems: 'center',
    justifyContent: 'center',
    borderRadius:50,
    height:'100%',
    backgroundColor:'white',
    shadowOffset: {width: 0, height: 4}, // 阴影偏移量
    shadowRadius: 8, // 阴影模糊半径
    shadowOpacity: 0.2, // 阴影不透明度
    shadowColor: '#5c5c5c', // 设置阴影色
  },
  title: {
    flexDirection:'row',
    position:'absolute',
    top:170,
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor:'transparent'
  },
  welcome_text: {
    fontSize: 45,
    fontWeight: 'bold',
    color:"white",
  },
  input: {
    width: '75%',
    top:-80,
  },
  input_text: {
    alignItems: 'center',
    justifyContent: 'center',
    textAlign: "left",
    paddingVertical: 5,
    fontSize: 20,
    marginTop:4,
  },
  link_text: {
    alignItems: 'flex-end',
    justifyContent: 'flex-end',
    textAlign: "right",
    paddingVertical: 5,
    fontSize: 12,
    textDecorationLine: 'none',
    color: 'gray',
    top:200,
    marginLeft:180,
  },
  botton: {
    width: 90,
    backgroundColor:"#F6F6F6",
    shadowOffset: {width: 0, height: 3}, // 阴影偏移量
    shadowRadius: 6, // 阴影模糊半径
    shadowOpacity: 0.2, // 阴影不透明度
    shadowColor: '#5c5c5c', // 设置阴影色
    borderRadius: 100,
    top:-250,
    alignItems:'center',
    alignContent:'center',
    height:90,
    flexDirection:'row',
  },
  separator: {
    marginVertical: 30,
    height: 1,
    width: '0%',
  },
});