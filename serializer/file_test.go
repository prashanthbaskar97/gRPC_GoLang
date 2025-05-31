package serializer_test

import (
	"pcbook/pb"
	"pcbook/sample"
	"pcbook/serializer"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()
	// Test writing and reading a laptop to/from a binary file and JSON file
	jsonFile := "../tmp/laptop.json"
	binaryFile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WrtieProtobuftoJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
