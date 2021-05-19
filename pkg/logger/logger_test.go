package logger

import "testing"

func TestName(t *testing.T) {

	printf(DEBUG, "test %s", "hello")

	Debugf("hello")
	Infof("hello")
	Warnf("hello")
	Errorf("hello")
	Fatalf("hello")

	Errorf("world")

}
