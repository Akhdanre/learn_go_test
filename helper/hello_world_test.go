package helper

import "testing"




func TestHelloWorld(t *testing.T){
	result := HelloWorld("akhdan")
	if result != "hello akhdan"{
		panic("result is not 'hello akhdan'")
	}
}
func TestHelloWorldRobbani(t *testing.T){
	result := HelloWorld("robbani")
	if result != "hello robbani"{
		panic("result is not 'hello robbani'")
	}
}