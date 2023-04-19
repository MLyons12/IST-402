//Matthew Lyons and Avinash Sookram
//Barebone Structures for Enigma Machine in GoLang

package EnigmaMachine

type Rotor struct {
	wiring      []int // wiring of the rotor
	offset      int   // current offset of the rotor
	notch       int   // position of the rotor notch
	turnoverPos int   // position at which the next rotor is turned over
}

type Reflector struct {
	wiring [26]int // wiring of the reflector
}

type Plugboard struct {
	pairs map[int]int // pairs of letters to swap
}

type InputRotor struct {
	wiring [26]int
}
type RotorSet struct {
	offset  int
	writing [26]int
}

type EnigmaMachine struct {
	rotors     []*Rotor    // rotors in the machine
	reflector  *Reflector  // reflector used in the machine
	plugboard  *Plugboard  // plugboard used in the machine
	inputRotor *InputRotor //rotor used for input
	rotorSet   *RotorSet   // steckering used in the machine
}
