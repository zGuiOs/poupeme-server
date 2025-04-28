package routes

import (
	"net/http"

	"github.com/zGuiOs/poupeme-server/src/controllers"
)

var usersRoutes = []Route{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Handler:  controllers.CreateUser,
		NeedAuth: false,
	},
	{
		URI:      "/users/{UUID}",
		Method:   http.MethodPut,
		Handler:  controllers.UpdateUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{UUID}",
		Method:   http.MethodDelete,
		Handler:  controllers.DeleteUser,
		NeedAuth: true,
	},
}
