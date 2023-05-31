import { Alert, FlatList, StyleSheet } from 'react-native';
import { Text, View } from '../../components/Themed';
import Card from '../../components/Card';
import { getData } from '../../components/Api';
import { SetStateAction, useEffect, useState } from 'react';
import type { messageBoxinfo } from '../../constants/DataType';
import { NavigationParamList, Props } from '../../constants/NavigationType';
import type { NativeStackNavigationProp } from '@react-navigation/native-stack';
import g from '../globaldata';

const onPress = (userid: number, username: string, visibility: number, navigation: NativeStackNavigationProp<NavigationParamList, "boxes", undefined>) => {
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

export default function BoxesScreen({ route, navigation }: Props<'boxes'>) {
  const [data, setdata] = useState(new Array<messageBoxinfo>(0));
  let page_num = 2;
  const getNextBoxes = () => {
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
  const getNewBoxes = (setdata: { (value: SetStateAction<messageBoxinfo[]>): void; (arg0: any): void; }) => {
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
  };
  useEffect(() => {getNewBoxes(setdata)}, []);
  
  return (
    <View style={styles.container}>
      <FlatList 
        style={styles.flat}
        data={data}
        renderItem={({ item }) => 
          <Card title={item.owner_name + " 的提问箱:"} text={item.title} onPress={() => {onPress(item.poster_id, item.poster_name, item.visibility, navigation)}}/>
        }
        refreshing={false}
        keyExtractor={(item) => item.id.toString()}
        onRefresh={() => {getNewBoxes(setdata)}}
        onEndReachedThreshold={0.01}
        onEndReached={() =>{getNextBoxes()}}
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
