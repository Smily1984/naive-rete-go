package rete

import (
	"container/list"
	"testing"
)

func Test_make_token(t *testing.T) {
	node := BetaMemory{
		items:    list.New(),
		parent:   nil,
		children: list.New(),
	}
	w := &WME{
		fields:     [4]string{"Object", "B1", "on", "table"},
		alpha_mems: list.New(),
		tokens:     list.New(),
	}
	token := make_token(&node, nil, w, nil)
	if token.node.GetNodeType() != BETA_MEMORY_NODE {
		t.Error("token node type error")
	}
	if token.wme != w {
		t.Error("token wme error")
	}
}
