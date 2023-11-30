package main

type Tile struct {
	Type TileType `json:"type"`
	X    int      `json:"x"`
	Y    int      `json:"y"`
}

func NewTile() Tile {
	return Tile{
		Type: WATER,
	}
}

type TileType int

const (
	WATER TileType = iota + 1
	SAND
	GRASS
	ROCK
)
