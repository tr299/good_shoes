package service

import (
    "context"
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/otel/trace"
    "gorm.io/gorm"

    "good_shoes/common/config"
    "good_shoes/common/model/model_product"
    "good_shoes/common/util"
    inventoryRepository "good_shoes/inventory/repository"
    "good_shoes/logger"
    "good_shoes/product/repository"
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

func (h *Handler) CreateProduct(c *gin.Context) {
    newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
    ctx, span := h.tracer.Start(newCtx, "CreateProduct")
    defer span.End()

    req := &model_product.CreateProductRequest{}

    if err := util.BindRequest(ctx, c, req); nil != err {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    logger.Infof("Enter CreateProduct, data request = ", req)

    if err := validateCreateProduct(req); nil != err {
        logger.Error(err)
        c.JSON(http.StatusBadRequest, err)
        return
    }

    repo := repository.NewRepository(h.database)
    product, err := repo.CreateProduct(prepareDataToCreateProduct(req))
    if nil != err {
        c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
        return
    }

    // create option
    if len(req.Options) > 0 {
        options, err := repo.CreateProductOptions(prepareDataToCreateOptions(req.Options, product.Id))
        if nil != err {
            c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
            return
        }

        mapOptionNameToId := map[string]string{}
        mapOptionIdToName := map[string]string{}
        for _, option := range options {
            mapOptionNameToId[option.Name] = option.Id
            mapOptionIdToName[option.Id] = option.Name
        }

        // thứ tự của variant name => VD: abc-40-Blue -> size = 40, color = Blue
        var variantNameSortOrder []string
        var listVariantNames []string
        for _, option := range req.Options {
            optionItems, err := repo.CreateProductOptionItems(prepareDataToCreateOptionItems(option.Items, mapOptionNameToId[option.Name]))
            if nil != err {
                c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
                return
            }

            mapOptionToNameOptionItems := map[string][]string{}
            for _, item := range optionItems {
                if len(mapOptionToNameOptionItems[item.OptionId]) > 0 {
                    mapOptionToNameOptionItems[item.OptionId] = append(mapOptionToNameOptionItems[item.OptionId], item.Label)
                    continue
                }
                mapOptionToNameOptionItems[item.OptionId] = []string{item.Label}
            }
            for k, nameOptionItems := range mapOptionToNameOptionItems {
                variantNameSortOrder = append(variantNameSortOrder, mapOptionIdToName[k])
                listVariantNames = prepareDataToCreateProductVariant(nameOptionItems, listVariantNames)
            }
        }

        // create variant product
        for _, variantName := range listVariantNames {
            _, err := repo.CreateProduct(prepareDataToCreateVariant(req, variantName, product.Id, variantNameSortOrder))
            if nil != err {
                c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
                return
            }
        }
    }

    // Sync quantity
    go func() {
        inventoryRepo := inventoryRepository.NewRepository(h.database)
        inventoryRepo.UpdateParentQty(product.Id)
    }()

    c.JSON(http.StatusOK, &model_product.CreateProductResponse{
        ProductId: product.Id,
        Message:   "Create product success",
    })
}

func (h *Handler) UpdateProduct(c *gin.Context) {
    newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
    ctx, span := h.tracer.Start(newCtx, "UpdateProduct")
    defer span.End()

    req := &model_product.UpdateProductRequest{}

    if err := util.BindRequest(ctx, c, req); nil != err {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    logger.Infof("Enter UpdateProduct, data request = ", req)

    if err := validateUpdateProduct(req); nil != err {
        logger.Error(err)
        c.JSON(http.StatusBadRequest, err)
        return
    }

    repo := repository.NewRepository(h.database)
    data, err := repo.UpdateProduct(prepareDataToUpdateProduct(req))
    if nil != err {
        c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
    }

    logger.Infof("UpdateProduct -- data.Type = %v", data.Type)

    go repo.UpdateVariant(data)

    c.JSON(http.StatusOK, &model_product.CreateProductResponse{
        ProductId: data.Id,
        Message:   "Update product success",
    })
}

func (h *Handler) DeleteProduct(c *gin.Context) {
    newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
    ctx, span := h.tracer.Start(newCtx, "DeleteProduct")
    defer span.End()

    req := &model_product.DeleteProductByIdRequest{}

    if err := util.BindRequest(ctx, c, req); nil != err {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    logger.Infof("Enter DeleteProduct, data request = ", req)

    repo := repository.NewRepository(h.database)
    err := repo.DeleteProduct(req.Id)
    if nil != err {
        c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
    }

    c.JSON(http.StatusOK, &model_product.CreateProductResponse{
        ProductId: req.Id,
        Message:   "Delete product success",
    })
}

func (h *Handler) ListProduct(c *gin.Context) {
    newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
    ctx, span := h.tracer.Start(newCtx, "ListProduct")
    defer span.End()

    req := &model_product.ListProductRequest{}

    if err := util.BindRequest(ctx, c, req); nil != err {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    logger.Infof("Enter GetProduct, data request = ", req)

    repo := repository.NewRepository(h.database)
    result, err := repo.ListProduct(req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
        return
    }

    totalProduct := repo.Count(req)

    data := prepareDataToResponseListProduct(result, totalProduct)
    c.JSON(http.StatusOK, data)
}

func (h *Handler) GetProduct(c *gin.Context) {
    newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
    ctx, span := h.tracer.Start(newCtx, "GetProduct")
    defer span.End()

    req := &model_product.GetProductByIdRequest{}

    if err := util.BindRequest(ctx, c, req); nil != err {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    logger.Infof("Enter GetProduct, data request = ", req)

    if err := validateGetProduct(req); nil != err {
        logger.Error(err)
        c.JSON(http.StatusBadRequest, err)
        return
    }

    repo := repository.NewRepository(h.database)
    data, err := repo.GetProductById(req)
    if nil != err {
        c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
        return
    }
    response := model_product.ConvertProductModelToProductResponse(data)

    // get product options
    options, _ := repo.ListProductOptionsByProductId(req.Id)
    _options, optionIds := model_product.ConvertOptionsModelToOptions(options)

    // get product option items
    optionItems, _ := repo.ListProductOptionItemsByOptionIds(optionIds)
    response.Options = prepareOptionToResponse(_options, model_product.ConvertOptionItemsModelToOptionItems(optionItems))

    c.JSON(http.StatusOK, response)
}
