import { Alert, FlatList, StyleSheet } from 'react-native';
import { Text, View } from '../../components/Themed';
import Card from '../../components/Card';
import { getData } from '../../components/Api';
import { SetStateAction, useEffect, useState } from 'react';
import type { postinfo } from '../../constants/DataType';
import { NavigationParamList, Props } from '../../constants/NavigationType';
import type { NativeStackNavigationProp } from '@react-navigation/native-stack';
import g from '../globaldata';
import FlatListTail from '../../components/FlatListTail';

const PAGE_SIZE = 10;

export default function HomePostsScreen({ route, navigation }: Props<'homeposts'>) {
  const onPress = () => {

  };

  const [data, setdata] = useState(new Array<postinfo>(0));
  let page_num = 2;
  const getNextPosts = () => {
    getData("/mypost?page_num="+page_num.toString()+"&page_size="+PAGE_SIZE.toString(),g.token).then(
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
  }
  const getNewPosts = (setdata: { (value: SetStateAction<postinfo[]>): void; (arg0: any): void; }) => {
    getData("/mypost?page_num=1&page_size="+PAGE_SIZE.toString(),g.token).then(
      ret => {
        if(ret.code != 0) {
          throw new Error(ret.err_msg);
        } else {
          setdata(ret.data.posts);
          page_num = 2;
        }
      }
    ).catch(
      err => {
        failToast(err+" 刷新失败");
      }
    )
  };
  useEffect(() => {getNewPosts(setdata)}, []);
  
  return (
    <View style={styles.container}>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <FlatList 
        style={styles.flat}
        data={data}
        renderItem={({ item }) => 
        <Card 
            title={"#"+item.id.toString()+(item.visibility==2?(" (匿名)"):(""))} 
            text={"\n"+item.content}
            onPress={() => {navigation.navigate("postmodal",{postinfo:item, refresh:()=>{getNewPosts(setdata)}, message_box_owner_id:-1})}}
        />
        }
        refreshing={false}
        keyExtractor={(item) => item.id.toString()}
        onRefresh={() => {getNewPosts(setdata)}}
        onEndReachedThreshold={0.01}
        onEndReached={() =>{getNextPosts()}}
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
