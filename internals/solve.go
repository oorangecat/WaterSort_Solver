// Contains game solver
package game

type SwapProposal struct {
	stateChange          StateChange
	heuristicImprovement int
}

type GamePath struct {
	Head *PathNode
}

/*
1. At each state we must generate possible state changes, eventually exploring only the most promising ones,
calculating the heuristic for each one.
2. Each pathnode should be in a max-heap so that the most promising one is always explored.
3. Discarded paths should be added to a set of explored states to avoid re-exploring
4. Exploration should continue from top of heap
5. Only back-tracking info is useful (= previous Node for path reconstruction)
6. If the new-found node is in the heap the shortest pathLen should be updated, using a map/set keeping the references
*/
type PathNode struct {
	gameState     GameState
	heuristic     int
	exploredPaths int
	availblePaths []StateChange
	prevNode      *PathNode
	pathLen       int
}

type Solver struct {
	CurrentGame     *Game
	CurrentSolution *GamePath
	SolutionLen     int
}

//n^2 search as the number of vials will be very small
/*
1. find 2 vials with the same head color
2. if they are both full, skip
3. count number of swappable units
	3.1 Find available spots in destination
	3.2 Count number of subsequent colors in source
	3.3 Move the minimum of the two in the other vial, creating a new VialState to be added to the StateChange of the SwapProposal
*/

func (s Solver) generateProposals(state GameState) []SwapProposal {
	var proposals []SwapProposal
	var vial1, vial2 VialState

	for pos1 := 0; pos1 < len(state.VialStates)-1; pos1++ {
		vial1 = state.VialStates[pos1]
		for pos2 := pos1; pos2 < len(state.VialStates); pos2++ {
			vial2 = state.VialStates[pos2]
			if vial2.Head == vial1.Head && vial1.Head == 0 { // if both are full, skip
				continue
			}

			if vial1.Content[vial1.Head] == vial2.Content[vial2.Head] { // Same top color
				var stateChange StateChange
				var swapLen int
				if vial1.Head != 0 {
					if vial2.Head != 0 { //vial1 -> vial2
						stateChange.SourceVial[0] = vial1
						stateChange.DestVial[0] = vial2
						swapLen = min(vial2.Head, vial1.TopLen)

						if swapLen == 0 { //TODO remove after testing, should be impossible
							continue
						}

						var newvial1, newvial2 VialState
						newvial1.Head = vial1.Head + swapLen
						copy(newvial1.Content[newvial1.Head:], vial1.Content[newvial1.Head:]) //Copy remaining vial1
						newvial2.Head = vial2.Head - swapLen

						copy(newvial2.Content[vial2.Head:], vial1.Content[vial2.Head:])                                //Copy old ones vial2
						copy(newvial1.Content[newvial2.Head:vial2.Head], vial1.Content[vial1.Head:vial1.Head+swapLen]) //Copy new ones vial1->vial2

						stateChange.SourceVial[1] = newvial1
						stateChange.SourceVial[2] = newvial2

						if newvial1.Head == VialSize {
							newvial1.TopLen = 0
						} else {
							var topcolor = newvial1.Content[newvial1.Head]
							var toplen = 0
							var pos = newvial1.Head
							for newvial1.Content[pos] == topcolor {
								toplen++
								pos++
							}
							newvial1.TopLen = toplen
						}

						var topcolor = newvial2.Content[newvial2.Head]
						var toplen = 0
						var pos = newvial2.Head
						for newvial2.Content[pos] == topcolor {
							toplen++
							pos++
						}
						newvial2.TopLen = toplen
					}
				}
			}
		}
	}

	return proposals
}

func doSwap(vial1, vial2 VialState) StateChange {
	var stateChange StateChange
	var newvial1, newvial2 VialState

	var swapLen = min(vial2.Head, vial1.TopLen)

	stateChange.SourceVial[0] = vial1
	stateChange.DestVial[0] = vial2

	newvial1.Head = vial1.Head + swapLen
	copy(newvial1.Content[newvial1.Head:], vial1.Content[newvial1.Head:]) //Copy remaining vial1

	newvial2.Head = vial2.Head - swapLen
	copy(newvial2.Content[vial2.Head:], vial1.Content[vial2.Head:])                                //Copy old ones vial2
	copy(newvial1.Content[newvial2.Head:vial2.Head], vial1.Content[vial1.Head:vial1.Head+swapLen]) //Copy new ones vial1->vial2

	stateChange.SourceVial[1] = newvial1
	stateChange.SourceVial[2] = newvial2

	//Update toplens of new vials
	if newvial1.Head == VialSize {
		newvial1.TopLen = 0
	} else {
		var topcolor = newvial1.Content[newvial1.Head]
		var toplen = 0
		var pos = newvial1.Head
		for newvial1.Content[pos] == topcolor {
			toplen++
			pos++
		}
		newvial1.TopLen = toplen
	}

	var topcolor = newvial2.Content[newvial2.Head]
	var toplen = 0
	var pos = newvial2.Head
	for newvial2.Content[pos] == topcolor {
		toplen++
		pos++
	}
	newvial2.TopLen = toplen

	return stateChange
}
