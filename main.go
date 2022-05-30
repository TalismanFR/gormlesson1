package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
	"gormlesson/persistence/migration"
	"gormlesson/persistence/model"
	"gormlesson/persistence/repository"
	"log"
	"os"
	"time"
)

func main() {
	db, err := gorm.Open("postgres", "user=web password=1234567 port=54320 dbname=payment sslmode=disable")
	db.LogMode(true)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(os.Args) > 1 {
		val := os.Args[1]
		switch val {
		case "migrate":
			migration.Migrate(db)
		case "insert":
			repo := repository.NewPayRepository(db)
			pay := &model.Pay{PayRequest: model.PayRequest{Description: "Оплата заказа", Amount: 10000, InvoiceId: uuid.NewV4()},
				CreatedAt: datatypes.Date(time.Now()), Uuid: uuid.NewV4()}
			_ = pay.MarshallPayRequest()
			err := repo.Save(pay)
			if err != nil {
				log.Fatalln(err)
				return
			}
		}
	}

	defer db.Close()

}
