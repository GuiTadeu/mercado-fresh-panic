package sellers

import (
	"errors"
	"testing"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/stretchr/testify/assert"
)

func Test_FindAll_OK(t *testing.T) {

	expectedResult := []db.Seller{
		{
			Id:          1,
			Cid:         1,
			CompanyName: "NIKE",
			Telephone:   "13997780814",
			Address:     "Rua Goiás, 37",
		},
		{
			Id:          2,
			Cid:         2,
			CompanyName: "adidas",
			Telephone:   "13997780813",
			Address:     "Rua Espirito Santo, 132",
		},
	}

	mockRepository := mockSellerRepository{
		result: expectedResult,
		err:    nil,
	}

	service := NewService(mockRepository)
	result, err := service.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, len(expectedResult), len(result))
}

func Test_FindAll_Error(t *testing.T) {

	expectedResult := []db.Seller{}
	expectedError := errors.New("Deu ruim no banco!")

	mockRepository := mockSellerRepository{
		result: expectedResult,
		err:    expectedError,
	}

	service := NewService(mockRepository)
	result, err := service.FindAll()

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError, err)
}

func Test_FindOne_Non_Existent(t *testing.T) {

	expectedResult := db.Seller{}

	nonExistentId := 2

	mockRepository := mockSellerRepository{
		result: expectedResult,
		err:    SellerNotFoundError,
	}

	service := NewService(mockRepository)
	result, err := service.FindOne(uint64(nonExistentId))

	assert.Equal(t, err, SellerNotFoundError)
	assert.Equal(t, db.Seller{}, result)
}

func Test_FindOne_Existent(t *testing.T) {

	expectedResult := db.Seller{
		Id:          1,
		Cid:         1,
		CompanyName: "NIKE",
		Telephone:   "13997780814",
		Address:     "Rua Goiás, 37",
	}

	existentId := 1

	mockRepository := mockSellerRepository{
		result:  expectedResult,
		err:     nil,
		getByID: expectedResult,
	}

	service := NewService(mockRepository)
	result, err := service.FindOne(uint64(existentId))

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Create_OK(t *testing.T) {

	expectedResult := db.Seller{
		Id:          1,
		Cid:         1,
		CompanyName: "NIKE",
		Telephone:   "13997780814",
		Address:     "Rua Goiás, 37",
	}

	mockRepository := mockSellerRepository{
		result:          expectedResult,
		err:             nil,
		existsSellerCid: false,
	}

	service := NewService(mockRepository)
	result, err := service.Create(expectedResult.Cid, expectedResult.CompanyName, expectedResult.Address, expectedResult.Telephone)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Create_Conflict(t *testing.T) {

	expectedResult := db.Seller{}

	mockRepository := mockSellerRepository{
		result:          expectedResult,
		err:             ExistsSellerCodeError,
		existsSellerCid: true,
	}

	service := NewService(mockRepository)
	result, err := service.Create(expectedResult.Cid, expectedResult.CompanyName, expectedResult.Address, expectedResult.Telephone)

	assert.Equal(t, err, ExistsSellerCodeError)
	assert.Equal(t, expectedResult, result)
}

func Test_Update_OK(t *testing.T) {

	sellerToUpdate := db.Seller{
		Id:          1,
		Cid:         1,
		CompanyName: "adidas",
		Telephone:   "13997782222",
		Address:     "Rua Goiania, 37",
	}

	expectedResult := db.Seller{
		Id:          1,
		Cid:         2,
		CompanyName: "NIKE",
		Telephone:   "13997780814",
		Address:     "Rua Goiás, 37",
	}

	mockRepository := mockSellerRepository{
		result:  expectedResult,
		err:     nil,
		getByID: sellerToUpdate,
	}

	service := NewService(mockRepository)
	result, err := service.Update(expectedResult.Id, expectedResult.Cid, expectedResult.CompanyName, expectedResult.Address, expectedResult.Telephone)

	assert.Equal(t, nil, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Update_Non_Existent(t *testing.T) {

	expectedResult := db.Seller{}

	mockRepository := mockSellerRepository{
		result:          expectedResult,
		err:             SellerNotFoundError,
		existsSellerCid: false,
		getByID:         db.Seller{},
	}

	service := NewService(mockRepository)
	result, err := service.Update(expectedResult.Id, expectedResult.Cid, expectedResult.CompanyName, expectedResult.Address, expectedResult.Telephone)

	assert.Equal(t, SellerNotFoundError, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Update_Cid_Existent_Error(t *testing.T) {

	expectedResult := db.Seller{}

	sellerGetByID := db.Seller{
		Id:          1,
		Cid:         1,
		CompanyName: "NIKE",
		Telephone:   "13997780814",
		Address:     "Rua Goiás, 37",
	}

	mockRepository := mockSellerRepository{
		result:          expectedResult,
		err:             ExistsSellerCodeError,
		existsSellerCid: true,
		getByID:         sellerGetByID,
	}

	service := NewService(mockRepository)
	result, err := service.Update(sellerGetByID.Id, sellerGetByID.Cid, sellerGetByID.CompanyName, sellerGetByID.Address, expectedResult.Telephone)

	assert.Equal(t, ExistsSellerCodeError, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Delete_Ok(t *testing.T) {

	mockRepository := mockSellerRepository{
		result: db.Seller{},
		err:    nil,
	}

	service := NewService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, nil, err)
}

func Test_Delete_Non_Existent(t *testing.T) {

	mockRepository := mockSellerRepository{
		result: db.Seller{},
		err:    SellerNotFoundError,
	}

	service := NewService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, SellerNotFoundError, err)
}
