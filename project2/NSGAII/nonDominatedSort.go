package NSGAII

const N = 500

type solution struct {
	genom          [N]rune
	fitCon, fitDif float64
}
type sillyFiller struct{}

func nonDominatedSort(P []*solution) {

	rank := make(map[int]int)
	F := make(map[int]map[int]sillyFiller)
	var domCount [N]int
	S := make(map[int]map[int]sillyFiller)
	for p, indp := range P {
		for q, indq := range P {
			if indq.fitCon < indp.fitCon && indq.fitDif < indp.fitDif {
				S[p][q] = struct{}{}
			} else if indp.fitCon < indq.fitCon && indp.fitDif < indq.fitDif {
				domCount[p] += 1
			}
		}
		if domCount[p] == 0 {
			F[0][p] = struct{}{}
			rank[p] = 0
		}
	}
	i := 0
	for len(F[i]) != 0 {
		i += 1
		for p, _ := range F[i-1] {
			for q, _ := range S[p] {
				domCount[q] -= 1
				if domCount[q] == 0 {
					rank[q] = i
					F[i][q] = struct{}{}
				}
			}
		}
	}
}
