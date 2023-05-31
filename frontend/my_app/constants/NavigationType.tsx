import type { NativeStackScreenProps } from '@react-navigation/native-stack';

type NavigationParamList = {
    user : { setisLogin : Function };
    login : { setisLogin : Function };
    register : { setisLogin : Function };
    tabs : undefined;
    tabwall: undefined;
    modal : undefined;
    walladdmodal : undefined;
    homepage : { userid : number, username : string };
    homepagemodal : { userid : number, username : string };
};

type Props<T extends keyof NavigationParamList> = NativeStackScreenProps<NavigationParamList, T>;

export type { NavigationParamList, Props};