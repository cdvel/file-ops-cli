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

func File(filename string) FileStatus {
	return FileStatus{true, nil, filename}
}

func (f FileStatus) CheckFile() FileStatus {
	return f.Exists().CheckNotDir()
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

func (f FileStatus) CheckNotDir() FileStatus {

	if f.Err != nil {
		return f
	}

	info, _ := os.Stat(f.Filename)

	if info.IsDir() {
		f.Status, f.Err = false, fmt.Errorf("error: Expected file got directory '%v'", f.Filename)
	}

	return f
}

func (f FileStatus) CheckNotBinary() FileStatus {

	if f.Err != nil {
		return f
	}

	mime, _ := mimetype.DetectFile(f.Filename)

	if mime.Is("application/x-mach-binary") || mime.Is("application/octet-stream") {
		f.Status, f.Err = false, fmt.Errorf("error: Cannot do linecount for binary file '%v'", f.Filename)
	}

	return f
}
