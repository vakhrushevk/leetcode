package duplicateinteger

func HashDuplicate(arr []int) bool {
	m := make(map[int]struct{})

	for i := 0; i < len(arr); i++ {
		if _, ok := m[arr[i]]; ok {
			return true
		}
		m[arr[i]] = struct{}{}

	}
	return false
}
