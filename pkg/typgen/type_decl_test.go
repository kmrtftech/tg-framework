package typgen_test

import (
	"testing"

	"github.com/kmrtftech/tg-framework/pkg/typgen"
	"github.com/stretchr/testify/require"
)

func TestTypeDecl(t *testing.T) {
	typeDecl := &typgen.TypeDecl{
		Name: "some-name",
		Docs: []string{"doc1", "doc2"},
	}
	require.Equal(t, "some-name", typeDecl.GetName())
	require.Equal(t, []string{"doc1", "doc2"}, typeDecl.GetDocs())
}
