package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesPosts = []Route{
	// SELECT All
	{
		URI:         "/posts",
		Method:      http.MethodGet,
		Function:    controllers.SelectPosts,
		AuthRequire: false,
	},
	// SELECT One
	{
		URI:         "/post/{postId}",
		Method:      http.MethodGet,
		Function:    controllers.SelectPost,
		AuthRequire: true,
	},
	// SEARCH *
	{
		URI:         "/user/{userId}/post",
		Method:      http.MethodGet,
		Function:    controllers.SearchPost,
		AuthRequire: true,
	},
	// INSERT
	{
		URI:         "/post",
		Method:      http.MethodPost,
		Function:    controllers.InsertPost,
		AuthRequire: false,
	},
	// UPDATE
	{
		URI:         "/post/{postId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdatePost,
		AuthRequire: true,
	},
	// DELETE
	{
		URI:         "/post/{postId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeletePost,
		AuthRequire: true,
	},
	// LIKE *
	{
		URI:         "/post/{postId}/like",
		Method:      http.MethodPost,
		Function:    controllers.LikePost,
		AuthRequire: true,
	},
	// UNLIKE *
	{
		URI:         "/post/{postId}/unlike",
		Method:      http.MethodPost,
		Function:    controllers.UnlikePost,
		AuthRequire: true,
	},
}
