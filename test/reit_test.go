package test

import (
	"../api"
	"../route"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"../models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)


type Reit struct {
	reitItems []models.ReitItem
	reitItem models.ReitItem
	reitFavorite []*models.FavoriteInfo
	err error
}

func (self Reit) GetUserFromToken(c echo.Context) (string, string) {
	panic("implement me")
}

func (self Reit) GetUserProfile(c echo.Context) error {
	token := c.Request().Header.Get("Authorization");
	if token == "Bearer token" {
		user := models.UserProfile{"1", "testUsername", "TestFullName", "Test@test.com", "testURL", "testSite"}
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusNoContent,"Not Found User")

}

func (self Reit) GetReitAll(c echo.Context) error {
	item := models.ReitItem{"1",
		"ทรัสต์เพื่อการลงทุนในสิทธิการเช่าอสังหาริมทรัพย์โกลเด้นเวนเจอร์",
		"Golden Ventures Leasehold Real Estate Investment Trust",
		"GVREIT",
		"บริษัทหลักทรัพย์จัดการกองทุน กสิกรไทย จำกัด",
		"",
		"8,046,150,000 บาท",
		"",
		"",
		"บริษัท ยูนิเวนเจอร์ รีท แมเนจเม้นท์ จำกัด",
		"1",
		"1",
		"1",
		"1",
		"1",
		"1",
		"1",
		"1",
		"1",
		"xx"}
	item2 := models.ReitItem{"2",
		"",
		"",
		"LHSC",
		"บริษัทหลักทรัพย์จัดการกองทุน ไทยพาณิชย์ จำกัด",
		"non-address",
		"",
		"",
		"",
		"",
		"2",
		"2",
		"2",
		"2",
		"2",
		"2",
		"2",
		"2",
		"2",
		"yy"}
	self.reitItems = []models.ReitItem{item,item2}
	return c.JSON(http.StatusOK,self.reitItems)
}

func (self Reit) GetReitBySymbol(c echo.Context) error {
	self.reitItem = models.ReitItem{"1",
		"ทรัสต์เพื่อการลงทุนในสิทธิการเช่าอสังหาริมทรัพย์โกลเด้นเวนเจอร์",
		"Golden Ventures Leasehold Real Estate Investment Trust",
		"GVREIT",
		"บริษัทหลักทรัพย์จัดการกองทุน กสิกรไทย จำกัด",
		"",
		"8,046,150,000 บาท",
		"",
		"",
		"บริษัท ยูนิเวนเจอร์ รีท แมเนจเม้นท์ จำกัด",
		"1",
		"1",
		"1",
		"1",
		"1",
		"1",
		"1",
		"1",
		"1",
		"xx"}
	symbol := c.Param("symbol");
	if(symbol == "GVREIT"){
		return c.JSON(http.StatusOK,self.reitItem)
	}
	return c.JSON(http.StatusNoContent,nil)
}

func (Reit) GetFavoriteReitAll(c echo.Context) error {
	panic("implement me")
}

func (Reit) DeleteFavoriteReit(c echo.Context) error {
	userID := c.FormValue("userId")
	ticker := c.FormValue("Ticker")
	if userID == "1" && ticker == "GVREIT" {
		return c.String(http.StatusOK, "success")

	}
	return c.String(http.StatusBadRequest, "fail")
}

func (Reit) SaveFavoriteReit(c echo.Context) error {
	userID := c.FormValue("userId")
	ticker := c.FormValue("Ticker")
	if userID == "1" && ticker == "GVREIT" {
		return c.String(http.StatusOK, "success")

	}
	return c.String(http.StatusBadRequest, "fail")
}



//func TestRouterReitWithSuccess(t *testing.T) {
//	r := route.Init()
//	ts := httptest.NewServer(r)
//	defer ts.Close()
//
//	res, err := http.Get(ts.URL + "/api/reit")
//	if err != nil {
//		t.Fatal(err)
//	}
//	assert.Equal(t, 200, res.StatusCode, "OK response is expected")
//}

func TestGetReitAll(t *testing.T)  {
	// Setup
	e := route.Init()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/reit")

	var userJSON = `[{"ID":"1","TrustNameTh":"ทรัสต์เพื่อการลงทุนในสิทธิการเช่าอสังหาริมทรัพย์โกลเด้นเวนเจอร์","TrustNameEn":"Golden Ventures Leasehold Real Estate Investment Trust","Symbol":"GVREIT","Trustee":"บริษัทหลักทรัพย์จัดการกองทุน กสิกรไทย จำกัด","Address":"","InvestmentAmount":"8,046,150,000 บาท","EstablishmentDate":"","RegistrationDate":"","ReitManager":"บริษัท ยูนิเวนเจอร์ รีท แมเนจเม้นท์ จำกัด","ParValue":"1","CeilingValue":"1","FloorValue":"1","PeValue":"1","ParNAV":"1","Policy":"1","PriceOfDay":"1","MaxPriceOfDay":"1","MinPriceOfDay":"1","NickName":"xx"},{"ID":"2","TrustNameTh":"","TrustNameEn":"","Symbol":"LHSC","Trustee":"บริษัทหลักทรัพย์จัดการกองทุน ไทยพาณิชย์ จำกัด","Address":"non-address","InvestmentAmount":"","EstablishmentDate":"","RegistrationDate":"","ReitManager":"","ParValue":"2","CeilingValue":"2","FloorValue":"2","PeValue":"2","ParNAV":"2","Policy":"2","PriceOfDay":"2","MaxPriceOfDay":"2","MinPriceOfDay":"2","NickName":"yy"}]
`

	var methodTest api.ReitController
	methodTest = Reit{}
	// Assertions
	if assert.NoError(t, methodTest.GetReitAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestGetReitBySymbol(t *testing.T)  {
	// Setup
	e := route.Init()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/reit/:symbol")
	c.SetParamNames("symbol")
	c.SetParamValues("GVREIT")
	userJSON := `{"ID":"1","TrustNameTh":"ทรัสต์เพื่อการลงทุนในสิทธิการเช่าอสังหาริมทรัพย์โกลเด้นเวนเจอร์","TrustNameEn":"Golden Ventures Leasehold Real Estate Investment Trust","Symbol":"GVREIT","Trustee":"บริษัทหลักทรัพย์จัดการกองทุน กสิกรไทย จำกัด","Address":"","InvestmentAmount":"8,046,150,000 บาท","EstablishmentDate":"","RegistrationDate":"","ReitManager":"บริษัท ยูนิเวนเจอร์ รีท แมเนจเม้นท์ จำกัด","ParValue":"1","CeilingValue":"1","FloorValue":"1","PeValue":"1","ParNAV":"1","Policy":"1","PriceOfDay":"1","MaxPriceOfDay":"1","MinPriceOfDay":"1","NickName":"xx"}
`
	var methodTest api.ReitController
	methodTest = Reit{}
	// Assertions
	if assert.NoError(t, methodTest.GetReitBySymbol(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestDeleteFavoriteReit(t *testing.T)  {
	// Setup
	e := route.Init()
	f := make(url.Values)
	f.Set("userId", "1")
	f.Set("Ticker", "GVREIT")
	req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(f.Encode()))
	req.Form = f
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/reitFavorite")


	var methodTest api.ReitController
	methodTest = Reit{}
	// Assertions
	if assert.NoError(t, methodTest.DeleteFavoriteReit(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestSaveFavoriteReit(t *testing.T)  {
	// Setup
	e := route.Init()
	f := make(url.Values)
	f.Set("userId", "1")
	f.Set("Ticker", "GVREIT")
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/reitFavorite")


	var methodTest api.ReitController
	methodTest = Reit{}
	// Assertions
	if assert.NoError(t, methodTest.SaveFavoriteReit(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUserProfile(t *testing.T)  {
	// Setup
	e := route.Init()
	f := make(url.Values)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	req.Header.Add("Authorization","Bearer token")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/profile")
	userJSON := `{"UserID":"1","UserName":"testUsername","FullName":"TestFullName","Email":"Test@test.com","Image":"testURL","Site":"testSite"}
`
	var methodTest api.ReitController
	methodTest = Reit{}
	// Assertions
	if assert.NoError(t, methodTest.GetUserProfile(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestGetUserProfileNotFound(t *testing.T)  {
	// Setup
	e := route.Init()
	f := make(url.Values)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	req.Header.Add("Authorization","Bearer")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/profile")
	userJSON := `"Not Found User"
`
	var methodTest api.ReitController
	methodTest = Reit{}
	// Assertions
	if assert.NoError(t, methodTest.GetUserProfile(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
