package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"good_shoes/common/model/model_order"
	"good_shoes/common/util"
	"good_shoes/logger"
	"good_shoes/order/repository"
	"gorm.io/gorm"
	"net/http"

	"good_shoes/common/config"
)

type Handler struct {
	config   *config.Config
	database *gorm.DB
	tracer   trace.Tracer
}

func NewHandler(config *config.Config, db *gorm.DB, tracer trace.Tracer) (*Handler, error) {
	return &Handler{
		config:   config,
		database: db,
		tracer:   tracer,
	}, nil
}

func (h *Handler) CreateSalesOrder(c *gin.Context) {

}

func (h *Handler) UpdateSalesOrder(c *gin.Context) {

}

func (h *Handler) ListSalesOrder(c *gin.Context) {
	newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
	ctx, span := h.tracer.Start(newCtx, "ListSalesOrder")
	defer span.End()

	req := &model_order.ListSalesOrderRequest{}

	if err := util.BindRequest(ctx, c, req); nil != err {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	logger.Infof("Enter ListSalesOrder, data request = ", req)

	repo := repository.NewRepository(h.database)
	result, err := repo.ListOrder(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	data := prepareDataToResponseListSalesOrder(result)
	c.JSON(http.StatusOK, data)
}

func (h *Handler) GetSalesOrder(c *gin.Context) {
	newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
	ctx, span := h.tracer.Start(newCtx, "GetSalesOrder")
	defer span.End()

	req := &model_order.GetOrderByIdRequest{}

	if err := util.BindRequest(ctx, c, req); nil != err {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	logger.Infof("Enter GetProduct, data request = ", req)

	//if err := validateGetOrder(req); nil != err {
	//	logger.Error(err)
	//	c.JSON(http.StatusBadRequest, err)
	//	return
	//}

	repo := repository.NewRepository(h.database)
	data, err := repo.GetOrderById(req.Id)
	if nil != err {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
		return
	}
	response := prepareDataToResponseGetSalesOrder(data, nil)

	// get product options
	//options, _ := repo.ListProductOptionsByProductId(req.Id)
	//_options, optionIds := model_product.ConvertOptionsModelToOptions(options)

	// get product option items
	//optionItems, _ := repo.ListProductOptionItemsByOptionIds(optionIds)
	//response.Options = prepareOptionToResponse(_options, model_product.ConvertOptionItemsModelToOptionItems(optionItems))

	c.JSON(http.StatusOK, response)
}
