package md5

import "testing"

func TestMd5(t *testing.T) {
	t.Log(Md5("hello"))
	t.Log(Md5("hello"))
}
