package model

import (
	"testing"
)

func TestPost_GetPosts(t *testing.T) {
	setup()
	post := &Post{}

	t.Log(post.GetPosts(0, 5))
}
