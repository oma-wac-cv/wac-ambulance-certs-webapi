package api

import (
    _ "embed"
    "net/http"

    "github.com/gin-gonic/gin"
)

//go:embed oma_wac_certs.openapi.yaml
var openapiSpec []byte

func HandleOpenApi(ctx *gin.Context) {
    ctx.Data(http.StatusOK, "application/yaml", openapiSpec)
}
