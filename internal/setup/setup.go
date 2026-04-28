package setup

import (
	"gorm.io/gorm"
	"simple-blog-system/config/db"
	"simple-blog-system/pkg/transaction"

	userHandler "simple-blog-system/internal/app/user/handler"
	userPorts "simple-blog-system/internal/app/user/port"
	userRepo "simple-blog-system/internal/app/user/repository"
	userService "simple-blog-system/internal/app/user/service"

	postHandler "simple-blog-system/internal/app/post/handler"
	postPorts "simple-blog-system/internal/app/post/port"
	postRepo "simple-blog-system/internal/app/post/repository"
	postService "simple-blog-system/internal/app/post/service"

	commentHandler "simple-blog-system/internal/app/comment/handler"
	commentPorts "simple-blog-system/internal/app/comment/port"
	commentRepo "simple-blog-system/internal/app/comment/repository"
	commentService "simple-blog-system/internal/app/comment/service"
)

// Struct Utama (Hanya ada satu di package setup)
type InternalAppStruct struct {
	Repositories initRepositoriesApp
	Services     initServicesApp
	Handler      InitHandlerApp
}

type initRepositoriesApp struct {
	userRepo    userPorts.IUserRepository
	postRepo    postPorts.IPostRepository
	commentRepo commentPorts.ICommentRepository
	TrxHandler  transaction.ISqlTransaction
	dbInstance  *gorm.DB
}

func initAppRepo(gormDB *db.GormDB, initializeApp *InternalAppStruct) {
	initializeApp.Repositories.userRepo = userRepo.NewRepository(gormDB)
	initializeApp.Repositories.postRepo = postRepo.NewRepository(gormDB)
	initializeApp.Repositories.commentRepo = commentRepo.NewRepository(gormDB)
	initializeApp.Repositories.TrxHandler = transaction.NewSqlTransaction(gormDB)
	initializeApp.Repositories.dbInstance = gormDB.DB
}

type initServicesApp struct {
	UserService    userPorts.IUserService
	PostService    postPorts.IPostService
	CommentService commentPorts.ICommentService
}

func initAppService(initializeApp *InternalAppStruct) {
	initializeApp.Services.UserService = userService.New(initializeApp.Repositories.userRepo)
	initializeApp.Services.PostService = postService.New(initializeApp.Repositories.postRepo, initializeApp.Repositories.userRepo)
	initializeApp.Services.CommentService = commentService.New(initializeApp.Repositories.commentRepo, initializeApp.Repositories.userRepo, initializeApp.Repositories.postRepo)
}

type InitHandlerApp struct {
	UserHandler    userPorts.IUserHandler
	PostHandler    postPorts.IPostHandler
	CommentHandler commentPorts.ICommentHandler
}

func initAppHandler(initializeApp *InternalAppStruct) {
	initializeApp.Handler.UserHandler = userHandler.New(initializeApp.Services.UserService)
	initializeApp.Handler.PostHandler = postHandler.New(initializeApp.Services.PostService)
	initializeApp.Handler.CommentHandler = commentHandler.New(initializeApp.Services.CommentService)
}