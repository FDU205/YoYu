import { StatusBar } from 'expo-status-bar';
import { Platform, Pressable, StyleSheet, useColorScheme, Switch, Alert } from 'react-native';

import { Text, View } from '../components/Themed';
import React, { useState } from 'react';
import LongTextBox from '../components/LongTextBox';
import FontAwesome from '@expo/vector-icons/FontAwesome';
import Colors from '../constants/Colors';
import Card from '../components/Card';
import {Keyboard} from 'react-native';
import { postData } from '../components/Api';
import { NavigationParamList, Props } from '../constants/NavigationType';
import { NativeStackNavigationProp } from '@react-navigation/native-stack';
import g from './globaldata';

export default function BoxAddModalScreen({ route, navigation }: Props<'boxaddmodal'>) {
  const colorScheme = useColorScheme();
  const [text, onChangeText] = useState('');

  return (
    <View style={styles.container}>
      <View style={{ flexDirection:'row', alignSelf:'flex-end'}}>
        <Text style={styles.title}>
          发布你的提问箱~
        </Text>
        <Pressable onPress={()=>{HandleCreateBox(text, navigation)}} style={{marginLeft:100}}>
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
      <LongTextBox placeholder='提问箱标题(最多50字)' onSubmit={()=>{}} onChangeOutterText={onChangeText}/>
      <Text style={{
        padding: 0,
        color: 'gray',
        alignSelf:'flex-end',
        fontSize: 15,
        marginRight: 30,
      }}>
          {text.length} / 50
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

function failToast(msg: string) {
  Alert.alert(msg);
}

const HandleCreateBox = async (content: string, navigation: NativeStackNavigationProp<NavigationParamList, "boxaddmodal", undefined>) => {
  if(content.length < 1) {
    failToast("提问箱标题不能为空！");
    return;
  }
  postData("/messageBox", {"title" : content}, g.token).then(
    ret => {
      if(ret.code != 0) {
        throw new Error(ret.err_msg);
      } else {
        navigation.goBack();
      }
    }
  ).catch(
    err => {
      failToast(err+" 创建失败");
    }
  )
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
