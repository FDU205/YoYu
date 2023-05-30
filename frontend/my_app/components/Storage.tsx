import Storage from 'react-native-storage';
import AsyncStorage from '@react-native-async-storage/async-storage';
 
const storage = new Storage({
    // maximum capacity, default 1000 key-ids
    size: 1000,

    // Use AsyncStorage for RN apps, or window.localStorage for web apps.
    // If storageBackend is not set, data will be lost after reload.
    storageBackend: AsyncStorage, // for web: window.localStorage

    // expire time, default: 1 day (1000 * 3600 * 24 milliseconds).
    // can be null, which means never expire.
    defaultExpires: 1000 * 3600 * 24,

    // cache data in the memory. default is true.
    enableCache: true,
});

const _storage = {

    // 使用key来保存数据。这些数据一般是全局独有的，常常需要调用的。
    // 除非你手动移除，这些数据会被永久保存，而且默认不会过期。
    save(key: any, obj: any) {
      storage.save({
        key: key,  // 注意: 请不要在key中使用_下划线符号!
        data: obj,
        // 如果不指定过期时间，则会使用defaultExpires参数
        // 如果设为null，则永不过期
        expires: 1000 * 3600 * 24
      })
    },
  
    // 取数据
    load(key: any, callBack: (arg0: any) => any) {
      storage.load({
        key: key,
        // autoSync(默认为true)意味着在没有找到数据或数据过期时自动调用相应的sync方法
        autoSync: false,
        // syncInBackground(默认为true)意味着如果数据过期，
        // 在调用sync方法的同时先返回已经过期的数据。
        // 设置为false的话，则始终强制返回sync方法提供的最新数据(当然会需要更多等待时间)。
        syncInBackground: true,
        // 你还可以给sync方法传递额外的参数
        syncParams: {
          extraFetchOptions: { // 各种参数
          },
          someFlag: true,
        }
      }).then(ret => {
        // 如果找到数据，则在then方法中返回
        // 注意：这是异步返回的结果（不了解异步请自行搜索学习）
        // 你只能在then这个方法内继续处理ret数据
        // 而不能在then以外处理
        callBack(ret)
        //return ret
      }).catch(err => {
        //如果没有找到数据且没有sync方法，
        //或者有其他异常，则在catch中返回
        //console.warn(err.message);
        switch (err.name) {
          case 'NotFoundError':
            // TODO
            break
          case 'ExpiredError':
            // TODO
            break
        }
        callBack(null)
      })
    },
  
    // 获取某个key下的所有id(仅key-id数据)
    getIdsForKey(id: string, callback: (arg0: string[]) => any) {
      storage.getIdsForKey(id).then(ids => {
        callback(ids)
      })
    },
  
    // 获取某个key下的所有数据(仅key-id数据)
    getAllDataForKey(key: string, callback: (arg0: any[]) => any) {
      storage.getAllDataForKey(key).then(users => {
        callback(users)
      })
    },
  
    // !! 清除某个key下的所有数据(仅key-id数据)
    clearMapForKey(key: string) {
      storage.clearMapForKey(key)
    },
  
    // 删除单个数据
    remove(key: any) {
      storage.remove({
        key: key
      })
    },
  
    // !! 清空map，移除所有"key-id"数据（但会保留只有key的数据）
    clearMap() {
      storage.clearMap()
    }
}
  
export {_storage as storage}
