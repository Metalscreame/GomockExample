package services

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/metalscreame/GomockExample/src/models"
	"github.com/metalscreame/GomockExample/src/mocks"
	"errors"
)

var userToSave = models.User{"MockedUserToSave"}

func TestSaver(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserPersistence := mocks.NewMockUserPersistence(mockCtrl)
	service := NewService(mockUserPersistence)

	mockUserPersistence.EXPECT().Save(userToSave).Return(nil)

	err := service.Saver(userToSave)
	if err != nil {
		t.Fatalf("Want %v but got %v", nil, err)
	}
}


func TestSaverWithError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserPersistence := mocks.NewMockUserPersistence(mockCtrl)
	service := NewService(mockUserPersistence)

	mockUserPersistence.EXPECT().Save(userToSave).Return(errors.New("error"))

	// Do the same as the example above
	//mockUserPersistance.EXPECT().Save(gomock.Any()).Return(errors.New("error"))

	err := service.Saver(userToSave)
	if err == nil {
		t.Fatalf("Want %v but got %v", errors.New("error"), nil)
	}
}

/*
Call order

callFirst := mockUserPersistence.EXPECT().Save(gomock.Any())
callA := mockUserPersistence.EXPECT().Delete(gomock.Any()).After(callFirst)
callB := mockUserPersistence.EXPECT().Update(gomock.Any()).After(callFirst)

or

gomock.InOrder(
	mockUserPersistence.EXPECT().Save(gomock.Any()),
	mockUserPersistence.EXPECT().Delete(gomock.Any()).After(callFirst),
	mockUserPersistence.EXPECT().Update(gomock.Any()).After(callFirst),
)


Mock actions

	mockUserPersistence.EXPECT().Save(gomock.Any()).Return(errors.New("error")).Do(func(user models.User) {
		fmt.Println(user.Name)
	})


Go generate

package models

//go:generate mockgen -destination=src/mocks/mock_UserPersistence.go -package=mocks GomockExample/src/models UserPersistence


 */