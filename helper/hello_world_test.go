package helper

import (
	"fmt"
	"testing"
)




func TestHelloWorld(t *testing.T){
	result := HelloWorld("akhdan")
	if result != "hello akhdan"{
		t.Error("result not same with 'hello akhdan'")
		// t.Fail("result not same with 'hello akhdan'")
	}
	fmt.Println("TestHelloWorld done")
}
func TestHelloWorldRobbani(t *testing.T){
	result := HelloWorld("robbani")
	if result != "hello robbani"{
		// t.FailNow("result not same with 'hello robbani'")
		t.Fatal("result not same with 'hello robbani'")
		
	}
	fmt.Println("TestHelloWorldRobbani done")
}