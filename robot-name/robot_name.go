package robotname

import (
	"math/rand"
	"time"
)

// Robot blabla
type Robot struct {
	name string
}

var robots = map[string]bool{}

// Name sets the name property of a Robot to a specific random value
func (r *Robot) Name() (string, error) {
	// if name is already in registry, return it
	if robots[r.name] {

		return r.name, nil
	}

	rand.Seed(time.Now().UTC().UnixNano())
	// otherwise generate new name
	name := createRobotName()
	for robots[name] {
		name = createRobotName()
	}

	r.name = name
	robots[name] = true

	return name, nil
}

// Reset again give the robot a new name
func (r *Robot) Reset() {
	r.name = ""
}

// createRobotName creates a robot string of the following format ^[A-Z]{2}\d{3}$
func createRobotName() string {
	rs := []rune{
		rune(rand.Intn(26) + 65),
		rune(rand.Intn(26) + 65),
		rune(rand.Intn(10) + 48),
		rune(rand.Intn(10) + 48),
		rune(rand.Intn(10) + 48),
	}

	return string(rs)
}
