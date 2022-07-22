package golukeprot

import (
	"errors"
	"fmt"

	"github.com/Pinablink/goluke/golukemsg"
	"github.com/Pinablink/goluke/golukeutil"
)

//
func okStreamMagic(ref []byte) ([]byte, error) {

	var byte0Available byte = ref[STREAM_MAGIC[INI_RANGE]]
	var byte1Available byte = ref[STREAM_MAGIC[FIN_RANGE]]

	if byte0Available == 0xAC && byte1Available == 0xED {
		return []byte{byte0Available, byte1Available}, nil
	}

	return nil, errors.New(golukemsg.MSG_ERROR_STREAM_MAGIC_OB_JAVA)
}

//
func okStreamVersion(ref []byte) ([]byte, error) {

	// DEVERÁ SER OBSERVADO
	var byte2Available byte = ref[STREAM_VERSION[INI_RANGE]]
	var byte3Available byte = ref[STREAM_VERSION[FIN_RANGE]]

	return []byte{byte2Available, byte3Available}, nil
}

//
func okTcObject(ref []byte) (byte, error) {

	var byte4Available byte = ref[TC_OBJECT[SINGLE]]

	if byte4Available == 0x73 {
		return byte4Available, nil
	}

	return 0xFF, nil
}

//
func okTcClassDesc(ref []byte) (byte, error) {

	var byte5Available byte = ref[TC_CLASSDESC[SINGLE]]

	if byte5Available == 0x72 {
		return byte5Available, nil
	}

	return 0xFF, nil
}

//
func okLenClassName(ref []byte) ([]byte, error) {

	var bRet []byte
	var byte6Available byte = ref[LEN_CLASS_NAME[INI_RANGE]]
	var byte7Available byte = ref[LEN_CLASS_NAME[FIN_RANGE]]

	valNum := golukeutil.Get16BitValNum(byte6Available, byte7Available)

	if valNum > 0 {
		bRet = make([]byte, 2)
		bRet[0] = byte6Available
		bRet[1] = byte7Available

		return bRet, nil
	}

	return nil, errors.New(golukemsg.MSG_ERROR_LEN_NAME_CLASS_OB_JAVA)
}

//
func pClassName(ref []byte, lenName uint) ([]byte, error) {
	var fin = cursor + lenName
	var byteName []byte = ref[cursor:fin]
	cursor = fin
	return byteName, nil
}

//
func pProtocolUUID(ref []byte) ([]byte, error) {
	var fin = cursor + BYTE_SERIAL_VERSION_UUID
	var byteSerial []byte = ref[cursor:fin]
	cursor = fin

	return byteSerial, nil
}

//
func okSuportSerialization(ref []byte) (byte, error) {

	var vSuport byte = ref[cursor]
	cursor = cursor + 1

	if vSuport == 0x02 {
		return vSuport, nil
	}

	return 0xFF, errors.New(golukemsg.MSG_ERROR_STREAM_NOT_SUPPORT_SERIAL)
}

//
func okNumField(ref []byte) ([]byte, error) {

	var byte8Available byte
	var byte9Available byte
	var numFields uint

	byte8Available = ref[cursor]
	cursor = cursor + 1
	byte9Available = ref[cursor]

	numFields = uint(golukeutil.Get16BitValNum(byte8Available, byte9Available))

	fmt.Printf("%x\n", byte8Available)
	fmt.Printf("%x\n", byte9Available)

	if numFields > 0 {
		var retByte []byte = make([]byte, 2)
		retByte[0] = byte8Available
		retByte[1] = byte9Available

		return retByte, nil
	}

	return nil, errors.New(golukemsg.MSG_ERROR_STREAM_OBJ_NOT_FIELD)
}
