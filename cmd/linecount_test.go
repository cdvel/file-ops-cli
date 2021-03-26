package cmd

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"runtime"
	"testing"
)

var basepath string

func TestMain(m *testing.M) {
	_, filename, _, _ := runtime.Caller(0)
	basepath = path.Join(path.Dir(filename), "..")
	exitVal := m.Run()
	os.Exit(exitVal)
}

func absPath(filepath string) string {
	return path.Join(basepath, filepath)
}

func TestCount(t *testing.T) {
	cases := []struct {
		filepath string
		expects  int
	}{
		{absPath("assets/empty.txt"), 0},
		{absPath("assets/lines.txt"), 3},
		{absPath("assets/myfile.txt"), 4},
	}

	for _, testcase := range cases {
		linecount, _ := count(testcase.filepath)
		assert.Equal(t, linecount, testcase.expects, fmt.Sprintf("Failed test with %v", testcase.filepath))
	}
}
