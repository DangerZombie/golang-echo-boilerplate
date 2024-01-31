package service_driver

import (
	"go-echo/helper/message"
	"go-echo/helper/static"
	"go-echo/helper/util"
	"go-echo/model/base"
	"go-echo/model/entity"
	"go-echo/model/request"
	"go-echo/model/response"

	"go.uber.org/zap"
)

// swagger:operation POST /driver/ Driver InsertDriverRequest
// Add Driver
//
// ---
// security:
// - Bearer: []
// responses:
//   '200':
//     description: Success Response.
//     schema:
//       properties:
//         meta:
//           $ref: '#/definitions/MetaSingleSuccessResponse'
//         data:
//           properties:
//             record:
//               $ref: '#/definitions/InsertDriverResponse'
//           type: object

func (s *driverServiceImpl) InsertDriver(req request.InsertDriverRequest) (*response.InsertDriverResponse, message.Message, map[string]string) {
	logger := s.logger.With(zap.String("DriverService", "InsertDriver"))
	errMsg := map[string]string{}

	data := entity.Driver{
		Name:          req.Name,
		LicenseNumber: req.LicenseNumber,
		IsAvailable:   req.IsAvailable,
		BaseModel: base.BaseModel{
			CreatedBy: "system",
			UpdatedBy: "system",
		},
	}

	tx := s.baseRepo.GetBegin()
	driver, err := s.driverRepo.InsertDriver(tx, data)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["driver"] = "error has been occured while inserting driver"
		return nil, message.FailedMsg, errMsg
	}

	resp := response.InsertDriverResponse{
		Id:            driver.ID,
		Name:          driver.Name,
		LicenseNumber: driver.LicenseNumber,
		IsAvailable:   driver.IsAvailable,
		CreatedAt:     util.UnixToFullDate(driver.CreatedAt, static.LayoutDefault),
		CreatedBy:     driver.CreatedBy,
		UpdatedAt:     util.UnixToFullDate(driver.UpdatedAt, static.LayoutDefault),
		UpdatedBy:     driver.UpdatedBy,
	}

	s.baseRepo.BeginCommit(tx)
	return &resp, message.SuccessMsg, nil
}

func (s *driverServiceImpl) GetListDrivers(req request.GetListDriversRequest) ([]response.InsertDriverResponse, *base.Pagination, message.Message, map[string]string) {
	logger := s.logger.With(zap.String("DriverService", "GetListDrivers"))
	errMsg := map[string]string{}

	if req.Page < 1 {
		req.Page = 1
	}

	if req.Limit < 1 {
		req.Limit = 10
	}

	filter := map[string]interface{}{
		"name": req.Name,
	}

	tx := s.baseRepo.GetBegin()
	drivers, pagiantion, err := s.driverRepo.GetListDrivers(tx, req.Limit, req.Page, req.Sort, req.Dir, filter)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["driver"] = "error has been occured while looking driver"
		return nil, nil, message.FailedMsg, errMsg
	}

	result := []response.InsertDriverResponse{}
	for _, val := range drivers {
		driver := response.InsertDriverResponse{
			Id:            val.ID,
			Name:          val.Name,
			LicenseNumber: val.LicenseNumber,
			IsAvailable:   val.IsAvailable,
			CreatedAt:     util.UnixToFullDate(val.CreatedAt, static.LayoutDefault),
			CreatedBy:     val.CreatedBy,
			UpdatedAt:     util.UnixToFullDate(val.UpdatedAt, static.LayoutDefault),
			UpdatedBy:     val.UpdatedBy,
		}

		result = append(result, driver)
	}

	s.baseRepo.BeginCommit(tx)
	return result, pagiantion, message.SuccessMsg, nil
}

func (s *driverServiceImpl) GetDriverByNumber(req request.GetDriverByNumber) (*response.GetDriverByNumberResponse, message.Message, map[string]string) {
	logger := s.logger.With(zap.String("DriverService", "GetDriverByNumber"))
	errMsg := map[string]string{}

	if req.Number == "" {
		logger.Error("log", zap.String("error", "number is empty"))
		errMsg["driver"] = "number is empty"
		return nil, message.FailedMsg, errMsg
	}

	tx := s.baseRepo.GetBegin()
	driver, err := s.driverRepo.GetDriverByNumber(tx, req.Number)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["driver"] = "error has been occured while looking driver"
		return nil, message.FailedMsg, errMsg
	}

	result := response.GetDriverByNumberResponse{
		Id:            driver.ID,
		Name:          driver.Name,
		LicenseNumber: driver.LicenseNumber,
		IsAvailable:   driver.IsAvailable,
		CreatedAt:     util.UnixToFullDate(driver.CreatedAt, static.LayoutDefault),
		CreatedBy:     driver.CreatedBy,
		UpdatedAt:     util.UnixToFullDate(driver.UpdatedAt, static.LayoutDefault),
		UpdatedBy:     driver.UpdatedBy,
	}

	return &result, message.SuccessMsg, nil
}

func (s *driverServiceImpl) UpdateDriverByNumber(req request.UpdateDriverByNumber) (*response.UpdateDriverByNumberResponse, message.Message, map[string]string) {
	logger := s.logger.With(zap.String("DriverService", "UpdateDriverByNumber"))
	errMsg := map[string]string{}

	if req.Number == "" {
		logger.Error("log", zap.String("error", "number is empty"))
		errMsg["driver"] = "number is empty"
		return nil, message.FailedMsg, nil
	}

	updateData := map[string]interface{}{
		"is_available": req.IsAvailable,
	}

	tx := s.baseRepo.GetBegin()
	driver, err := s.driverRepo.UpdateDriverByNumber(tx, req.Number, updateData)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["driver"] = "error has been occured while updating driver"
		return nil, message.FailedMsg, nil
	}

	result := response.UpdateDriverByNumberResponse{
		Id:            driver.ID,
		Name:          driver.Name,
		LicenseNumber: driver.LicenseNumber,
		IsAvailable:   driver.IsAvailable,
		CreatedAt:     util.UnixToFullDate(driver.CreatedAt, static.LayoutDefault),
		CreatedBy:     driver.CreatedBy,
		UpdatedAt:     util.UnixToFullDate(driver.UpdatedAt, static.LayoutDefault),
		UpdatedBy:     driver.UpdatedBy,
	}

	return &result, message.SuccessMsg, nil
}

func (s *driverServiceImpl) DeleteDriverByNumber(req request.DeleteDriverByNumber) (*response.DeleteDriverByNumberResponse, message.Message, map[string]string) {
	logger := s.logger.With(zap.String("DriverService", "UpdateDriverByNumber"))
	errMsg := map[string]string{}

	if req.Number == "" {
		logger.Error("log", zap.String("error", "number is empty"))
		errMsg["driver"] = "number is empty"
		return nil, message.FailedMsg, errMsg
	}

	tx := s.baseRepo.GetBegin()
	err := s.driverRepo.DeleteDriverByNumber(tx, req.Number)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["driver"] = "error has been occured while deleting driver"
		return nil, message.FailedMsg, errMsg
	}

	result := response.DeleteDriverByNumberResponse{
		Message: "Success",
	}

	return &result, message.SuccessMsg, nil
}
