package rotate

func rotate_ints(slice []int) {
	if len(slice) == 0 {
		return
	}
	first := slice[0]
	copy(slice, slice[1:])
	slice[len(slice)-1] = first
}
