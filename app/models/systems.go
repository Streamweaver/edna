package models

import (
	"encoding/json"
	"math"
	"os"
)

// System data about a star system.
type System struct {
	ID   uint32  `json:"id"`
	Name string  `json:"name"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Z    float64 `json:"z"`
}

// Distance to target system t in lightyears
func (s System) Distance(t System) float64 {
	return Distance(s, t)
}

// LoadSystems from the json export.
func LoadSystems() (*[]System, error) {
	var u = struct {
		Systems []System
	}{}
	f, err := os.Open("data/systems.json")
	if err != nil {
		return &u.Systems, err
	}
	decoder := json.NewDecoder(f)
	if err = decoder.Decode(&u.Systems); err != nil {
		return &u.Systems, err
	}

	return &u.Systems, nil
}

// BuildSysIndex is a hack to return a map of systems by ID number.
// TODO remove this after implementing GORM models.
func BuildSysIndex(systems []System) map[uint32]System {
	var idx map[uint32]System
	idx = make(map[uint32]System)
	for _, s := range systems {
		idx[s.ID] = s
	}
	return idx
}

// Distance between two star systems in light years.
func Distance(s1 System, s2 System) float64 {
	xd := s2.X - s1.X
	x := xd * xd
	yd := s2.Y - s1.Y
	y := yd * yd
	zd := s2.Z - s1.Z
	z := zd * zd
	return math.Sqrt(x + y + z)
}
