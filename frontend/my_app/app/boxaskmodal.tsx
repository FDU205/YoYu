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

export default function BoxAskModalScreen({ route, navigation }: Props<'boxaskmodal'>) {
  const colorScheme = useColorScheme();
  const [text, onChangeText] = useState('');
  const [visibility, onChangeVisibility] = useState(true);

  const HandleAddAsk = async () => {
    if(text.length < 1) {
      failToast("提问不能为空！");
      return;
    }
    await postData("/post", 
      {"message_box_id" : route.params.box.id, "content":text, "visibility":(visibility?1:2)}, 
      g.token).then(
      ret => {
        if(ret.code != 0) {
          throw new Error(ret.err_msg);
        } else {
          navigation.goBack();
          route.params.onSubmit();
        }
      }
    ).catch(
      err => {
        failToast(err+" 创建失败");
      }
    )
  };

  return (
    <View style={styles.container}>

      <View style={{ flexDirection:'row', alignSelf:'flex-end'}}>
        <Text style={styles.title}>
          向TA提问
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
        placeholder='请不要用提问箱功能欺负别人哦~' 
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
      <Text style={{
        padding: 10,
        color: 'gray',
        alignSelf:'flex-end',
        fontSize: 15,
        marginRight: 20,
      }}>
          {visibility?"实名":"匿名"}发表
      </Text>
      <Switch 
        style={styles.switch} 
        value={visibility} 
        onValueChange={onChangeVisibility} 
        trackColor={{true: Colors.light.tint}}
      />


      <Text style={styles.text}>
          预览：
      </Text>
      <Card title={"#"} text={"\n"+text} onPress={() => {Keyboard.dismiss()}}/>
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
