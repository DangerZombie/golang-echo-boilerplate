package service_driver

import (
	"go-echo/helper/message"
	"go-echo/model/base"
	"go-echo/model/request"
	"go-echo/model/response"
	"go-echo/repository"
	"go-echo/repository/repository_driver"

	"go.uber.org/zap"
)

type driverServiceImpl struct {
	logger     *zap.Logger
	baseRepo   repository.BaseRepository
	driverRepo repository_driver.DriverRepository
}

type DriverService interface {
	InsertDriver(req request.InsertDriverRequest) (*response.InsertDriverResponse, message.Message, interface{})
	GetListDrivers(req request.GetListDriversRequest) ([]response.InsertDriverResponse, *base.Pagination, message.Message, interface{})
	GetDriverByNumber(req request.GetDriverByNumber) (*response.GetDriverByNumberResponse, message.Message, interface{})
	UpdateDriverByNumber(req request.UpdateDriverByNumber) (*response.UpdateDriverByNumberResponse, message.Message, interface{})
	DeleteDriverByNumber(req request.DeleteDriverByNumber) (*response.DeleteDriverByNumberResponse, message.Message, interface{})
}

func NewDriverService(
	lg *zap.Logger,
	br repository.BaseRepository,
	dr repository_driver.DriverRepository,
) DriverService {
	return &driverServiceImpl{lg, br, dr}
}
