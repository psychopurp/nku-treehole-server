package id_generator

import "testing"

func TestGenerateID(t *testing.T) {
	for i := 0; i < 20; i++ {
		t.Log(GenerateID())
	}
}
