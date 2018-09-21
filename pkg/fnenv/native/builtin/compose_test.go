package builtin

import (
	"testing"

	"github.com/fission/fission-workflows/pkg/types"
	"github.com/fission/fission-workflows/pkg/types/typedvalues"
)

func TestFunctionComposePassInput(t *testing.T) {
	expected := "ComposeCompose"
	internalFunctionTest(t,
		&FunctionCompose{},
		&types.TaskRunSpec{
			Inputs: map[string]*types.TypedValue{
				ComposeInput: typedvalues.MustParse(expected),
			},
		},
		expected)
}

func TestFunctionComposeEmpty(t *testing.T) {
	internalFunctionTest(t,
		&FunctionCompose{},
		&types.TaskRunSpec{
			Inputs: map[string]*types.TypedValue{},
		},
		nil)
}

func TestFunctionComposeObject(t *testing.T) {
	internalFunctionTest(t,
		&FunctionCompose{},
		&types.TaskRunSpec{
			Inputs: map[string]*types.TypedValue{
				"foo":     typedvalues.MustParse(true),
				"bar":     typedvalues.MustParse(false),
				"default": typedvalues.MustParse("hello"),
			},
		},
		map[string]interface{}{
			"foo":     true,
			"bar":     false,
			"default": "hello",
		})
}
