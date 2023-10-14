// Package v1 package v1
package v1

import (
	"net/http"
	"sort"
	"time"

	"tmt/global"
	"tmt/internal/entity"
	"tmt/internal/usecase/cases/analyze"

	"github.com/gin-gonic/gin"
)

type analyzeRoutes struct {
	t analyze.Analyze
}

func NewAnalyzeRoutes(handler *gin.RouterGroup, t analyze.Analyze) {
	r := &analyzeRoutes{t}

	h := handler.Group("/analyze")
	{
		h.GET("/reborn", r.getRebornTargets)
	}
}

type reborn struct {
	Date   string         `json:"date"`
	Stocks []entity.Stock `json:"stocks"`
}

// getRebornTargets -.
//
//	@Summary		getRebornTargets
//	@Description	getRebornTargets
//	@ID				getRebornTargets
//	@Tags			analyze
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]reborn{}
//	@Router			/analyze/reborn [get]
func (r *analyzeRoutes) getRebornTargets(c *gin.Context) {
	mapData := r.t.GetRebornMap(c.Request.Context())
	result := []reborn{}
	dateArr := []time.Time{}
	for date := range mapData {
		dateArr = append(dateArr, date)
	}

	sort.Slice(dateArr, func(i, j int) bool {
		return dateArr[i].After(dateArr[j])
	})

	for _, date := range dateArr {
		result = append(result, reborn{
			Date:   date.Format(global.ShortTimeLayout),
			Stocks: mapData[date],
		})
	}
	c.JSON(http.StatusOK, result)
}
