package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func Test_checkSQL(t *testing.T) {
	tests := []struct {
		query    string
		expected bool
	}{
		// 有効なクエリ
		{"SELECT name FROM users LIMIT 1;", true},
		{"SELECT age FROM employees LIMIT 1;", true},
		{"SELECT address FROM customers WHERE id = 1 LIMIT 1;", true},

		// 無効なクエリ
		{"SELECT * FROM users LIMIT 1;", false},                 // ワイルドカード使用
		{"SELECT name, age FROM users LIMIT 1;", false},         // 複数カラム指定
		{"SELECT name FROM users;", false},                      // LIMIT 1がない
		{"SELECT name FROM users LIMIT 1", false},               // セミコロンがない
		{"INSERT INTO users (name) VALUES ('John');", false},    // SELECT以外のクエリ
		{"UPDATE users SET name = 'Jane' WHERE id = 1;", false}, // SELECT以外のクエリ
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Testing query: %s", test.query), func(t *testing.T) {
			result := checkSQL(test.query)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v for query: %s", test.expected, result, test.query)
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
	// Mockデータベースと期待される動作のセットアップ
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %s", err)
	}
	defer db.Close()

	// テストケース 1: 正常なクエリ
	t.Run("Valid query", func(t *testing.T) {
		query := "SELECT name FROM users WHERE id = 1"
		mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("John Doe"))

		result, err := queryExec(db, query)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
		if result != "John Doe" {
			t.Errorf("Expected 'John Doe', but got '%s'", result)
		}
	})

	// テストケース 2: 行が見つからない場合
	t.Run("No rows", func(t *testing.T) {
		query := "SELECT name FROM users WHERE id = 2"
		mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)

		result, err := queryExec(db, query)
		if err == nil {
			t.Errorf("Expected an error, but got none")
		}
		if result != "" {
			t.Errorf("Expected empty result, but got '%s'", result)
		}
	})

	// テストケース 3: クエリの構文エラー
	t.Run("Query syntax error", func(t *testing.T) {
		query := "SELECT name FROM non_existing_table"
		mock.ExpectQuery(query).WillReturnError(sqlmock.ErrCancelled)

		result, err := queryExec(db, query)
		if err == nil {
			t.Errorf("Expected an error, but got none")
		}
		if result != "" {
			t.Errorf("Expected empty result, but got '%s'", result)
		}
	})

	// Mockの期待通りに動作したか確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unmet expectations: %s", err)
	}
}
