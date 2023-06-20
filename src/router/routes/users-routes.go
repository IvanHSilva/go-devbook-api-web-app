package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	// SELECT All
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.SelectUsers,
		AuthRequire: false,
	},
	// SELECT One
	{
		URI:         "/user/{userId}",
		Method:      http.MethodGet,
		Function:    controllers.SelectUser,
		AuthRequire: true,
	},
	// SEARCH *
	{
		URI:         "/user",
		Method:      http.MethodGet,
		Function:    controllers.SearchUser,
		AuthRequire: true,
	},
	// INSERT
	{
		URI:         "/user/{userId}",
		Method:      http.MethodPost,
		Function:    controllers.InsertUser,
		AuthRequire: false,
	},
	// UPDATE
	{
		URI:         "/user/{userId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		AuthRequire: true,
	},
	// DELETE
	{
		URI:         "/user/{userId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		AuthRequire: true,
	},
	// FOLLOW *
	{
		URI:         "/user/{userId}/follow",
		Method:      http.MethodPost,
		Function:    controllers.FollowUser,
		AuthRequire: true,
	},
	// UNFOLLOW *
	{
		URI:         "/user/{userId}/unfollow",
		Method:      http.MethodPost,
		Function:    controllers.UnfollowUser,
		AuthRequire: true,
	},
	// LIST FOLLOWERS *
	{
		URI:         "/user/{userId}/followers",
		Method:      http.MethodGet,
		Function:    controllers.FollowersOfUser,
		AuthRequire: true,
	},
	// LIST WHO FOLLOWS *
	{
		URI:         "/user/{userId}/whofollows",
		Method:      http.MethodGet,
		Function:    controllers.WhoUserFollows,
		AuthRequire: true,
	},
}
