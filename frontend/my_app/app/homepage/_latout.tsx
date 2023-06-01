import { Alert, Button, Pressable, useColorScheme } from 'react-native';
import { createMaterialTopTabNavigator } from '@react-navigation/material-top-tabs'
import { Icon } from '../../components/FontAwesomeIcon';
import { Text, View } from '../../components/Themed';
import { StyleSheet } from 'react-native';
import HomeWallScreen from './wall';
import HomePostsScreen from './homeposts';
import HomeBoxesScreen from './homeboxes';
import { useState } from 'react';
import { Link } from 'expo-router';
import { Props } from '../../constants/NavigationType'; 
import g from '../globaldata';
import FontAwesome from '@expo/vector-icons/FontAwesome';
import { postData, getData } from '../../components/Api';

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

const Tabs = createMaterialTopTabNavigator();

export default function HomePageLayout({ route, navigation } : Props<'homepage'>) {
  const colorScheme = useColorScheme();
  const [ismine, setismine] = useState(route.params.username == g.username);
  const [follownum, setfollownum] = useState(0);
  const [fansnum, setfansnum] = useState(0);
  const [isfollow, setisfollow] = useState(false);

  // 获取关注数
  getData("/user/followcount",g.token).then(
    ret => {
      if(ret.code != 0) {
        throw new Error(ret.err_msg);
      } else {
        setfollownum(ret.count);
      }
    }
  ).catch(
    err => {
      Alert.alert(err+" 获取关注数失败");
    }
  )

  // 获取粉丝数
  getData("/user/fanscount",g.token).then(
    ret => {
      if(ret.code != 0) {
        throw new Error(ret.err_msg);
      } else {
        setfansnum(ret.count);
      }
    }
  ).catch(
    err => {
      Alert.alert(err+" 获取粉丝数失败");
    }
  )

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

  return (
    <View style={styles.container}>
      <Text style={styles.title}>{route.params.username}</Text>
      <View style={{alignContent:'center', alignSelf:'center'}}>
        {!ismine?
          (isfollow?
              (<Button title="已关注" disabled={true}></Button>):
              (<Button title="关注" onPress={() => {Follow(route.params.userid);}}></Button>)
          ):
          (<></>)
        }
      </View>
      <View>
        <View style={{flexDirection: 'row', padding: 15, alignContent:'center', alignSelf:'center'}}>
          <Link href="/followmodal" style = {styles.text}>  关注  </Link><Text style = {styles.text}>{follownum}</Text>
          <Text>               </Text>
          <Link href="/fansmodal" style = {styles.text}>  粉丝  </Link><Text style = {styles.text}>{fansnum}</Text>
        </View>
        
      </View>
      <Tabs.Navigator>
        {ismine?(<>
        <Tabs.Screen
          name="homeboxes"
          component={HomeBoxesScreen}
          options={{
            title: '我的提问箱',
            tabBarIcon: ({ color }) => <FontAwesome size={25} style={{ marginBottom: -3 }} name="dropbox" color={color} />,
          }} 
          initialParams={{owner_id:route.params.userid}}
        />
        <Tabs.Screen
          name="homeposts"
          component={HomePostsScreen}
          options={{
            title: '我的帖子',
            tabBarIcon: ({ color }) => <FontAwesome size={25} style={{ marginBottom: -3 }} name="file-text-o" color={color} />,
          }} 
        />
        <Tabs.Screen
          name="homewall"
          component={HomeWallScreen}
          options={{
            title: '我的表白',
            tabBarIcon: ({ color }) => <FontAwesome size={24} style={{ marginBottom: -3 }} name="heart-o" color={color} />,
          }} 
        /></>):(<>
          
          </>)}
        
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
