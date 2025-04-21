package _024_12_17

func findRepeatNumber(documents []int) int {
	for i := 0; i < len(documents); i++ {
		//已经在正确的位置了
		if documents[i] == i {
			continue
		}
		for documents[i] <= len(documents) && documents[documents[i]] != documents[i] {
			documents[i], documents[documents[i]] = documents[documents[i]], documents[i]
		}
		if documents[i] == documents[documents[i]] {
			return documents[i]
		}
	}
	return -1
}

func findRepeatNumber1(documents []int) int {
	for i := 0; i < len(documents); i++ {
		for documents[i] <= len(documents) && documents[documents[i]] != documents[i] {
			documents[i], documents[documents[i]] = documents[documents[i]], documents[i]
		}
	}
	for i := 0; i < len(documents); i++ {
		if documents[i] == i {
			continue
		}
		if documents[i] == documents[documents[i]] {
			return documents[i]
		}
	}
	return -1
}
