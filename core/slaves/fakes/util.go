package fakes 



func AllFakes(src int) int {
	for _, count := range FakeSlaves {
		src += count
	}; return src
}