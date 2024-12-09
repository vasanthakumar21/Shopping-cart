package inits

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=ep-rough-rice-a1smj83m.ap-southeast-1.aws.neon.tech user=neondb_owner password=HFRZW23DKrja dbname=shoppingcart port=5432 sslmode=require"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil  {
		log.Fatal("Failed to connect to the Database")
	}
}