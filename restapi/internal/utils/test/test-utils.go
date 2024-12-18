package test

func ToPtr[T any](v T) *T {
	return &v
}

func ConvertToSet(arr []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, v := range arr {
		set[v] = struct{}{}
	}
	return set
}
