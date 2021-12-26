package arrays

func CopySlice(src []int) []int {
	dest := make([]int, len(src), (cap(src)+1)*2)
	// for i := range src {
	// 	dest[i] = src[i]
	// }
    copy(dest, src)
	return dest
}

func AppendToSlice(slice []int, data ...int) []int {
    m := len(slice)
    n := m + len(data)

    if n > cap(slice) {
        newSlice := make([]int, (n + 1) * 2)
        copy(newSlice, slice)
        slice = newSlice
    }

    slice = slice[0:n]
    copy(slice[m:n], data)
    return slice
}

func SliceFilter(s []int, fn func(int) bool) []int {
    var p []int // == nil
    for _, v := range s {
        if fn(v) {
            p = append(p, v)
        }
    }
    return p
}
