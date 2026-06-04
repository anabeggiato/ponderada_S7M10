package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/anabeggiato/pond3_M10S07/domain"
	"github.com/anabeggiato/pond3_M10S07/service"
	"github.com/gin-gonic/gin"
)

type FigurinhaHandler struct {
	svc service.FigurinhaService
}

func NewFigurinhaHandler(svc service.FigurinhaService) *FigurinhaHandler {
	return &FigurinhaHandler{svc: svc}
}

func (h *FigurinhaHandler) respond(c *gin.Context, data interface{}, err error) {
	if err == nil {
		status := http.StatusOK
		if c.Request.Method == http.MethodPost {
			status = http.StatusCreated
		}
		c.JSON(status, data)
		return
	}

	switch {
	case errors.Is(err, service.ErrFigurinhaNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno"})
	}
}

// Create cria uma figurinha.
// @Summary Criar figurinha
// @Description Cria uma nova figurinha a partir das informações enviadas no corpo da requisição.
// @Tags figurinhas
// @Accept json
// @Produce json
// @Param figurinha body domain.CreateFigurinhaRequest true "Dados da figurinha"
// @Success 201 {object} domain.Figurinha
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /figurinhas [post]
func (h *FigurinhaHandler) Create(c *gin.Context) {
	var req domain.CreateFigurinhaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	figurinha, err := h.svc.CreateFigurinha(req)
	h.respond(c, figurinha, err)
}

// List lista todas as figurinhas.
// @Summary Listar figurinhas
// @Description Retorna todas as figurinhas cadastradas.
// @Tags figurinhas
// @Produce json
// @Success 200 {array} domain.Figurinha
// @Failure 500 {object} domain.ErrorResponse
// @Router /figurinhas [get]
func (h *FigurinhaHandler) List(c *gin.Context) {
	figurinhas, err := h.svc.ListFigurinhas()
	h.respond(c, figurinhas, err)
}

// GetByID busca uma figurinha pelo ID.
// @Summary Buscar figurinha por ID
// @Description Retorna uma figurinha específica a partir do ID enviado na rota.
// @Tags figurinhas
// @Produce json
// @Param id path int true "ID da figurinha"
// @Success 200 {object} domain.Figurinha
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /figurinhas/{id} [get]
func (h *FigurinhaHandler) GetByID(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	figurinha, err := h.svc.GetFigurinha(id)
	h.respond(c, figurinha, err)
}

// Update atualiza uma figurinha.
// @Summary Atualizar figurinha
// @Description Atualiza os campos enviados de uma figurinha já cadastrada.
// @Tags figurinhas
// @Accept json
// @Produce json
// @Param id path int true "ID da figurinha"
// @Param figurinha body domain.UpdateFigurinhaRequest true "Dados para atualizar"
// @Success 200 {object} domain.Figurinha
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /figurinhas/{id} [put]
func (h *FigurinhaHandler) Update(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var req domain.UpdateFigurinhaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	figurinha, err := h.svc.UpdateFigurinha(id, req)
	h.respond(c, figurinha, err)
}

// Delete remove uma figurinha.
// @Summary Deletar figurinha
// @Description Remove uma figurinha a partir do ID enviado na rota.
// @Tags figurinhas
// @Produce json
// @Param id path int true "ID da figurinha"
// @Success 200 {string} string "null"
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /figurinhas/{id} [delete]
func (h *FigurinhaHandler) Delete(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	err = h.svc.DeleteFigurinha(id)
	h.respond(c, nil, err)
}

func parseID(c *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	return uint(id), err
}
