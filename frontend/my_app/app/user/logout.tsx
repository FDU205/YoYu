import React from 'react';
import { Text } from 'react-native';
import { storage } from '../../components/Storage';
import g from '../globaldata';

export default function Logout() { 
  g.userid = 0;
  g.username = "";
  g.token = "";
  storage.clearMapForKey('token');
  storage.clearMapForKey('username');
  storage.clearMapForKey('userid');
}