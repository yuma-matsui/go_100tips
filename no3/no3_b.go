package no3

import "database/sql"

/* init関数ではなく通常の関数で定義することでいかが可能になる
* エラー処理の責任を呼び出し元に任せることができる
 */

func Solution(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
