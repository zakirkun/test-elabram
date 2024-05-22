package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zakirkun/test-elabram/app/controller"
	"github.com/zakirkun/test-elabram/app/middleware"
)

func NewRouter() http.Handler {

	// Init Router
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20
	r.Static("/public", "./public")

	// Cors
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"*"},
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	}))

	// Controller
	companyController := controller.NewCompanyController()
	adminController := controller.NewAdminController()
	employeController := controller.NewEmployeController()
	claimController := controller.NewClaimController()

	// Health Checker
	r.GET("/", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong " + fmt.Sprint(time.Now().Unix()),
		})
	})

	api := r.Group("/api")
	{
		// Auth
		auth := api.Group("/auth")
		{
			// Empploye
			auth.POST("/employe/login", employeController.LoginEmploy)
			auth.POST("/employe/register", employeController.RegisterEmploye)

			// Admin
			auth.POST("/admin/login", adminController.LoginAdmin)
			auth.POST("/admin/register", adminController.RegisterAdmin)
		}

		// Company
		company := api.Group("/company")
		{
			company.POST("/create", companyController.CreateCompany)
			company.GET("/", companyController.GetAllCompany)
		}

		// Admin Menu
		admin := api.Group("/admin", middleware.AdminMiddleware)
		{
			admin.POST("/create", adminController.CreateAdmin)
			admin.GET("/", adminController.GetAllAdmin)
			admin.GET("/:id", adminController.GetAdmin)
			admin.PUT("/update", adminController.UpdateAdmin)
			admin.DELETE("/:id", adminController.DeleteAdmin)

			// Employe Menu
			_employe := admin.Group("/employe")
			{
				_employe.POST("/create", employeController.CreateEmploye)
				_employe.GET("/", employeController.GetAllEmploye)
				_employe.GET("/:id", employeController.GetEmploye)
				_employe.PUT("/update", employeController.UpdateEmploye)
				_employe.DELETE("/:id", employeController.DeleteEmploye)
			}

			// Claim Menu
			_claim := admin.Group("/claim")
			{
				_claim.GET("/", claimController.GetAllClaim)
				_claim.GET("/company/:id", claimController.GetByCompany)
				_claim.PUT("/approve", claimController.ApproveClaim)
				_claim.PUT("/reject", claimController.RejectClaim)
			}
		}

		// Employe menu
		employe := api.Group("/employe", middleware.EmployeMiddleware)
		{
			employe.POST("/claim/create", claimController.CreateClaim)
			employe.GET("/claim", claimController.GetByEmploye)
			employe.GET("/claim/:id", claimController.DetailClaim)
			employe.DELETE("/claim/:id", claimController.DeleteClaim)
			employe.PUT("/claim", claimController.UpdateClaim)
		}
	}

	return r
}
