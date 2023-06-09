import { StatusBar } from 'expo-status-bar';
import { Platform, Pressable, StyleSheet, useColorScheme, Alert } from 'react-native';

import { Text, View } from '../components/Themed';
import React, { useEffect, useState } from 'react';
import LongTextBox from '../components/LongTextBox';
import FontAwesome from '@expo/vector-icons/FontAwesome';
import Colors from '../constants/Colors';
import Card from '../components/Card';
import {Keyboard} from 'react-native';
import { putData, deleteData } from '../components/Api';
import { Props } from '../constants/NavigationType';
import g from './globaldata';
import PopUpModal from '../components/PopUpModal';

function failToast(msg: string) {
  Alert.alert(msg);
}

export default function BoxModifyModalScreen({ route, navigation }: Props<'boxmodifymodal'>) {
  const colorScheme = useColorScheme();
  const [text, onChangeText] = useState('');
  const [ifshowdeletemodal, onChangeifshowdeletemodal] = useState(false);

  const HandleModifyBox = async () => {
    if(pretext == text) {
      return;
    }
    if(text.length < 1) {
      failToast("提问箱标题不能为空！");
      return;
    }
    await putData("/messageBox/"+route.params.box.id.toString(), {"title" : text}, g.token).then(
      ret => {
        if(ret.code != 0) {
          throw new Error(ret.err_msg);
        } else {
          route.params.changebox({
            id:route.params.box.id,owner_id:route.params.box.owner_id,title:text,owner_name:route.params.box.owner_name
          });
          navigation.goBack();
        }
      }
    ).catch(
      err => {
        failToast(err+" 创建失败");
      }
    )
  };

  const HandleDeleteBox = async () => {
    await deleteData("/messageBox/"+route.params.box.id.toString(), {}, g.token).then(
      ret => {
        if(ret.code != 0) {
          throw new Error(ret.err_msg);
        } else {
          route.params.changebox({
            id:0,owner_id:0,title:"",owner_name:"",
          });
          navigation.goBack();
          route.params.setinbox(false);
        }
      }
    ).catch(
      err => {
        failToast(err+" 删除失败");
      }
    )
  };

  useEffect(()=>{onChangeText(route.params.box.title)},[]);
  let pretext = route.params.box.title;
  return (
    <View style={styles.container}>
      <PopUpModal 
        ifshow={ifshowdeletemodal} 
        info={"确认删除提问箱?"} 
        onCancel={()=>{onChangeifshowdeletemodal(false);}} 
        onSubmit={()=>{onChangeifshowdeletemodal(false);HandleDeleteBox();}}
      />
      

      <View style={{ flexDirection:'row', alignSelf:'flex-end'}}>
        <Text style={styles.title}>
          修改你的提问箱
        </Text>
        <Pressable onPress={()=>{HandleModifyBox()}} style={{marginLeft:100}}>
            {({ pressed }) => (
              <FontAwesome
                name="paper-plane"
                size={25}
                color={Colors[colorScheme ?? 'light'].text}
                style={{ marginRight: 15, opacity: pressed ? 0.5 : 1 , marginTop: 30}}
              />
            )}
        </Pressable>
      </View>
      
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <LongTextBox 
        placeholder='' 
        onSubmit={()=>{}} 
        onChangeOutterText={onChangeText}
        defaulttext={text}
      />
      <Text style={{
        padding: 0,
        color: 'gray',
        alignSelf:'flex-end',
        fontSize: 15,
        marginRight: 30,
      }}>
          {text.length} / 50
      </Text>

      <Text style={{
        padding: 0,
        color: 'red',
        alignSelf:'flex-end',
        fontSize: 15,
        marginRight: 30,
        marginTop:5,
      }}
        onPress={()=>{onChangeifshowdeletemodal(true);}}
      >
          删除提问箱
      </Text>
      
      <Text style={styles.text}>
          预览：
      </Text>
      <Card title={text} text={"\n\n"+g.username+" 的提问箱→"} onPress={() => {Keyboard.dismiss()}}/>
      {/* Use a light status bar on iOS to account for the black space above the modal */}
      <StatusBar style={Platform.OS === 'ios' ? 'light' : 'auto'} />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  text: {
    padding: 10,
    alignSelf:'flex-start',
    fontSize: 20,
    borderLeftWidth:20,
  },
  title: {
    alignSelf:'center',
    marginTop: 30,
    fontSize: 20,
    fontWeight: 'bold',
  },
  separator: {
    alignSelf:'center',
    marginVertical: 30,
    height: 1,
    width: '80%',
  },
  switch: {
    padding: 10,
    alignSelf: 'flex-end',
    marginRight: 28,
  }
});
