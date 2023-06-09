import { Text, View } from '../components/Themed';
import { StyleSheet } from 'react-native';

export default function FlatListTail() {
    return(
        <View style={styles.container}>
          <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
          <Text style={{alignSelf:'center',padding:10,color:"#C0C0C0"}}>总得有底吧~</Text>
          <View style={styles.gap} lightColor="transparent" darkColor="rgba(255,255,255,0.1)" />
        </View>
    )
}

const styles = StyleSheet.create({
  container: {
    flex:1,
    marginTop: 15,
  },
  separator: {
    alignSelf: 'center',
    marginTop: 15,
    height: 1,
    width: '80%',
  },
  gap: {
    alignSelf: 'center',
    marginTop: 15,
    height: 30,
    width: '80%',
  },
}
)