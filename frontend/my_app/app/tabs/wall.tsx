import { Alert, FlatList, StyleSheet } from 'react-native';
import { Text, View } from '../../components/Themed';
import Card from '../../components/Card';
import { getData } from '../../components/Api';
import { SetStateAction, useEffect, useState } from 'react';
import type { wallpost } from '../../constants/DataType';
import { NavigationParamList, Props } from '../../constants/NavigationType';
import type { NativeStackNavigationProp } from '@react-navigation/native-stack';
import g from '../globaldata';

const TEST_DATA = [
  {
    id: "1",
    poster_id: "1",
    content: "First Itemdasdsiuhdshjdfsiujhfhsiaufdghuishjdfsiujhfhsiaufdghuiasfghshjdfsiujhfhsiaufdghuiasfghshjdfsiujhfhsiaufdghuiasfghshjdfsiujhfhsiaufdghuiasfghshjdfsiujhfhsiaufdghuiasfghasfghfgsduyfgasduifhsduigfsduigfsdugfsduifsduyfsduyfgsduihgfsduighfasdghfsdu",
    visibility: 1,
  },
  {
    id: "2",
    poster_id: "2",
    content: "我好想做嘉然小姐的狗啊",
    visibility: 1,
  },
  {
    id: "3",
    poster_id: "3",
    content: "Third Item",
    visibility: 1,
  },
  {
    id: "4",
    poster_id: "4",
    content: "Four Item",
    visibility: 1,
  },
  {
    id: "5",
    poster_id: "5",
    content: "Five Item",
    visibility: 1,
  },
  {
    id: "6",
    poster_id: "6",
    content: "Six Item",
    visibility: 1,
  },
];

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
    getData("/wall?page_num="+page_num.toString()+"&page_size="+PAGE_SIZE.toString(),g.token).then(
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
    getData("/wall?page_num=1&page_size="+PAGE_SIZE.toString(),g.token).then(
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
  useEffect(() => {getNewWall(setdata)}, []);
  
  return (
    <View style={styles.container}>
      <Text style={styles.title}>{getDateNow()}</Text>
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

function getDateNow():string {
  let date = new Date();
  return date.getFullYear().toString() + "年" + (date.getMonth()+1).toString() + "月" + date.getDate() + "日";
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
