package service

import "github.com/namburisnehitha/IssueTracker/domain"

type ActivityService struct {
	activityRepository domain.ActivityRepository
}

func NewActivityService(activityRepository domain.ActivityRepository) ActivityService {
	return ActivityService{
		activityRepository: activityRepository,
	}
}

func (a *ActivityService) CreateActivity(id string, issueid string, userid string, description string, action domain.ActivityType) error {
	activity := domain.NewActivity(id, issueid, userid, description, action)
	return a.activityRepository.Save(activity)

}

func (a *ActivityService) GetById(id string) (domain.Activity, error) {
	return a.activityRepository.GetById(id)
}

func (a *ActivityService) GetByUserId(userid string) ([]domain.Activity, error) {
	return a.activityRepository.GetByUserId(userid)
}

func (a *ActivityService) GetByIssueId(issueid string) ([]domain.Activity, error) {
	return a.activityRepository.GetByIssueId(issueid)
}

func (a *ActivityService) GetByAction(action domain.ActivityType) ([]domain.Activity, error) {
	return a.activityRepository.GetByAction(action)
}

func (a *ActivityService) ActivityList() ([]domain.Activity, error) {
	return a.activityRepository.ActivityList()
}
