package test

import (
	"../models"
	"../services"
	"testing"
)

type ReitMock struct {}

func (self ReitMock) GetUserProfileByCriteria(userId string, site string) models.UserProfile {
	panic("implement me")
}

func (self ReitMock) SaveUserProfile(profile *models.UserProfile) string {
	panic("implement me")
}

func (self ReitMock) GetReitAll() ([]*models.ReitItem, error) {
	return nil, nil
}

func (self ReitMock) GetReitBySymbol(symbol string) (models.ReitItem, error) {
	return models.ReitItem{}, nil
}

func (self ReitMock) SaveReitFavorite(userId string, symbol string) error{
	return nil
}

func (self ReitMock) DeleteReitFavorite(userId string, ticker string) error {
	return nil
}


func (self ReitMock) GetReitFavoriteByUserIDJoin(userId string) []*models.FavoriteInfo {
	return nil
}

func TestGetReitAllProcess(t *testing.T) {
	result,_ := services.GetReitAllProcess(ReitMock{})
	if result == nil {
		println("OK")
	}else{
		println("Not OK")
	}
}

func TestGetReitBySymbolProcess(t *testing.T) {
	result,_ := services.GetReitAllProcess(ReitMock{})
	if result == nil {
		println("OK")
	}else{
		println("Not OK")
	}
}

func TestSaveReitFavoriteProcess(t *testing.T) {
	err := services.SaveReitFavoriteProcess(ReitMock{},"","")
	if err == nil {
		println("OK")
	}else{
		println("Not OK")
	}
}

func TestDeleteReitFavoriteProcess(t *testing.T) {
	err := services.DeleteReitFavoriteProcess(ReitMock{},"","")
	if err == nil {
		println("OK")
	}else{
		println("Not OK")
	}
}

func TestGetReitFavoriteByUserIDJoinProcess(t *testing.T) {
	result := services.GetReitFavoriteByUserIDJoinProcess(ReitMock{},"")
	if result == nil {
		println("OK")
	}else{
		println("Not OK")
	}
}