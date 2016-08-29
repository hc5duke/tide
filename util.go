package tide

func synchSafeInt(buf []byte) (r int) {
	r = 0
	for _, v := range buf {
		r <<= 7
		r += (int(v) % 128)
	}

	return r
}
