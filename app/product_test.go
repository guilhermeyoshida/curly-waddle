package app

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
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
			name: "test_method_product_enable_ok",
			fields: fields{
				ID:     "1",
				Name:   "hello",
				Price:  10,
				Status: DISABLED,
			},
			wantErr: false,
		},
		{
			name: "test_method_product_enable_not_ok",
			fields: fields{
				ID:     "1",
				Name:   "hello",
				Price:  0,
				Status: DISABLED,
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
			if err := p.Enable(); (err != nil) != tt.wantErr {
				require.Nil(t, err)
			}
		})
	}
}
