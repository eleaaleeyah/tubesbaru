package main

import "fmt"

func main() {

	namaRoti := [10]string{
		"Coffee Bun", "Chocolate Bun ", "Butter Croissant", "Triple Cheese Croissant ", "Pizza",
		"Pain Suisse", "Chicken Floss Bun ", "Slice Brownies", "Pain Au Chocolate", "Donat  Kampung"}

	hargaRoti := [10]int{
		15000, 16000, 18000, 23000, 27000, 20000, 15000, 15000, 22000, 9000}

	namaCoffee := [10]string{
		"Americano", "Cappuccino", "Latte", "Vanilla Latte", "Mochaccino", "Pistachio Latte",
		"Strawberry Shake", "Matcha Latte", "Thai Tea", "Chocolate"}

	hargaCoffee := [10]int{
		15000, 20000, 22000, 25000, 23000, 25000, 20000, 28000, 23000, 24000}

	var pilihan [50]int
	var jumlah [50]int
	var kategori [50]int
	banyak := 0

	selesai := false

	for !selesai {

		fmt.Println("=================================")
		fmt.Println("       TOKO ROTI GrandIDN        ")
		fmt.Println("1. Menu Roti")
		fmt.Println("2. Menu Coffee")
		fmt.Println("3. Bayar")
		fmt.Println("=================================")
		fmt.Print("Pilih menu: ")

		var menu int
		fmt.Scan(&menu)
		fmt.Println()

		if menu == 1 {
			fmt.Println("===== MENU ROTI =====")
			for i := 0; i < 10; i++ {
				fmt.Println(i+1, namaRoti[i], "-", hargaRoti[i])
			}
			fmt.Println()

			loopRoti := 0
			for loopRoti == 0 {

				var pilih int
				fmt.Print("pilih Roti (0 = kembali): ")
				fmt.Scan(&pilih)

				if pilih == 0 {
					loopRoti = 1

				} else if pilih < 1 || pilih > 10 {
					fmt.Println("pilih menu yang tersedia!\n")
				} else {
					var jml int
					fmt.Print("jumlah: ")
					fmt.Scan(&jml)

					pilihan[banyak] = pilih - 1
					jumlah[banyak] = jml
					kategori[banyak] = 1
					banyak++

					fmt.Println("\ntelah ditambahkan:", namaRoti[pilih-1], "x", jml)
				}
			}
		}

		if menu == 2 {
			fmt.Println("===== MENU COFFEE =====")
			for i := 0; i < 10; i++ {
				fmt.Println(i+1, namaCoffee[i], "-", hargaCoffee[i])
			}
			fmt.Println()

			loopCof := 0
			for loopCof == 0 {

				var pilih int
				fmt.Print("pilih Coffe (0 = kembali): ")
				fmt.Scan(&pilih)

				if pilih == 0 {
					loopCof = 1

				} else if pilih < 1 || pilih > 10 {
					fmt.Println("pilih menu yang tersedia!\n")

				} else {
					var jml int
					fmt.Print("jumlah: ")
					fmt.Scan(&jml)

					pilihan[banyak] = pilih - 1
					jumlah[banyak] = jml
					kategori[banyak] = 2
					banyak++

					fmt.Println("\ntelah ditambahkan:", namaCoffee[pilih-1], "x", jml)
				}
			}
		}

		if menu == 3 {

			if banyak == 0 {
				fmt.Println("belum ada yang ditambahkan!\n")
			} else {

				totalBelanja := 0
				totalRoti := 0
				totalCoffee := 0
				diskon := 0

				fmt.Println("=========== STRUK BELANJA ===========")

				for i := 0; i < banyak; i++ {

					if kategori[i] == 1 {
						nama := namaRoti[pilihan[i]]
						harga := hargaRoti[pilihan[i]]
						sub := harga * jumlah[i]

						fmt.Println(nama, "x", jumlah[i], "=", sub)
						totalBelanja += sub
						totalRoti += jumlah[i]

					} else if kategori[i] == 2 {
						nama := namaCoffee[pilihan[i]]
						harga := hargaCoffee[pilihan[i]]
						sub := harga * jumlah[i]

						fmt.Println(nama, "x", jumlah[i], "=", sub)
						totalBelanja += sub
						totalCoffee += jumlah[i]
					}
				}

				if totalBelanja > 250000 {
					diskon = totalBelanja / 10
				}

				fmt.Println("-------------------------------------")
				fmt.Println("Total Belanja =", totalBelanja)
				fmt.Println("Total Diskon  =", diskon)
				fmt.Println("Total Bayar =", totalBelanja-diskon)

				if totalRoti%10 == 0 && totalRoti != 0 {
					fmt.Println("You Got 2 Bun Free")
				}
				if totalCoffee%10 == 0 && totalCoffee != 0 {
					fmt.Println("You Got 1 Thai Tea")
				}
				if totalBelanja > 250000 {
					diskon = totalBelanja / 10
				}
				fmt.Println("            Terima kasih!            ")

				fmt.Println("=====================================\n")

				banyak = 0
				selesai = true
			}
		}
	}
}
