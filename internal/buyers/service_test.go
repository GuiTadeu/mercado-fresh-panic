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
	result, _ := service.Update(11, "22", "Meli", "Developers")

	assert.Equal(t, expectedResult, result)
}

func Test_Create_conflict(t *testing.T) {

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
