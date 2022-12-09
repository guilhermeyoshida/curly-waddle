package app

import (
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	type fields struct {
		Name   string
		Price  float64
		Status string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test_Method_Product_Enable",
			fields: fields{
				Name:   "Hello",
				Price:  10,
				Status: DISABLED,
			},
			wantErr: false,
		},
		{
			name: "Test_Method_Product_Enable_Error",
			fields: fields{
				Name:   "Hello",
				Price:  0,
				Status: DISABLED,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{
				Name:   tt.fields.Name,
				Price:  tt.fields.Price,
				Status: tt.fields.Status,
			}
			if err := p.Enable(); (err != nil) != tt.wantErr {
				t.Errorf("Enable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProduct_Disable(t *testing.T) {
	type fields struct {
		ID     string
		Name   string
		Price  float64
		Status string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test_Method_Product_Disable",
			fields: fields{
				Name:   "Hello",
				Price:  0,
				Status: ENABLED,
			},
			wantErr: false,
		},
		{
			name: "Test_Method_Product_Disable_Error",
			fields: fields{
				Name:   "Hello",
				Price:  10,
				Status: ENABLED,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{
				ID:     tt.fields.ID,
				Name:   tt.fields.Name,
				Price:  tt.fields.Price,
				Status: tt.fields.Status,
			}
			if err := p.Disable(); (err != nil) != tt.wantErr {
				t.Errorf("Disable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProduct_IsValid(t *testing.T) {
	type fields struct {
		ID     string
		Name   string
		Price  float64
		Status string
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "Test_Method_Product_IsValid",
			fields: fields{
				ID:     uuid.NewV4().String(),
				Name:   "Hello",
				Price:  10,
				Status: ENABLED,
			},
			wantErr: false,
			want:    true,
		},
		{
			name: "Test_Method_Product_IsValid_Error_INVALID",
			fields: fields{
				ID:     uuid.NewV4().String(),
				Name:   "Hello",
				Price:  10,
				Status: "INVALID",
			},
			wantErr: true,
			want:    false,
		},
		{
			name: "Test_Method_Product_IsValid_Error_LESS_THAN_ZERO",
			fields: fields{
				ID:     uuid.NewV4().String(),
				Name:   "Hello",
				Price:  -10,
				Status: ENABLED,
			},
			wantErr: true,
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{
				ID:     tt.fields.ID,
				Name:   tt.fields.Name,
				Price:  tt.fields.Price,
				Status: tt.fields.Status,
			}
			got, err := p.IsValid()
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsValid() got = %v, want %v", got, tt.want)
			}
		})
	}
}
