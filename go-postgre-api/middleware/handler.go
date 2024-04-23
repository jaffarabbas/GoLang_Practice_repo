package middleware

import (
	"database/sql"
	"encoding/json"
	"go-postgre-api/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"omitempty"`
	Message string `json:"message,omitempty"`
}

func GetConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRE_URL"))

	if err != nil {
		panic("Error connecting to the database")
	}

	err = db.Ping()

	if err != nil {
		panic("Error pinging the database")
	}

	return db
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Error getting the ID from the request %v", err)
	}

	stock, err := getStock(int64(id))

	if err != nil {
		log.Fatalf("Error getting the stock %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func GetStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := getStocks()

	if err != nil {
		log.Fatalf("Error getting the stocks %v", err)
	}

	json.NewEncoder(w).Encode(stocks)
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatal("Error reading the request body %v", err)
	}
	insertID := insertStock(stock)

	res := response{
		ID:      insertID,
		Message: "Stock created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Error getting the ID from the request %v", err)
	}

	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Error reading the request body %v", err)
	}

	updatedRows := updateStock(int64(id), stock)

	res := response{
		ID:      updatedRows,
		Message: "Stock updated successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Error getting the ID from the request %v", err)
	}

	deleteRows := deleteStock(int64(id))

	res := response{
		ID:      deleteRows,
		Message: "Stock deleted successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func insertStock(stock models.Stock) int64 {
	db := GetConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO stocks (name, price, company) VALUES ($1, $2, $3) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatalf("Error inserting the stock %v", err)
	}

	return id
}

func getStock(id int64) (models.Stock, error) {
	db := GetConnection()

	defer db.Close()

	var stock models.Stock

	sqlStatement := `SELECT * FROM stocks WHERE id=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&stock.ID, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		log.Printf("No stock with the ID %d", id)
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Error getting the stock %v", err)
	}

	return stock, err
}

func getStocks() ([]models.Stock, error) {
	db := GetConnection()

	defer db.Close()

	var stocks []models.Stock

	sqlStatement := `SELECT * FROM stocks`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Error getting the stocks %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var stock models.Stock

		err = rows.Scan(&stock.ID, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Fatalf("Error scanning the row %v", err)
		}

		stocks = append(stocks, stock)
	}

	return stocks, err
}

func updateStock(id int64, stock models.Stock) int64 {
	db := GetConnection()

	defer db.Close()

	sqlStatement := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Fatalf("Error updating the stock %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error getting the rows affected %v", err)
	}

	return rowsAffected
}

func deleteStock(id int64) int64 {
	db := GetConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM stocks WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Error deleting the stock %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error getting the rows affected %v", err)
	}

	return rowsAffected
}
