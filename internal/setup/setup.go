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

	// Import untuk Company
	compHandler "simple-blog-system/internal/app/company/handler"
	compRepo "simple-blog-system/internal/app/company/repository"
	compService "simple-blog-system/internal/app/company/service"

	// --- TAMBAHAN IMPORT BANK ACCOUNT ---
	bankHandler "simple-blog-system/internal/app/bank_account/handler"
	bankRepo "simple-blog-system/internal/app/bank_account/repository"
	bankService "simple-blog-system/internal/app/bank_account/service"
)

// Struct Utama
type InternalAppStruct struct {
	Repositories initRepositoriesApp
	Services     initServicesApp
	Handler      InitHandlerApp
}

type initRepositoriesApp struct {
	userRepo    userPorts.IUserRepository
	postRepo    postPorts.IPostRepository
	commentRepo commentPorts.ICommentRepository
	
	// Tambahan Repo Company & Bank Account
	companyRepo     *compRepo.CompanyRepository
	bankAccountRepo *bankRepo.BankAccountRepository // Baru

	TrxHandler transaction.ISqlTransaction
	dbInstance *gorm.DB
}

func initAppRepo(gormDB *db.GormDB, initializeApp *InternalAppStruct) {
	initializeApp.Repositories.userRepo = userRepo.NewRepository(gormDB)
	initializeApp.Repositories.postRepo = postRepo.NewRepository(gormDB)
	initializeApp.Repositories.commentRepo = commentRepo.NewRepository(gormDB)
	
	// Ambil database SQL murni dari GORM untuk repository yang pakai sql.DB
	sqlDB, _ := gormDB.DB.DB() 
	
	// Init Repo Company
	initializeApp.Repositories.companyRepo = compRepo.NewCompanyRepository(sqlDB)
	
	// --- INIT REPO BANK ACCOUNT ---
	initializeApp.Repositories.bankAccountRepo = bankRepo.NewBankAccountRepository(sqlDB)

	initializeApp.Repositories.TrxHandler = transaction.NewSqlTransaction(gormDB)
	initializeApp.Repositories.dbInstance = gormDB.DB
}

type initServicesApp struct {
	UserService    userPorts.IUserService
	PostService    postPorts.IPostService
	CommentService commentPorts.ICommentService
	
	// Service Company & Bank Account
	CompanyService     *compService.CompanyService
	BankAccountService *bankService.BankAccountService // Baru
}

func initAppService(initializeApp *InternalAppStruct) {
	initializeApp.Services.UserService = userService.New(initializeApp.Repositories.userRepo)
	initializeApp.Services.PostService = postService.New(initializeApp.Repositories.postRepo, initializeApp.Repositories.userRepo)
	initializeApp.Services.CommentService = commentService.New(initializeApp.Repositories.commentRepo, initializeApp.Repositories.userRepo, initializeApp.Repositories.postRepo)
	
	// Init Service Company
	initializeApp.Services.CompanyService = compService.NewCompanyService(initializeApp.Repositories.companyRepo)
	
	// --- INIT SERVICE BANK ACCOUNT ---
	initializeApp.Services.BankAccountService = bankService.NewBankAccountService(initializeApp.Repositories.bankAccountRepo)
}

type InitHandlerApp struct {
	UserHandler    userPorts.IUserHandler
	PostHandler    postPorts.IPostHandler
	CommentHandler commentPorts.ICommentHandler
	
	// Handler Company & Bank Account
	CompanyHandler     *compHandler.CompanyHandler
	BankAccountHandler *bankHandler.BankAccountHandler // Baru
}

func initAppHandler(initializeApp *InternalAppStruct) {
	initializeApp.Handler.UserHandler = userHandler.New(initializeApp.Services.UserService)
	initializeApp.Handler.PostHandler = postHandler.New(initializeApp.Services.PostService)
	initializeApp.Handler.CommentHandler = commentHandler.New(initializeApp.Services.CommentService)
	
	// Init Handler Company
	initializeApp.Handler.CompanyHandler = compHandler.NewCompanyHandler(initializeApp.Services.CompanyService)
	
	// --- INIT HANDLER BANK ACCOUNT ---
	initializeApp.Handler.BankAccountHandler = bankHandler.NewBankAccountHandler(initializeApp.Services.BankAccountService)
}