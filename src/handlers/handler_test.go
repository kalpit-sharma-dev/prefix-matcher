package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"prefix-matcher/src/mocks"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestHandleSearchWord(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIService(ctrl)
	searchHandler := NewHandler(mockService)

	url := fmt.Sprintf("/search/prefix/%s", "kalpit")

	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err, "get longest common prefix request")
	req = mux.SetURLVars(req, map[string]string{
		"word": "kalpit",
	})

	tests := []struct {
		name           string
		setupMocks     func()
		urls           string
		rq             *http.Request
		expectedStatus int
	}{
		{

			name: "Success",
			setupMocks: func() {
				mockService.EXPECT().SearchWord(gomock.Any()).Return("", nil)
			},
			expectedStatus: 200,
			urls:           url,
			rq:             req,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			rr := httptest.NewRecorder()

			searchHandler.HandleSearchWord(rr, tt.rq)
			if !reflect.DeepEqual(rr.Code, tt.expectedStatus) {
				t.Errorf("%s got unexpected result: got %v want %v", tt.name, rr.Code, tt.rq)
			}
		})
	}
}
