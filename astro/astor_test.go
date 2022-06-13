package astro_test

import (
	"fmt"
	"testing"

	"github.com/nugrohoac/tlx.toki/astro"
)

func TestSort(t *testing.T) {
	t.Run("sort", func(t *testing.T) {
		//fmt.Println(astro.Sort([]int32{3, 6, 4, 4, 5, 5}))
		//fmt.Println(astro.Sort([]int32{8, 5, 5, 5, 5, 1, 1, 1, 4, 4}))
		//astro.Sort2([]int{8, 1, 2, 2, 1, 1, 3, 3})
		//fmt.Println(astro.Sort([]int{4, 1, 3, 2, -1}))
		//fmt.Println(astro.Sort([]int{7, 8, 6, -1, 3, -2}))
		fmt.Println(astro.MainSort([]int32{3, 1, 2, 2, 3, 1, 1, 8}))
	})
}
