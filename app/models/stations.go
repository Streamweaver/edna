package models

type Station struct {
	Id           uint32 `json:"id"`
	SystemId     uint32 `json:"system_id"`
	Name         string `json:"name"`
	Blackmarket  bool   `json:"has_blackmarket"`
	LandingSize  int    `json:"max_landing_pad_size"`
	StarDistance int    `json:"distance_to_star"`
}
