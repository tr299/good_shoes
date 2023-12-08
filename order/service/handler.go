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
    newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
    ctx, span := h.tracer.Start(newCtx, "CreateSalesOrder")
    defer span.End()

    req := &model_order.CreateSalesOrderRequest{}

    if err := util.BindRequest(ctx, c, req); nil != err {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    logger.Infof("Enter CreateSalesOrder, data request = ", req)

    //if err := validateCreateProduct(req); nil != err {
    //    logger.Error(err)
    //    c.JSON(http.StatusBadRequest, err)
    //    return
    //}

    // create sales order
    repo := repository.NewRepository(h.database)
    data, err := repo.CreateSalesOrder(prepareDataToCreateSalesOrder(req))
    if nil != err || nil == data {
        c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
        return
    }

    // create sales order items
    _, err = repo.CreateSalesOrderItems(prepareDataToCreateSalesOrderItems(req.Items, data.Id))
    if nil != err {
        c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
        return
    }

    c.JSON(http.StatusOK, &model_order.CreateSalesOrderResponse{
        Message: "Create Sales Order Success",
        OrderId: data.Id,
    })
}

func (h *Handler) UpdateSalesOrderStatus(c *gin.Context) {
    newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
    ctx, span := h.tracer.Start(newCtx, "UpdateSalesOrderState")
    defer span.End()

    req := &model_order.UpdateOrderStatusRequest{}

    if err := util.BindRequest(ctx, c, req); nil != err {
        c.JSON(http.StatusBadRequest, err.Error())
        return
    }

    logger.Infof("Enter ListSalesOrder, data request = ", req)

    if err := validateUpdateOrderStatus(req); nil != err {
        logger.Error(err)
        c.JSON(http.StatusBadRequest, err.Error())
    }

    repo := repository.NewRepository(h.database)
    data, err := repo.UpdateSalesOrderStatus(req.Id, req.Status)
    if err != nil {
        c.JSON(http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, &model_order.UpdateOrderStatusResponse{
        OrderId: data.Id,
        Status:  data.Status,
        Message: "Update sales order status success",
    })
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

    // get sales order
    repo := repository.NewRepository(h.database)
    order, err := repo.GetOrderById(req.Id)
    if nil != err || nil == order {
        c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
        return
    }

    // list sales order items
    orderItems, err := repo.ListOrderItems(order.Id)
    if nil != err {
        logger.Error(err)
    }

    response := prepareDataToResponseGetSalesOrder(order, orderItems)

    c.JSON(http.StatusOK, response)
}
