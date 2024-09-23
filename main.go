package main

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/mackerelio/checkers"
)

type options struct {
	DB_USER    string `short:"u" long:"user" description:"mysql user" default:"root" required:"false"`
	DB_HOST    string `short:"h" long:"host" description:"mysql host" default:"localhost" required:"true"`
	DB_PORT    string `short:"p" long:"port" description:"mysql port" default:"3306" required:"false"`
	QUERY_FILE string `short:"f" long:"file" description:"query file" required:"true"`
	VALUE      string `short:"v" long:"value" description:"value" required:"true"`
}

var opts options

func newDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s@tcp(%s:%s)/", opts.DB_USER, opts.DB_HOST, opts.DB_PORT))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func isSingleColumnSelect(query string) bool {
	// 正規表現でSELECT句を解析し、カラム数を確認
	re := regexp.MustCompile(`(?i)^SELECT\s+([^,]+?)\s+FROM`)
	match := re.FindStringSubmatch(query)
	if len(match) == 0 {
		return false // SELECT句が不正またはカラム数が1でない場合
	}

	// カラム名にワイルドカード(*)が含まれていないかを確認
	column := strings.TrimSpace(match[1])
	if column == "*" {
		return false // ワイルドカードが使用されている場合
	}

	return true
}

func isValidSQL(db *sql.DB, query string) bool {
	// クエリをトリムして大文字に変換し、SELECTで始まるかを確認
	trimmedQuery := strings.TrimSpace(query)
	if !strings.HasPrefix(strings.ToUpper(trimmedQuery), "SELECT") {
		fmt.Printf("Rejected non-SELECT query: %s\n", query)
		return false
	}

	// クエリを検証するためにPrepareを使用
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Printf("Invalid SQL: %s\nError: %s\n", query, err)
		return false
	}
	// ステートメントを閉じる
	stmt.Close()

	// クエリがSELECTで1カラムのみを指定しているかを確認
	if !isSingleColumnSelect(trimmedQuery) {
		fmt.Printf("Rejected SELECT query with more than 1 column: %s\n", query)
		return false
	}

	return true
}

// ファイルからクエリを読み込んで返す
func readQueryFromFile(filename string) (string, error) {
	// ファイルを開く
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// ファイルの内容を全て読み取る
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	if !isValidSQL(nil, string(content)) {
		return "", fmt.Errorf("Invalid SQL: %s", content)
	}

	return string(content), nil
}

// クエリを実行して取得した文字列を返す
func queryExec(db *sql.DB, query string) (string, error) {
	var result string
	err := db.QueryRow(query, "condition_value").Scan(&result)
	if err != nil {
		return "", err
	}
	return result, nil
}

func Do() *checkers.Checker {
	// ファイルからクエリを読み込む
	query, err := readQueryFromFile(opts.QUERY_FILE)
	if err != nil {
		return checkers.Critical("Failed to read query file")
	}

	db, err := newDB()
	if db == nil && err != nil {
		return checkers.Critical("Failed to connect to database")
	}
	defer db.Close()

	result, err := queryExec(db, query)
	if err != nil {
		return checkers.Critical("Failed to execute query")
	}

	if result != opts.VALUE {
		return checkers.Critical("Not expected value")
	}

	return checkers.Ok("OK")
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		checkers.Critical("not correct value")
		fmt.Println(opts)
		os.Exit(1)
	}

	if chk := Do(); chk != nil {
		chk.Exit()
	}
}
