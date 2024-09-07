package ticket

type Ticket struct {
	ID      string
	Site_ID string
	User_ID string
	Content string
}

func (t Ticket) GetID() string {
	return t.ID
}

type IDHolder interface {
	GetID() string
}
