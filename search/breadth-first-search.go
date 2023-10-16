package search

func BreadthFirstSearch(graph map[string][]string, startNode string) (bool, string) {
	queue := make([]string, 0)
	queue = append(queue, graph[startNode]...)

	alreadyVerified := map[string]bool{}

	for len(queue) > 0 {
		seller := queue[0]
		queue = queue[1:]

		if !alreadyVerified[seller] {
			if isMangoSeller(seller) {
				return true, seller
			} else {
				queue = append(queue, graph[seller]...)
				alreadyVerified[seller] = true
			}
		}

	}

	return false, ""
}

func isMangoSeller(seller string) bool {
	return seller == "thom"
}
