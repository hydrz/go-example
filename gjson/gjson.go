package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/tidwall/gjson"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/hydrz/go-example/gjson/conf"
)

// this json block is poorly formed on purpose.
var basicJSON = `{
    "log": {
        "level": "info"
    },
    "server": {
        "http": {
            "addr": "0.0.0.0:8000",
            "timeout": "1s"
        },
        "grpc": {
            "addr": "0.0.0.0:9000",
            "timeout": "1s"
        }
    },
    "data": {
        "database": {
            "driver": "mysql",
            "source": "root:root@tcp(mysql:3306)/test"
        },
        "redis": {
            "addr": "mysql:6379",
            "read_timeout": "0.2s",
            "write_timeout": "0.2s"
        }
    }
}`

func main() {
	b := gjson.Parse(basicJSON).Value().(map[string]interface{})

	var conf conf.Conf

	ms, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           &conf,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			ProtoHook(),
		),
	})

	err := ms.Decode(b)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(conf)

}

func ProtoHook() mapstructure.DecodeHookFunc {
	return func(from reflect.Value, to reflect.Value) (interface{}, error) {
		if from.Kind() != reflect.String {
			return from.Interface(), nil
		}

		if m, ok := to.Interface().(proto.Message); ok {
			mi := m.ProtoReflect().New().Interface()

			b := []byte(from.String())

			protojson.Unmarshal(b, mi)

			return mi, nil
		}

		if to.Kind() == reflect.TypeOf(durationpb.Duration{}).Kind() {

			d, err := time.ParseDuration(from.String())
			if err != nil {
				return nil, err
			}

			return durationpb.New(d), nil
		}

		return from.Interface(), nil
	}
}
