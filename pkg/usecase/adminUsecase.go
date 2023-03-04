package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCae struct {
	adminRepo interfaces.AdminRepository
}

func NewAdminUsecase(adminRepo interfaces.AdminRepository) services.AdminUsecase {
	return &adminUseCae{
		adminRepo: adminRepo,
	}
}

func (c *adminUseCae) CreateAdmin(ctx context.Context, admin helperStruct.CreateAdmin, createrId int) (response.AdminData, error) {
	IsSuper, err := c.adminRepo.IsSuperAdmin(createrId)
	if err != nil {
		return response.AdminData{}, err
	}
	if !IsSuper {
		return response.AdminData{}, fmt.Errorf("not a super admin")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return response.AdminData{}, err
	}
	admin.Password = string(hash)
	adminData, err := c.adminRepo.CreateAdmin(admin)

	return adminData, err
}

func (c *adminUseCae) AdminLogin(admin helperStruct.LoginReq) (string, error) {
	adminData, err := c.adminRepo.AdminLogin(admin.Email)
	if err != nil {
		return "", err
	}

	if adminData.Email == "" {
		return "", fmt.Errorf("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminData.Password), []byte(admin.Password))
	if err != nil {
		return "", err
	}

	if adminData.IsBlocked {
		return "", fmt.Errorf("user is blocked")
	}

	claims := jwt.MapClaims{
		"id":  adminData.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return ss, nil
}
