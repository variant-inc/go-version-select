package handlers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessVersions(t *testing.T) {
	mockCtx := context.Background()
	// failed at parse version
	mockVersionList := "test.tst.r"
	_, err := ProcessVersions(mockCtx, mockVersionList, ">=1.0.0")
	assert.NotNil(t, err)

	// fail at parse version constraint
	mockVersionList = "1.0.3,2.5.4"
	mockVersionConstraint := "rrr.ttt.r"
	_, err = ProcessVersions(mockCtx, mockVersionList, mockVersionConstraint)
	assert.NotNil(t, err)

	// no match found
	mockVersionConstraint = "=1.0.0"
	result, err := ProcessVersions(mockCtx, mockVersionList, mockVersionConstraint)
	assert.Nil(t, err)
	assert.Empty(t, result)

	// match found
	mockVersionConstraint = "^1.0.0"
	result, err = ProcessVersions(mockCtx, mockVersionList, mockVersionConstraint)
	assert.Nil(t, err)
	assert.Equal(t, result, "1.0.3")
}
