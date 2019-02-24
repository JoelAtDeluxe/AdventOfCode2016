package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	UP   = 1
	DOWN = -1
)

func main() {
	logic()
}

type Generator struct {
	Material string
	Floor    int
}

type Microchip struct {
	Material string
	Floor    int
}

type MCGenPair struct {
	Material string
	NickName string
	MCFloor  int
	GenFloor int
}

type State struct {
	Generators []Generator
	Mircochips []Microchip
	Pairs      []MCGenPair
	PrintOrder []int
	Elevator   int
}

func (s *State) nicksOnElevator() []string {
	return s.nicksOnFloor(s.Elevator)
}

func (mgp *MCGenPair) asMcNick() string {
	return fmt.Sprintf("%vM", mgp.NickName)
}

func (mgp *MCGenPair) asGenNick() string {
	return fmt.Sprintf("%vG", mgp.NickName)
}

func (mgp *MCGenPair) getFloorRefByNick(s string) *int {
	if s == mgp.asGenNick() {
		return &mgp.GenFloor
	} else if s == mgp.asMcNick() {
		return &mgp.MCFloor
	}
	return nil
}

func (s *State) nicksOnFloor(floor int) []string {
	rtn := []string{}
	for _, p := range s.Pairs {
		if p.GenFloor == floor {
			rtn = append(rtn, p.asGenNick())
		}
		if p.MCFloor == floor {
			rtn = append(rtn, p.asMcNick())
		}
	}
	return rtn
}

func (s *State) isFloorSafe(floor int) bool {
	hasGenerator := false
	for _, p := range s.Pairs {
		if p.GenFloor == floor {
			hasGenerator = true
			break
		}
	}
	if hasGenerator {
		for _, p := range s.Pairs {
			if p.MCFloor == floor && p.GenFloor != floor {
				return false
			}
		}
	}
	return true
}

func makeGen(material string, floor int) Generator {
	return Generator{material, floor}
}

func logic() {
	// opting to just encode this problem rather than parse it out
	state := State{
		Generators: []Generator{{"polonium", 1}, {"promethium", 1}, {"thulium", 1}, {"ruthenium", 1}, {"cobalt", 1}},
		Mircochips: []Microchip{{"polonium", 2}, {"promethium", 2}, {"thulium", 1}, {"ruthenium", 1}, {"cobalt", 1}},
		Pairs: []MCGenPair{
			{"polonium", "Po", 2, 1},
			{"promethium", "Pr", 2, 1},
			{"thulium", "Th", 1, 1},
			{"ruthenium", "Ru", 1, 1},
			{"cobalt", "Co", 1, 1},
		},
		Elevator: 1,
	}
	state.PrintOrder = []int{4, 0, 1, 3, 2} // TODO: sort names alphabetically, and record the index

	state.drawState()

	checkMove := func(err error) {
		if err != nil {
			fmt.Println(err)
		}
	}

	checkMove(state.move("CoM", "RuM", UP))
	checkMove(state.move("CoM", "RuM", UP))
	checkMove(state.move("CoM", "", DOWN))
	checkMove(state.move("CoM", "PrM", UP))
	checkMove(state.move("CoM", "RuM", UP))
	checkMove(state.move("CoM", "", DOWN))
	checkMove(state.move("CoM", "", DOWN))
	checkMove(state.move("CoM", "", DOWN))

	checkMove(state.move("PoG", "PrG", UP))
	checkMove(state.move("PoG", "PrG", UP))
	checkMove(state.move("PoG", "", DOWN))
	checkMove(state.move("PoG", "", DOWN))

	checkMove(state.move("PoG", "RuG", UP))
	checkMove(state.move("PoG", "RuG", UP))
	checkMove(state.move("PoG", "RuG", UP))
	checkMove(state.move("PoG", "", DOWN))




	// checkMove(state.move("PG", "", DOWN))

	// checkMove(state.move("PoG", "", DOWN))
	//checkMove(state.move("PoG", "PoM", UP))



	// checkMove(state.move("PoG", "PrG", UP))

	// checkMove(state.move("CoM", "CoG", UP))
	// checkMove(state.move("PoM", "", DOWN))

	// checkMove(state.move("C", "", UP))




	state.drawState()

	// fmt.Println(state)
}

func (s *State) move(left, right string, direction int) error {
	if (direction == UP && s.Elevator == 4) || (direction == DOWN && s.Elevator == 1) {
		return errors.New("Can't move in that direction")
	}
	if len(left) == 0 && len(right) == 0 {
		return errors.New("Can't move the elevator without an element")
	}
	elements := strings.Join(s.nicksOnElevator(), " ")
	if !strings.Contains(elements, left) || !strings.Contains(elements, right) {
		return errors.New("Element does not exist on this level")
	}

	for i := range s.Pairs {
		leftFloor := s.Pairs[i].getFloorRefByNick(left)
		if leftFloor != nil {
			(*leftFloor) += direction
		}

		rightFloor := s.Pairs[i].getFloorRefByNick(right)
		if rightFloor != nil {
			(*rightFloor) += direction
		}
	}
	s.Elevator += direction
	if !s.isFloorSafe(s.Elevator) || !s.isFloorSafe(s.Elevator - direction) {
		return errors.New("Some state is now in error!")
	}

	return nil
}

// 1 CoM, _ u
// 2 CoM, PoM,  d
// 3 CoM d
// 4 RuM, ThM, u
// 5

func (s *State) drawState() {
	empty := ".  "
	floorState := make(map[int][]string, 4)

	valOrEmpty := func(i, j int, name string) string {
		if i == j {
			return name
		}
		return empty
	}

	for i := 4; i > 0; i-- {
		ele := valOrEmpty(i, s.Elevator, "E  ")
		floorState[i] = append(floorState[i], fmt.Sprintf("F%v  %v", i, ele))
	}

	addElementToFloors := func(nickname string, onFloor int) {
		for i := 4; i > 0; i-- {
			floorState[i] = append(floorState[i], valOrEmpty(i, onFloor, nickname))
		}
	}

	for _, pairIndex := range s.PrintOrder {
		pair := &s.Pairs[pairIndex]
		addElementToFloors(pair.asGenNick(), pair.GenFloor)
		addElementToFloors(pair.asMcNick(), pair.MCFloor)
	}

	fmt.Println()
	for i := 4; i >= 0; i-- {
		if s.isFloorSafe(i) {
			fmt.Println(asColor(strings.Join(floorState[i], " "), LightGreen))
		} else {
			fmt.Println(asColor(strings.Join(floorState[i], " "), Red))
		}
	}
}
