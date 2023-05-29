import { useColorScheme } from "react-native";
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import LoginScreen from "./login";

const Stack = createNativeStackNavigator();

export default function UserLayout() {
  return (
    <Stack.Navigator>
      <Stack.Screen 
        name="login"
        component={LoginScreen}
        options={{ headerShown: false }} 
      />
      {/* <Stack.Screen name="register" options={{ headerShown: false }} /> */}
    </Stack.Navigator>
  );
}