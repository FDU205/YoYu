import FontAwesome from '@expo/vector-icons/FontAwesome';
import { DarkTheme, DefaultTheme, ThemeProvider } from '@react-navigation/native';
import { useFonts } from 'expo-font';
import { SplashScreen } from 'expo-router';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { useEffect, useState } from 'react';
import { useColorScheme } from 'react-native';
import UserLayout from './user/_layout'
import TabLayout from './tabs/_layout';
import HomePageLayout from './homepage/_latout';
import ModalScreen from './modal';
import WallAddModalScreen from './walladdmodal';
import { storage } from '../components/Storage';
import type { NavigationParamList, Props } from '../constants/NavigationType';
import { LogBox } from 'react-native';
import g from './globaldata';

LogBox.ignoreLogs([
  'Non-serializable values were found in the navigation state',
  'The navigation state parsed from the URL contains routes not present in the root navigator.',
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
  storage.load("token", (ret)=>{g.token=ret});
  storage.load("username", (ret)=>{g.username=ret});
  storage.load("userid", (ret)=>{g.userid=ret});
  useEffect(() => {()=>{
    if(g.token != null) {
      setisLoggedin(true);
    }
  }}, []);
  
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
              <Stack.Screen
                name="homepage"
                component={HomePageLayout}
              />
            </Stack.Group>
          )
        }        
      </Stack.Navigator>
    </ThemeProvider>
  );
}

