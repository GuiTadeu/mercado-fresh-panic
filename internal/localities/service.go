package localities

import (
	"errors"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

var (
	ExistsLocalityId      = errors.New("locality id already exists")
	ExistsProvinceIdError = errors.New("province id does not exist")
	LocalityNotFoundError = errors.New("locality not found")
)

type Service interface {
	Create(localityId string, localityName string, provinceId uint64) (database.Locality, error)
	GetLocalityInfo(localityId string) ([]LocalityInfo, error)
}

type service struct {
	repo Repository
}

func (s service) Create(localityId string, localityName string, provinceId uint64) (database.Locality, error) {

	isUsedLocalityId := s.repo.FindLocalityId(localityId)

	if isUsedLocalityId {
		return database.Locality{}, ExistsLocalityId
	}

	isProvinceIdValid := s.repo.ExistsProvinceId(provinceId)

	if !isProvinceIdValid {
		return database.Locality{}, ExistsProvinceIdError
	}

	localityData, err := s.repo.Create(localityId, localityName, provinceId)

	if err != nil {
		return database.Locality{}, err
	}
	return localityData, nil
}

func (s service) GetLocalityInfo(localityId string) ([]LocalityInfo, error) {
	isLocalityIdFound := s.repo.FindLocalityId(localityId)

	if !isLocalityIdFound {
		return []LocalityInfo{}, LocalityNotFoundError
	}

	localityInfo, err := s.repo.GetLocalityInfo(localityId)

	if err != nil {
		return []LocalityInfo{}, err
	}

	return localityInfo, nil
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}
