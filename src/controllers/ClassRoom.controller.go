package controllers

import "github.com/gin-gonic/gin"

type classRoomController struct{}

func NewClassRoomController() *classRoomController {
	return &classRoomController{}
}

func (srv *classRoomController) Recommend(r *gin.Context) {

}

func (srv *classRoomController) ForYou(r *gin.Context) {

}

func (srv *classRoomController) Categories(r *gin.Context) {

}

func (srv *classRoomController) ClassPurchase(r *gin.Context) {

}

func (srv *classRoomController) ClassPurchased(r *gin.Context) {

}
