package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSub(b *testing.B){
	b.Run("akhdan",func(b *testing.B) {
		for i := 0; i < b.N ; i++{
			HelloWorld("Akhdan")
		}
	})
	b.Run("oukenzeumasio",func(b *testing.B) {
		for i := 0; i < b.N ; i++{
			HelloWorld("oukenzeumasio")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B){
	for i := 0; i < b.N ; i++{
		HelloWorld("Akhdan")
	}
}
func BenchmarkHelloWorldRobbani(b *testing.B){
	for i := 0; i < b.N ; i++{
		HelloWorld("Robbani")
	}
}

func TestTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(akhdan)",
			request:  "akhdan",
			expected: "hello akhdan",
		},
		{
			name:     "HelloWorld(robbani)",
			request:  "robbani",
			expected: "hello robbani",
		},
		{
			name:     "HelloWorld(oukenze)",
			request:  "oukenze",
			expected: "hello oukenze",
		},
		{
			name:     "HelloWorld(akeon)",
			request:  "akoen",
			expected: "hello akeon",
		},
		{
			name:     "HelloWorld(zeuma)",
			request:  "zeuma",
			expected: "hello zeuma",
		},
	}
	for _, value := range tests {
		t.Run(value.name, func(t *testing.T) {
			result := HelloWorld(value.request)
			assert.Equal(t, value.expected, result, "result not same with '%v'", value.expected)

		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("akhdan", func(t *testing.T) {
		result := HelloWorld("akhdan")
		assert.Equal(t, "hello akhdan", result, "result not same with 'hello akhdan")
	})
	t.Run("robbani", func(t *testing.T) {
		result := HelloWorld("robbanu")
		assert.Equal(t, "hello robbani", result, "result not same with 'hello robbani")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Sebelum Unit Test")

	m.Run()

	fmt.Println("Sesudah Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("can't run in linux")
	}
	result := HelloWorld("akhdan")
	assert.Equal(t, "hello akhdan", result, "result not same with 'hello akhdan")
	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("akhdan")
	assert.Equal(t, "hello akhdan", result, "result not same with 'hello akhdan")
	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("akhdan")
	require.Equal(t, "hello akhdan", result, "result not same with 'hello akhdan")
	fmt.Println("TestHelloWorld done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("akhdan")
	if result != "hello akhdan" {
		t.Error("result not same with 'hello akhdan'")
		// t.Fail("result not same with 'hello akhdan'")
	}
	fmt.Println("TestHelloWorld done")
}
func TestHelloWorldRobbani(t *testing.T) {
	result := HelloWorld("robbani")
	if result != "hello robbani" {
		// t.FailNow("result not same with 'hello robbani'")
		t.Fatal("result not same with 'hello robbani'")

	}
	fmt.Println("TestHelloWorldRobbani done")
}
