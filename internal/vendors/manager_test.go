package vendors

import (
	"esusu-sample1/internal/vendors/mocks"
	"testing"
)

func TestList(t *testing.T) {
	// Arrange
	mockedDBConnector := &mocks.DBConnector{}
	mockedDBConnector.On("").
	mgr := NewManager(mockedDBConnector)

	// Act
	res, err := mgr.List()

	// Assert
	// res should be equal expected x
	// err should be expected nil

}
