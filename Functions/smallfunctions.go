package lemin

func CopySlice(dest, src []string) {
	n := len(dest)
	for i := 0; i < n; i++ {
		dest[i] = src[i]
	}
}

func Contains(path []string, node string) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}
