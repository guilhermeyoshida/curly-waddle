package app_test

import (
	"github.com/golang/mock/gomock"
	"github.com/guilhermeyoshida/curly-waddle/app"
	mockApp "github.com/guilhermeyoshida/curly-waddle/app/mocks"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"reflect"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	persistence := mockApp.NewMockProductPersistenceInterface(ctrl)
	product := mockApp.NewMockProductInterface(ctrl)
	type fields struct {
		Persistence *gomock.Call
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
			fields:  fields{Persistence: persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()},
			args:    args{id: uuid.NewV4().String()},
			want:    mockApp.NewMockProductInterface(ctrl),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &app.ProductService{
				Persistence: persistence,
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

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Persistence *mockApp.MockProductPersistenceInterface
	}
	type args struct {
		name  string
		price float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *mockApp.MockProductInterface
		wantErr bool
	}{
		{
			name:    "Test_Method_Product_Service_Create",
			fields:  fields{Persistence: mockApp.NewMockProductPersistenceInterface(ctrl)},
			args:    args{name: "Product", price: rand.Float64()},
			want:    mockApp.NewMockProductInterface(ctrl),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Persistence.EXPECT().Save(gomock.Any()).Return(tt.want, nil).AnyTimes()
			p := &app.ProductService{
				Persistence: tt.fields.Persistence,
			}
			got, err := p.Create(tt.args.name, tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductService_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Persistence *mockApp.MockProductPersistenceInterface
	}
	type args struct {
		product *mockApp.MockProductInterface
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *mockApp.MockProductInterface
		wantErr bool
	}{
		{
			name:    "Test_Method_Product_Service_Enable",
			fields:  fields{Persistence: mockApp.NewMockProductPersistenceInterface(ctrl)},
			args:    args{product: mockApp.NewMockProductInterface(ctrl)},
			want:    mockApp.NewMockProductInterface(ctrl),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Persistence.EXPECT().Save(gomock.Any()).Return(tt.want, nil).AnyTimes()
			p := &app.ProductService{
				Persistence: tt.fields.Persistence,
			}
			tt.args.product.EXPECT().Enable().Return(nil)
			got, err := p.Enable(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enable() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductService_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Persistence *mockApp.MockProductPersistenceInterface
	}
	type args struct {
		product *mockApp.MockProductInterface
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *mockApp.MockProductInterface
		wantErr bool
	}{
		{
			name:    "Test_Method_Product_Service_Disable",
			fields:  fields{Persistence: mockApp.NewMockProductPersistenceInterface(ctrl)},
			args:    args{product: mockApp.NewMockProductInterface(ctrl)},
			want:    mockApp.NewMockProductInterface(ctrl),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Persistence.EXPECT().Save(gomock.Any()).Return(tt.want, nil).AnyTimes()
			p := &app.ProductService{
				Persistence: tt.fields.Persistence,
			}
			tt.args.product.EXPECT().Disable().Return(nil)
			got, err := p.Disable(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("Disable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Disable() got = %v, want %v", got, tt.want)
			}
		})
	}
}
