package gee

import (
	"reflect"
	"testing"
)

func TestSplitTrieNodePath(t *testing.T) {
	ok := reflect.DeepEqual(SplitTrieNodePath("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(SplitTrieNodePath("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(SplitTrieNodePath("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test ParseSplits failed")
	}
}
