package test

import (
	"encoding/hex"
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	v1 "helloworld/api/helloworld/v1"
	"testing"
)

func TestMarshal(t *testing.T) {
	reply := &v1.HelloReply{}
	a, err := anypb.New(&v1.Data{
		FieldA: "A",
		FieldB: "B",
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(a)

	reply.Any = a
	marshal, err := proto.Marshal(reply)
	if err != nil {
		t.Error(err)
	}

	bytes, err := protojson.Marshal(reply)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(bytes))
	t.Log(string(marshal))
	t.Log(hex.EncodeToString(marshal))
}

func TestAnyPbUnmarshal(t *testing.T) {
	decodeString, err := hex.DecodeString("12300a26747970652e676f6f676c65617069732e636f6d2f68656c6c6f776f726c642e76312e4461746112060a0141120142")
	if err != nil {
		t.Error(err)
	}

	bytes := decodeString
	reply := &v1.HelloReply{}
	err = proto.Unmarshal(bytes, reply)
	if err != nil {
		t.Error(err)
	}

	unmarshalNew, err := reply.Any.UnmarshalNew()
	if err != nil {
		t.Error(err)
	}

	t.Log(unmarshalNew.(*v1.Data))

	t.Logf("%T %v", unmarshalNew, unmarshalNew)
}

func TestOneOf(t *testing.T) {
	v := &v1.DataOneOf{
		Data: &v1.DataOneOf_A{A: 1},
	}

	marshal, err := json.Marshal(v)
	t.Logf("%s %v", marshal, err)

	marshal, err = protojson.Marshal(v)
	t.Logf("%s %v", marshal, err)

}
