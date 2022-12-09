package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vansh284/CVWOWebForumAPI/pkg/controllers"
)

func InitThreadRoutes(app *fiber.App) {
	app.Get("/users/:username", controllers.GetUser)                                    //Retrives user :username
	app.Post("/users", controllers.CreateUser)                                          //Creates user
	app.Patch("/users/username/:username", controllers.UpdateUserUsername)              //Edits username of user with {username}
	app.Patch("/users/password/:username", controllers.UpdateUserPassword)              //Edits password of user with {username}
	app.Delete("/users/:username", controllers.DeleteUser)                              //Deletes user with {username}
	app.Get("/threads", controllers.GetThreads)                                         //Retrives list of threads
	app.Post("/threads", controllers.CreateThread)                                      //Creates thread
	app.Put("/threads/:id<int>", controllers.EditThread)                                //Edits Thread {id}. Thread content and tag required in body.
	app.Delete("/threads/:id<int>", controllers.DeleteThread)                           //Deletes Thread {id}
	app.Get("/threads/:thread_id<int>/comments", controllers.GetCommentsT)              //Retrieves comments in thread {thread_id}
	app.Post("/threads/:thread_id<int>/comments", controllers.CreateComment)            //Creates comment in thread {thread_id}
	app.Put("/threads/:thread_id<int>/comments/:id<int>", controllers.EditComment)      //Edits comment {id} in thread {thread_id}. Comment content required in body.
	app.Delete("/threads/:thread_id<int>/comments/:id<int>", controllers.DeleteComment) //Deletes comment {id} in thread{thread_id}
}
