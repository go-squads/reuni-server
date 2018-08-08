package organization

import (
	"github.com/go-squads/reuni-server/helper"
)

type repository interface {
	createNewOrganization(organization_name string) error
}

type mainRepository struct {
	execer helper.QueryExecuter
}

const (
	createNewOrganizationQuery = "INSERT INTO organization(name) VALUES($1)"
)

func initRepository(execer helper.QueryExecuter) *mainRepository {
	return &mainRepository{
		execer: execer,
	}
}

func (s *mainRepository) createNewOrganization(organization_name string) error {
	_, err := s.execer.DoQuery(createNewOrganizationQuery, organization_name)
	return err
}
