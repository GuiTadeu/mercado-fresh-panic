package localities

import "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"

type mockLocalityRepository struct {
	result         any
	err            error
	findProvinceId bool
	findLocalityId bool
}

func (m mockLocalityRepository) FindCid(cid uint64) bool {
	return m.err != nil
}

func (m mockLocalityRepository) Create(localityId string, localityName string, provinceId uint64) (database.Locality, error) {
	if m.err != nil || m.findLocalityId || !m.findProvinceId {
		return database.Locality{}, m.err
	}
	return m.result.(database.Locality), nil
}

func (m mockLocalityRepository) GetLocalityInfo(localityId string) ([]LocalityInfo, error) {
	if m.err != nil {
		return []LocalityInfo{}, m.err
	}
	return m.result.([]LocalityInfo), nil
}

func (m mockLocalityRepository) FindLocalityId(localityId string) bool {
	return m.findLocalityId
}

func (m mockLocalityRepository) ExistsProvinceId(provinceId uint64) bool {
	return m.findProvinceId
}
