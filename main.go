package main

import (
	"fmt"
	"log"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	// Memeriksa apakah URL permintaan sama dengan "/test"
	if r.URL.Path != "/test" {
		// Jika URL tidak sesuai, kirim respons 404 Not Found
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Memeriksa apakah metode permintaan adalah GET
	if r.Method != "GET" {
		// Jika metode bukan GET, kirim respons "method is not supported" dengan status 405 Method Not Allowed
		http.Error(w, "metode tidak didukung", http.StatusMethodNotAllowed)
		return
	}

	// Jika semua pemeriksaan berhasil, kirim file "static/test.html" sebagai respons
	http.ServeFile(w, r, "static/test.html")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// Memeriksa apakah parsing form data dari permintaan berhasil
	if err := r.ParseForm(); err != nil {
		// Jika parsing gagal, kirim pesan kesalahan ke client
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}

	// Mengirim respons bahwa permintaan POST berhasil
	fmt.Fprintf(w, "Permintaan POST berhasil")

	// Mendapatkan nilai-nilai dari form data yang dikirim dalam permintaan
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Mengirim data nama dan alamat ke client sebagai bagian dari respons
	fmt.Fprintf(w, "Nama = %s\n", name)
	fmt.Fprintf(w, "Alamat = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Memeriksa apakah URL permintaan sesuai dengan "/hello"
	if r.URL.Path != "/hello" {
		// Jika URL tidak sesuai, kirim respons 404 Not Found
		http.Error(w, "404 tidak ditemukan", http.StatusNotFound)
		return
	}

	// Memeriksa apakah metode permintaan adalah GET
	if r.Method != "GET" {
		// Jika metode bukan GET, kirim respons "metode tidak didukung" dengan status 405 Method Not Allowed
		http.Error(w, "metode tidak didukung", http.StatusMethodNotAllowed)
		return
	}

	// Jika semua pemeriksaan berhasil, kirim respons "hello!" ke client
	fmt.Fprintf(w, "Halo!")
}

func main() {
	// Membuat server file untuk menangani file statis dalam direktori "./static"
	fileServer := http.FileServer(http.Dir("./static"))
	// Mengarahkan permintaan root "/" ke server file
	http.Handle("/", fileServer)

	// Menghubungkan fungsi-fungsi penanganan khusus ke rute tertentu
	http.HandleFunc("/form", formHandler)   // Menghubungkan "/form" ke fungsi formHandler
	http.HandleFunc("/hello", helloHandler) // Menghubungkan "/hello" ke fungsi helloHandler
	http.HandleFunc("/test", testHandler)   // Menghubungkan "/test" ke fungsi testHandler

	// Menampilkan pesan bahwa server sedang berjalan pada port 8080
	fmt.Printf("Memulai server di port 8080\n")

	// Mengeksekusi server HTTP pada port 8080, dan menangani kesalahan jika ada
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
