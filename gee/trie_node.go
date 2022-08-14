package gee

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	leaf     string
	item     string
	children []*TrieNode
	strict   bool
}

func (node *TrieNode) Find(item string) *TrieNode {
	for _, cur := range node.children {
		if cur.item == item || !cur.strict {
			return cur
		}
	}
	return nil
}

func (node *TrieNode) Search(item string) []*TrieNode {
	result := make([]*TrieNode, 0)
	for _, cur := range node.children {
		if cur.item == item || !cur.strict {
			result = append(result, cur)
		}
	}

	return result
}

func (node *TrieNode) Append(path string, splits []string, deep int) {
	if len(splits) == deep {
		node.leaf = path
		return
	}

	item := splits[deep]
	hint := node.Find(item)

	if hint == nil {
		hint = &TrieNode{item: item, strict: item[0] != ':' && item[0] != '*'}
		node.children = append(node.children, hint)
	}

	hint.Append(path, splits, deep+1)
}

func (node *TrieNode) Match(splits []string, deep int) *TrieNode {
	if len(splits) == deep || strings.HasPrefix(node.item, "*") {
		if node.leaf == "" {
			return nil
		}
		return node
	}

	split := splits[deep]
	list := node.Search(split)

	for _, cur := range list {
		hint := cur.Match(splits, deep+1)
		if hint != nil {
			return hint
		}
	}
	return nil
}

func (node *TrieNode) Print(deep int) {
	fmt.Print(strings.Repeat("--", deep) + node.item)
	if node.leaf == "" {
		fmt.Println()
	} else {
		fmt.Println("@@" + node.leaf)
	}
	for _, item := range node.children {
		item.Print(deep + 1)
	}
}
