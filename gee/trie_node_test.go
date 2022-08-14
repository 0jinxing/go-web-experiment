package gee

import "testing"

func TestAppend(t *testing.T) {
	node := &TrieNode{item: "/", strict: false}
	node.Append("/hello", []string{"hello"}, 0)
	node.Append("/hello/world", []string{"hello", "world"}, 0)
	node.Append("/hello/world1", []string{"hello", "world1"}, 0)

	ok := len(node.children) == 1

	if !ok {
		t.Fatal("node children len should eq 1")
	}

	node.Append("/hi/world", []string{"hi", "world"}, 0)

	ok = len(node.children) == 2

	if !ok {
		t.Fatal("node children len should eq 2")
	}

	node.Print(0)
}
