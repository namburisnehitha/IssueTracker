package service

import "github.com/namburisnehitha/IssueTracker/domain"

type LabelService struct {
	labelRepository domain.LabelRepository
}

func NewLabelService(labelrepository domain.LabelRepository) LabelService {
	return LabelService{
		labelRepository: labelrepository,
	}
}

func (l *LabelService) CreateLabel(id string, name string, description string, colour string) error {
	label, err := domain.NewLabel(id, name, description, colour)
	if err != nil {
		return err
	}
	return l.labelRepository.Save(label)
}

func (l *LabelService) GetById(id string) (domain.Label, error) {
	return l.labelRepository.GetById(id)
}

func (l *LabelService) GetByName(name string) (domain.Label, error) {
	return l.labelRepository.GetByName(name)
}

func (l *LabelService) GetByColour(colour string) ([]domain.Label, error) {
	return l.labelRepository.GetByColour(colour)
}

func (l *LabelService) UpdateLabel(label domain.Label) error {
	return l.labelRepository.UpdateLabel(label)
}

func (l *LabelService) DeleteLabel(label domain.Label) error {
	return l.labelRepository.DeleteLabel(label)
}

func (l *LabelService) LabelList() ([]domain.Label, error) {
	return l.labelRepository.LabelList()
}
