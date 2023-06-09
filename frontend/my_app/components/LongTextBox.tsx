import { useState } from 'react';
import { KeyboardAvoidingView, Platform, TextInput, StyleSheet, ScrollView } from 'react-native';
import { Text } from '../components/Themed';
import Colors from '../constants/Colors';

interface ILongTextBoxProps {
    placeholder: string;
    onSubmit: () => void;
    onChangeOutterText: (text: string) => void;
    defaulttext?: string;
}

const LongTextBox: React.FC<ILongTextBoxProps> = ({
    placeholder,
    onSubmit,
    onChangeOutterText,
    defaulttext="",
}) => {
    const [text, onChangeText] = useState('');
    return(
        <ScrollView scrollEnabled={false} style={{padding:0, maxHeight:225}}>
             <KeyboardAvoidingView 
                style={styles.container}
                behavior='padding' 
                keyboardVerticalOffset={90}
            > 
                <TextInput 
                    placeholder={placeholder}
                    defaultValue={defaulttext}
                    style={styles.input} 
                    multiline={true} 
                    maxLength={200}
                    onChangeText={ text => {
                    onChangeText(text);
                    onChangeOutterText(text);
                }}
                />
            </KeyboardAvoidingView>
        </ScrollView>
    );
};

const styles = StyleSheet.create({
    container: {
        borderStyle: 'solid',
        borderColor: 'white',
        backgroundColor: 'rgba(255,255,255,1)',
        borderWidth:1,
        marginTop: 10,
        marginHorizontal:30,
        borderRadius:20,
        shadowOffset: {width: 0, height: 3}, // 阴影偏移量
        shadowRadius: 6, // 阴影模糊半径
        shadowOpacity: 0.2, // 阴影不透明度
        shadowColor: '#5c5c5c', // 设置阴影色
        padding:10,
        height:200,
        elevation:4,
    },
    input: {
        alignItems:'flex-start',
        borderWidth:0,
        borderColor:'#666',
        textAlign:'left',
        textAlignVertical:'top',
        fontSize: 20,
        padding: 0,
        minHeight: 200,
    },
});

export default LongTextBox;
