package helper                                                                                                                 

import(
        "testing"
)                                                                                                                               
func forEach(data []int, f func(int)) {
        for _, i := range data {
                f(i)
        }
}

func checkFor3s(tb testing.TB, data []int) {
        tb.Helper()
        forEach(data, func(i int) {
                tb.Helper()
                if i != 3 {
                        tb.Error("Saw", i, "instead of 3")
                }
        })
}

func TestHelper(t *testing.T) {
        checkFor3s(t, []int{3, 3, 3, 3})
        checkFor3s(t, []int{3, 3, 2, 3})
        checkFor3s(t, []int{3, 3, 3, 3})
}
