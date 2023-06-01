import { Alert, Pressable, FlatList, StyleSheet, useColorScheme, Modal } from 'react-native';
import FontAwesome from '@expo/vector-icons/FontAwesome';
import Colors from '../constants/Colors';
import { Text, View } from '../components/Themed';
import Card from '../components/Card';
import { deleteData, getData } from '../components/Api';
import { SetStateAction, useEffect, useState } from 'react';
import type { thread, postdetail } from '../constants/DataType';
import { NavigationParamList, Props } from '../constants/NavigationType';
import type { NativeStackNavigationProp } from '@react-navigation/native-stack';
import g from './globaldata';
import PopUpModal from '../components/PopUpModal';
import { TouchableOpacity } from 'react-native-gesture-handler';
import FlatListTail from '../components/FlatListTail';
import LongTextBox from '../components/LongTextBox';

const onPress = (userid: number, username: string, navigation: NativeStackNavigationProp<NavigationParamList, "fansmodal", undefined>) => {
  navigation.navigate(
    'homepagemodal',{userid: userid, username: username},
  );
  return;
};

const PAGE_SIZE = 10;

export default function PostScreen({ route, navigation }: Props<'postmodal'>) {
  const colorScheme = useColorScheme();
  const [data, setdata] = useState(new Array<thread>(0));
  const [post_detail, setpostdetail] = useState({
    id: 0,
    poster_id: 0,
    poster_name: '',
    content: '',
    visibility: 0,
    message_box_id: 0
  });
  const [ifshowdeletemodal, onChangeifshowdeletemodal] = useState(false);
  const [ifshowthreadmodal, onChangeifshowthreadmodal] = useState(false);

  const getNew = async (setdata: { (value: SetStateAction<thread[]>): void; (arg0: any): void; }) => {
    await getData("/post/"+route.params.postinfo.id, g.token).then(
      ret => {
        if(ret.code != 0) {
          throw new Error(ret.err_msg);
        } else {
          post_detail.id = ret.data.id;
          post_detail.poster_id = ret.data.poster_id;
          post_detail.poster_name = ret.data.poster_name;
          post_detail.content = ret.data.content;
          post_detail.visibility = ret.data.visibility;
          post_detail.message_box_id = ret.data.message_box_id;
          setdata(ret.data.threads);
        }
      }
    ).catch(
      err => {
        failToast(err+" 刷新失败");
      }
    )
  };

  const HandleDeletePost = async () => {
    await deleteData("/post/"+post_detail.id.toString(), {}, g.token).then(
      ret => {
        if(ret.code != 0) {
          throw new Error(ret.err_msg);
        } else {
          navigation.goBack();
          route.params.refresh();
        }
      }
    ).catch(
      err => {
        failToast(err+" 删除失败");
      }
    )
  };
  

  useEffect(() => {getNew(setdata)}, []);
  
  let treadtext = "";

  return (
    <View style={styles.container}>
      
      <PopUpModal 
        ifshow={ifshowdeletemodal} 
        info={"确认删除此问题?"} 
        onCancel={()=>{onChangeifshowdeletemodal(false);}} 
        onSubmit={()=>{onChangeifshowdeletemodal(false);HandleDeletePost();}}
      />
      <Text></Text>
      <View style={{ flexDirection:'row', alignSelf:'flex-end'}}>
        <Text style={styles.title} onPress={()=>{
          if(g.userid!=post_detail.poster_id){
            navigation.navigate("homepagemodal",{userid:post_detail.poster_id, username:post_detail.poster_name});
          }
        }}
        >
          {post_detail.poster_name} 的问题
        </Text>
        {post_detail.poster_id==g.userid?(
        <Pressable onPress={()=>{onChangeifshowdeletemodal(true);}} style={{marginLeft:100}}>
            {({ pressed }) => (
              <FontAwesome
                name="trash"
                size={26}
                color={"#DC143C"}
                style={{ marginRight: 20, opacity: pressed ? 0.5 : 1 , marginTop: 20}}
              />
            )}
        </Pressable>  
        ):(<Text style={{marginLeft:130}}></Text>)}
      </View>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <Card title={'问题:'} text={"\n"+post_detail.content}/>
      <View style={styles.separator} lightColor="white" darkColor="rgba(255,255,255,0.1)" />

      
      {g.userid==post_detail.poster_id?(
        <Pressable
          style={{ ...styles.openButton, backgroundColor: Colors.light.tint }}
          onPress={() => {
            navigation.navigate("threadaddmodal",{type:1, post_id:post_detail.id, refresh:()=>{getNew(setdata)}});
          }}
        >
          <Text style={styles.textStyle}>追问</Text>
        </Pressable>
      ):(
        <></>
      )}

      {g.userid==route.params.message_box_owner_id?(
        <Pressable
          style={{ ...styles.openButton, backgroundColor: Colors.light.tint }}
          onPress={() => {
            navigation.navigate("threadaddmodal",{type:2, post_id:post_detail.id, refresh:()=>{getNew(setdata)}});
          }}
        >
          <Text style={styles.textStyle}>回答</Text>
        </Pressable>
      ):(
        <></>
      )}

      <FlatList 
        style={styles.flat}
        data={data}
        renderItem={({ item }) =>
          <Card 
            title={item.type==1?"追问":"回答"} 
            text={"\n"+item.content} 
            onPress={() => {}}
          />
        }
        refreshing={false}
        keyExtractor={(item) => item.id.toString()}
        onRefresh={() => {getNew(setdata)}}
        onEndReachedThreshold={0.01}
        ListFooterComponent={<FlatListTail/>}
      />
    </View>
  );
}

function failToast(msg: string) {
  Alert.alert(msg);
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
  },
  title: {
    alignSelf: 'center',
    marginTop: 15,
    fontSize: 20,
    marginRight: 8,
  },
  name: {
    justifyContent:'center',
    textAlign: 'left',
    marginLeft: '10%',
    marginTop: '5%',
    fontSize: 20,
  },
  button: {
    justifyContent: 'center',
    marginLeft: '50%',
    marginTop: 'auto',
  },
  separator: {
    alignSelf: 'center',
    marginTop: 15,
    height: 1,
    width: '80%',
  },
  flat: {
    height: '100%',
    zIndex:-1,
  },
  openButton: {
    alignSelf:'center',
    backgroundColor: Colors.light.tint,
    borderRadius: 17,
    padding: 15,
    top:700,
    width:200,
    position:'absolute',
    shadowOffset: {width: 0, height: 3}, // 阴影偏移量
    shadowRadius: 6, // 阴影模糊半径
    shadowOpacity: 0.5, // 阴影不透明度
    shadowColor: '#5c5c5c', // 设置阴影色
  },
  textStyle: {
    color: "white",
    fontWeight: "bold",
    textAlign: "center",
    fontSize:17,
  },
  centeredView: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    marginTop: 0,
  },
  modalView: {
    margin: 20,
    backgroundColor: "white",
    borderRadius: 20,
    paddingBottom: 25,
    paddingTop: 30,
    padding: 40,
    alignItems: "center",
    shadowColor: "black",
    shadowOffset: {
        width: 0,
        height: 3,
    },
    shadowOpacity: 0.2,
    shadowRadius: 6,
    elevation: 5,
},
});
