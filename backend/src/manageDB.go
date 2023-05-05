package src

import (
	"backend/src/database"
	"database/sql"
	"regexp"
)

func addPertanyaan(str string, db *sql.DB) (string, bool) {
	pattern := `Tambahkan pertanyaan\s*(.+)\s*dengan jawaban\s*(.+)\s*`

	// Compile the regex pattern
	re := regexp.MustCompile(pattern)

	// Find the matching X and Y substrings
	matches := re.FindStringSubmatch(str)

	if len(matches) == 0 {
		return "", false
	} else {
		isUpdate := false
		x := matches[1]
		y := matches[2]
		conn, _ := db.Query("SELECT PERTANYAAN FROM PRODUCTS")
		for conn.Next() {
			var q database.Product
			// conn.scan(&q.Pertanyaan)
			if q.Pertanyaan == x {
				isUpdate = true
				break
			}
		}

		if isUpdate == true {
			db.Query("UPDATE PRODUCTS SET JAWABAN = ? WHERE PERTANYAAN = ?", y, x)
			return "Pertanyaan" + x + "sudah ada!" + "jawaban diupdate ke " + y, true

		} else {
			db.Query("INSERT INTO PRODUCTS (PERTANYAAN, JAWABAN) VALUES (?, ?)", x, y)
			return "Pertanyaan" + x + "telah ditambah", false
		}
	}
}

// func hapusPertanyaan(str string) {

// }
