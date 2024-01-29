package initialization

import (
	"go-echo/helper/database"
	transport "go-echo/http"
	"go-echo/repository"
	"go-echo/repository/user"
	"go-echo/service"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

func DbInit() (*gorm.DB, error) {
	// Init DB connection
	driver := viper.GetString("database.driver")
	dbname := viper.GetString("database.dbname")
	host := viper.GetString("database.host")
	user := viper.GetString("database.username")
	password := viper.GetString("database.password")
	port := viper.GetInt("database.port")

	db, err := database.NewDBConnection(driver, dbname, host, user, password, port)
	if err != nil {
		return nil, err
	}

	// _ = db.AutoMigrate(&model.Driver{})
	// _ = db.AutoMigrate(&model.Vehicle{})

	return db, nil
}

func ServerInit(log *zap.Logger, db *gorm.DB) {
	driverSvc := service.NewDriverService(log, repository.NewBaseRepository(db), repository.NewDriverRepository(repository.NewBaseRepository(db)))
	userSvc := service.NewUserService(log, repository.NewBaseRepository(db), user.NewUserRepository(repository.NewBaseRepository(db)))

	r := echo.New()
	apiGroupDriver := r.Group("/api/driver")
	apiGroupUser := r.Group("/api/user")
	transport.DriverHandler(apiGroupDriver, driverSvc)
	transport.UserHandler(apiGroupUser, userSvc)
	transport.SwaggerHttpHandler(r)

	r.Start(":9000")
}

func NewZapLogger(filename string) (*zap.Logger, error) {
	config := zap.NewProductionConfig()

	// if filename not empty
	if filename != "" {
		config.OutputPaths = append(config.OutputPaths, filename)
		config.ErrorOutputPaths = append(config.ErrorOutputPaths, filename)
	}

	config.Encoding = "json"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
	logger, err := config.Build()
	defer logger.Sync()

	return logger, err
}
