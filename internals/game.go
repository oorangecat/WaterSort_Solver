// Contains game mechanics
package game

import "strings"

type VialState struct {
	Content [VialSize]byte
	Head    int //0->4, also represents the empty units
	TopLen  int //cache length of the top color to avoid multiple re-explorations
}

type GameState struct {
	VialStates []VialState
}

type StateChange struct {
	SourceVial []VialState //[starting state, final state]
	DestVial   []VialState
}

type Game struct {
	InitialState   *GameState
	NumberOfColors int
}

func ValidateStateChange(sc *StateChange) bool {
	// source musta have a color to move, dest must have space
	if sc.SourceVial[0].Head == VialSize || sc.DestVial[0].Head == 0 {
		return false
	}

	sourceColor := sc.SourceVial[0].Content[sc.SourceVial[0].Head]

	if sc.DestVial[0].Head != VialSize {
		destColor := sc.DestVial[0].Content[sc.DestVial[0].Head]

		if sourceColor != destColor { //Heads must be of the same color, dest must have a free slot
			return false
		}
	}

	var destIdx int = sc.DestVial[1].Head
	var maxDestSpace int = sc.DestVial[0].Head

	for i := 0; i < VialSize; i++ { //Will hardly ever fully loop

		//Every possible drop of color must have been moved to the new vial
		if sc.SourceVial[0].Content[sc.SourceVial[0].Head+i] != sourceColor || i > maxDestSpace {
			break
		}
		if sc.DestVial[1].Content[destIdx] != sourceColor {
			return false
		}
		destIdx++
	}

	return true
}

func CheckWin(gs *GameState) bool {
	var emptyVials int = 0

	for i := 0; i < len(gs.VialStates); i++ {
		if gs.VialStates[i].Head != 0 {
			if gs.VialStates[i].Head != VialSize {
				return false // A vial is still half empty
			} else if emptyVials < MaxEmptyVials {
				emptyVials++
			} else {
				return false // More than 2 fully empty vials
			}
		}

		for j := 1; j < VialSize; j++ {
			if gs.VialStates[i].Content[j] != gs.VialStates[i].Content[j-1] {
				return false // All colors in the vials must be the same
			}
		}
	}

	return true
}

func (vs VialState) VialHash() string {
	return string(vs.Content[vs.Head:])
}

func (state GameState) StateHash() string {
	var builder strings.Builder

	for pos := 0; pos != len(state.VialStates); pos++ {
		builder.WriteString(state.VialStates[pos].VialHash())
	}

	return builder.String()
}
