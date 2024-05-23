package oma_wac_certs

import (
  "net/http"
  "os"
  "fmt"
  
  "github.com/google/uuid"
  "github.com/gin-gonic/gin"
  "github.com/oma-wac-cv/wac-ambulance-certs-webapi/internal/db_service"

)

// AddCertification - Add a new certification
func (this *implAmbulanceStaffCertificationsAPI) AddCertification(ctx *gin.Context) {
  value, exists := ctx.Get("db_service_cert")
  if !exists {
    return
  }
  db, ok := value.(db_service.DbService[Certification])
  if !ok {
    return
  }

  cert := Certification{}
  err := ctx.BindJSON(&cert)
  cert.Id = uuid.New().String()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
    return
  }

  err = db.CreateDocument(ctx, cert.Id, &cert)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"message": "duplicate probably"})
    return
  }

  ctx.JSON(http.StatusCreated, cert)
}

// DeleteCertification - Delete a certification
func (this *implAmbulanceStaffCertificationsAPI) DeleteCertification(ctx *gin.Context) {
  value, exists := ctx.Get("db_service_cert")
  if !exists {
    ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
    return
  }
  db, ok := value.(db_service.DbService[Certification])
  if !ok {
    ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
    return
  }

  value, exists = ctx.Get("db_service_user_cert")
  if !exists {
    ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
    return
  }
  db_user_cert, ok := value.(db_service.DbService[UserCertification])
  if !ok {
    ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
    return
  }

  id := ctx.Param("certificationId")

  // check if exists
  _, err := db.FindDocument(ctx, id)
  if err != nil {
    ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
    return
  }

  // delete all relations between this cert and users
  userCerts, err := db_user_cert.FindAllDocuments(ctx)
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
    return
  }

  for _, userCert := range userCerts {
    if userCert.CertificationId == id {
      err = db_user_cert.DeleteDocument(ctx, userCert.Id)
      if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
        return
      }
    }
  }


  err = db.DeleteDocument(ctx, id)
  if err != nil {
    ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
    return
  }

  ctx.JSON(http.StatusNoContent, nil)
}

// GetCertifications - Get all certifications
func (this *implAmbulanceStaffCertificationsAPI) GetCertifications(ctx *gin.Context) {
  value, exists := ctx.Get("db_service_cert")
  if !exists {
    return
  }
  db, ok := value.(db_service.DbService[Certification])
  if !ok {
    return
  }

  users, err := db.FindAllDocuments(ctx)
  if err != nil {
    fmt.Println(err)
    ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
    return
  }

  ctx.JSON(http.StatusOK, users)
}

// GetUsers - Get all users
func (this *implAmbulanceStaffCertificationsAPI) GetUsers(ctx *gin.Context) {
  value, exists := ctx.Get("db_service_user")
  if !exists {
    return
  }
  
  db, ok := value.(db_service.DbService[User])
  if !ok {
    return
  }

  users, err := db.FindAllDocuments(ctx)
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
    return
  }

  ctx.JSON(http.StatusOK, users)
}

// UpdateUser - Update a user
func (this *implAmbulanceStaffCertificationsAPI) UpdateUser(ctx *gin.Context) {
  value, exists := ctx.Get("db_service_user")
  if !exists {
    ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
    return
  }
  
  col_user, ok := value.(db_service.DbService[User])
  if !ok {
    ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
    return
  }

  value, exists = ctx.Get("db_service_cert")
  if !exists {
    ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
    return
  }

  col_cert, ok := value.(db_service.DbService[Certification])
  if !ok {
    ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
    return
  }

  // parse body
  user := User{}
  err := ctx.BindJSON(&user)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
    return
  }

  id := ctx.Param("userId")
  if id != user.Id {
    ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
    return
  }

  // check whether user exists
  _, err = col_user.FindDocument(ctx, id)
  if err != nil {
    ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
    return
  }


  // check whether certification exists
  for _, cert := range user.Certifications {
    _, err = col_cert.FindDocument(ctx, cert.CertificationId)
    if err != nil {
      message := fmt.Sprintf("Certification with id %s not found", cert.CertificationId)
      ctx.JSON(http.StatusBadRequest, gin.H{"message": message})
      return
    }
  }

  err = col_user.UpdateDocument(ctx, id, &user)


  // // remove old certifications
  // for _, cert := range user_old.Certifications {
  //   // i do not want to check if it deletes it or no, i dont care
  //   col_user_cert.DeleteDocument(ctx, cert.Id)
  // }

  // // add new certifications
  // for _, cert := range user.Certifications {
  //   // check whether certification exists
  //   _, err = col_cert.FindDocument(ctx, cert.CertificationId)
  //   if err != nil {
  //     message := fmt.Sprintf("Certification with id %s not found", cert.CertificationId)
  //     ctx.JSON(http.StatusBadRequest, gin.H{"message": message})
  //     return
  //   }

  //   cert.Id = uuid.New().String()
  //   cert.UserId = user.Id
  //   col_user_cert.CreateDocument(ctx, cert.Id, &cert)
  // }

  ctx.JSON(http.StatusCreated, user)
}

// SeedDatabase - Seed the database
func (this *implAmbulanceStaffCertificationsAPI) SeedDatabase(ctx *gin.Context) {
  // get secret passphase from env
  secret := os.Getenv("AMBULANCE_API_MONGODB_SEED_PASSPHRASE")

  // check for authorization header
  if ctx.GetHeader("Authorization") != secret {
    ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
    return
  }

  value, exists := ctx.Get("db_service_user")
  if !exists {
    return
  }
  
  db, ok := value.(db_service.DbService[User])
  if !ok {
    return
  }

  ids := []string{
    "e6698483-7ef6-4432-acd9-baeb68830dae",
    "ab965b31-7f6d-42d4-9dde-366ae53d8d48",
    "c2f21653-98bc-489a-863e-e610688e8b00",
    "b8cf9a57-12ba-487f-a588-10d67d602709",
    "5f1de7c2-eafc-4ac4-9ade-8384e88f29a4",
  }

  names:= []string{
    "Natalie J. Capps",
    "Joan M. Thorp",
    "Victoria M. Armstrong",
    "Miki M. Goff",
    "Dick K. Roberts",
  }
  
  // iterate over all names and create user for each
  for i := 0; i < len(ids); i++ {
    user := User{
      Id: ids[i],
      Name: names[i],
      Certifications: []UserCertification{},
    }

    el, err := db.FindDocument(ctx, user.Id)
    if el != nil {
      fmt.Println("Database alerady seeded.")
      ctx.JSON(http.StatusBadRequest, gin.H{"message": "Database already seeded"})
      return
    }

    err = db.CreateDocument(ctx, user.Id, &user)
    if err != nil {
      fmt.Println(err)
      ctx.JSON(http.StatusBadRequest, gin.H{"message": "Database already seeded"})
      return
    }
  }
  ctx.JSON(http.StatusOK, gin.H{"message": "Database seeded"})
}
