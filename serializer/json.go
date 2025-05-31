package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtobuftoJSON converts a protobuf buffer message to a JSON string.
// func ProtobuftoJSON(message proto.Message) (string, error) {
// 	marshaler := jsonpb.Marshaler{
// 		EnumsAsInts:  false,
// 		EmitDefaults: true,
// 		Indent:       "  ",
// 		OrigName:     true,
// 	}

// 	return marshaler.MarshalToString(message)
// }

func ProtobuftoJSON(message proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		Multiline:       true,
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}

	jsonBytes, err := marshaler.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
