package traversal

import (
	"testing"
	"fmt"
	. "github.com/NewbMiao/AlgorithmCodeNote/kit/btree"
)

var sbt *Node

func init() {
	sbt = &Node{Value: 1}
	sbt.Left = &Node{Value: 2}
	sbt.Right = &Node{Value: 3}
	sbt.Left.Left = &Node{Value: 4}
	sbt.Right.Right = &Node{Value: 5}
}

func TestSerialByLevel(t *testing.T) {
	want := "1!2!3!4!#!#!5!#!#!#!#!"
	t.Log(fmt.Sprintf("input btree=%v, SerialByLevel want: %v", sbt, want))
	r := SerialByLevel(sbt)
	if r == want {
		t.Log("SerialByLevel is ok")
	} else {
		t.Error("SerialByLevel is not ok, result:", r)
	}
}

//go test  -run TestReconByLevel -v src/btree/traversal/*
func TestReconByLevel(t *testing.T) {
	r := ReconByLevel("1!2!3!4!#!#!5!#!#!#!#!")
	PrintTree(r)
}

func TestSerialByPre(t *testing.T) {
	want := "1!2!4!#!#!#!3!#!5!#!#!"
	t.Log(fmt.Sprintf("input btree=%v, SerialByPre want: %v", sbt, want))
	r := SerialByPre(sbt)
	if r == want {
		t.Log("SerialByPre is ok")
	} else {
		t.Error("SerialByPre is not ok, result:", r)
	}
}

//go test  -run TestReconByLevel -v src/btree/traversal/*
func TestReconByPre(t *testing.T) {
	r := ReconByPre("1!2!4!#!#!#!3!#!5!#!#!")
	PrintTree(r)
}
