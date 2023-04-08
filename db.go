package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

func main() {
	// Connect to DB
	conn, err := sql.Open("pgx", "host=localhost dbname=todolist user=postgres password= postgres port=5432")
	if err != nil {
		log.Fatal(err)
	}
	//always close the connection
	defer conn.Close()
	//let check the pool
	err = conn.Ping()
	if err != nil {
		log.Fatal("Connection Error", err)
	}
	fmt.Println("Connected to the DB Successfully ")

	//Create (Insert) Row DB
	//qry := `insert into products(model, company, price) values ($1, $2, $3)`
	//_, err = conn.Exec(qry, "Golang is Awesome", "Golang", 25666556)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println("Data inserted")

	////Update DB "Company"
	//qry := `update products set company = $1 where id = $2`
	//_, err = conn.Exec(qry, "Go is Awesome Updated", 2)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println("Data Updated")
	//get all row: R

	// Delete Row DB
	//qry := `delete from products where company = $1` //номер
	//_, err = conn.Exec(qry, "Golang")                //номер id
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println("Data deleted")

	err = FetchAllPosts(conn)
	if err != nil {
		log.Println(err)
	}
}

func FetchAllPosts(conn *sql.DB) error {
	qry := `select * from products` //id, model, company, price
	rows, err := conn.Query(qry)
	if err != nil {
		log.Println(err)
		return err
	}

	for rows.Next() {
		var id, price int
		var model, company string
		//Scan result into variables
		err := rows.Scan(&id, &model, &company, &price)
		if err != nil {
			log.Println("err while scaning from row")
		}
		fmt.Printf("#%d Model: %s Company: %s Price: %s\n", id, model, company, price)
	}
	// because we are fetching multiple rows
	defer rows.Close()
	return nil
}
