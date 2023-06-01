import { StatusBar } from 'expo-status-bar';
import { Platform, Pressable, StyleSheet, useColorScheme, Alert, Switch } from 'react-native';

import { Text, View } from '../components/Themed';
import React, { useState } from 'react';
import LongTextBox from '../components/LongTextBox';
import FontAwesome from '@expo/vector-icons/FontAwesome';
import Colors from '../constants/Colors';
import Card from '../components/Card';
import {Keyboard} from 'react-native';
import { postData } from '../components/Api';
import { Props } from '../constants/NavigationType';
import g from './globaldata';

function failToast(msg: string) {
  Alert.alert(msg);
}

export default function ThreadAddModalScreen({ route, navigation }: Props<'threadaddmodal'>) {
  const colorScheme = useColorScheme();
  const [text, onChangeText] = useState('');

  const HandleAddAsk = async () => {
    if(text.length < 1) {
      failToast((route.params.type==1?("追问"):("回答"))+"不能为空！");
      return;
    }
    await postData("/post/channel", {"post_id":route.params.post_id,"content":text,"type":route.params.type}, g.token).then(
      ret => {
        if(ret.code != 0) {
          throw new Error(ret.err_msg);
        } else {
          navigation.goBack();
          route.params.refresh();
        }
      }
    ).catch(
      err => {
        failToast(err+" "+(route.params.type==1?("追问"):("回答"))+"失败");
      }
    )
  };

  return (
    <View style={styles.container}>

      <View style={{ flexDirection:'row', alignSelf:'flex-end'}}>
        <Text style={styles.title}>
          向TA{route.params.type==1?("追问"):("回答")}
        </Text>
        <Pressable onPress={()=>{HandleAddAsk()}} style={{marginLeft:100}}>
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
        placeholder={"友善"+(route.params.type==1?("追问"):("回答"))+"~" }
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
          {text.length} / 200
      </Text>

      <Text style={styles.text}>
          预览：
      </Text>
      <Card title={(route.params.type==1?("追问"):("回答"))} text={"\n"+text} onPress={() => {Keyboard.dismiss()}}/>
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
    marginRight:20,
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
