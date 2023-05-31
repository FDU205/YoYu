import { FlatList, StyleSheet } from 'react-native';

import EditScreenInfo from '../../components/EditScreenInfo';
import { Text, View } from '../../components/Themed';
import Card from '../../components/Card';

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

const onPress = () => {
  return;
};

export default function TabWallScreen() {
  return (
    <View style={styles.container}>
      <Text style={styles.title}>{getDateNow()}</Text>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <FlatList 
        style={styles.flat}
        data={TEST_DATA}
        renderItem={({ item }) => 
          <Card title={item.id} text={item.content} onPress={onPress}/>
        }
        keyExtractor={(item) => item.id}
      />
    </View>
  );
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
