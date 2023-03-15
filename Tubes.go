// Nama: Muhammad Ilman Dewantara
// NIM: 1303213069
package main

import "fmt"

func cumlaude(ipk float64, masaStudi int, publikasi bool) bool {
	if ipk >= 3.51 && ipk <= 4.00 && masaStudi <= 8 && publikasi {
		return true
	} else {
		return false
	}
}

func sangatMemuaskan(ipk float64, masaStudi int, publikasi bool) bool {
	if ipk >= 3.00 && ipk < 3.51 && masaStudi <= 14 {
		return true
	} else {
		return false
	}
}

func memuaskan(ipk float64, masaStudi int, publikasi bool) bool {
	if ipk >= 2.00 && ipk < 2.75 && masaStudi <= 14 {
		return true
	} else {
		return false
	}
}

func main() {
	var ipk float64
	var masaStudi int
	var publikasi bool
	fmt.Scan(&ipk, &masaStudi, &publikasi)

	if cumlaude(ipk, masaStudi, publikasi) {
		fmt.Println("Cumlaude")
	} else if sangatMemuaskan(ipk, masaStudi, publikasi) {
		fmt.Println("Sangat Memuaskan")
	} else if memuaskan(ipk, masaStudi, publikasi) {
		fmt.Println("Memuaskan")
	}
}
