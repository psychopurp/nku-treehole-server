package service

import (
	"testing"
)

func TestPostService_GetPosts(t *testing.T) {
	setup()
	service := &PostService{}

	res, err := service.GetPosts(0, 10)
	t.Log(res, err)

	res, err = service.GetPosts(1, 10)
	t.Log(res, err)

}
