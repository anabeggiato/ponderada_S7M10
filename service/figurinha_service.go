package service

import (
	"errors"

	"github.com/anabeggiato/pond3_M10S07/domain"
	"github.com/anabeggiato/pond3_M10S07/repository"
	"gorm.io/gorm"
)

var (
	ErrFigurinhaNotFound = errors.New("figurinha não encontrada")
)

type FigurinhaService interface {
	CreateFigurinha(req domain.CreateFigurinhaRequest) (*domain.Figurinha, error)
	ListFigurinhas() ([]domain.Figurinha, error)
	GetFigurinha(id uint) (*domain.Figurinha, error)
	UpdateFigurinha(id uint, req domain.UpdateFigurinhaRequest) (*domain.Figurinha, error)
	DeleteFigurinha(id uint) error
}

type figurinhaServiceImpl struct {
	repo repository.FigurinhaRepository
}

func NewFigurinhaService(repo repository.FigurinhaRepository) FigurinhaService {
	return &figurinhaServiceImpl{repo: repo}
}

func (s *figurinhaServiceImpl) CreateFigurinha(req domain.CreateFigurinhaRequest) (*domain.Figurinha, error) {
	figurinha := &domain.Figurinha{
		Numero:  req.Numero,
		Tipo:    req.Tipo,
		Posicao: req.Posicao,
	}
	if err := s.repo.Create(figurinha); err != nil {
		return nil, err
	}
	return figurinha, nil
}

func (s *figurinhaServiceImpl) ListFigurinhas() ([]domain.Figurinha, error) {
	return s.repo.FindAll()
}

func (s *figurinhaServiceImpl) GetFigurinha(id uint) (*domain.Figurinha, error) {
	figurinha, err := s.repo.FindByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrFigurinhaNotFound
	}
	return figurinha, err
}

func (s *figurinhaServiceImpl) UpdateFigurinha(id uint, req domain.UpdateFigurinhaRequest) (*domain.Figurinha, error) {
	figurinha, err := s.GetFigurinha(id)
	if err != nil {
		return nil, err
	}

	if req.Numero != nil {
		figurinha.Numero = *req.Numero
	}
	if req.Tipo != nil {
		figurinha.Tipo = *req.Tipo
	}
	if req.Posicao != nil {
		figurinha.Posicao = *req.Posicao
	}

	if err := s.repo.Update(figurinha); err != nil {
		return nil, err
	}
	return figurinha, nil
}

func (s *figurinhaServiceImpl) DeleteFigurinha(id uint) error {
	if _, err := s.GetFigurinha(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
