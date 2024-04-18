package endpoints

import (
	"context"
	"errors"
	"fmt"
	"github.com/engagerocketco/templates-api-svc/internal/handler/transport/models"
	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	"github.com/go-kit/kit/endpoint"
)

func MakeEmployeeEndpoint(service templateservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		employees, ok := request.([]*models.Employee)
		if !ok {
			return nil, errors.New("invalid request format")
		}

		if len(employees) == 0 {
			return EmptyResponseSample{}, nil
		}

		for _, employeePtr := range employees {
			employee := templateservice.BaseEmployeesDTO{
				Email:         employeePtr.Email,
				PreferredName: employeePtr.Preferred,
				FullName:      employeePtr.FullName,
				UniqueID:      employeePtr.UniqueIdentifier,
				ManagerEmail:  employeePtr.ManagersEmail,
			}

			err, _ := service.CreateEmployees(ctx, employee)
			if err != nil {
				return nil, fmt.Errorf("failed to create employee: %w", err)
			}
		}

		return "Employees received successfully", nil
	}
}
