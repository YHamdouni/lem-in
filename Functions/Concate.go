package lemin

func ConcatMaps(m1, m2 map[string][]string) map[string][]string {
	result := make(map[string][]string)
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		if existing, ok := result[k]; ok {
			result[k] = append(existing, v...)
		} else {
			result[k] = v
		}
	}
	return result
}
