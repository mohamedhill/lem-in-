package funcs

type Room struct {
	Name    string
	X       int
	Y       int
	Links   []*Room
	IsStart bool
	IsEnd   bool
}
type Graph struct {
	Rooms map[string]*Room
	Start *Room
	End   *Room
	Ants  int
}
