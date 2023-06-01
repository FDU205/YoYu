import type { NativeStackScreenProps } from '@react-navigation/native-stack';
import type { messageBoxinfo, postinfo } from './DataType';

type NavigationParamList = {
    user : { setisLogin : Function };
    login : { setisLogin : Function };
    register : { setisLogin : Function };
    tabs : { setisLogin : Function };
    tabwall: undefined;
    homewall: undefined;
    homeposts: undefined;
    boxes: undefined;
    homeboxes: {owner_id:number};
    modal : undefined;
    walladdmodal : undefined;
    boxaddmodal : undefined;
    threadaddmodal : {post_id:number, type:number, refresh:()=>void};
    boxaskmodal : { box:messageBoxinfo, onSubmit: ()=>void,};
    boxmodifymodal : { box:messageBoxinfo, changebox: (t:messageBoxinfo)=>void, setinbox: (t:boolean)=>void,};
    followmodal : undefined;
    fansmodal : undefined;
    postmodal : { postinfo:postinfo, refresh:()=>void, message_box_owner_id:number };
    homepage : { userid : number, username : string };
    homepagemodal : { userid : number, username : string };
};

type Props<T extends keyof NavigationParamList> = NativeStackScreenProps<NavigationParamList, T>;

export type { NavigationParamList, Props};