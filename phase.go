package cah

type Phase uint8

const (
	NotStarted Phase = iota
	SinnersPlaying
	CzarChoosingWinner
	Finished
)

var phases = [...]string{
	"Not started",
	"Sinners playing their cards",
	"Czar is choosing winner",
	"Finished",
}

func (p Phase) String() string {
	return phases[p]
}
