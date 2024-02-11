package paldex

type Pal struct {
	Id          int           `json:"id"`
	Key         string        `json:"key"`
	Image       string        `json:"image"`
	Name        string        `json:"name"`
	Wiki        string        `json:"wiki"`
	Types       []string      `json:"types"`
	ImageWiki   string        `json:"image_wiki"`
	Suitability []Suitability `json:"suitability"`
	Drops       []string      `json:"drops"`
	Aura        Aura          `json:"aura"`
	Skills      []Skill       `json:"skills"`
	Stats       Stats         `json:"stats"`
	Asset       string        `json:"asset"`
	Genus       string        `json:"genus"`
	Rarity      int           `json:"rarity"`
	Price       int           `json:"price"`
	Size        string        `json:"size"`
	Maps        Maps          `json:"maps"`
}

type Suitability struct {
	Type  string `json:"type"`
	Image string `json:"image"`
	Level string `json:"level"`
}

type Aura struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Tech        string `json:"tech"`
}

type Skill struct {
	Level       int    `json:"level"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Cooldown    int    `json:"cooldown"`
	Power       int    `json:"power"`
	Description string `json:"description"`
}

type Stats struct {
	Hp      int     `json:"hp"`
	Attack  Attacks `json:"attack"`
	Defense int     `json:"defense"`
	Speed   Speeds  `json:"speed"`
	Stamina int     `json:"stamina"`
	Support int     `json:"support"`
	Food    int     `json:"food"`
}

type Attacks struct {
	Melee  int `json:"melee"`
	Ranged int `json:"ranged"`
}

type Speeds struct {
	Ride int `json:"ride"`
	Run  int `json:"run"`
	Walk int `json:"walk"`
}

type Maps struct {
	Day   string `json:"day"`
	Night string `json:"night"`
}
