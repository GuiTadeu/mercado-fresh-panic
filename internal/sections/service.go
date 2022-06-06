package sections

import db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"

type SectionService interface {
	GetAll() ([]db.Section, error)
	Get(id uint64) (db.Section, error)
	Create(id uint64, number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64) (db.Section, error)
	Update(id uint64, number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64) (db.Section, error)
	Delete(id uint64) error
	GetNextId() uint64
}

type sectionService struct {
	sectionRepository SectionRepository
}

func NewService(r SectionRepository) SectionService {
	return &sectionService{
		sectionRepository: r,
	}
}

func (s *sectionService) GetAll() ([]db.Section, error) {
	return s.sectionRepository.GetAll()
}

func (s *sectionService) Get(id uint64) (db.Section, error) {
	return s.sectionRepository.Get(id)
}

func (s *sectionService) Create(id uint64, number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64) (db.Section, error) {
	return s.sectionRepository.Create(id, number, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId)
}

func (s *sectionService) Update(id uint64, number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64) (db.Section, error) {
	return s.sectionRepository.Update(id, number, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId)
}

func (s *sectionService) Delete(id uint64) error {
	return s.sectionRepository.Delete(id)
}

func (s *sectionService) GetNextId() uint64 {
	sections, err := s.sectionRepository.GetAll()
	if err != nil {
		return 1
	}

	if len(sections) == 0 {
		return 1
	}

	return sections[len(sections)-1].Id + 1
}
