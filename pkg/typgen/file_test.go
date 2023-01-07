package typgen_test

import (
	"testing"

	"github.com/kmrtftech/tg-framework/pkg/typgen"
	"github.com/stretchr/testify/require"
)

func TestFile_SourceCode(t *testing.T) {
	testCases := []struct {
		TestName string
		File     *typgen.File
		Expected string
	}{
		{
			File: &typgen.File{
				Name: "some package",
				Imports: []*typgen.Import{
					{Name: "", Path: "fmt"},
					{Name: "a", Path: "github.com/kmrtftech/tg-framework"},
				},
			},
			Expected: `package some package

import (
	"fmt"
	a "github.com/kmrtftech/tg-framework"
)`,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.File.Code())
		})
	}
}
