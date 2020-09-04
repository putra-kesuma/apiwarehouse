package routes

import (
	"apiwarehouse/controllers"
	"apiwarehouse/repositories"
	"apiwarehouse/usecases"
	"database/sql"
	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB){
	warehouseRepo := repositories.InitWarehouseRepoImpl(db)
	reportWarehouseRepo := repositories.InitReportWarehouseRepoImpl(db)
	warehouseStorageRepo := repositories.InitWarehouseStorageRepoImpl(db)
	itemRepo := repositories.InitItemRepoImpl(db)
	userRepo := repositories.InitUserRepoImpl(db)
	moviesRepo := repositories.InitMoviesRepoImpl(db)

	warehouseUseCase := usecases.InitWarehouseUsecase(warehouseRepo)
	reportWarehouseUseCase := usecases.InitReportWarehouseUsecase(reportWarehouseRepo)
	warehouseStorageUseCase := usecases.InitWarehouseStorageUsecase(warehouseStorageRepo)
	itemUseCase := usecases.InitItemUsecase(itemRepo)
	userUseCase := usecases.InitUserUsecase(userRepo)
	moviesUseCase := usecases.InitMoviesUsecase(moviesRepo)

	controllers.WarehouseController(r,warehouseUseCase)
	controllers.ReportWarehouseController(r,reportWarehouseUseCase)
	controllers.WarehouseStorageController(r,warehouseStorageUseCase)
	controllers.ItemController(r,itemUseCase)
	controllers.UserController(r,userUseCase)
	controllers.MoviesController(r,moviesUseCase)
}
