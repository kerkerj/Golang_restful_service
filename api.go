package main

import (
	"database/sql"
	"fmt"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func root(r render.Render) {
	r.JSON(200, map[string]string{
		"path": "root",
		"say":  "hello!",
	})
}

func notFound(r render.Render) {
	r.JSON(404, map[string]string{
		"error": "404",
	})
}

func getPageAll(params martini.Params, r render.Render, db *sql.DB) {
	stmtOut, err := db.Query("SELECT * FROM mess_data")

	PanicIf(err, r)
	defer stmtOut.Close()

	wholeData := make([]*Page, 0)

	for stmtOut.Next() {
		dbRow := NewPage()

		err = stmtOut.Scan(
			&dbRow.Id,
			&dbRow.Test1,
			&dbRow.Test2,
			&dbRow.Test3,
			&dbRow.Created_at,
		)
		PanicIf(err, r)

		fmt.Printf("\nData fetched from DB:\n%v\n\n", dbRow)
		wholeData = append(wholeData, dbRow)
	}

	r.JSON(200, map[string]interface{}{
		"data": wholeData,
	})
}

func getPage(params martini.Params, r render.Render, db *sql.DB) {
	wholeData := getPageDB(params, r, db)

	//fmt.Printf("%v\n", wholeData)
	r.JSON(200, map[string]interface{}{
		"data": wholeData,
	})
}

func getPageDB(params martini.Params, r render.Render, db *sql.DB) []*Page {
	stmtOut, err := db.Query("SELECT * FROM mess_data WHERE id = ?", params["id"])

	PanicIf(err, r)
	defer stmtOut.Close()

	wholeData := make([]*Page, 0)

	for stmtOut.Next() {
		dbRow := NewPage()

		err = stmtOut.Scan(
			&dbRow.Id,
			&dbRow.Test1,
			&dbRow.Test2,
			&dbRow.Test3,
			&dbRow.Created_at,
		)
		PanicIf(err, r)

		fmt.Printf("\nData fetched from DB:\n%v\n\n", dbRow)
		wholeData = append(wholeData, dbRow)
	}

	return wholeData
}

func postPage(page Page, r render.Render, db *sql.DB) {
	stmt, err := db.Prepare("INSERT mess_data SET test1 = ?, test2 = ?, test3 = ?")
	PanicIf(err, r)

	res, err := stmt.Exec(page.Test1, page.Test2, page.Test3)
	PanicIf(err, r)

	id, err := res.LastInsertId()
	PanicIf(err, r)

	fmt.Println(id)
	params := map[string]string{
		"id": fmt.Sprintf("%v", id),
	}
	wholeData := getPageDB(params, r, db)

	r.JSON(200, map[string]interface{}{
		"data": wholeData,
	})
}

func putPage(page Page, params martini.Params, r render.Render, db *sql.DB) {
	stmt, err := db.Prepare("UPDATE mess_data SET test1 = ?, test2 = ?, test3 = ? WHERE id = ?")
	PanicIf(err, r)

	res, err := stmt.Exec(page.Test1, page.Test2, page.Test3, params["id"])
	PanicIf(err, r)

	_, err = res.LastInsertId()
	PanicIf(err, r)

	//fmt.Println(id)
	//params := map[string]string{
	//	"id": fmt.Sprintf("%v", id),
	//}
	wholeData := getPageDB(params, r, db)

	r.JSON(200, map[string]interface{}{
		"data": wholeData,
	})
}

func deletePage(page Page, params martini.Params, r render.Render, db *sql.DB) {
	stmt, err := db.Prepare("DELETE FROM mess_data WHERE id = ?")
	PanicIf(err, r)

	res, err := stmt.Exec(params["id"])
	PanicIf(err, r)

	affect, err := res.RowsAffected()
	PanicIf(err, r)

	fmt.Println(affect)

	r.JSON(200, map[string]interface{}{
		"delete_id": params["id"],
		"status":    affect,
	})
}
