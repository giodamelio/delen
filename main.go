package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/giodamelio/delen/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", r)
}

func dbStuff() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	boil.SetDB(db)
	boil.DebugMode = true

	ctx := context.Background()

	fmt.Println("Insert an item")
	var newItem models.Item
	newItem.Name = "Gio"
	newItem.Contents = null.BytesFrom([]byte{1, 2, 3})
	err = newItem.InsertG(ctx, boil.Infer())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	items, err := models.Items().AllG(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Listing all items")
	for _, item := range items {
		fmt.Println(item)
	}

	fmt.Println("Get item by id")
	item, err := models.Items(Where("id == ?", 1)).OneG(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(item)

	fmt.Println("Count the items")
	count, err := models.Items().CountG(ctx)
	fmt.Println(count)

	fmt.Println("Insert a item")
	var anotherNewitem models.Item
	anotherNewitem.Name = "Gio"
	anotherNewitem.Contents = null.BytesFrom([]byte{1, 2, 3})
	err = anotherNewitem.InsertG(ctx, boil.Infer())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Count the items")
	secondCount, err := models.Items().CountG(ctx)
	fmt.Println(secondCount)

	fmt.Println("Delete all items named Gio")
	rowsAffected, err := models.Items(models.ItemWhere.Name.EQ("Gio")).DeleteAllG(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Deleted %d items\n", rowsAffected)
}
