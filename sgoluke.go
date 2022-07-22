package goluke

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/Pinablink/goluke/golukemsg"
	"github.com/Pinablink/goluke/golukeprot"
	uuid "github.com/satori/go.uuid"
)

// Goluke
type Goluke struct {
	GolukeUUID       uuid.UUID
	PathJavaStream   string
	NameStruct       string
	dataStreamObJava []byte
}

//
func NewGoluke(pathJavaStream string) (*Goluke, error) {

	if len(pathJavaStream) <= 0 {
		return nil, errors.New(golukemsg.MSG_UNKNOWN_PATH_STREAM_JAVA)
	}

	dataStream, iErr := readSourceObStream(pathJavaStream)

	if iErr != nil {
		return nil, iErr
	}

	uuid, _ := uuid.NewV4()
	return &Goluke{
		GolukeUUID:       uuid,
		PathJavaStream:   pathJavaStream,
		dataStreamObJava: dataStream,
	}, nil

}

//
func (refGoluke *Goluke) ParsePlease(targetStruct interface{}) error {

	var golukeStreamByteData *golukeprot.GolukeStreamByteData = golukeprot.NewGolukeStreamByteData(refGoluke.dataStreamObJava)
	return golukeStreamByteData.ToSliceHeader()
}

//
func readSourceObStream(refPathStr string) ([]byte, error) {
	var streamDataObJava []byte
	var iError error

	streamDataObJava, iError = ioutil.ReadFile(refPathStr)

	if iError != nil {
		var strMessage string = fmt.Sprintf("%s - %s", golukemsg.MSG_ERROR_READ_PATH_STREAM_JAVA, iError)
		return nil, errors.New(strMessage)
	}

	if len(streamDataObJava) == 0 {
		return nil, errors.New(golukemsg.MSG_ERROR_LEN_BYTE_STREAM_JAVA)
	}

	return streamDataObJava, nil
}

//
func extractNameStruct(refGoluke *Goluke, target interface{}) error {
	var val reflect.Value = reflect.Indirect(reflect.ValueOf(target))
	refGoluke.NameStruct = val.Type().Name()

	return nil
}
