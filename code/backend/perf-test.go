package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "dicdoc_service"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sizes := []int{10, 50, 100, 250, 500, 1000, 2500, 5000, 7500, 10000, 25000, 50000, 100000, 250000, 500000, 750000, 1000000}
	// Генерация данных и измерение времени выполнения запросов
	for _, v := range sizes {
		insertData(db, v)
		measureQueryPerformance(db, v)
		deleteData(db)
	}
}

func insertData(db *sql.DB, count int) {
	for i := 0; i < count; i++ {
		db.Exec("INSERT INTO public.appointments VALUES ((SELECT MAX(id) FROM appointments) + 1, 200, 200, '2024-12-03 19:00')")
	}
}

func deleteData(db *sql.DB) {
	db.Exec("DELETE FROM public.appointments")
}

func measureQueryPerformance(db *sql.DB, count int) {
	var rows *sql.Rows
	var err error

	start := time.Now()
	for i := 0; i < 10000; i++ {
		rows, err = db.Query("SELECT *, random() as rand_val FROM appointments WHERE doctorid = $1 and patientid = $2", 200, 200)
		rows.Close()
	}
	elapsed := time.Since(start)
	if err != nil {
		log.Fatal(err)
	}

	var resultCount int
	for rows.Next() {
		resultCount++
	}

	fmt.Printf("Time taken for %d records: %s\n", count, elapsed/1000)
}
