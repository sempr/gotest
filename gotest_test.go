package gotest

import (
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil { //try a unit test on function
		t.Error("divide not pass, result should be 3")
	} else {
		t.Log("test 1 passed!")
	}
}

func Test_Division_2(t *testing.T) {
	//t.Log("第二个测试也通过了....")
	t.Error("haha you failed the test")
}
