package cmd

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestGetChecksum(t *testing.T) {
	cases := []struct {
		filepath  string
		algorithm string
		expects   string
	}{
		{absPath("assets/myfile.txt"), "md5", "821348cb57676aff720f5c3b4a6f3c7b"},
		{absPath("assets/myfile.txt"), "sha1", "159abd5efb0d815114e5933de86206a70bb0a8ae"},
		{absPath("assets/myfile.txt"), "sha256", "6b5bb6a926b59c46f80b375a4bf709e343bf1dbcc819d92996b9d31cfd6d0b5c"},
	}

	for _, testcase := range cases {
		contents, _ := ioutil.ReadFile(testcase.filepath)
		sum := GetChecksum(contents, testcase.algorithm)
		assert.Equal(t, sum, testcase.expects, fmt.Sprintf("Failed test with %v", testcase.filepath))
	}
}

func TestChecksumRun(t *testing.T) {
	cmd := ChecksumCmd()
	b := bytes.NewBufferString("")
	assert.Error(t, cmd.RunE(cmd, []string{}))
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--file", absPath("assets/myfile.txt"), "--md5"})
	cmd.Execute()
	_, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
}
