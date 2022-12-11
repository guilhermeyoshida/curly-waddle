package db

import (
	"database/sql"
	"github.com/guilhermeyoshida/curly-waddle/app"
	"log"
	"reflect"
	"testing"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc", "Product Test", 0, "disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	type fields struct {
		db *sql.DB
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *app.Product
		wantErr bool
	}{
		{
			name:   "Test_Method_Product_Db_Get",
			fields: fields{db: Db},
			args:   args{id: "abc"},
			want: &app.Product{
				ID:     "abc",
				Name:   "Product Test",
				Price:  0,
				Status: app.DISABLED,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProductDb{
				db: tt.fields.db,
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

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	product := app.NewProduct()
	product.Name = "Product Test"
	product.Price = 25
	type fields struct {
		db *sql.DB
	}
	type args struct {
		product *app.Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *app.Product
		wantErr bool
	}{
		{
			name:    "Test_Method_Product_Db_Save",
			fields:  fields{db: Db},
			args:    args{product: product},
			want:    product,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProductDb{
				db: tt.fields.db,
			}
			got, err := p.Save(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Save() got = %v, want %v", got, tt.want)
			}
		})
	}
}
