package filesys

import (
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"os"
)

type FileStatus struct {
	Status   bool
	Err      error
	Filename string
}

func (f FileStatus) Validate() FileStatus {
	file := FileStatus{true, nil, f.Filename}.Exists().IsNotDir().IsNotBinary()
	return file
}

func (f FileStatus) Exists() FileStatus {

	if f.Err != nil {
		return f
	}
	_, err := os.Stat(f.Filename)

	if os.IsNotExist(err) {
		f.Status, f.Err = false, fmt.Errorf("error: No such file '%v'", f.Filename)
	}

	return f
}

func (f FileStatus) IsNotDir() FileStatus {

	if f.Err != nil {
		return f
	}

	info, _ := os.Stat(f.Filename)

	if info.IsDir() {
		f.Status, f.Err = false, fmt.Errorf("error: Expected file got directory '%v'", f.Filename)
	}

	return f
}

func (f FileStatus) IsNotBinary() FileStatus {

	if f.Err != nil {
		return f
	}

	mime, _ := mimetype.DetectFile(f.Filename)

	if mime.Is("application/x-mach-binary") {
		f.Status, f.Err = false, fmt.Errorf("error: Cannot do linecount for binary file '%v'", f.Filename)
	}

	return f
}
