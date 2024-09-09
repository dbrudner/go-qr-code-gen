package ticket

type Ticket struct {
	ID        string
	UserID    int
	CreatedAt string
	UpdatedAt string
	Content   string
	SiteID    string
}

func (t Ticket) GetID() string {
	return t.ID
}

type IDHolder interface {
	GetID() string
}
