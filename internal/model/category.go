package model

import (
	"errors"
	"fmt"
	"github.com/dembygenesis/local.tools/internal/sysconsts"
	"github.com/dembygenesis/local.tools/internal/utilities/errs"
	"github.com/dembygenesis/local.tools/internal/utilities/validationutils"
	"strings"
)

type UpdateCategory struct {
	Id                int    `json:"id" validate:"required,greater_than_zero"`
	CategoryTypeRefId int    `json:"category_type_ref_id"`
	Name              string `json:"name"`
}

type DeleteCategory struct {
	ID int `json:"id" validate:"required,greater_than_zero"`
}

type RestoreCategory struct {
	ID int `json:"id" validate:"required,greater_than_zero"`
}

func (c *UpdateCategory) Validate() error {
	var errList errs.List
	if err := validationutils.Validate(c); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	hasAtLeastOneUpdateParameters := false
	if c.CategoryTypeRefId > 0 {
		hasAtLeastOneUpdateParameters = true
	} else {
		errList.Add(sysconsts.ErrCapturePageSetId)
	}

	if strings.TrimSpace(c.Name) != "" {
		hasAtLeastOneUpdateParameters = true
	} else {
		errList.Add(sysconsts.ErrCategoryTypeRefIdInvalid)
	}

	if !hasAtLeastOneUpdateParameters {
		return errors.New(sysconsts.ErrHasNotASingleValidateUpdateParameter)
	}

	return nil
}

func (c *UpdateCategory) ToCategory() *Category {
	return &Category{
		Id:                c.Id,
		CategoryTypeRefId: c.CategoryTypeRefId,
		Name:              c.Name,
	}
}

type CreateCategory struct {
	CategoryTypeRefId int    `json:"category_type_ref_id" validate:"required,greater_than_zero"`
	Name              string `json:"name" validate:"required"`
}

func (c *CreateCategory) ToCategory() *Category {
	return &Category{
		Name:              c.Name,
		CategoryTypeRefId: c.CategoryTypeRefId,
	}
}

func (c *CreateCategory) Validate() error {
	if err := validationutils.Validate(c); err != nil {
		return fmt.Errorf("validate: %v", err)
	}
	return nil
}

type Category struct {
	Id                int    `json:"id" boil:"id"`
	CategoryTypeRefId int    `json:"category_type_ref_id" boil:"category_type_ref_id" swaggerignore:"true"`
	Name              string `json:"name" boil:"name"`
	CategoryType      string `json:"category_type" boil:"category_type"`
}

func (c *Category) Validate() error {
	if err := validationutils.Validate(c); err != nil {
		return fmt.Errorf("validate: %v", err)
	}
	return nil
}

type CategoryType struct {
	Id   int    `json:"id" boil:"id"`
	Name string `json:"name" validate:"required" boil:"name"`
}

type Categories []Category

type PaginatedCategories struct {
	Categories []Category  `json:"categories"`
	Pagination *Pagination `json:"pagination"`
}

type CategoryFilters struct {
	CategoryNameIn         []string `query:"category_name_in" json:"category_name_in"`
	CategoryTypeNameIn     []string `query:"category_type_name_in" json:"category_type_name_in"`
	CategoryTypeIdIn       []int    `query:"category_type_id_in" json:"category_type_id_in"`
	CategoryIsActive       []int    `query:"is_active" json:"is_active"`
	IdsIn                  []int    `query:"ids_in" json:"ids_in"`
	PaginationQueryFilters `swaggerignore:"true"`
}

func (c *CategoryFilters) Validate() error {
	if err := c.ValidatePagination(); err != nil {
		return fmt.Errorf("pagination: %v", err)
	}
	if err := validationutils.Validate(c); err != nil {
		return fmt.Errorf("category filters: %v", err)
	}

	return nil
}
