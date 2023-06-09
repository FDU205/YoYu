import FontAwesome from '@expo/vector-icons/FontAwesome';
import { Link } from 'expo-router';
import { Button, Pressable, useColorScheme } from 'react-native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import TabOneScreen from './one'
import TabWallScreen from './wall';
import BoxesScreen from './boxes';
import HomePageLayout from '../homepage/_latout';

import Colors from '../../constants/Colors';
import { Icon } from '../../components/FontAwesomeIcon';
import { Props } from '../../constants/NavigationType';
import g from '../globaldata';
import Logout from '../user/logout';

const Tabs = createBottomTabNavigator();

export default function TabLayout({ route, navigation }: Props<'tabs'>) {
  const colorScheme = useColorScheme();
  return (
    <Tabs.Navigator
      screenOptions={{
        tabBarActiveTintColor: Colors[colorScheme ?? 'light'].tint,
        headerTitleAlign:'center'
      }}
    >
      <Tabs.Screen
        name="boxes"
        component={BoxesScreen}
        options={{
          title: '提问箱',
          tabBarIcon: ({ color }) => <Icon name="envelope-square" color={color} />,
          headerRight: () => (
            <Link href="/boxaddmodal" asChild>
              <Pressable>
                {({ pressed }) => (
                  <FontAwesome
                    name="plus-circle"
                    size={25}
                    color={Colors[colorScheme ?? 'light'].text}
                    style={{ marginRight: 15, opacity: pressed ? 0.5 : 1 }}
                  />
                )}
              </Pressable>
            </Link>
          ),
        }}
      />
      <Tabs.Screen
        name="tabwall"
        component={TabWallScreen}
        options={{
          title: '表白墙',
          tabBarIcon: ({ color }) => <Icon name="heart-o" color={color} />,
          headerRight: () => (
            <Link href="/walladdmodal" asChild>
              <Pressable>
                {({ pressed }) => (
                  <FontAwesome
                    name="plus-circle"
                    size={25}
                    color={Colors[colorScheme ?? 'light'].text}
                    style={{ marginRight: 15, opacity: pressed ? 0.5 : 1 }}
                  />
                )}
              </Pressable>
            </Link>
          ),
        }}
      />
      <Tabs.Screen
        name="homepage"
        component={HomePageLayout}
        options={{
          title: '我的主页',
          tabBarLabel: '我的',
          tabBarIcon: ({ color }) => (<Icon name="user-circle-o" color={color} />),
          headerRight: () => (
            <Button
              title="退出登录"
              onPress={()=>{ route.params.setisLogin(false); Logout;}}
            />
          ),
        }}
        initialParams={{ userid: g.userid, username: g.username.slice(0) }}
      />
    </Tabs.Navigator>
  );
}
