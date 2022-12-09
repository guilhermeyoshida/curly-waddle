package app_test

import (
	"github.com/golang/mock/gomock"
	"github.com/guilhermeyoshida/curly-waddle/app"
	mockApp "github.com/guilhermeyoshida/curly-waddle/app/mocks"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Persistence *mockApp.MockProductPersistenceInterface
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *mockApp.MockProductInterface
		wantErr bool
	}{
		{
			name:    "Test_Method_Product_Service_Get",
			fields:  fields{Persistence: mockApp.NewMockProductPersistenceInterface(ctrl)},
			args:    args{id: uuid.NewV4().String()},
			want:    mockApp.NewMockProductInterface(ctrl),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Persistence.EXPECT().Get(gomock.Any()).Return(tt.want, nil).AnyTimes()
			p := &app.ProductService{
				Persistence: tt.fields.Persistence,
			}
			got, err := p.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
