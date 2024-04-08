package generics

func Max[T comparable](values []T) T {
	if len(values) == 0 {
		panic("empty slice")
	}
	max := values[0]
	for _, v := range values[1:] {
		if v.(comparable) > max.(comparable) {
			max = v
		}
	}
	return max
}
func main() {

}
