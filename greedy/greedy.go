package greedy

type Line struct {
	User        string
	Priority    int
	Performance float32
}

func NewLine(user string, priority int, performance float32) *Line {
	return &Line{user, priority, performance}
}

func GetUserComp(lines []*Line, alocatedUsers map[string]bool) ([]*Line, int) {
	userList := make([]*Line, 5)
	count := 0
	comparacoes := 0
	iters := 0
	for i := 0; count < 5 || iters < len(lines); i++ {
		line := lines[i%len(lines)]

		comparacoes++
		if alocatedUsers[line.User] {
			iters++
			continue
		}

		comparacoes++
		if userList[line.Priority] == nil {
			userList[line.Priority] = line
			alocatedUsers[line.User] = true
			count++
		} else if userList[line.Priority].Performance > line.Performance {
			comparacoes++
			alocatedUsers[userList[line.Priority].User] = false
			userList[line.Priority] = line
			alocatedUsers[line.User] = true
		}

		iters++
	}

	return userList, comparacoes
}
