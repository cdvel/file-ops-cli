package filesys

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

func init() {
	_, filename, _, _ := runtime.Caller(0)
	basepath := path.Join(path.Dir(filename), "..")
	err := os.Chdir(basepath)
	if err != nil {
		panic(err)
	}
}

func absPath(filepath string) string {
	return path.Join(basepath, filepath)
}

func TestCheckFile(t *testing.T) {
	cases := []struct {
		filepath string
		expects  bool
	}{
		{absPath("missing-file.txt"), false},
		{absPath("assets/empty.txt"), true},
		{absPath("assets/lines.txt"), true},
		{absPath("assets/myfile.txt"), true},
		{absPath("assets/temp"), false},
	}

	f := File("")
	for _, testcase := range cases {

		f = File(testcase.filepath).CheckFile()
		assert.Equal(t, f.Status, testcase.expects, fmt.Sprintf("Failed test with %v", f.Filename))
	}
}

func TestExists(t *testing.T) {
	cases := []struct {
		filepath string
		expects  bool
	}{
		{absPath("missing-file.txt"), false},
		{absPath("assets/myfile.txt"), true},
	}

	f := File("")
	for _, testcase := range cases {

		f = File(testcase.filepath).Exists()
		assert.Equal(t, f.Status, testcase.expects, fmt.Sprintf("Failed test with %v", f.Filename))
	}
}

func TestCheckNotDir(t *testing.T) {
	f := File(absPath("assets/tmp")).CheckNotDir()
	assert.Equal(t, f.Status, false, fmt.Sprintf("Failed test with %v", f.Filename))
}

func TestCheckNotBinary(t *testing.T) {
	f := File(absPath("assets/binary.dat")).CheckNotBinary()
	assert.Equal(t, f.Status, false, fmt.Sprintf("Failed test with %v", f.Filename))
}
