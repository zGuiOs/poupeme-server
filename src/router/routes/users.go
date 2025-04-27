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
		URI:      "/users",
		Method:   http.MethodGet,
		Handler:  controllers.FetchUsers,
		NeedAuth: false,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodGet,
		Handler:  controllers.FetchUserById,
		NeedAuth: false,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodPut,
		Handler:  controllers.UpdateUser,
		NeedAuth: false,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodDelete,
		Handler:  controllers.DeleteUser,
		NeedAuth: false,
	},
}
