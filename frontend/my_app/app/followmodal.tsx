import { Alert, Button, FlatList, StyleSheet } from 'react-native';
import { Text, View } from '../components/Themed';
import Card from '../components/Card';
import { deleteData, getData } from '../components/Api';
import { SetStateAction, useEffect, useState } from 'react';
import type { follows } from '../constants/DataType';
import { NavigationParamList, Props } from '../constants/NavigationType';
import type { NativeStackNavigationProp } from '@react-navigation/native-stack';
import g from './globaldata';
import FlatListTail from '../components/FlatListTail';

const onPress = (userid: number, username: string, navigation: NativeStackNavigationProp<NavigationParamList, "followmodal", undefined>) => {
  navigation.navigate(
    'homepagemodal',{userid: userid, username: username},
  );
  return;
};

// 取消关注
const cancelFollow = (followid: number) => {
  deleteData("/user/unfollow",{"follow_id":followid},g.token).then(
    ret => {
      if(ret.code != 0) {
        throw new Error(ret.err_msg);
      } else {
        Alert.alert("取消关注成功");
      }
    }
  ).catch(
    err => {
      Alert.alert(err+" 取消关注失败");
    }
  )
}

const PAGE_SIZE = 10;

export default function FollowScreen({ route, navigation }: Props<'followmodal'>) {
  const [data, setdata] = useState(new Array<follows>(0));
  let page_num = 2;
  const getNext = () => {
    getData("/user/followlist?page_num="+page_num.toString()+"&page_size="+PAGE_SIZE.toString(),g.token).then(
      ret => {
        if(ret.code != 0) {
          throw new Error(ret.err_msg);
        } else {
          let temp = ret.data.follows;
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
  const getNew = (setdata: { (value: SetStateAction<follows[]>): void; (arg0: any): void; }) => {
    getData("/user/followlist?page_num=1&page_size="+PAGE_SIZE.toString(),g.token).then(
      ret => {
        if(ret.code != 0) {
          throw new Error(ret.err_msg);
        } else {
          setdata(ret.data.follows);
        }
        page_num = 2;
      }
    ).catch(
      err => {
        failToast(err+" 刷新失败");
      }
    )
  };
  useEffect(() => {getNew(setdata)}, []);
  
  return (
    <View style={styles.container}>
      <Text></Text>
      <Text style={styles.title}>关注列表</Text>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <FlatList 
        style={styles.flat}
        data={data}
        renderItem={({ item }) =>
          <View style={{ alignContent:'center' }}>
            <View style={{flexDirection: 'row', alignContent:'center' }}>
              <Text style = {styles.name} onPress={() => {onPress(item.user_id, item.username, navigation)}}>{item.username}</Text>
              <Text style = {styles.button} onPress={() => {cancelFollow(item.user_id);getNew(setdata)}}>取消关注</Text>
            </View>
            <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
          </View>
        }
        refreshing={false}
        keyExtractor={(item) => item.user_id.toString()}
        onRefresh={() => {getNew(setdata)}}
        onEndReachedThreshold={0.01}
        onEndReached={() =>{getNext()}}
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
  },
  name: {
    justifyContent:'center',
    textAlign: 'left',
    marginLeft: '10%',
    marginTop: 'auto',
    fontSize: 20,
  },
  button: {
    color: 'red',
    justifyContent: 'center',
    marginLeft: '50%',
    marginTop: 10,

    borderStyle: 'solid',
    borderColor: 'red',    
    borderWidth: 1,
    borderRadius: 20, // 圆角
    padding: 8,
  },
  separator: {
    alignSelf: 'center',
    marginTop: 15,
    height: 1,
    width: '80%',
  },
  flat: {
    height: '100%'
  }
});
