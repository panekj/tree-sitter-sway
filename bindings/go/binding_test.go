package tree_sitter_sway_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-sway"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_sway.Language())
	if language == nil {
		t.Errorf("Error loading Sway grammar")
	}
}
