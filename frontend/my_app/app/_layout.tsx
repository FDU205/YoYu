import FontAwesome from '@expo/vector-icons/FontAwesome';
import { DarkTheme, DefaultTheme, ThemeProvider } from '@react-navigation/native';
import { useFonts } from 'expo-font';
import { SplashScreen } from 'expo-router';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { useEffect, useState } from 'react';
import { useColorScheme } from 'react-native';
import UserLayout from './user/_layout'
import TabLayout from './tabs/_layout';
import HomePageModalLayout from './homepagemodal';
import ModalScreen from './modal';
import WallAddModalScreen from './walladdmodal';
import FollowScreen from './followmodal';
import FansScreen from './fansmodal';
import { storage } from '../components/Storage';
import type { NavigationParamList, Props } from '../constants/NavigationType';
import { LogBox } from 'react-native';
import g from './globaldata';
import BoxAddModalScreen from './boxaddmodal';
import BoxModifyModalScreen from './boxmodifymodal';
import BoxAskModalScreen from './boxaskmodal';
import PostScreen from './postmodal';
import ThreadAddModalScreen from './threadaddmodal';

LogBox.ignoreLogs([
  'Non-serializable values were found in the navigation state',
  'The navigation state parsed from the URL contains routes not present in the root navigator.',
  'Sending `onAnimatedValueUpdate`'
]);

const Stack = createNativeStackNavigator<NavigationParamList>();

export {
  // Catch any errors thrown by the Layout component.
  ErrorBoundary,
} from 'expo-router';

export default function RootLayout() {
  const [loaded, error] = useFonts({
    SpaceMono: require('../assets/fonts/SpaceMono-Regular.ttf'),
    ...FontAwesome.font,
  });

  // Expo Router uses Error Boundaries to catch errors in the navigation tree.
  useEffect(() => {
    if (error) throw error;
  }, [error]);

  return (
    <>
      {/* Keep the splash screen open until the assets have loaded. In the future, we should just support async font loading with a native version of font-display. */}
      {!loaded && <SplashScreen />}
      {loaded && <RootLayoutNav />}
    </>
  );
}

function RootLayoutNav() {
  const colorScheme = useColorScheme();
  const [isLoggedin, setisLoggedin] = useState(false);

  const setisLogin =(t: boolean | ((prevState: boolean) => boolean)) => {setisLoggedin(t)};
  useEffect(() => {
    Promise.all([
      storage.load({key:"token"}),
      storage.load({key:"username"}),
      storage.load({key:"userid"}),
    ]).then(([token, username, userid]) => {
      // 将数据存储到全局变量中
      g.token = token;
      g.username = username;
      g.userid = userid;

      console.log(g.token);

      // 检查是否已登录
      if (g.token != null) {
        setisLoggedin(true);
      }
    }).catch((err) => {
      console.log(err);
    });
  }, []);
  
  return (
    <ThemeProvider value={colorScheme === 'dark' ? DarkTheme : DefaultTheme}>
      <Stack.Navigator>
        {
          !isLoggedin ? (
            <Stack.Group screenOptions={{ headerShown: false }}>
              <Stack.Screen 
                name="user" 
                component={UserLayout}
                options={{ headerShown: false }}
                initialParams={{ setisLogin:setisLogin }}
              />
            </Stack.Group>
          ) : (
            <Stack.Group screenOptions={{ headerShown: false }}>
              <Stack.Screen 
                name="tabs" 
                component={TabLayout}
                options={{ headerShown: false }}
                initialParams={{ setisLogin:setisLogin }}
              />
              <Stack.Screen 
                name="modal" 
                component={ModalScreen} 
                options={{ presentation: 'modal' }} 
              />
              <Stack.Screen 
                name="walladdmodal" 
                component={WallAddModalScreen} 
                options={{ presentation: 'modal' }} 
              />
              <Stack.Screen //???
                name="homepagemodal" 
                component={HomePageModalLayout} 
                options={{ presentation: 'modal' }} 
              />
              <Stack.Screen 
                name="boxaddmodal" 
                component={BoxAddModalScreen} 
                options={{ presentation: 'modal' }} 
              />
              <Stack.Screen 
                name="boxmodifymodal" 
                component={BoxModifyModalScreen} 
                options={{ presentation: 'modal' }} 
              />
              <Stack.Screen 
                name="boxaskmodal" 
                component={BoxAskModalScreen} 
                options={{ presentation: 'modal' }} 
              />
              <Stack.Screen 
                name="postmodal" 
                component={PostScreen} 
                options={{ presentation: 'modal' }} 
              />
              <Stack.Screen
                name="followmodal" 
                component={FollowScreen} 
                options={{ presentation: 'modal' }} 
              />
              <Stack.Screen 
                name="fansmodal" 
                component={FansScreen} 
                options={{ presentation: 'modal' }} 
              />
              <Stack.Screen 
                name="threadaddmodal" 
                component={ThreadAddModalScreen} 
                options={{ presentation: 'modal' }} 
              />
            </Stack.Group>
          )
        }        
      </Stack.Navigator>
    </ThemeProvider>
  );
}

