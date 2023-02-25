package helper

import (
	"fmt"
	"runtime"
	"testing"


	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)




func Benchmarksub(b *testing.B) { 
	b.Run("wanta", func (b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("wanta")
		}
	})
	b.Run("fikri", func (b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("fikri")
		}
	})

}

func BenchmarHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("wanta")
	}

}
func BenchmarHelloWorld1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("fikri")
	}

}


func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string 
		request  string 
		expected string 
	}{
		{
		name: "wanta",
		request: "wanta",
		expected: "Hello wanta",
		},
		{
			name: "zulfikri",
			request: "zulfikri",
			expected: "Hello zulfikri",
		},
		{
			name: "fikri", 
			request: "fikri",
			expected: "Hello fikri",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
			})
	}
}
func BenchmarkTable(b *testing.B) {

	tests := []struct {
		name     string 
		request  string 
		expected string 
	}{
		{
		name: "wanta",
		request: "wanta",
		expected: "Hello wanta",
		},
		{
			name: "zulfikri",
			request: "zulfikri",
			expected: "Hello zulfikri",
		},
		{
			name: "fikri", 
			request: "fikri",
			expected: "Hello fikri",
		},
	}
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B){
			result := HelloWorld(test.request)
			require.Equal(b, test.expected, result)
			})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("wanta", func (t *testing.T)  {
		result := HelloWorld("wanta")
		require.Equal(t, "HelloWanta", result, "Result must be 'Hello Wanta'")
	})
	t.Run("Zulfikri", func(t *testing.T){
		result := HelloWorld("Zulfikri")
		require.Equal(t,"HelloZulfikri", result, "rsult must be 'hello wanta'")
	})
}

func TestMain(m *testing.M){
	fmt.Println("sebelum unit test")

	m.Run()

	fmt.Println("SESUDAH UNIT TEST")
}
func TestSkip(t * testing.T){
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test tidak bisa jalan di mac")
	}

	result := HelloWorld("wanta")
	require.Equal(t, "Hello Wanta", result)
}
func TestHelloWorldrequire(t * testing.T) {
	result := HelloWorld("wanta")
	require.Equal(t, "Hellowanta", result, "result must be 'hello wanta' ")
	fmt.Println("TestHelloWorld with Assert Done ")

}
func TestHelloWorldAssert(t * testing.T) {
	result := HelloWorld("wanta")
	assert.Equal(t, "Hellowanta", result, "result must be 'hello wanta' ")
	fmt.Println("TestHelloWorld with Assert Done ")

}



func TestHelloWorld(t *testing.T){
	result := HelloWorld ("wanta")
	if result != "Hello wanta" {
		// unit test failed
		t.Error("Result masbi wanta")
	}
	fmt.Println("TestHelloWorld wanta")
}
func TestHelloWorldfikri(t *testing.T){
	result := HelloWorld ("wanta fikri")
	if result != "Hello wanta" {
		// unit test failed
		t.Fatal("result wanta zulfikri")
	}
	fmt.Println("HELLOWORLD WANTA")
}
