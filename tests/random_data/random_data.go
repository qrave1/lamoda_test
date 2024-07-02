package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/qrave1/lamoda_test/config"
	"github.com/qrave1/lamoda_test/internal/infrastructure/persistence/postgres"
)

func main() {
	arg := os.Args[1]

	cfg, err := config.ReadConfig("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	cfg.Database.Host = "localhost"

	db, err := postgres.NewConnect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	switch arg {
	case "insert":
		err = insert(tx)
	case "truncate":
		err = truncate(tx)
	default:
		log.Fatalf("unknown command: %s", arg)
	}
}

func insert(tx *sql.Tx) (txErr error) {
	_, txErr = tx.Exec(`insert into warehouses(id, name, is_available) VALUES 
                                                  (1, 'sklad1', true),     
                                                  (2, 'sklad2', false),  
                                                  (3, 'sklad3', true),      
                                                  (4, 'sklad4', false),     
                                                  (5, 'sklad5', true)`,
	)
	if txErr != nil {
		log.Println(txErr)
		return
	}

	_, txErr = tx.Exec(`insert into products(id, name, code, size) VALUES 
									   			(1, 'tovar1', 1234, 123),
									   			(2, 'tovar2', 444, 45),
										   (3, 'tovar3', 322, 100),
										   (4, 'tovar4', 1337, 15),
										   (5, 'tovar5', 1984, 1000),
										   (6, 'tovar6', 21, 4401)`,
	)

	if txErr != nil {
		log.Println(txErr)
		return
	}

	_, txErr = tx.Exec(`insert into product_warehouse(warehouse_id, product_id, quantity) VALUES 
                                                                    	(1, 1, 15),
																		(1, 2, 20),
																		(2, 3, 15),
																		(3, 4, 1),
																		(3, 5, 20),
																		(3, 6, 100),
																		(4, 1, 6),
																		(5, 1, 1)
																		`,
	)

	_, txErr = tx.Exec(`insert into product_warehouse(warehouse_id, product_id, quantity, reserved_quantity) VALUES
                                                                                        (3, 1, 100, 80) `,
	)
	if txErr != nil {
		log.Println(txErr)
		return
	}

	return nil
}

func truncate(tx *sql.Tx) (txErr error) {
	_, txErr = tx.Exec(`truncate table product_warehouse, products, warehouses`)
	if txErr != nil {
		log.Println(txErr)
	}

	return
}
