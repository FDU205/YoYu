import { Pressable, useColorScheme } from 'react-native';
import { createMaterialTopTabNavigator } from '@react-navigation/material-top-tabs'
import { Icon } from '../../components/FontAwesomeIcon';
import { Text, View } from '../../components/Themed';
import { StyleSheet } from 'react-native';
import TabWallScreen from './wall';
import { useState } from 'react';
import { Link } from 'expo-router';
import { Props } from '../../constants/NavigationType'; 
import g from '../globaldata';

const Tabs = createMaterialTopTabNavigator();

export default function HomePageLayout({ route, navigation } : Props<'homepage'>) {
  const colorScheme = useColorScheme();
  const [ismine, setismine] = useState(route.params.username == g.username);
  const [follownum, setfollownum] = useState(0);
  const [fansnum, setfansnum] = useState(0);

  return (
    <View style={styles.container}>
      <Text style={styles.title}>{route.params.username}</Text>
      <View>
        <Link href="/modal" style = {styles.text}>关注</Link>
        <Link href="/walladdmodal" style = {styles.text}>粉丝</Link>
      </View>
      <Tabs.Navigator>
        <Tabs.Screen
          name="wall1"
          component={TabWallScreen}
          options={{
            title: ismine?'我的提问箱':'ta的提问箱',
          }}
        />
        <Tabs.Screen
          name="wall2"
          component={TabWallScreen}
          options={{
            title: ismine?'我的帖子':'ta的帖子',
            tabBarIcon: ({ color }) => <Icon name="heart-o" color={color} />,
          }}
        />
        <Tabs.Screen
          name="wall3"
          component={TabWallScreen}
          options={{
            title: ismine?'我的表白':'ta的表白',
            tabBarIcon: ({ color }) => <Icon name="heart-o" color={color} />,
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
