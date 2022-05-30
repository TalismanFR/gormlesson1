package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/gormigrate.v1"
	"gorm.io/datatypes"
	"log"
)

func Migrate(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20220530_01",
			Migrate: func(db *gorm.DB) error {
				type Pay struct {
					Uuid      uuid.UUID      `gorm:"primary_key;type:uuid;"`
					CreatedAt datatypes.Date `gorm:"type:timestamp"`
				}
				return db.AutoMigrate(&Pay{}).Error
			},
			Rollback: func(db *gorm.DB) error {
				return db.DropTable("pays").Error
			},
		},
		{
			ID: "20220530_02",
			Migrate: func(db *gorm.DB) error {

				type Pay struct {
					PayRequest postgres.Jsonb `gorm:"type:jsonb;column:pay_request"`
				}

				return db.AutoMigrate(&Pay{}).Error
			},
			Rollback: func(db *gorm.DB) error {
				return db.Table("pays").DropColumn("pay_request").Error
			},
		},
		{
			ID: "20220530_03",
			Migrate: func(db *gorm.DB) error {
				return db.Exec("CREATE INDEX \"pay_requestInd\" ON \"public\".\"pays\" USING gin (pay_request);").
					Error
			},
			Rollback: func(db *gorm.DB) error {
				return db.Table("pays").RemoveIndex("pay_requestInd").Error
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalln(err)
	}

	log.Println("Migration success")
}
