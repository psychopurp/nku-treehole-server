package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// md5 encryption
func Md5(src string) string {
	m := md5.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
