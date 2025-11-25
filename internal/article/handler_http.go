package article

import (
	"net/http"
	"strconv"

	"github.com/bagusshndr/linknau-article-test/pkg/response"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	svc Service
}

func NewHTTPHandler(s Service) *HTTPHandler {
	return &HTTPHandler{svc: s}
}

func (h *HTTPHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/articles", h.Create)
	rg.GET("/articles", h.List)
	rg.GET("/articles/:id", h.GetByID)
	rg.PUT("/articles/:id", h.Update)
	rg.DELETE("/articles/:id", h.Delete)
}

func (h *HTTPHandler) Create(c *gin.Context) {
	var req ArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	art, err := h.svc.Create(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Created(c, ToArticleResponse(art))
}

func (h *HTTPHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	items, total, err := h.svc.List(c.Request.Context(), page, size)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	res := make([]ArticleResponse, 0, len(items))
	for i := range items {
		res = append(res, *ToArticleResponse(&items[i]))
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  res,
		"page":  page,
		"size":  size,
		"total": total,
	})
}

func (h *HTTPHandler) GetByID(c *gin.Context) {
	id, err := parseIDParam(c)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	art, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c, ToArticleResponse(art))
}

func (h *HTTPHandler) Update(c *gin.Context) {
	id, err := parseIDParam(c)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	var req ArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	art, err := h.svc.Update(c.Request.Context(), id, &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c, ToArticleResponse(art))
}

func (h *HTTPHandler) Delete(c *gin.Context) {
	id, err := parseIDParam(c)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

func parseIDParam(c *gin.Context) (uint, error) {
	raw := c.Param("id")
	i, err := strconv.Atoi(raw)
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}
