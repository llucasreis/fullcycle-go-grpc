package service

import (
	"context"

	"github.com/llucasreis/fullcycle-go-grpc/internal/database"
	"github.com/llucasreis/fullcycle-go-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var response []*pb.Category

	for _, c := range categories {
		categoryResponse := &pb.Category{
			Id:          c.ID,
			Name:        c.Name,
			Description: c.Description,
		}

		response = append(response, categoryResponse)
	}

	return &pb.CategoryList{Categories: response}, nil
}
