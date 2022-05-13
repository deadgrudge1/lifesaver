package rest

import ( 
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"lifesaver/internal/lifesaver/service"
	"lifesaver/pkg/models"
)

// var userService service.userService

func GetUser(context *gin.Context) {
	log.Println("[UserController] BEGIN")

	userId := context.Param("userId")

	user, err := service.GetUser(userId)
	if(err != nil) {
		context.IndentedJSON(err.Code, err)
	} else {
		context.IndentedJSON(http.StatusOK, user)
	}
}

func SaveUser(context *gin.Context) {
	var user models.User
    if bindError := context.BindJSON(&user); bindError != nil {
        context.Status(http.StatusBadRequest)
    }

	newUserId, err := service.CreateUser(user)
	if(err != nil) {
		context.IndentedJSON(err.Code, err)
	} else {
		context.JSON(http.StatusOK, gin.H{ "id" : newUserId } )
	}
}

func UpdateUser(context *gin.Context) {
	var user *models.User
    if bindError := context.BindJSON(&user); bindError != nil {
        context.Status(http.StatusBadRequest)
    }
	user.Id = context.Param("userId")

	err := service.UpdateUser(user)
	if(err != nil) {
		context.IndentedJSON(err.Code, err)
	} else {
		context.Status(http.StatusOK)
	}
}

func DeleteUser(context *gin.Context) {
	userId := context.Param("userId")

	err := service.RemoveUser(userId)
	if(err != nil) {
		context.IndentedJSON(err.Code, err)
	} else {
		context.Status(http.StatusOK)
	}
}