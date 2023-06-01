import React from "react";
import {
  Modal,
  StyleSheet,
  Text,
  TouchableOpacity,
  View
} from "react-native";
import { Platform } from "react-native";
import Colors from '../constants/Colors';

interface IPopUpModalProps {
    ifshow: boolean;
    info: string;
    onCancel: () => void;
    onSubmit: () => void;
}

const PopUpModal: React.FC<IPopUpModalProps> = ({
    ifshow,
    info,
    onCancel,
    onSubmit,
}) => {
    return(
        <Modal
            animationType="slide"
            transparent={true}
            visible={ifshow}
            onRequestClose={() => {
                onCancel();
            }}
        >
            <View style={styles.centeredView}>
                <View style={styles.modalView}>
                    <Text style={styles.modalText}>{info}</Text>
                    <View style={{flexDirection:'row',}}>
                        <TouchableOpacity
                            style={{ ...styles.openButton, backgroundColor: Colors.light.tint }}
                            onPress={() => {
                                onCancel();
                            }}
                            >
                            <Text style={styles.textStyle}>取消</Text>
                        </TouchableOpacity>
                        <TouchableOpacity
                            style={{ ...styles.openButton, backgroundColor: Colors.light.tint }}
                            onPress={() => {
                                onSubmit();
                            }}
                            >
                            <Text style={styles.textStyle}>确认</Text>
                        </TouchableOpacity>
                    </View>
                    
                </View>
            </View>
        </Modal>
    );
};

const styles = StyleSheet.create({
    centeredView: {
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
        marginTop: 0,
    },
    modalView: {
        margin: 20,
        backgroundColor: "white",
        borderRadius: 20,
        paddingBottom: 25,
        paddingTop: 30,
        padding: 40,
        alignItems: "center",
        shadowColor: "black",
        shadowOffset: {
            width: 0,
            height: 3,
        },
        shadowOpacity: 0.2,
        shadowRadius: 6,
        elevation: 5,
    },
    openButton: {
        backgroundColor: "#F194FF",
        borderRadius: 17,
        padding: 10,
        marginHorizontal:10,
        marginTop:10,
        width:70
    },
    textStyle: {
        color: "white",
        fontWeight: "bold",
        textAlign: "center",
    },
    modalText: {
        marginBottom: 15,
        textAlign: "center",
        fontSize:18,
        fontWeight: "bold",
        color:"#696969",
    }
});

export default PopUpModal;

