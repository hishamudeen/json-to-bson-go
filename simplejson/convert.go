package simplejson

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
	"github.com/sindbach/json-to-bson-go/jsonutil"
)

func Convert(jsonStr []byte) (string, error) {
	input, err := jsonutil.Unmarshal(jsonStr)
	if err != nil {
		return "", err
	}

	var fields []Code
	fields = append(fields, Id("u").Float32())
	for key, val := range input {
		elem := Id(key)
		switch val.(type) {
		case string:
			elem.Add(String())
		case bool:
			elem.Add(Bool())
		case int32:
			elem.Add(Int32())
		case int64:
			elem.Add(Int64())
		case float64:
			elem.Add(Float64())
		default:
			return "", fmt.Errorf("value for key %q has unrecognized type %T", key, val)
		}

		fields = append(fields, elem)
	}

	output := NewFile("main").Type().Id("AutoGenerated").Struct(fields...)
	return output.GoString(), nil
}
