package groupsanagrams

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v", test.strs)
		t.Run(testName, func(t *testing.T) {
			got := groupAnagrams(test.strs)

			// Sort slice
			sort.Slice(got, func(i, j int) bool {
				return len(got[i]) < len(got[j])
			})

			// Sort element in the group
			for i := range got {
				sort.Strings(got[i])
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}
