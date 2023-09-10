package game

import "testing"

func TestSmth(t *testing.T) {

}

func TestObjects(t *testing.T) {
	var state1, state2 VialState

	for i := 0; i < VialSize; i++ {
		state1.Content[i] = 'a'
		state2.Content[i] = 'b'
	}

	for i := 0; i < VialSize; i++ {
		if state1.Content[i] != 'a' || state2.Content[i] != 'b' {
			t.Error("Error in VialState internals. Content does not match")
		}
	}

	var game GameState
	game.VialStates = append(game.VialStates, state1, state2)

	if game.VialStates[0].Content[0] != 'a' || game.VialStates[1].Content[1] != 'b' {
		t.Error("Error in GameState internals. VialStates do not mach")
	}

	var change StateChange

	change.SourceVial[0] = state1
	change.DestVial[0] = state2

	for i := 0; i < VialSize; i++ {
		if change.SourceVial[0].Content[i] != 'a' || change.DestVial[0].Content[i] != 'b' {
			t.Error("Error in StateChange internals. VialStates Content do not match")
		}
	}

}

var goodstate GameState

var vial0 VialState = VialState{
	Content: [4]byte{'A', 'A', 'A', 'A'},
	Head:    0,
	TopLen:  4,
}

var vial1 VialState = VialState{
	Content: [4]byte{'B', 'B', 'B', 'B'},
	Head:    0,
	TopLen:  4,
}

var vial2 VialState = VialState{
	Content: [4]byte{'C', 'C', 'C', 'C'},
	Head:    0,
	TopLen:  4,
}
var vial3 VialState = VialState{
	Content: [4]byte{'A', 'B', 'A', 'C'},
	Head:    4,
	TopLen:  0,
}

var vial4 VialState = VialState{
	Content: [4]byte{'A', 'A', 'A', 'C'},
	Head:    4,
	TopLen:  0,
}

func initGoodState() {
	goodstate.VialStates = append(goodstate.VialStates, vial0, vial1, vial2, vial3, vial4)
}

/*
[A,A,B,C] -> [E,E,B,C]
[E,E,A,C] -> [A,A,A,C]
*/
func TestValidateStateChange(t *testing.T) {
	source0 := VialState{
		Content: [4]byte{'A', 'A', 'B', 'C'},
		Head:    0,
		TopLen:  2,
	}

	source1 := VialState{
		Content: [4]byte{'A', 'A', 'B', 'C'},
		Head:    2,
		TopLen:  1,
	}

	dest0 := VialState{
		Content: [4]byte{'C', 'C', 'A', 'C'},
		Head:    2,
		TopLen:  1,
	}

	dest1 := VialState{
		Content: [4]byte{'A', 'A', 'A', 'C'},
		Head:    0,
		TopLen:  3,
	}

	change := StateChange{
		SourceVial: []VialState{source0, source1},
		DestVial:   []VialState{dest0, dest1},
	}

	if !ValidateStateChange(&change) {
		t.Error("ValidateStateChange wrongly reported a correct change")
	}

	// TESTING ERROR IN DESTINATION
	wrongdest := dest1

	wrongdest.Head = 1
	change.DestVial[1] = wrongdest

	if ValidateStateChange(&change) {
		t.Error("ValidateStateChange did not detect a wrong dest head")
	}

	wrongdest.TopLen = 4
	change.DestVial[1] = wrongdest

	if ValidateStateChange(&change) {
		t.Error("ValidateStateChange did not detect a wrong toplen")
	}

	//TESTING ERROR IN SOURCE

	wrongsource := source0

	wrongsource.Content[0] = 'B'
	wrongsource.Content[1] = 'B'

	change.DestVial[1] = dest1

	change.SourceVial[0] = wrongsource

	if ValidateStateChange(&change) {
		t.Error("ValidateStateChange did not detect a wrong source content")
	}

	wrongsource = source0
	wrongsource.Head = 2
	change.SourceVial[0] = wrongsource

	if ValidateStateChange(&change) {
		t.Error("ValidateStateChange did not detect a wrong source head")
	}

	wrongsource = source1

	wrongsource.Head = 0
	change.SourceVial[0] = source0
	change.SourceVial[1] = wrongsource

	if ValidateStateChange(&change) {
		t.Error("ValidateStateChange did not detect a wrong source head post move")
	}
}

func TestCheckWin(t *testing.T) {
	initGoodState()

	gamestate := goodstate

	if !CheckWin(&gamestate) {
		t.Error("CheckWin did not recognize a real win")
	}

	wrongvial := vial4
	wrongvial.Head = 1
	gamestate.VialStates[4] = wrongvial

	if CheckWin(&gamestate) {
		t.Error("CheckWin did not recognize an half full vial")
	}

	gamestate.VialStates[4] = vial4

	wrongvial = vial0
	wrongvial.Content[2] = 'C'
	gamestate.VialStates[0] = wrongvial

	if CheckWin(&gamestate) {
		t.Error("CheckWin did not recognize an uneven vial")
	}

}

func TestStateHash(t *testing.T) {
	initGoodState()
	gamestate := goodstate

	var hash string = gamestate.StateHash()
	var expected string = "AAAABBBBCCCCEEEEEEEE"

	if hash != expected {
		print(gamestate.StateHash())
		t.Errorf("StateHash returned a wrong hash. Output: %s. Expected: %s", hash, expected)
	}
}
