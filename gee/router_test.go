package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(SplitTrieNodePath("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(SplitTrieNodePath("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(SplitTrieNodePath("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test ParseSplits failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := NewRouter()

	r.AddRoute("GET", "/", nil)
	r.AddRoute("GET", "/hello/:name", nil)
	r.AddRoute("GET", "/hello/b/c", nil)
	r.AddRoute("GET", "/hi/:name", nil)
	r.AddRoute("GET", "/assets/*filepath", nil)

	node, params := r.GetRoute("GET", "/hello/geektutu")

	if node == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if node.leaf != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if params["name"] != "geektutu" {
		t.Fatal(`name should be equal to "geektutu"`)
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", node.leaf, params["name"])
}
