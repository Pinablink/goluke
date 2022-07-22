package golukeprot

import (
	"fmt"

	"github.com/Pinablink/goluke/golukeutil"
	uuid "github.com/satori/go.uuid"
)

//
type objJavaSerialProt [2]int8

const (
	SINGLE                   int  = 0
	INI_RANGE                int  = 0
	FIN_RANGE                int  = 1
	POS_UNMAPPED_INI         uint = 8
	BYTE_SERIAL_VERSION_UUID uint = 8
)

var STREAM_MAGIC objJavaSerialProt = objJavaSerialProt{0, 1}
var STREAM_VERSION objJavaSerialProt = objJavaSerialProt{2, 3}
var TC_OBJECT objJavaSerialProt = objJavaSerialProt{4}
var TC_CLASSDESC objJavaSerialProt = objJavaSerialProt{5}
var LEN_CLASS_NAME objJavaSerialProt = objJavaSerialProt{6, 7}
var cursor uint = POS_UNMAPPED_INI

//
type GolukeFieldIn struct {
}

//
type GolukeObIn struct {
}

//
type GolukeStreamByteData struct {
	GolukeStreamByteDataUUID uuid.UUID
	principalByteData        []byte
	ProtolStreamMagic        []byte
	ProtolStreamVersion      []byte
	ProtolTcObjectData       byte
	ProtolTcClassDesc        byte
	ProtolLenClassName       []byte
	ProtolClassName          []byte
	ProtolSerialUUID         []byte
	ProtolFlagPermSupport    byte
	ProtolNumField           []byte
	Fields                   []GolukeObIn
	Obs                      []GolukeObIn
}

//
func NewGolukeStreamByteData(byteStreamJavaOb []byte) *GolukeStreamByteData {
	uuid, _ := uuid.NewV4()
	return &GolukeStreamByteData{
		GolukeStreamByteDataUUID: uuid,
		principalByteData:        byteStreamJavaOb,
	}
}

//
func (golukeStreamByteData *GolukeStreamByteData) ToSliceHeader() error {

	byteMagic, err := okStreamMagic(golukeStreamByteData.principalByteData)

	if err != nil {
		return err
	}

	golukeStreamByteData.ProtolStreamMagic = byteMagic

	byteVersion, err1 := okStreamVersion(golukeStreamByteData.principalByteData)

	if err1 != nil {
		return err1
	}

	golukeStreamByteData.ProtolStreamVersion = byteVersion

	byteObjectData, err2 := okTcObject(golukeStreamByteData.principalByteData)

	if err2 != nil {
		return err2
	}

	golukeStreamByteData.ProtolTcObjectData = byteObjectData

	byteTcClassDesc, err3 := okTcClassDesc(golukeStreamByteData.principalByteData)

	if err3 != nil {
		return err3
	}

	golukeStreamByteData.ProtolTcClassDesc = byteTcClassDesc

	byteLenClassName, err4 := okLenClassName(golukeStreamByteData.principalByteData)

	if err4 != nil {
		return err4
	}

	golukeStreamByteData.ProtolLenClassName = byteLenClassName
	cursor = 8
	// fmt.Printf("%x\n", golukeStreamByteData)

	byteClassName, err5 := pClassName(golukeStreamByteData.principalByteData,
		uint(golukeutil.Get16BitValNum(golukeStreamByteData.ProtolLenClassName[0],
			golukeStreamByteData.ProtolLenClassName[1])))

	if err5 != nil {
		return err5
	}

	golukeStreamByteData.ProtolClassName = byteClassName

	byteSerialUUID, err6 := pProtocolUUID(golukeStreamByteData.principalByteData)

	if err6 != nil {
		return err6
	}

	golukeStreamByteData.ProtolSerialUUID = byteSerialUUID

	byteFlagSupportSerial, err7 := okSuportSerialization(golukeStreamByteData.principalByteData)

	if err7 != nil {
		return err7
	}

	golukeStreamByteData.ProtolFlagPermSupport = byteFlagSupportSerial

	byteNumField, err8 := okNumField(golukeStreamByteData.principalByteData)

	if err8 != nil {
		return err8
	}

	golukeStreamByteData.ProtolNumField = byteNumField

	testerScan(golukeStreamByteData)

	return nil
}

// RETIRAR NA FINALIZAÇÃO
func testerScan(golukeStreamByteData *GolukeStreamByteData) {

	fmt.Printf("%x\n", golukeStreamByteData.ProtolStreamMagic)
	fmt.Printf("%x\n", golukeStreamByteData.ProtolStreamVersion)
	fmt.Printf("%x\n", golukeStreamByteData.ProtolTcObjectData)
	fmt.Printf("%x\n", golukeStreamByteData.ProtolTcClassDesc)
	fmt.Printf("%x\n", golukeStreamByteData.ProtolLenClassName)
	fmt.Printf("%x\n", golukeStreamByteData.ProtolClassName)
	fmt.Printf("%x\n", golukeStreamByteData.ProtolSerialUUID)
	fmt.Printf("%x\n", golukeStreamByteData.ProtolFlagPermSupport)
	fmt.Printf("%x\n", golukeStreamByteData.ProtolNumField)
}
