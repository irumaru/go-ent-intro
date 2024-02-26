package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/irumaru/go-ent-intro/ent"
	"github.com/irumaru/go-ent-intro/ent/migrate"
)

func main() {
	//jst, _ := time.LoadLocation("Asia/Tokyo")

	dcs := mysql.Config{
		Addr:                 "db:3306",
		User:                 "dev",
		DBName:               "dev",
		Passwd:               "dev",
		Net:                  "tcp",
		AllowNativePasswords: true,
		Collation:            "utf8mb4_general_ci",
	}

	//fmt.Println(dcs.FormatDSN())

	// Open the connection to the database (MySQL).
	client, err := ent.Open("mysql", dcs.FormatDSN())
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Create a new user.
	u, _ := CreateUser(context.Background(), client)
	log.Println(u)
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(16).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}
