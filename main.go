package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

type Product struct {
  gorm.Model
  Code  string `gorm:"column:sku"`
  Price uint
}

func printProduct(mesg string, product Product) {
    println(mesg, product.ID, product.Code, product.Price)
}

func main() {
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&Product{})
  println("BD creada")

  // Create
  product1 := Product{Code: "D42", Price: 100}
  db.Create(&product1)

  product2 := Product{Code: "C10", Price: 30}
  db.Create(&product2)
  //db.Create(&Product{Code: "D42", Price: 100})
  println("Producto creado")

  // Read
  var product Product
  result := db.First(&product, 1) // find product with integer primary key
  println("# Registros", result.RowsAffected)
  printProduct("Producto s1", product)

  var products2 Product
  // Se debe indicar el nombre de la columna sku 
  db.First(&products2, "sku = ?", "C10") // find product with code D42
  printProduct("Producto s2", products2)

  // Update - update product's price to 200
  db.Model(&product).Update("Price", 200)
  printProduct("Producto update simple", product)
  // Update - update multiple fields
  db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
  printProduct("Producto update complex 1", product)
  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
  printProduct("Producto update complex 2", product)
  // Delete - delete product
  result2 := db.Delete(&product, 1)
  print("Registros borrados", result2.RowsAffected)
}
