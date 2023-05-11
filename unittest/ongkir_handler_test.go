package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Bhinneka/go-rajaongkir"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"ebook/internal/application/usecase"
	"ebook/internal/application/usecase/mocks"
	"ebook/internal/delivery/http/handler"
)

func TestOngkirHandler_CalculateOngkir(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock OrderUsecase
	mockUsecase := mocks.NewMockOrderUsecase(ctrl)
	mockUsecase.EXPECT().GetOrder(gomock.Any()).Return(&usecase.Order{}, nil)

	handler := handler.OngkirHandler{
		OrderUsecase: mockUsecase,
	}

	// Set up Echo
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up request body
	body := `{
		"origin": "Jakarta",
		"shipping_address": "Bandung",
		"total_shipping": 10,
		"courier": "jne"
	}`
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Body = http.NoBody
	if body != "" {
		req.Body = httptest.NewRecorder().Body
		req.Body.Write([]byte(body))
		req.Body = http.NoBody
	}

	// Set up RajaOngkir API
	ro := rajaongkir.New("736cb684012893b91477f360458ea29d", 10*time.Second)

	// Mock the response from RajaOngkir API
	mockCost := ro.Cost{
		Code:    "mock_code",
		Name:    "mock_name",
		Costs:   nil,
		Results: nil,
	}
	mockResult := rajaongkir.QueryResult{
		DestinationDetails: nil,
		OriginDetails:      nil,
		Code:               200,
		Status:             rajaongkir.SuccessStatus,
	}
	mockResponse := rajaongkir.QueryResponse{
		Costs:  []rajaongkir.Cost{mockCost},
		Result: &mockResult,
	}

	// Mock the call to RajaOngkir API
	mockRajaOngkir := mocks.NewMockRajaOngkir(ctrl)
	mockRajaOngkir.EXPECT().GetCost(gomock.Any()).Return(rajaongkir.QueryResponse{
		Costs: []rajaongkir.Cost{mockCost},
		Result: &rajaongkir.QueryResult{
			DestinationDetails: nil,
			OriginDetails:      nil,
			Code:               200,
			Status:             rajaongkir.SuccessStatus,
		},
	}, nil)

	// Call the handler function
	err := handler.CalculateOngkir()(c)
	assert.NoError(t, err)

	// Check the response
	assert.Equal(t, http.StatusOK, rec.Code)
}
