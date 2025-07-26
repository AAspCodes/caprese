package cut

func Cut(wholeTomato string) []rune {
	tomSlices := []rune{}
	for _, piece := range wholeTomato {
		tomSlices = append(tomSlices, piece)
	}
	return tomSlices
}
