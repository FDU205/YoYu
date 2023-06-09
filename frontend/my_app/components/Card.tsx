import { useState } from "react";
import { View, Text, StyleSheet, Pressable } from "react-native";
import { Platform } from "react-native";

interface ICardProps {
    title: string;
    text: string;
    onPress?: () => void;
}

const Card: React.FC<ICardProps> = ({
    title,
    text,
    onPress,
    ...rest
}) => {
    const [init, setinit] = useState(true);
    const [numlines, setnumlines] = useState(0);
    const [showOpenDetail, setshowOpenDetail] = useState(false);
    return(
        <Pressable 
            onPress={onPress}
            style={({ pressed }) => [
                !pressed ? styles.container : styles.containerpressed,
            ]}
            onLongPress={() => {
                if(!init) {
                    if(showOpenDetail) {
                        setnumlines(0);
                        setshowOpenDetail(false);
                    } else {
                        setnumlines(3);
                        setshowOpenDetail(true);
                    }
                }
            }}
            {...rest}
        >
            {
                Platform.OS === 'android' ? (
                <View style={styles.card} elevation={4}>
                    <Text style={styles.cardTitle}>
                        {title}
                    </Text>
                    <Text 
                        style={styles.cardText}
                        allowFontScaling={false}
                        numberOfLines={numlines}
                        onLayout={(event: { nativeEvent: { layout: { height: any; }; }; }) => {
                            if(init && !showOpenDetail) {
                                const height = Math.floor(event.nativeEvent.layout.height || 0);
                                if(height >= 90) {
                                    setnumlines(3);
                                    setshowOpenDetail(true);
                                    setinit(false);
                                }
                            }
                        }}
                    >
                        {text}
                    </Text>
                    {showOpenDetail 
                    && (<Text style={styles.opt}>长按展开</Text>)}
                </View>
                ) : (
                <View style={styles.card}>
                    <Text style={styles.cardTitle}>
                        {title}
                    </Text>
                    <Text 
                        style={styles.cardText}
                        allowFontScaling={false}
                        numberOfLines={numlines}
                        onLayout={(event: { nativeEvent: { layout: { height: any; }; }; }) => {
                            if(init && !showOpenDetail) {
                                const height = Math.floor(event.nativeEvent.layout.height || 0);
                                if(height >= 90) {
                                    setnumlines(3);
                                    setshowOpenDetail(true);
                                    setinit(false);
                                }
                            }
                        }}
                    >
                        {text}
                    </Text>
                    {showOpenDetail 
                    && (<Text style={styles.opt}>长按展开</Text>)}    
                </View>) 
            } 
        </Pressable>
    );
};

const styles = StyleSheet.create({
    container: {
        marginTop: 15,
        marginLeft: 16,
        marginRight: 16,
        opacity: 1,
    },
    containerpressed: {
        marginTop: 15,
        marginLeft: 16,
        marginRight: 16,
        opacity: 0.7,
    },
    card: {
        borderStyle: 'solid',
        borderColor: 'white',
        backgroundColor: 'rgba(255,255,255,1)',
        
        borderWidth: 2,
        shadowOffset: {width: 0, height: 3}, // 阴影偏移量
        shadowRadius: 6, // 阴影模糊半径
        shadowOpacity: 0.2, // 阴影不透明度
        borderRadius: 20, // 圆角
        shadowColor: '#5c5c5c', // 设置阴影色
        padding: 15,
    },
    cardTitle: {
        fontSize: 20,
        color: "#5c5c5c",
        alignItems: "center",
        alignContent: "center",
        justifyContent: "center",
        fontWeight: 'bold',
    },
    cardText: {
        fontSize: 20,
        color: "#5c5c5c",
        alignItems: "center",
        alignContent: "center",
        justifyContent: "center",
    },
    opt: {
        marginTop: 3,
        fontSize: 15,
        color: "#5c5c5c",
        opacity: 0.5,
        alignItems: "center",
    }
});

export default Card;