package organization

import (
	"github.com/go-squads/reuni-server/appcontext"
)

var activeRepository repository

func getRepository() repository {
	if activeRepository == nil {
		activeRepository = initRepository(appcontext.GetHelper())
	}
	return activeRepository
}

func createNewOrganizationProcessor(organization_name string, userId int64) error {
	id, err := getRepository().createNewOrganization(organization_name)
	if err != nil {
		return err
	}
	return getRepository().addUser(id, userId, "Admin")
}
