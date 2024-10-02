package main

import "fmt"

func main() {
	var (
		nama string
		prodi string = "S1-IF"
		kelas string = "IF-12-01"
		nim int
	)
	fmt.Println("Masukkan nama")
	fmt.Scan(&nama)

	fmt.Println("Masukkan kelas")
	fmt.Scan(&kelas)

	fmt.Println("Masukkan NIM")
	fmt.Scan(&nim)

	fmt.Println("Perkenalkan saya adalah", nama, "salah satu mahasiswa prodi", prodi, "dari kelas", kelas, "dengan NIM", nim)
}