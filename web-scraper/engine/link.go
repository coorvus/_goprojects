package engine

var Links []Link

type Link struct {
	Url    string `json:"url" validate:"required"`
	IsDead bool   `json:"is_dead"`
}

func (l *Link) Scrape() {
	_, err := GetSource(l.Url)
	if err != nil {
		l.IsDead = true
	}
}
