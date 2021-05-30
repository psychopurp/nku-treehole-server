package handler

import (
	"fmt"
	"testing"
)

func TestGetPosts(t *testing.T) {

	tests := []struct {
		limit      int
		totalCount int
		totalPage  int
	}{
		// TODO: Add test cases.
		{5, 20, 4},
		{1, 10, 10},
		{5, 4, 1},
		{5, 5, 1},
		{5, 6, 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			totalPage := (tt.totalCount + tt.limit - 1) / tt.limit
			if tt.totalPage != totalPage {
				t.Error(tt, totalPage)
			}
		})
	}
}
