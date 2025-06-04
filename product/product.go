package product

import (
	"app/db"
	"fmt"
)

func AddProduct(name string) {
	_, err := db.DB.Exec("INSERT INTO products (name) VALUES ($1)", name)
	if err != nil {
		fmt.Println("❌ Failed to add product:", err)
		return
	}
	fmt.Println("✅ Product added:", name)
}

func ShowProducts() {
	rows, err := db.DB.Query("SELECT name FROM products")
	if err != nil {
		fmt.Println("❌ Failed to fetch products:", err)
		return
	}
	defer rows.Close()

	fmt.Println("📦 Products:")
	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Println("-", name)
	}
}
