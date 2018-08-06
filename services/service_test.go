package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsServiceEmptyReturnTrue(t *testing.T) {
	service := service{}
	assert.True(t, service.IsEmpty())
}
func TestIsServiceEmptyReturnFalse(t *testing.T) {
	service := service{Name: "hello-world"}
	assert.False(t, service.IsEmpty())
}

func TestIsServiceTokenEmptyReturnTrue(t *testing.T) {
	token := serviceToken{}
	assert.True(t, token.IsEmpty())
}
func TestIsServiceTokenEmptyReturnFalse(t *testing.T) {
	token := serviceToken{Token: "loremipsum"}
	assert.False(t, token.IsEmpty())
}

func TestSliceEmptyReturnTrue(t *testing.T) {
	var slices []service
	assert.True(t, isSliceEmpty(slices))
}

func TestSliceEmptyReturnTrueWhenAllDataEmpty(t *testing.T) {
	slices := []service{
		MockServiceStruct(0, ""),
		MockServiceStruct(0, ""),
	}
	assert.True(t, isSliceEmpty(slices))
}

func TestSliceEmptyReturnFalseWhenDataNotEmpty(t *testing.T) {
	slices := []service{
		MockServiceStruct(1, "go-pay"),
		MockServiceStruct(2, "go-ride"),
	}
	assert.False(t, isSliceEmpty(slices))
}
