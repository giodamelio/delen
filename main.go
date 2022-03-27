package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/giodamelio/delen/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	boil.SetDB(db)
	boil.DebugMode = true

	ctx := context.Background()
	users, err := models.Users().AllG(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Listing all users")
	for _, user := range users {
		fmt.Println(user)
	}

	fmt.Println("Get user by id")
	user, err := models.Users(Where("id == ?", 4)).OneG(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(user)

	fmt.Println("Count the users")
	count, err := models.Users().CountG(ctx)
	fmt.Println(count)

	fmt.Println("Insert a user")
	var newUser models.User
	newUser.Name = null.StringFrom("Gio")
	newUser.Email = null.StringFrom("giodamelio@gmail.com")
	err = newUser.InsertG(ctx, boil.Infer())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Count the users")
	secondCount, err := models.Users().CountG(ctx)
	fmt.Println(secondCount)

	fmt.Println("Delete all users named Gio")
	// rowsAffected, err := models.Users(Where("name == Gio")).DeleteAll(ctx, db)
	rowsAffected, err := models.Users(models.UserWhere.Name.EQ(null.StringFrom("Gio"))).DeleteAllG(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Deleted %d users\n", rowsAffected)
}
