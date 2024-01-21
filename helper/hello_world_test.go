package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)




func TestHelloWorldAssert(t *testing.T){
	result := HelloWorld("akhdan")
	assert.Equal(t, "hello akhdan", result, "result not same with 'hello akhdan")
	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("akhdan")
	require.Equal(t, "hello akhdan", result, "result not same with 'hello akhdan")
	fmt.Println("TestHelloWorld done")
}

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