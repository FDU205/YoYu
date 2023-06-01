import { Alert, FlatList, StyleSheet } from 'react-native';
import { Text, View } from '../../components/Themed';
import Card from '../../components/Card';
import { getData } from '../../components/Api';
import { SetStateAction, useEffect, useState } from 'react';
import type { wallpost } from '../../constants/DataType';
import { NavigationParamList, Props } from '../../constants/NavigationType';
import type { NativeStackNavigationProp } from '@react-navigation/native-stack';
import g from '../globaldata';

const onPress = (userid: number, username: string, visibility: number, navigation: NativeStackNavigationProp<NavigationParamList, "tabwall", undefined>) => {
  if(visibility == 1) {
    if(userid == g.userid) {
      navigation.navigate(
        'homepage',{userid: userid, username: username},
      );
    } else {
      navigation.navigate(
        'homepagemodal',{userid: userid, username: username},
      );
    }
  }
  return;
};

const PAGE_SIZE = 10;

export default function TabWallScreen({ route, navigation }: Props<'tabwall'>) {
  const [data, setdata] = useState(new Array<wallpost>(0));
  let page_num = 2;
  const getNextWall = () => {
    getData("/wall/mywall?page_num="+page_num.toString()+"&page_size="+PAGE_SIZE.toString(),g.token).then(
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
  const getNewWall = (setdata: { (value: SetStateAction<wallpost[]>): void; (arg0: any): void; }) => {
    getData("/wall/mywall?page_num=1&page_size="+PAGE_SIZE.toString(),g.token).then(
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
  useEffect(() => {getNewWall(setdata)}, []);
  
  return (
    <View style={styles.container}>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <FlatList 
        style={styles.flat}
        data={data}
        renderItem={({ item }) => 
          <Card title={item.poster_name} text={item.content} onPress={() => {onPress(item.poster_id, item.poster_name, item.visibility, navigation)}}/>
        }
        refreshing={false}
        keyExtractor={(item) => item.id.toString()}
        onRefresh={() => {getNewWall(setdata)}}
        onEndReachedThreshold={0.01}
        onEndReached={() =>{getNextWall()}}
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
