package serializer

import (
	"fmt"
	"io/ioutil"
	
	"google.golang.org/protobuf/proto"
)

func WrtieProtobuftoJSONFile(message proto.Message, filename string) error {
	data, err := ProtobuftoJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to JSON file: %w", err)

	}
	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON data to file: %w", err)
	}

	return nil
}

// WriteProtobufToBinaryFile serializes a protobuf message to a binary file
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary file: %w", err)

	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}

	return nil
}

// ReadProtobufFromBinaryFile deserializes a protobuf message from a binary file
func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Cannot read binary data from file : %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary data to proto message: %w", err)
	}
	return nil

}
