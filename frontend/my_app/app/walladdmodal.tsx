import { StatusBar } from 'expo-status-bar';
import { Platform, StyleSheet } from 'react-native';

import { Text, View } from '../components/Themed';
import React, { useState } from 'react';
import LongTextBox from '../components/LongTextBox';

export default function WallAddModalScreen() {
  const [text, onChangeText] = useState('');
  return (
    <View style={styles.container}>
      <Text style={styles.title}>发布你的表白</Text>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <LongTextBox placeholder='文明表白' onSubmit={()=>{}} onChangeOutterText={onChangeText}/>
      

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
    alignSelf:'center',
    fontSize: 20,
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
});
