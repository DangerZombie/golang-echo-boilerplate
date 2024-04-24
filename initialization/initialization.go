package initialization

import (
	"go-echo/helper/auth"
	"go-echo/helper/database"
	transport "go-echo/http"
	"go-echo/repository"
	"go-echo/repository/repository_teacher"
	"go-echo/repository/repository_user"
	"go-echo/service/service_teacher"
	"go-echo/service/service_user"

	"github.com/labstack/echo/v4"
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

	// _ = db.AutoMigrate(&entity.Role{})
	// _ = db.AutoMigrate(&entity.JobTitle{})
	// _ = db.AutoMigrate(&entity.Teacher{})
	// _ = db.AutoMigrate(&entity.User{})

	return db, nil
}

func ServerInit(log *zap.Logger, db *gorm.DB) {
	// base repository
	baseRepository := repository.NewBaseRepository(db)
	teacherRepository := repository_teacher.NewTeacherRepository(baseRepository)
	userRepository := repository_user.NewUserRepository(baseRepository)

	// auth helper
	authHelper := auth.NewAuthHelper(
		baseRepository,
		userRepository,
	)

	// service
	teacherSvc := service_teacher.NewTeacherService(
		log,
		baseRepository,
		teacherRepository,
	)

	userSvc := service_user.NewUserService(
		log,
		authHelper,
		baseRepository,
		userRepository,
	)

	r := echo.New()

	// group endpoint
	apiGroupTeacher := r.Group("/api/v1/teacher")
	apiGroupUser := r.Group("/api/v1/user")

	// transport
	transportHandler := transport.NewHttp(
		authHelper,
	)

	transportHandler.SwaggerHttpHandler(r)
	transportHandler.TeacherHandler(apiGroupTeacher, teacherSvc)
	transportHandler.UserHandler(apiGroupUser, userSvc)

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
