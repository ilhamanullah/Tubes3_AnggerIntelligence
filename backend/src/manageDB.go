package src

import (
	"backend/src/database"
	"database/sql"
	"regexp"
)

func AddPertanyaan(str string, db *sql.DB) string {
	pattern := `[t|T]ambahkan pertanyaan (.+) dengan jawaban (.+)`

	// Compile the regex pattern
	re := regexp.MustCompile(pattern)

	// Find the matching X and Y substrings
	matches := re.FindStringSubmatch(str)

	if len(matches) == 0 {
		return ""
	} else {
		isUpdate := false
		x := matches[1]
		y := matches[2]
		conn, _ := db.Query("SELECT PERTANYAAN FROM PRODUCTS")
		for conn.Next() {
			var q database.Product
			conn.Scan(&q.Pertanyaan)
			if q.Pertanyaan == x {
				isUpdate = true
				break
			}
		}

		if isUpdate {
			db.Query("UPDATE PRODUCTS SET JAWABAN = ? WHERE PERTANYAAN = ?", y, x)
			return "Pertanyaan " + x + " sudah ada! jawaban diupdate ke " + y

		} else {
			db.Query("INSERT INTO PRODUCTS (PERTANYAAN, JAWABAN) VALUES (?, ?)", x, y)
			return "Pertanyaan " + x + " telah ditambah"
		}
	}
}

func hapusPertanyaan(str string, db *sql.DB) string {
	pattern := `[H|h]apus pertanyaan (.+)`

	// Compile the regex pattern
	re := regexp.MustCompile(pattern)

	// Find the matching X and Y substrings
	matches := re.FindStringSubmatch(str)

	if len(matches) == 0 {
		return ""
	} else {
		x := matches[1]
		conn, _ := db.Query("SELECT PERTANYAAN FROM PRODUCTS")
		for conn.Next() {
			var q database.Product
			conn.Scan(&q.Pertanyaan)
			// fmt.Print(q.Pertanyaan)
			// fmt.Println("++++++++")
			if q.Pertanyaan == x {
				db.Query("DELETE FROM PRODUCTS WHERE PERTANYAAN = ?", x)
				return "Pertanyaan " + x + " telah dihapus"
			}
		}
		return "Tidak ada pertanyaan " + x + " pada database"
	}
}

func findJawaban(str string, db *sql.DB, isKmp bool) string {
	if isKmp {
		// fmt.Println("KMP")
		conn, _ := db.Query("SELECT PERTANYAAN, JAWABAN FROM PRODUCTS")
		str = str[1:]
		for conn.Next() {
			var q database.Product
			conn.Scan(&q.Pertanyaan, &q.Jawaban)
			if kmpMatch(q.Pertanyaan, str) == 0 {
				return q.Jawaban
				conn, _ = db.Query("SELECT JAWABAN FROM PRODUCTS WHERE PERTANYAAN = ?", str)
			}
		}
		for conn.Next() {
			var q database.Product
			conn.Scan(&q.Pertanyaan, &q.Jawaban)
			if distance(q.Pertanyaan, str) >= 90 {
				return q.Jawaban
				conn, _ = db.Query("SELECT JAWABAN FROM PRODUCTS WHERE PERTANYAAN = ?", str)
			}
		}
	} else {
		// fmt.Println("BM")
		conn, _ := db.Query("SELECT PERTANYAAN, JAWABAN FROM PRODUCTS")
		str = str[1:]
		for conn.Next() {
			var q database.Product
			conn.Scan(&q.Pertanyaan, &q.Jawaban)
			if bmMatch(q.Pertanyaan, str) == 0 {
				return q.Jawaban
				conn, _ = db.Query("SELECT JAWABAN FROM PRODUCTS WHERE PERTANYAAN = ?", str)
			}
		}
		for conn.Next() {
			var q database.Product
			conn.Scan(&q.Pertanyaan, &q.Jawaban)
			if distance(q.Pertanyaan, str) >= 90 {
				return q.Jawaban
				conn, _ = db.Query("SELECT JAWABAN FROM PRODUCTS WHERE PERTANYAAN = ?", str)
			}
		}
	}
	conn, _ := db.Query("SELECT PERTANYAAN, JAWABAN FROM PRODUCTS")
	var s string
	s = "Pertanyaan tidak ditemukan di database.Apakah maksud anda : "
	for conn.Next() {
		var q database.Product
		conn.Scan(&q.Pertanyaan, &q.Jawaban)
		if distance(q.Pertanyaan, str) >= 40 {
			s += q.Pertanyaan + ", "
		}
	}
	return s
}
