package propagation



func AllSlaves(src int) int {
	for _, count := range Reps {
		src += count
	}; return src
}