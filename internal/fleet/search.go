package fleet

func Search(vehicles []Vehicle, branch, date string) []Vehicle {
	matches := []Vehicle{}
	for _, v := range vehicles {
		if v.Branch == branch && Available(v, date) {
			matches = append(matches, v)
		}
	}
	return matches
}
