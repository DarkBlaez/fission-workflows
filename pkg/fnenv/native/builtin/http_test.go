package builtin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fission/fission-workflows/pkg/types"
	"github.com/fission/fission-workflows/pkg/types/typedvalues"
	"github.com/stretchr/testify/assert"
)

func TestFunctionHttp_Invoke(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			panic("incorrect method")
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		if v, ok := r.Header["Foo"]; !ok || len(v) == 0 || v[0] != "Bar" {
			panic("Header 'Foo: Bar' not present")
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(data))
	}))
	defer ts.Close()

	fn := &FunctionHTTP{}
	body := "body"
	out, err := fn.Invoke(&types.TaskRunSpec{
		Inputs: map[string]*types.TypedValue{
			types.InputMethod: typedvalues.MustParse(http.MethodPost),
			HttpInputUrl:      typedvalues.MustParse(ts.URL),
			types.InputMain:   typedvalues.MustParse(body),
			types.InputHeaders: typedvalues.MustParse(map[string]interface{}{
				"Foo": "Bar",
			}),
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, body, typedvalues.MustFormat(out))
}

func TestFunctionHttp_Invoke_Invalid(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "expected error", 500)
	}))
	defer ts.Close()

	fn := &FunctionHTTP{}
	body := "body"
	out, err := fn.Invoke(&types.TaskRunSpec{
		Inputs: map[string]*types.TypedValue{
			types.InputMethod: typedvalues.MustParse(http.MethodDelete),
			HttpInputUrl:      typedvalues.MustParse(ts.URL),
			types.InputMain:   typedvalues.MustParse(body),
			types.InputHeaders: typedvalues.MustParse(map[string]interface{}{
				"Foo": "Bar",
			}),
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, "expected error\n", typedvalues.MustFormat(out))
}
