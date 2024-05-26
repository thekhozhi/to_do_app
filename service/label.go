package service

import (
	"context"
	"fmt"
	"to_do_app/api/models"
	"to_do_app/storage"
)

type labelService struct {
	storage storage.IStorage
}

func NewLabelService(storage storage.IStorage) labelService {
	return labelService{
		storage: storage,
	}
}

func (l labelService) Create(ctx context.Context, req models.CreateLabel) (models.Label, error) {
	label, err := l.storage.Label().Create(ctx,req)
	if err != nil{
		fmt.Println("Error in service layer, while creating label!", err.Error())
		return models.Label{}, err
	}

	return label, nil
}

func (l labelService) Get(ctx context.Context, req models.PrimaryKey) (models.Label, error) {
	label, err := l.storage.Label().Get(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, while getting label!", err.Error())
		return models.Label{}, err
	}

	return label, nil
}

func (l labelService) Update(ctx context.Context, req models.UpdateLabel) (models.UpdateResponse, error) {
	label, err := l.storage.Label().Update(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, while updating label!", err.Error())
		return models.UpdateResponse{},err
	}

	return label, nil
}

func (l labelService) Delete(ctx context.Context, req models.PrimaryKey) (error) {
	err := l.storage.Label().Delete(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, while deleting label!", err.Error())
		return err
	}
	
	return nil
}