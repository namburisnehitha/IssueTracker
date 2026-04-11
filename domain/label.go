package domain

type Label struct {
	Id          string
	Name        string
	Description string
	Colour      string
}

func NewLabel(Id string, Name string, Description string, Colour string) (Label, error) {
	if Name == "" {
		return Label{}, ErrInvalidLabelData
	}
	return Label{
		Id:          Id,
		Name:        Name,
		Description: Description,
		Colour:      Colour,
	}, nil
}

func (l *Label) UpdateName(Name string) {
	l.Name = Name
}

func (l *Label) UpdateColour(Colour string) {
	l.Colour = Colour
}
