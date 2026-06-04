package domain

import (
	"errors"
	"time"
)

type TipoFigurinha string

const (
	TipoComum        TipoFigurinha = "comum"
	TipoBrilhante    TipoFigurinha = "brilhante"
	TipoLegendOuro   TipoFigurinha = "legend_ouro"
	TipoLegendBronze TipoFigurinha = "legend_bronze"
)

type PosicaoJogadorFigurinha string

const (
	PosicaoGoleiro   PosicaoJogadorFigurinha = "goleiro"
	PosicaoZagueiro  PosicaoJogadorFigurinha = "zagueiro"
	PosicaoMeioCampo PosicaoJogadorFigurinha = "meio-campista"
	PosicaoAtacante  PosicaoJogadorFigurinha = "atacante"
)

type Figurinha struct {
	ID        uint                    `json:"id"        gorm:"primaryKey"`
	Numero    string                  `json:"number"    gorm:"not null"`
	Tipo      TipoFigurinha           `json:"type"      gorm:"not null"`
	Posicao   PosicaoJogadorFigurinha `json:"position"  gorm:"not null"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
}

type CreateFigurinhaRequest struct {
	Numero  string                  `json:"number"   binding:"required"`
	Tipo    TipoFigurinha           `json:"type"     binding:"required"`
	Posicao PosicaoJogadorFigurinha `json:"position" binding:"required"`
}

type UpdateFigurinhaRequest struct {
	Numero  *string                  `json:"number,omitempty"`
	Tipo    *TipoFigurinha           `json:"type,omitempty"`
	Posicao *PosicaoJogadorFigurinha `json:"position,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

var ErrFigurinhaNaoEncontrada = errors.New("figurinha não encontrado")
