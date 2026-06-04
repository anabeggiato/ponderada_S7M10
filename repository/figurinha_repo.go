package repository

import (
	"github.com/anabeggiato/pond3_M10S07/domain"
	"gorm.io/gorm"
)

type FigurinhaRepository interface {
	Create(figurinha *domain.Figurinha) error
	FindAll() ([]domain.Figurinha, error)
	FindByID(id uint) (*domain.Figurinha, error)
	Update(figurinha *domain.Figurinha) error
	Delete(id uint) error
}

type figurinhaRepositoryImpl struct {
	db *gorm.DB
}

func NewFigurinhaRepository(db *gorm.DB) FigurinhaRepository {
	return &figurinhaRepositoryImpl{db: db}
}

func (r *figurinhaRepositoryImpl) Create(figurinha *domain.Figurinha) error {
	return r.db.Create(figurinha).Error
}

func (r *figurinhaRepositoryImpl) FindAll() ([]domain.Figurinha, error) {
	var figurinhas []domain.Figurinha
	err := r.db.Find(&figurinhas).Error
	return figurinhas, err
}

func (r *figurinhaRepositoryImpl) FindByID(id uint) (*domain.Figurinha, error) {
	var figurinha domain.Figurinha
	err := r.db.First(&figurinha, id).Error
	if err != nil {
		return nil, err
	}
	return &figurinha, nil
}

func (r *figurinhaRepositoryImpl) Update(figurinha *domain.Figurinha) error {
	return r.db.Save(figurinha).Error
}

func (r *figurinhaRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Figurinha{}, id).Error
}
