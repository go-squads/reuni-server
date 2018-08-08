package organization

import (
	"github.com/go-squads/reuni-server/appcontext"
)

type processor interface {
	createNewOrganizationProcessor(orginizationName string, userId int64) error
}

type mainProcessor struct{}

var activeRepository repository

func getRepository() repository {
	if activeRepository == nil {
		activeRepository = initRepository(appcontext.GetHelper())
	}
	return activeRepository
}

func (s *mainProcessor) createNewOrganizationProcessor(organizationName string, userId int64) error {
	id, err := getRepository().createNewOrganization(organizationName)
	if err != nil {
		return err
	}
	return getRepository().addUser(id, userId, "Admin")
}
