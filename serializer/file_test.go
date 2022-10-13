package serializer_test

import (
	"testing"

	"github.com/BarunBlog/Pc_book_manager/pb"
	"github.com/BarunBlog/Pc_book_manager/sample"
	"github.com/BarunBlog/Pc_book_manager/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) { // must start with test prefix
	t.Parallel()

	// we want to serializer the object to laptop.bin file
	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtoBufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
