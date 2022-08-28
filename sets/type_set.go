package sets

type Set interface {
	[]any
}

func NewSet[T comparable, S Set](members ...T) (set S) {
	appendToSet := func(m T) {
		set = append(set, m)
	}
	checkIfNew := func(m T) (first bool) {
		first = true
		for _, el := range set {
			if el == m {
				first = false
			}
		}
		return
	}

	for _, member := range members {
		switch len(set) {
		case 0:
			appendToSet(member)
		default:
			if checkIfNew(member) {
				appendToSet(member)
			}
		}
	}
	return
}
