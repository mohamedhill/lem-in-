package funcs

type Room struct {
	Name    string
	x       int
	y       int
	links   []*Room
	IsStart bool
	IsEnd   bool
}
type Graph struct {
	Rooms map[string]*Room
	Start *Room
	End   *Room
	Ants  int
}
