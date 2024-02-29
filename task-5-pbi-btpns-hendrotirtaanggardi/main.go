package main

import (
	"fmt"
	"net/http"
	"pbi-hendrotirta-btpns/mod/database"
	"pbi-hendrotirta-btpns/mod/router"
)

func main() {
	// Inisialisasi koneksi database
	database.InitDB()
	fmt.Println("Koneksi basis data berhasil diinisialisasi")

	// Inisialisasi router
	router.InitRouter()
	fmt.Println("Router berhasil diinisialisasi")

	fmt.Println("Aplikasi berhasil dimulai")

	// Jalankan server
	http.ListenAndServe(":8000", nil)
}
