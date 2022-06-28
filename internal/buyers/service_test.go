package buyers

import (
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Create_ok(t *testing.T) {
	expectedResult := db.Buyer{
		Id:           11,
		CardNumberId: "22",
		FirstName:    "Meli",
		LastName:     "Developers",
	}

	mockBuyerRepository := mockBuyerRepository{
		result:                  expectedResult,
		err:                     nil,
		existsBuyerCardNumberId: false,
	}

	service := NewBuyerService(mockBuyerRepository)
	result, err := service.Create("22", "Meli", "Developers")

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Create_Conflict(t *testing.T) {
	expectedResult := db.Buyer{
		Id:           11,
		CardNumberId: "22",
		FirstName:    "Meli",
		LastName:     "Developers",
	}

	mockBuyerRepository := mockBuyerRepository{
		result:                  expectedResult,
		err:                     ExistsBuyerCardNumberIdError,
		existsBuyerCardNumberId: true,
	}

	service := NewBuyerService(mockBuyerRepository)
	_, err := service.Create("22", "Meli", "Developers")

	assert.Equal(t, ExistsBuyerCardNumberIdError, err)
}

func Test_GetAll_Ok(t *testing.T) {

	expectResult := []db.Buyer{{}, {}, {}}

	mockRepository := mockBuyerRepository{
		result: expectResult,
		err:    nil,
	}

	service := NewBuyerService(mockRepository)
	result, err := service.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, expectResult, result)
}

func Test_Get_Id_Non_Existent(t *testing.T) {
	mockBuyerRepository := mockBuyerRepository{
		err:                     BuyerNotFoundError,
		existsBuyerCardNumberId: false,
	}

	service := NewBuyerService(mockBuyerRepository)
	_, err := service.Get(10)

	assert.Equal(t, BuyerNotFoundError, err)
}

func Test_Get_Ok(t *testing.T) {
	expectedResult := db.Buyer{
		Id:           11,
		CardNumberId: "22",
		FirstName:    "Meli",
		LastName:     "Developers",
	}

	mockBuyerRepository := mockBuyerRepository{
		getById: expectedResult,
		err:     nil,
	}

	service := NewBuyerService(mockBuyerRepository)
	result, err := service.Get(11)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Update_Ok(t *testing.T) {
	expectedResult := db.Buyer{
		Id:           11,
		CardNumberId: "22",
		FirstName:    "Meli",
		LastName:     "Developers",
	}

	getById := db.Buyer{
		Id:           11,
		CardNumberId: "20",
		FirstName:    "Name",
		LastName:     "LastName",
	}

	mockBuyerRepository := mockBuyerRepository{
		result:  expectedResult,
		err:     nil,
		getById: getById,
	}

	service := NewBuyerService(mockBuyerRepository)
	result, err := service.Update(11, "22", "Meli", "Developers")

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Update_Non_Existent(t *testing.T) {
	mockBuyerRepository := mockBuyerRepository{
		err: BuyerNotFoundError,
	}

	service := NewBuyerService(mockBuyerRepository)
	_, err := service.Update(11, "22", "Meli", "Developers")

	assert.Equal(t, BuyerNotFoundError, err)
}

func Test_Update_Conflict(t *testing.T) {
	getById := db.Buyer{
		Id:           11,
		CardNumberId: "20",
		FirstName:    "Name",
		LastName:     "LastName",
	}

	mockBuyerRepository := mockBuyerRepository{
		err:                     ExistsBuyerCardNumberIdError,
		getById:                 getById,
		existsBuyerCardNumberId: true,
	}

	service := NewBuyerService(mockBuyerRepository)
	_, err := service.Update(11, "22", "Meli", "Developers")

	assert.Equal(t, ExistsBuyerCardNumberIdError, err)
}

func Test_Delete_Ok(t *testing.T) {
	mockBuyerRepository := mockBuyerRepository{
		err: nil,
	}

	service := NewBuyerService(mockBuyerRepository)
	err := service.Delete(11)

	assert.Nil(t, err)
}

func Test_Delete_Non_Existent(t *testing.T) {
	mockBuyerRepository := mockBuyerRepository{
		err: BuyerNotFoundError,
	}

	service := NewBuyerService(mockBuyerRepository)
	err := service.Delete(11)

	assert.Equal(t, BuyerNotFoundError, err)
}
