package models

import (
	"capstone/app/database"
	"capstone/app/entity"
	"errors"

	"github.com/labstack/echo/v4"
)

func SaveModule(c echo.Context, u *entity.Module) (entity.Module, error) {
	db := database.GetDB(c)

	err := db.Debug().Create(&u)
	if err != nil {
		return entity.Module{}, err.Error
	}

	if err.RowsAffected == 0 {
		return entity.Module{}, errors.New("Failed to add Module")
	}

	return *u, nil
}

func FindAllModules(c echo.Context) ([]entity.Module, error) {
	var modules []entity.Module
	db := database.GetDB(c)

	err := db.Debug().Limit(100).Find(&modules)
	if err.Error != nil {
		return modules, err.Error
	}
	return modules, nil
}

func FindModuleById(c echo.Context, id uint32) (entity.Module, error) {
	var u entity.Module
	db := database.GetDB(c)

	err := db.Debug().First(&u, "id = ?", id)
	if err.Error != nil {
		return entity.Module{}, errors.New("Module not Found")
	}

	return u, nil
}

func UpdateModule(c echo.Context, id string, module *entity.Module) (int64, error) {
	db := database.GetDB(c)

	err := db.Model(&entity.Module{}).Where("id = ?", id).Updates(module)

	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func DeleteModule(c echo.Context, id string) (int64, error) {
	db := database.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.Module{})
	if err.Error != nil || err.RowsAffected == 0 {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
