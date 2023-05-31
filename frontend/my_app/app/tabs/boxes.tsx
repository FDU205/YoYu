import { Alert, FlatList, StyleSheet } from 'react-native';
import { Text, View } from '../../components/Themed';
import Card from '../../components/Card';
import { getData } from '../../components/Api';
import { SetStateAction, useEffect, useState } from 'react';
import type { messageBoxinfo } from '../../constants/DataType';
import { NavigationParamList, Props } from '../../constants/NavigationType';
import type { NativeStackNavigationProp } from '@react-navigation/native-stack';
import g from '../globaldata';
import { TextInput } from 'react-native-gesture-handler';
import FontAwesome from '@expo/vector-icons/FontAwesome';

const onPress = (userid: number, username: string, navigation: NativeStackNavigationProp<NavigationParamList, "boxes", undefined>) => {
  if(userid == g.userid) {
    navigation.navigate(
      'homepage',{userid: userid, username: username},
    );
  } else {
    navigation.navigate(
      'homepagemodal',{userid: userid, username: username},
    );
  }
  return;
};

const PAGE_SIZE = 10;
let searchtext = "";
let presearchtext = "";
let chosenbox: messageBoxinfo = {
  id: 0,
  owner_id: 0,
  title: '',
  owner_name: ''
};

export default function BoxesScreen({ route, navigation }: Props<'boxes'>) {
  const [data, setdata] = useState(new Array<messageBoxinfo>(0));
  const [isinbox, setisinbox] = useState(false);
  let page_num = 2;
  const getNextBoxes = () => {
    if(presearchtext.length != 0) {
      HandleNextSearch(presearchtext);
      page_num++;
    } else {
      getData("/messageBoxes?page_num="+page_num.toString()+"&page_size="+PAGE_SIZE.toString(),g.token).then(
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
    }
  }
  const getNewBoxes = (setdata: { (value: SetStateAction<messageBoxinfo[]>): void; (arg0: any): void; }) => {
    page_num = 1;
    if(presearchtext.length != 0) {
      HandleNewSearch(presearchtext);
    } else {
      getData("/messageBoxes?page_num=1&page_size="+PAGE_SIZE.toString(),g.token).then(
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
    }
  };
  const HandleNewSearch = (searchtext: string) => {
    getData("/messageBoxes?page_num="+page_num.toString()+"&page_size="+PAGE_SIZE.toString()+"&title="+searchtext,g.token).then(
      ret => {
        if(ret.code != 0) {
          throw new Error(ret.err_msg);
        } else {
          setdata(ret.data.messageBoxes);
        }
      }
    ).then(()=>{
      page_num++;
    }).catch(
      err => {
        failToast(err+" 搜索失败");
      }
    )
  };
  const HandleNextSearch = (searchtext: string) => {
    getData("/messageBoxes?page_num="+page_num.toString()+"&page_size="+PAGE_SIZE.toString()+"&title="+searchtext,g.token).then(
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
        failToast(err+" 搜索失败");
      }
    )
  };
  useEffect(() => {getNewBoxes(setdata)}, []);

  return (
    isinbox?(
      <></>
    ):(
      <View style={styles.container}>
        <View style={styles.search}>
          <FontAwesome name='search' color={"#D3D3D3"} size={28} style={{ marginTop: -3 , marginBottom: -3, marginRight: 10 }} />
          <TextInput 
            style={styles.searchinput} 
            onSubmitEditing={()=>{
              page_num = 1;
              HandleNewSearch(searchtext);
              presearchtext = searchtext.slice(0);
            }} 
            onChangeText={text => {searchtext=text;}}
          />
        </View>
        <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
        <FlatList 
          style={styles.flat}
          data={data}
          renderItem={({ item }) => 
            <Card title={item.title} text={"\n\n"+item.owner_name + " 的提问箱↑"} onPress={() => {onPress(item.owner_id, item.owner_name, navigation)}}/>
          }
          refreshing={false}
          keyExtractor={(item) => item.id.toString()}
          onRefresh={() => {getNewBoxes(setdata)}}
          onEndReachedThreshold={0.01}
          onEndReached={() =>{getNextBoxes()}}
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
});
