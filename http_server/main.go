package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Name string
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func newPost(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Post{})

	// Create
	// db.Create(&Post{Name: "marcelo"})

	// Read
	var posts []Post

	db.Find(&posts)

	// db.First(&product, 1)
	// db.First(&product, "name = ?", "marcelo")

	// Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	// db.Model(&product).Updates(Post{Name: "caio"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	// db.Delete(&product, 1)

	http.HandleFunc("/newpost", newPost)

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
