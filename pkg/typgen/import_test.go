package typgen_test

import (
	"go/ast"
	"testing"

	"github.com/kmrtftech/tg-framework/pkg/typgen"
	"github.com/stretchr/testify/require"
)

func TestImports(t *testing.T) {
	f := &ast.File{
		Imports: []*ast.ImportSpec{
			{
				Name: &ast.Ident{Name: "f"},
				Path: &ast.BasicLit{Value: "\"fmt\""},
			},
			{
				Name: &ast.Ident{Name: "s"},
				Path: &ast.BasicLit{Value: "\"strings\""},
			},
		},
	}
	expected := []*typgen.Import{
		{Name: "f", Path: "fmt"},
		{Name: "s", Path: "strings"},
	}
	require.Equal(t, expected, typgen.CreateImports(f))
}
