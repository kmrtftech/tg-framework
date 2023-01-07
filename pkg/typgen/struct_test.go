package typgen_test

import (
	"go/ast"
	"reflect"
	"testing"

	"github.com/kmrtftech/tg-framework/pkg/typgen"
	"github.com/stretchr/testify/require"
)

func TestStructTag(t *testing.T) {
	testcases := []struct {
		TestName string
		Tag      *ast.BasicLit
		Expected reflect.StructTag
	}{
		{
			Tag:      &ast.BasicLit{Value: "``"},
			Expected: reflect.StructTag(""),
		},
		{
			Tag:      &ast.BasicLit{Value: "`key1=value1 key2=value2`"},
			Expected: reflect.StructTag("key1=value1 key2=value2"),
		},
		{
			Tag:      &ast.BasicLit{},
			Expected: reflect.StructTag(""),
		},
		{
			Tag:      &ast.BasicLit{Value: "`"},
			Expected: reflect.StructTag(""),
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, typgen.StructTag(tt.Tag))
		})
	}
}
