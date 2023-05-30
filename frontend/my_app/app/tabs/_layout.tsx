import FontAwesome from '@expo/vector-icons/FontAwesome';
import { Link } from 'expo-router';
import { Pressable, useColorScheme } from 'react-native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import TabOneScreen from './one'
import TabTwoScreen from './two';
import TabThreeScreen from './three';

import Colors from '../../constants/Colors';
import { Icon } from '../../components/FontAwesomeIcon';

const Tabs = createBottomTabNavigator();

export default function TabLayout() {
  const colorScheme = useColorScheme();

  return (
    <Tabs.Navigator
      screenOptions={{
        tabBarActiveTintColor: Colors[colorScheme ?? 'light'].tint,
      }}>
      <Tabs.Screen
        name="index"
        component={TabOneScreen}
        options={{
          title: 'Tab Oneeeee',
          tabBarIcon: ({ color }) => <Icon name="code" color={color} />,
          headerRight: () => (
            <Link href="/modal" asChild>
              <Pressable>
                {({ pressed }) => (
                  <FontAwesome
                    name="info-circle"
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
        name="two"
        component={TabTwoScreen}
        options={{
          title: 'Tab Two',
          tabBarIcon: ({ color }) => <Icon name="code" color={color} />,
        }}
      />
      <Tabs.Screen
        name="three"
        component={TabThreeScreen}
        options={{
          tabBarLabel: 'Tab Three',
          tabBarIcon: ({ color }) => (<Icon name="code" color={color} />),
        }}
      />
    </Tabs.Navigator>
  );
}
