import { Pressable, useColorScheme } from 'react-native';
import { createMaterialTopTabNavigator } from '@react-navigation/material-top-tabs'
import { Icon } from '../components/FontAwesomeIcon';
import { Text, View } from '../components/Themed';
import { StyleSheet } from 'react-native';
import TabWallScreen from './homepage/wall';
import { storage } from '../components/Storage';
import React, { useState } from 'react';
import { Link } from 'expo-router';
import { Props } from '../constants/NavigationType'; 
import g from './globaldata';
import FontAwesome from '@expo/vector-icons/FontAwesome';

const Tabs = createMaterialTopTabNavigator();

export default function HomePageModalLayout({ route, navigation } : Props<'homepagemodal'>) {
  const colorScheme = useColorScheme();
  const [ismine, setismine] = useState(route.params.username == g.username);
  const [follownum, setfollownum] = useState(0);
  const [fansnum, setfansnum] = useState(0);

  return (
    <View style={styles.container}>
      <Text style={styles.title}>{route.params.username}</Text>
      <View>
        <View style={{flexDirection: 'row', padding: 15, alignContent:'center', alignSelf:'center'}}>
          <Link href="/modal" style = {styles.text}>  关注  </Link>
          <Text>               </Text>
          <Link href="/walladdmodal" style = {styles.text}>  粉丝  </Link>
        </View>
        
      </View>
      <Tabs.Navigator>
        <Tabs.Screen
          name="wall1"
          component={TabWallScreen}
          options={{
            title: ismine?'我的提问箱':'ta的提问箱',
            tabBarIcon: ({ color }) => <FontAwesome size={25} style={{ marginBottom: -3 }} name="dropbox" color={color} />,
          }}
        />
        <Tabs.Screen
          name="wall2"
          component={TabWallScreen}
          options={{
            title: ismine?'我的帖子':'ta的帖子',
            tabBarIcon: ({ color }) => <FontAwesome size={25} style={{ marginBottom: -3 }} name="file-text-o" color={color} />,
          }}
        />
        <Tabs.Screen
          name="wall3"
          component={TabWallScreen}
          options={{
            title: ismine?'我的表白':'ta的表白',
            tabBarIcon: ({ color }) => <FontAwesome size={24} style={{ marginBottom: -3 }} name="heart-o" color={color} />,
          }}
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
