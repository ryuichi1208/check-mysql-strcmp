package main

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mackerelio/checkers"
)

func Test_newDB(t *testing.T) {
	tests := []struct {
		name    string
		want    *sql.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("newDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSingleColumnSelect(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSingleColumnSelect(tt.args.query); got != tt.want {
				t.Errorf("isSingleColumnSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidSQL(t *testing.T) {
	type args struct {
		db    *sql.DB
		query string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSQL(tt.args.db, tt.args.query); got != tt.want {
				t.Errorf("isValidSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readQueryFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readQueryFromFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("readQueryFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readQueryFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryExec(t *testing.T) {
	type args struct {
		db    *sql.DB
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := queryExec(tt.args.db, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryExec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("queryExec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDo(t *testing.T) {
	tests := []struct {
		name string
		want *checkers.Checker
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
