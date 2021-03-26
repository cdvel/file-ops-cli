package algorithm

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

type Algorithm interface {
	GetType() string
	GetChecksum(data []byte) string
}

type MD5 struct {
	HashType string
}

type SHA1 struct {
	HashType string
}

type SHA256 struct {
	HashType string
}

func GetMD5() MD5 {
	return MD5{"md5"}
}

func GetSHA1() SHA1 {
	return SHA1{"sha1"}
}

func GetSHA256() SHA256 {
	return SHA256{"sha256"}
}

func (MD5) GetType() string {
	return "md5"
}

func (SHA1) GetType() string {
	return "sha1"
}

func (SHA256) GetType() string {
	return "sha256"
}

func (MD5) GetChecksum(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

func (SHA1) GetChecksum(data []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(data))
}

func (SHA256) GetChecksum(data []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(data))
}
