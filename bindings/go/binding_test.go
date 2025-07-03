package tree_sitter_epics_msi_template_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_epics_msi_template "github.com/minijackson/tree-sitter-epics-msi-template/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_epics_msi_template.Language())
	if language == nil {
		t.Errorf("Error loading EpicsMsiTemplate grammar")
	}
}
