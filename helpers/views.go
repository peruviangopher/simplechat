package helpers

func GetChatRoomsForView(rooms int) []int {
	nums := make([]int, rooms)
	for i := 0; i < rooms; i++ {
		nums[i] = i + 1
	}

	return nums
}
