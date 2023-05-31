import type { NativeStackScreenProps } from '@react-navigation/native-stack';

type NavigationParamList = {
    user : { setisLogin : Function };
    login : { setisLogin : Function };
    register : { setisLogin : Function };
    tabs : undefined;
    modal : undefined;
    walladdmodal : undefined;
};

type Props<T extends keyof NavigationParamList> = NativeStackScreenProps<NavigationParamList, T>;

export type { NavigationParamList, Props};