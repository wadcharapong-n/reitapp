package test

import (
	"../models"
	"../services"
	"testing"
)

type Reit_Mock struct {}

func (self Reit_Mock) GetReitAll() ([]*models.ReitItem, error) {
	return nil, nil
}

func (self Reit_Mock) GetReitBySymbol(symbol string) (models.ReitItem, error) {
	return models.ReitItem{}, nil
}

func (self Reit_Mock) SaveReitFavorite(userId string, symbol string) error{
	return nil
}

func (self Reit_Mock) DeleteReitFavorite(userId string, ticker string) error {
	return nil
}


func (self Reit_Mock) GetReitFavoriteByUserIDJoin(userId string) []*models.FavoriteInfo {
	return nil
}

func TestGetReitAllProcess(t *testing.T) {
	result,_ := services.GetReitAllProcess(Reit_Mock{})
	if result == nil {
		println("OK")
	}else{
		println("Not OK")
	}
}

func TestGetReitBySymbolProcess(t *testing.T) {
	result,_ := services.GetReitAllProcess(Reit_Mock{})
	if result == nil {
		println("OK")
	}else{
		println("Not OK")
	}
}

func TestSaveReitFavoriteProcess(t *testing.T) {
	err := services.SaveReitFavoriteProcess(Reit_Mock{},"","")
	if err == nil {
		println("OK")
	}else{
		println("Not OK")
	}
}

func TestDeleteReitFavoriteProcess(t *testing.T) {
	err := services.DeleteReitFavoriteProcess(Reit_Mock{},"","")
	if err == nil {
		println("OK")
	}else{
		println("Not OK")
	}
}

func TestGetReitFavoriteByUserIDJoinProcess(t *testing.T) {
	result := services.GetReitFavoriteByUserIDJoinProcess(Reit_Mock{},"")
	if result == nil {
		println("OK")
	}else{
		println("Not OK")
	}
}