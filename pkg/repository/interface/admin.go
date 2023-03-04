package interfaces

import (
	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	"github.com/akshayur04/project-ecommerce/pkg/domain"
)

type AdminRepository interface {
	IsSuperAdmin(createrId int) (bool, error)
	CreateAdmin(admin helperStruct.CreateAdmin) (response.AdminData, error)
	AdminLogin(email string) (domain.Admins, error)
}
