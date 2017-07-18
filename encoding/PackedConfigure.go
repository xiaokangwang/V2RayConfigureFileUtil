package encoding

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/sha3"
)

func (e *Encoder) PackV2RayConfigureIntoPackedForm(FileName string) ([]byte, error) {

	//Read the file
	f, err := os.Open(FileName)
	if err != nil {
		return nil, err
	}
	ft, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return e.PackV2RayConfigureIntoPackedFormB(FileName, ft)
}

func (e *Encoder) PackV2RayConfigureIntoPackedFormB(FileName string, ft []byte) ([]byte, error) {
	PbRepre := new(LibV2RayPackedConfig)
	//Guess File Type
	if strings.HasSuffix(FileName, "LibV2RaySimpleProtoV1.pb") {
		PbRepre.ConfigType = LibV2RayPackedConfig_LibV2RaySimpleProtoV1
	} else if strings.HasSuffix(FileName, ".pb") {
		PbRepre.ConfigType = LibV2RayPackedConfig_FullProto
	} else if strings.HasSuffix(FileName, ".json") {
		PbRepre.ConfigType = LibV2RayPackedConfig_FullJsonFile
	} else {
		return nil, errors.New("Cannot Guess Config Type")
	}

	//Calc CheckSum
	result := sha3.Sum256(ft)
	PbRepre.CheckSum = result[:]
	PbRepre.Payload = ft

	out, err := proto.Marshal(PbRepre)

	if err != nil {
		return nil, err
	}
	return out, nil
}
