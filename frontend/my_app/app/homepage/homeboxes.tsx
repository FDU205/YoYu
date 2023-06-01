import { Alert, FlatList, StyleSheet } from 'react-native';
import { Text, View } from '../../components/Themed';
import Card from '../../components/Card';
import { getData } from '../../components/Api';
import { Dispatch, SetStateAction, useEffect, useState } from 'react';
import type { messageBoxinfo, postinfo } from '../../constants/DataType';
import { NavigationParamList, Props } from '../../constants/NavigationType';
import type { NativeStackNavigationProp } from '@react-navigation/native-stack';
import g from '../globaldata';
import { TextInput } from 'react-native-gesture-handler';
import FontAwesome from '@expo/vector-icons/FontAwesome';
import Colors from '../../constants/Colors';
import { TouchableOpacity } from 'react-native';
import FlatListTail from '../../components/FlatListTail';

const PAGE_SIZE = 10;
let searchtext = "";
let presearchtext = "";
let chosenbox: messageBoxinfo = {
  id: 0,
  owner_id: 0,
  title: '',
  owner_name: ''
};

export default function HomeBoxesScreen({ route, navigation }: Props<'homeboxes'>) {
  const [data, setdata] = useState(new Array<messageBoxinfo>(0));
  const [postsdata, setpostsdata] = useState(new Array<postinfo>(0));
  const [isinbox, setisinbox] = useState(false);
  let page_num = 2;
  let padismine = "&owner="+route.params.owner_id.toString();
  const getNextBoxes = () => {
    getData("/messageBoxes?page_num="+page_num.toString()+"&page_size="+PAGE_SIZE.toString()+padismine,g.token).then(
    ret => {
        if(ret.code != 0) {
            throw new Error(ret.err_msg);
        } else {
            let temp = ret.data.messageBoxes;
            data.push.apply(data, temp);
        }
    }
    ).then(()=>{
        page_num++;
        }).catch(
        err => {
            failToast(err+" 刷新失败");
        }
    )
  };
  const getNewBoxes = (setdata: { (value: SetStateAction<messageBoxinfo[]>): void; (arg0: any): void; }) => {
    page_num = 1;
    getData("/messageBoxes?page_num=1&page_size="+PAGE_SIZE.toString()+padismine,g.token).then(
    ret => {
        if(ret.code != 0) {
            throw new Error(ret.err_msg);
        } else {
            setdata(ret.data.messageBoxes);
        }
    }
    ).catch(
        err => {
            failToast(err+" 刷新失败");
        }
    )
  };

  const getNextPosts = () => {
    getData("/posts?page_num="+page_num.toString()+"&page_size="+PAGE_SIZE.toString()+"&message_box_id="+ chosenbox.id.toString(),g.token).then(
      ret => {
        if(ret.code != 0) {
          throw new Error(ret.err_msg);
        } else {
          let temp = ret.data.posts;
          data.push.apply(data, temp);
        }
      }
    ).then(()=>{
      page_num++;
    }).catch(
      err => {
        failToast(err+" 刷新失败");
      }
    )
  };
  const getNewPosts = (setdata: { (value: SetStateAction<postinfo[]>): void; (arg0: any): void; }) => {
    page_num = 1;
    getData("/posts?page_num=1&page_size="+PAGE_SIZE.toString()+"&message_box_id="+ chosenbox.id.toString(),g.token).then(
      ret => {
        if(ret.code != 0) {
          throw new Error(ret.err_msg);
        } else {
          setdata(ret.data.posts);
        }
      }
    ).catch(
      err => {
        failToast(err+" 刷新失败");
      }
    )
  };

  const onPress = (box: messageBoxinfo, setisinbox: Dispatch<SetStateAction<boolean>>) => {
    chosenbox.id = box.id;
    chosenbox.owner_id = box.owner_id;
    chosenbox.owner_name = box.owner_name;
    chosenbox.title = box.title;
    setisinbox(true);
    getNewPosts(setpostsdata);
    return;
  };

  useEffect(() => {isinbox? getNewPosts(setpostsdata): getNewBoxes(setdata)}, []);

  return (
    isinbox?(
      <View style={styles.container}>
        <TouchableOpacity style={styles.back} onPress={() => {setisinbox(false);getNewBoxes(setdata);}}>
          <FontAwesome 
            name='chevron-left' 
            color={Colors.light.tint} size={18} 
            style={{ marginTop: 3 , marginBottom: 0, marginRight: 10 }}
          />
          <Text 
            style={styles.boxheader}
          >
            {chosenbox.owner_name} 的提问箱
          </Text>
        </TouchableOpacity>
        <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
        <Card 
          title={chosenbox.title} 
          text={chosenbox.owner_id==g.userid?"\n点击修改或删除提问箱":"\n点击向TA提问"} 
          onPress={() => {
            if(chosenbox.owner_id == g.userid) {
              navigation.navigate("boxmodifymodal",
              {
                box:chosenbox, 
                changebox:(t)=>{chosenbox=t;getNewPosts(setpostsdata);}, 
                setinbox:(t)=>{setisinbox(t);getNewBoxes(setdata);},
              });
            } else {
              navigation.navigate("boxaskmodal", {box: chosenbox, onSubmit:()=>{getNewPosts(setpostsdata);}});
            }
            }
          }
        />
        
        <View style={styles.gap} lightColor="transparent" darkColor="rgba(255,255,255,0.1)" />
        <FlatList 
          style={styles.flat}
          data={postsdata}
          renderItem={({ item }) => 
            <Card 
              title={"#"+item.id.toString()} 
              text={"\n"+item.content}
              onPress={() => {navigation.navigate("postmodal",{postinfo:item, refresh:()=>{getNewPosts(setpostsdata)}, message_box_owner_id:chosenbox.owner_id})}}
            />
          }
          refreshing={false}
          keyExtractor={(item) => item.id.toString()}
          onRefresh={() => {getNewPosts(setpostsdata)}}
          onEndReachedThreshold={0.01}
          onEndReached={() =>{getNextPosts()}}
          ListFooterComponent={<FlatListTail/>}
        />
      </View>
    ):(
      <View style={styles.container}>
        
        <FlatList 
          style={styles.flat}
          data={data}
          renderItem={({ item }) => 
            <Card 
              title={item.title} 
              text={"\n\n"+item.owner_name + " 的提问箱→"} 
              onPress={() => {onPress(item, setisinbox)}}
            />
          }
          refreshing={false}
          keyExtractor={(item) => item.id.toString()}
          onRefresh={() => {getNewBoxes(setdata)}}
          onEndReachedThreshold={0.01}
          onEndReached={() =>{getNextBoxes()}}
          ListFooterComponent={<FlatListTail/>}
        />
      </View>
    )
    
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
  },
  separator: {
    alignSelf: 'center',
    marginTop: 0,
    height: 1,
    width: '80%',
  },
  gap: {
    alignSelf: 'center',
    marginTop: 0,
    height: 13,
    width: '80%',
  },
  flat: {
    height: '100%'
  },
  search: {
    borderStyle: 'solid',
    borderColor: '#D3D3D3',
    borderWidth: 1,
    borderRadius: 25,
    padding: 15,
    marginHorizontal: 15,
    marginVertical: 20,
    flexDirection: 'row',
  },
  searchicon: {
    flex: 1,
  },
  searchinput: {
    flex: 8,
    fontSize: 15,
    padding: 0,
  },
  boxheader: {
    flex: 8,
    fontSize: 18,
    padding: 0,
  },
  back: {
    padding: 15,
    marginHorizontal: 15,
    marginVertical: 20,
    flexDirection: 'row',
  },
});
