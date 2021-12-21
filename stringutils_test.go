package stringutils_test

import (
	"testing"

	"github.com/serhatdmr/stringutils-demo"
)

func TestJoin(t *testing.T) {

	tcs := map[string]struct {
		input  []string
		input2 string
		want   string
	}{
		"with argument": {
			input:  []string{"Serhat", "Ahmet"},
			input2: ",",
			want:   "Serhat, Ahmet",
		},
		"null argument": {
			input:  nil,
			input2: ",",
			want:   "",
		},
	}

	for name, tc := range tcs {

		t.Run(name, func(t *testing.T) {

			got := stringutils.Join(tc.input, tc.input2)

			if tc.want != got {
				t.Errorf("want: %v got: %v", tc.want, got)
			}
		})
	}
}

var gs string

func BenchmarkJoin(b *testing.B) {

	var s string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = stringutils.Join([]string{"Serhat", "Ahmet", "Mehmet"}, "-")
	}
	gs = s

}
