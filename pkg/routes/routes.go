package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vansh284/CVWOWebForumAPI/pkg/controllers"
)

func InitThreadRoutes(app *fiber.App) {
	app.Get("/users", controllers.GetUser)                                              //Retrives user :username
	app.Post("/users", controllers.CreateUser)                                          //Creates user
	app.Post("/login", controllers.Login)                                               //Logs user in
	app.Post("/logout", controllers.Logout)                                             //Logs user out
	app.Patch("/users/username", controllers.UpdateUserUsername)                        //Edits username of user
	app.Patch("/users/password", controllers.UpdateUserPassword)                        //Edits password of user
	app.Delete("/users", controllers.DeleteUser)                                        //Deletes user
	app.Get("/threads", controllers.GetThreads)                                         //Retrives list of threads
	app.Post("/threads", controllers.CreateThread)                                      //Creates thread
	app.Put("/threads/:id<int>", controllers.EditThread)                                //Edits Thread {id}. Thread content and tag required in body.
	app.Delete("/threads/:id<int>", controllers.DeleteThread)                           //Deletes Thread {id}
	app.Get("/threads/:thread_id<int>/comments", controllers.GetCommentsT)              //Retrieves comments in thread {thread_id}
	app.Post("/threads/:thread_id<int>/comments", controllers.CreateComment)            //Creates comment in thread {thread_id}
	app.Put("/threads/:thread_id<int>/comments/:id<int>", controllers.EditComment)      //Edits comment {id} in thread {thread_id}. Comment content required in body.
	app.Delete("/threads/:thread_id<int>/comments/:id<int>", controllers.DeleteComment) //Deletes comment {id} in thread{thread_id}
}
