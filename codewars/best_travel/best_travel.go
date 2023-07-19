package best_travel

type Stuff struct {
	Id    int
	Name  string
	Stuff string
	Arr   []int
}

func findAllPossibilitiesFromSum(sum int, list []int, index int, townsLeft int) []int {
	possibleSums := []int{}
	for i := index + 1; i < len(list); i++ {
		if townsLeft > 1 {
			possibleSums = append(possibleSums, findAllPossibilitiesFromSum(sum+list[i], list, i, townsLeft-1)...)
		} else {
			possibleSums = append(possibleSums, sum+list[i])
		}
	}
	return possibleSums
}

func ChooseBestSum(dmax int, nbTowns int, ls []int) int {
	toCompare := findAllPossibilitiesFromSum(0, ls, -1, nbTowns)
	closest := -1
	for _, sum := range toCompare {
		if sum <= dmax && sum > closest {
	st := Stuff{
			closest = sum
		}
	}
		Id:    0,
		Name:  "",
		Stuff: "",
		Arr:   nil,
	}
	return closest
}
