package yaml

import (
	"fmt"
	"testing"
)

func TestDecodeOnNil(t *testing.T) {
	txt := []byte(`
a: [123, "foo"]
b:
 - x
 - y
 - 123
c: true
d: "xyz"
`)
	var x struct {
		D string `yaml:"d"`
	}
	var d interface{} = x
	err := Unmarshal(txt, &d)
	fmt.Printf("TestDecodeOnNil: d=%#v err=%v\n", d, err)
	// output:
	// panic: reflect: call of reflect.Value.Type on zero Value [recovered]
	// panic: reflect: call of reflect.Value.Type on zero Value [recovered]
	// panic: reflect: call of reflect.Value.Type on zero Value
}

func TestDecodeOnNonPointer(t *testing.T) {
	txt := []byte(`123`)
	var x int
	var d interface{} = x
	err := Unmarshal(txt, d)
	fmt.Printf("TestDecodeOnNonPointer: d=%#v err=%v\n", d, err)
	// output:
	// panic: reflect: reflect.Value.Set using unaddressable value [recovered]
	// panic: reflect: reflect.Value.Set using unaddressable value [recovered]
	// panic: reflect: reflect.Value.Set using unaddressable value
}
func TestDecodeOnNilPointer(t *testing.T) {
	txt := []byte(`123`)
	var d interface{} = (*int)(nil)
	err := Unmarshal(txt, d)
	fmt.Printf("TestDecodeOnNil: d=%#v err=%v\n", d, err)
	// output:
	// panic: reflect: reflect.Value.Set using unaddressable value [recovered]
	// panic: reflect: reflect.Value.Set using unaddressable value [recovered]
	// panic: reflect: reflect.Value.Set using unaddressable value
}
