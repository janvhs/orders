package article

type State = int

const (
	Open State = iota
	Ordered
	Rejected
)

type Article struct {
	ID int64

	Name       string
	Amount     int64
	URL        string
	Price      float64
	CostCentre string
	Status     State

	// TODO: Replace with User / Abteilung
	For string
}
