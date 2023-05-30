import { createNativeStackNavigator } from '@react-navigation/native-stack';
import LoginScreen from "./login";
import type { NavigationParamList, Props } from '../../constants/NavigationType';

const Stack = createNativeStackNavigator<NavigationParamList>();

export default function UserLayout({ route, navigation }: Props<'user'>) {
  return (
    <Stack.Navigator>
      <Stack.Screen 
        name="login"
        component={LoginScreen}
        options={{ headerShown: false }} 
        initialParams={{ setisLogin : route.params.setisLogin }}
      />
      {/* <Stack.Screen name="register" options={{ headerShown: false }} /> */}
    </Stack.Navigator>
  );
}