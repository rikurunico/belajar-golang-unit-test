package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	//Before
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	fmt.Println("AFTER UNIT TEST")
	//After
}

func TestSubTest(t *testing.T) {
	t.Run("Nico", func(t *testing.T) {
		result := HelloWorld("Nico")
		require.Equal(t, "Hello Nico", result, "Result should be Hello Nico")
	})

	t.Run("Dwi", func(t *testing.T) {
		result := HelloWorld("Dwi")
		require.Equal(t, "Hello Dwi", result, "Result should be Hello Dwi")
	})

	//Command Untuk Hanya Menjalankan 1 Sub Test
	//go test -v -run TestSubTest/Nico
}  

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Nico")
	if result != "Hello Nico" {
		t.Errorf("HelloNico(\"Nico\") = %s; want \"Hello Nico\"", result)
	}
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on Windows")
	}

	result := HelloWorld("Nico")
	require.Equal(t, "Hello Nico", result, "Result should be Hello Nico")
	fmt.Println("Test Hello World Require Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Nico")
	require.Equal(t, "Hello Nico", result, "Result should be Hello Nico")
	fmt.Println("Test Hello World Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Nico")
	assert.Equal(t, "Hello Nico", result, "Result should be Hello Nico")
	fmt.Println("Test Hello World Assert Done")
}

func TestHelloWorldAfay(t *testing.T) {
	result := HelloWorld("Afay")
	if result != "Hello Afay" {
		t.Fail()
	}
}

// Run Command:
// go test (Untuk melihat hasil test)
// go test -v (Untuk melihat detail test)
// go test -cover (Untuk melihat coverage test)
// go test -v -run TestHelloWorldAfay (Untuk melihat detail test dari function tertentu)
