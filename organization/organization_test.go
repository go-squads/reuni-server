package organization

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRoleValidShouldReturnFalse(t *testing.T) {
	member := &Member{
		Role: "adasdas",
	}
	assert.False(t, member.isRoleValid())
}

func TestIsRoleValidShouldReturnTrue(t *testing.T) {
	member := &Member{
		Role: "Admin",
	}
	assert.True(t, member.isRoleValid())
}
