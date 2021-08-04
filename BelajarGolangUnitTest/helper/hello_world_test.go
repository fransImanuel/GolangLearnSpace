package helper

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B){
	benchmarks := []struct{
		name string
		request string
	}{
		{
			name: "Eko",
			request: "Eko",
		},{
			name: "Kurniawan",
			request: "Kurniawan",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B){
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B){
	b.Run("Eko", func(b *testing.B){
		for i := 0; i < b.N; i++ {
			HelloWorld("Eko")
		}
	})
	b.Run("Kurniawan", func(b *testing.B){
		for i := 0; i < b.N; i++ {
			HelloWorld("Eko")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B){
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}

func BenchmarkHelloWorldKurniawan(b *testing.B){
	for i := 0; i < b.N; i++ {
		HelloWorld("Kurniawan")
	}
}

func TestHelloWorld(t *testing.T){
	result := HelloWorld("Eko")

	if result != "Hello Eko"{
		//error
		// panic("Result is not 'Hello Eko'")
		// t.Fail()
		t.Error("result must be Hello Eko")
	}

	fmt.Println("Test HelloWorld Done")
}

func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with require done")
}

func TestHelloWorldAssert(t *testing.T){
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with assert done")
}

func TestSkip(t *testing.T){
	// if runtime.GOOS == "darwin"{
		t.Skip("can not run in macos")
	// }
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with assert done")
}

func TestHelloWorldFrans(t *testing.T){
	result := HelloWorld("Frans")
	
	if result != "Hello Frans"{
		//error
		// panic("Result is not 'Hello Frans'")
		// t.FailNow()
		t.Error("result must be Hello Frans")
	}


	
	fmt.Println("Test TestHelloWorldFrans Done")
}

func TestMain(m *testing.M){
	//before
	fmt.Println("Sebelum unittest")

	m.Run()

	//after
	fmt.Println("Setelah unittest")
}

func TestSubTest(t *testing.T){
	t.Run("Eko", func(t *testing.T){
		result := HelloWorld("Eko")
		require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	})
	
	t.Run("Kurniawan", func(t *testing.T){
		result := HelloWorld("Kurniawan")
		require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	})
}

func TestTableHelloWorld(t *testing.T){
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name : "Eko",
			request: "Eko",
			expected: "Hello Eko",
		},{
			name : "Kurniawan",
			request: "Kurniawan",
			expected: "Hello Kurniawan",
		},{
			name : "Khannedy",
			request: "Khannedy",
			expected: "Hello Khannedy",
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T){
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}