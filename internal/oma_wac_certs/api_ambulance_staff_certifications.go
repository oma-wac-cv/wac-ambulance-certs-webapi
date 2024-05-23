/*
 * Waiting List Api
 *
 * Ambulance Waiting List API
 *
 * API version: 1.0.3
 * Contact: xmartinkao@stuba.sk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

 package oma_wac_certs

import (
   "net/http"

   "github.com/gin-gonic/gin"
)

type AmbulanceStaffCertificationsAPI interface {

   // internal registration of api routes
   addRoutes(routerGroup *gin.RouterGroup)

    // AddCertification - Add a new certification
   AddCertification(ctx *gin.Context)

    // DeleteCertification - Delete a certification
   DeleteCertification(ctx *gin.Context)

    // GetCertifications - Get all certifications
   GetCertifications(ctx *gin.Context)

    // GetUsers - Get all users
   GetUsers(ctx *gin.Context)

    // SeedDatabase - Seed the database
   SeedDatabase(ctx *gin.Context)

    // UpdateUser - Update a user
   UpdateUser(ctx *gin.Context)

}

// partial implementation of AmbulanceStaffCertificationsAPI - all functions must be implemented in add on files
type implAmbulanceStaffCertificationsAPI struct {

}

func newAmbulanceStaffCertificationsAPI() AmbulanceStaffCertificationsAPI {
  return &implAmbulanceStaffCertificationsAPI{}
}

func (this *implAmbulanceStaffCertificationsAPI) addRoutes(routerGroup *gin.RouterGroup) {
  routerGroup.Handle( http.MethodPost, "/certifications", this.AddCertification)
  routerGroup.Handle( http.MethodDelete, "/certifications/:certificationId", this.DeleteCertification)
  routerGroup.Handle( http.MethodGet, "/certifications", this.GetCertifications)
  routerGroup.Handle( http.MethodGet, "/users", this.GetUsers)
  routerGroup.Handle( http.MethodPost, "/seed", this.SeedDatabase)
  routerGroup.Handle( http.MethodPut, "/users/:userId", this.UpdateUser)
}

// Copy following section to separate file, uncomment, and implement accordingly
// // AddCertification - Add a new certification
// func (this *implAmbulanceStaffCertificationsAPI) AddCertification(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // DeleteCertification - Delete a certification
// func (this *implAmbulanceStaffCertificationsAPI) DeleteCertification(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetCertifications - Get all certifications
// func (this *implAmbulanceStaffCertificationsAPI) GetCertifications(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetUsers - Get all users
// func (this *implAmbulanceStaffCertificationsAPI) GetUsers(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // SeedDatabase - Seed the database
// func (this *implAmbulanceStaffCertificationsAPI) SeedDatabase(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // UpdateUser - Update a user
// func (this *implAmbulanceStaffCertificationsAPI) UpdateUser(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
