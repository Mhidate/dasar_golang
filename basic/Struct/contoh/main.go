package main

import "fmt"

// Product merepresentasikan data barang
type Product struct {
	ID    int
	Nama  string
	Harga int
}

// Customer merepresentasikan data pelanggan
type Customer struct {
	ID   int
	Nama string
}

// OrderItem merepresentasikan item dalam order
type OrderItem struct {
	Product  Product
	Quantity int
}

// Order merepresentasikan pesanan
type Order struct {
	ID         int
	Customer   Customer
	OrderItems []OrderItem
}

// Method untuk menghitung total harga order
func (o Order) TotalHarga() int {
	total := 0
	for _, item := range o.OrderItems {
		total += item.Product.Harga * item.Quantity
	}
	return total
}

func main() {
	// Data produk
	produk1 := Product{ID: 1, Nama: "Laptop", Harga: 10000000}
	produk2 := Product{ID: 2, Nama: "Mouse", Harga: 150000}

	// Data customer
	customer1 := Customer{ID: 1, Nama: "Budi"}

	// Order Item
	item1 := OrderItem{Product: produk1, Quantity: 1}
	item2 := OrderItem{Product: produk2, Quantity: 2}

	// Order
	order1 := Order{
		ID:       1,
		Customer: customer1,
		OrderItems: []OrderItem{
			item1, item2,
		},
	}

	// Cetak detail order
	fmt.Println("Order ID:", order1.ID)
	fmt.Println("Customer:", order1.Customer.Nama)
	fmt.Println("Detail Produk:")

	for _, item := range order1.OrderItems {
		fmt.Printf("- %s (x%d) : Rp %d\n", item.Product.Nama, item.Quantity, item.Product.Harga*item.Quantity)
	}

	fmt.Println("Total Harga: Rp", order1.TotalHarga())
}
