package helper

import "testing"




func TestHelloWorld(t *testing.T){
	result := HelloWorld("akhdan")
	if result != "hello akhdan"{
		panic("result is not 'hello akhdan'")
	}
}