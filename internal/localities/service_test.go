package localities

import (
	"errors"
	"testing"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/stretchr/testify/assert"
)

func Test_GetLocalityInfo_OK(t *testing.T) {

	expectedResult := []LocalityInfo{
		{
			LocalityId:   "11065001",
			LocalityName: "Santos",
			SellersCount: 1,
		},
	}

	mockRepository := mockLocalityRepository{
		result:         expectedResult,
		err:            nil,
		findLocalityId: true,
	}

	service := NewService(mockRepository)
	result, err := service.GetLocalityInfo("11065001")

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, len(expectedResult), len(result))
}

func Test_GetLocalityInfo_NotFoundError(t *testing.T) {

	expectedResult := []LocalityInfo{}
	expectedError := LocalityNotFoundError

	mockRepository := mockLocalityRepository{
		result:         expectedResult,
		err:            expectedError,
		findLocalityId: false,
	}

	service := NewService(mockRepository)
	result, err := service.GetLocalityInfo("11065000")

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError, err)
}

func Test_GetLocalityInfo_InternalServerError(t *testing.T) {

	expectedResult := []LocalityInfo{}
	expectedError := errors.New("internal server error")

	mockRepository := mockLocalityRepository{
		result:         expectedResult,
		err:            expectedError,
		findLocalityId: true,
	}

	service := NewService(mockRepository)
	result, err := service.GetLocalityInfo("11065000")

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError, err)
}

func Test_CreateLocality_OK(t *testing.T) {

	expectedResult := []LocalityInfo{
		{
			LocalityId:   "11065001",
			LocalityName: "Santos",
			SellersCount: 1,
		},
	}

	mockRepository := mockLocalityRepository{
		result:         expectedResult,
		err:            nil,
		findLocalityId: true,
	}

	service := NewService(mockRepository)
	result, err := service.GetLocalityInfo("11065001")

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, len(expectedResult), len(result))
}

func Test_Create_SameLocalityIdError(t *testing.T) {

	expectedResult := database.Locality{}
	expectedError := ExistsLocalityId

	mockRepository := mockLocalityRepository{
		result:         expectedResult,
		err:            expectedError,
		findLocalityId: true,
	}

	service := NewService(mockRepository)
	result, err := service.Create("11065000", "Santos", 1)

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError, err)
}

func Test_Create_ProvinceIdNotExistsError(t *testing.T) {

	expectedResult := database.Locality{}
	expectedError := ExistsProvinceIdError

	mockRepository := mockLocalityRepository{
		result:         expectedResult,
		err:            expectedError,
		findLocalityId: false,
		findProvinceId: false,
	}

	service := NewService(mockRepository)
	result, err := service.Create("11065000", "Santos", 1)

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError, err)
}

func Test_Create_OK(t *testing.T) {

	expectedResult := database.Locality{
		Id:         "11065000",
		Name:       "Santos",
		ProvinceId: 1,
	}

	mockRepository := mockLocalityRepository{
		result:         expectedResult,
		err:            nil,
		findLocalityId: false,
		findProvinceId: true,
	}

	service := NewService(mockRepository)
	result, err := service.Create("11065000", "Santos", 1)

	assert.Equal(t, expectedResult, result)
	assert.Nil(t, err)
}
