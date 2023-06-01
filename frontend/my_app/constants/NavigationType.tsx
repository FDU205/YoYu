import type { NativeStackScreenProps } from '@react-navigation/native-stack';
import type { messageBoxinfo } from './DataType';

type NavigationParamList = {
    user : { setisLogin : Function };
    login : { setisLogin : Function };
    register : { setisLogin : Function };
    tabs : undefined;
    tabwall: undefined;
    boxes: undefined;
    modal : undefined;
    walladdmodal : undefined;
    boxaddmodal : undefined;
    boxaskmodal : { box:messageBoxinfo, onSubmit: ()=>void,};
    boxmodifymodal : { box:messageBoxinfo, changebox: (t:messageBoxinfo)=>void, setinbox: (t:boolean)=>void,};
    followmodal : undefined;
    fansmodal : undefined;
    homepage : { userid : number, username : string };
    homepagemodal : { userid : number, username : string };
};

type Props<T extends keyof NavigationParamList> = NativeStackScreenProps<NavigationParamList, T>;

export type { NavigationParamList, Props};