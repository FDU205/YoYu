import { Alert, Button, Pressable, useColorScheme } from 'react-native';
import { createMaterialTopTabNavigator } from '@react-navigation/material-top-tabs'
import { Icon } from '../components/FontAwesomeIcon';
import { Text, View } from '../components/Themed';
import { StyleSheet } from 'react-native';
import HomePostScreen from './homepage/wall';
import { storage } from '../components/Storage';
import React, { useEffect, useState } from 'react';
import { Link } from 'expo-router';
import { Props } from '../constants/NavigationType'; 
import g from './globaldata';
import FontAwesome from '@expo/vector-icons/FontAwesome';
import { deleteData, getData, postData } from '../components/Api';
import HomeBoxesScreen from './homepage/homeboxes';

// 关注
const Follow = (userid: number) => {
  postData("/user/follow",{"follow_id":userid},g.token).then(
    ret => {
      if(ret.code != 0) {
        throw new Error(ret.err_msg);
      } else {
        Alert.alert("关注成功");
      }
    }
  ).catch(
    err => {
      Alert.alert(err+" 关注失败");
    }
  )
}

// 取消关注
const cancelFollow = (followid: number) => {
  deleteData("/user/unfollow",{"follow_id":followid},g.token).then(
    ret => {
      if(ret.code != 0) {
        throw new Error(ret.err_msg);
      } else {
        Alert.alert("取消关注成功");
      }
    }
  ).catch(
    err => {
      Alert.alert(err+" 取消关注失败");
    }
  )
}

const Tabs = createMaterialTopTabNavigator();

export default function HomePageLayout({ route, navigation } : Props<'homepagemodal'>) {
  const colorScheme = useColorScheme();
  const [ismine, setismine] = useState(route.params.username == g.username);
  const [isfollow, setisfollow] = useState(false);

  useEffect(() => {
    // 获取是否关注
    getData("/user/isfollow?follow_id="+route.params.userid,g.token).then(
      ret => {
        if(ret.code != 0) {
          throw new Error(ret.err_msg);
        } else {
          setisfollow(ret.yes);
        }
      }
    ).catch(
      err => {
        Alert.alert(err+" 获取是否关注失败");
      }
    )
  },[])

  return (
    <View style={styles.container}>
      <View style={styles.title}></View>
      <View style={styles.title}></View>  
      <Text style={styles.title}>{route.params.username}</Text>
      <View style={styles.text}></View>
      <View style={{alignContent:'center', alignSelf:'center'}}>
        {!ismine?
          (isfollow?
              (<Text style={styles.button2} onPress={() => {cancelFollow(route.params.userid);setisfollow(false);}}>已关注</Text>):
              (<Text style={styles.button} onPress={() => {Follow(route.params.userid);setisfollow(true)}}>关注</Text>)
          ):
          (<></>)
        }
      </View>
      <Tabs.Navigator>
        <Tabs.Screen
          name="homeboxes"
          component={HomeBoxesScreen}
          options={{
            title: ismine?'我的提问箱':'ta的提问箱',
            tabBarIcon: ({ color }) => <FontAwesome size={25} style={{ marginBottom: -3 }} name="dropbox" color={color} />,
          }}
          initialParams={{owner_id:route.params.userid}}
        />
      </Tabs.Navigator>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
  },
  title: {
    alignSelf: 'center',
    marginTop: 15,
    fontSize: 20,
  },
  text: {
    justifyContent:'center',
    alignSelf: 'center',
    marginTop: 15,
    fontSize: 15,
  },
  button: {
    borderStyle: 'solid',
    borderColor: 'blue',    
    color: 'blue',
    borderWidth: 1,
    borderRadius: 20, // 圆角
    padding: 8,
    marginVertical:2,
  },
  button2: {
    borderStyle: 'solid',
    borderColor: 'black', 
    color: 'black',  
    borderWidth: 1,
    borderRadius: 20, // 圆角
    padding: 8,
    marginVertical:2,
  },
  separator: {
    alignSelf: 'center',
    marginTop: 15,
    height: 1,
    width: '80%',
  },
  flat: {
    height: '100%'
  }
});
