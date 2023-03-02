package service

import (
	"context"
	"prefix-matcher/src/mocks"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSearchWord(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIRepository(ctrl)
	service := NewService(mockRepo)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		word       string
		setupMocks func()
		wantResp   string
		wantErr    bool
	}{
		{
			name: "Success",
			word: "kalpit",
			setupMocks: func() {
				mockRepo.EXPECT().SearchWord(gomock.Any()).Return(1)
			},
			wantResp: "k",
			wantErr:  false,
		},
		{
			name: "Failure",
			word: "kalpit",
			setupMocks: func() {
				mockRepo.EXPECT().SearchWord(gomock.Any()).Return(0)
			},
			wantResp: "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			gotResp, err := service.SearchWord(tt.word)

			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("got = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}

}
