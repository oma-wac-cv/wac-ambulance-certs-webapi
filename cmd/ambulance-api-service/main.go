package main

import (
    "log"
    "os"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/oma-wac-cv/wac-ambulance-certs-webapi/api"
    "github.com/oma-wac-cv/wac-ambulance-certs-webapi/internal/oma_wac_certs"
    "github.com/oma-wac-cv/wac-ambulance-certs-webapi/internal/db_service"
    "context"
    "time"
    "github.com/gin-contrib/cors"
)

func main() {
    log.Printf("Server started")
    port := os.Getenv("AMBULANCE_API_PORT")
    if port == "" {
        port = "8080"
    }
    environment := os.Getenv("AMBULANCE_API_ENVIRONMENT")
    if !strings.EqualFold(environment, "production") { // case insensitive comparison
        gin.SetMode(gin.DebugMode)
    }
    engine := gin.New()
    engine.Use(gin.Recovery())

    corsMiddleware := cors.New(cors.Config{
      AllowOrigins:     []string{"*"},
      AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH"},
      AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
      ExposeHeaders:    []string{""},
      AllowCredentials: false,
      MaxAge: 12 * time.Hour,
    })
    engine.Use(corsMiddleware)

    dbService := db_service.NewMongoService[oma_wac_certs.User](db_service.MongoServiceConfig{Collection: "users"})
    dbServiceCert := db_service.NewMongoService[oma_wac_certs.Certification](db_service.MongoServiceConfig{Collection: "certifications"})
    dbServiceUserCert := db_service.NewMongoService[oma_wac_certs.UserCertification](db_service.MongoServiceConfig{Collection: "user_certifications"})

    defer dbService.Disconnect(context.Background())
    defer dbServiceCert.Disconnect(context.Background())
    defer dbServiceUserCert.Disconnect(context.Background())

    engine.Use(func(ctx *gin.Context) {
        ctx.Set("db_service_user", dbService)
        ctx.Set("db_service_cert", dbServiceCert)
        ctx.Set("db_service_user_cert", dbServiceUserCert)
        ctx.Next()
    })

    oma_wac_certs.AddRoutes(engine)
    engine.GET("/openapi", api.HandleOpenApi)
    engine.Run(":" + port)
}
