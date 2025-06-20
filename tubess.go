package main
import "fmt"

const NMAX int = 100
type item struct{
	nama string
	kategori string
	cuaca string
	acara string
}

type ootd [NMAX]item

func main(){
	var outfit ootd
	var nData int
	var pilih int
	
	selesai := false
	fmt.Println("=======================================")
	fmt.Println("    👗👕 WELCOME TO OOTD PLANNER 👖👟  ")
	fmt.Println("=======================================")
	fmt.Println("        /\\             /\\")
    fmt.Println("       /  \\___________/  \\")
    fmt.Println("      /   /           \\   \\")
    fmt.Println("     /   /             \\   \\")
    fmt.Println("    /   /|             |\\   \\")
    fmt.Println("   /   / |             | \\   \\")
    fmt.Println("  /___/  |             |  \\___\\")
    fmt.Println("  |  |   |     o o     |   |  |")
    fmt.Println("  |  |   |             |   |  |")
    fmt.Println("  |  |   |     o o     |   |  |")
    fmt.Println("  |__|   |_____________|   |__|")
    fmt.Println("        /               \\")
    fmt.Println("       /_________________\\")
	fmt.Println("=======================================")
	fmt.Println("✨ Aplikasi manajemen outfit harianmu ✨")

	fmt.Println("📅 Buat, simpan, dan atur gaya kamu setiap hari!")

	fmt.Println("Memuat aplikasi...")

	fmt.Println("\n✅ Siap digunakan!")
	for !selesai{
		fmt.Println("\n--- OOTD Planner ---")
		fmt.Println("1. Tambah Item Fashion")
		fmt.Println("2. Lihat Semua Item")
		fmt.Println("3. Rekomendasi Outfit")
		fmt.Println("4. Urutkan data")
		fmt.Println("5. Edit Item Fashion")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)
		
		if pilih == 1{
			inputData(&outfit, &nData)
		}else if pilih == 2{
			cetakData(outfit, nData)
		}else if pilih == 3{
			rekomendasi(outfit, nData)
		}else if pilih == 4{
			urutData(&outfit, nData)
			cetakData(outfit, nData)
		}else if pilih == 5{
			editData(&outfit, &nData)
		} else if pilih == 6 {
			selesai = true
		}
	}
	fmt.Println("Terima kasih! Hope you enjoy your outfit. See you!🥰")
}

func inputData(A *ootd, n *int){
	var tambah int
	fmt.Println("Berapa banyak item yang akan ditambahkan?")
	fmt.Scan(&tambah) 
	for i := 0; i < tambah; i++{
		fmt.Println("Nama item(gunakan _ jika lebih dari 2 kata), Kategori(Atasan/Bawahan/Sepatu), Dan Acara(Formal/Kasual/Pesta): ")
		fmt.Scan(&A[*n].nama, &A[*n].kategori, &A[*n].acara)
		fmt.Println("Input berhasil 😊")
		*n++
	}
	fmt.Printf("Total item sekarang : %d\n", *n)
}

func cetakData(A ootd, n int){
   fmt.Printf("%-20s %-15s %-15s\n", "Nama Item", "Kategori", "Acara")
	for i := 0; i < 51; i++ {
        fmt.Print("-")
    }
	fmt.Println()
    for i := 0; i < n; i++ {
        fmt.Printf("%-20s %-15s %-15s\n", A[i].nama, A[i].kategori, A[i].acara)
    }
}

func urutData(A *ootd, n int) {
	var pilihan string
	fmt.Printf("Diurutkan berdasarkan apa? (Nama / Kategori / Acara): ")
	fmt.Scan(&pilihan)

	if pilihan == "Nama" {
		// Selection sort berdasarkan nama
		for i := 0; i < n-1; i++ {
			idxMin := i
			for j := i + 1; j < n; j++ {
				if A[j].nama < A[idxMin].nama {
					idxMin = j
				}
			}
			if idxMin != i {
				temp := A[i]
				A[i] = A[idxMin]
				A[idxMin] = temp
			}
		}
	} else if pilihan == "Kategori" {
		// Insertion sort berdasarkan kategori
		for i := 1; i < n; i++ {
			temp := A[i]
			j := i
			for j > 0 && temp.kategori < A[j-1].kategori {
				A[j] = A[j-1]
				j--
			}
			A[j] = temp
		}
	} else if pilihan == "Acara" {
		// Selection sort berdasarkan acara
		for i := 0; i < n-1; i++ {
			idxMin := i
			for j := i + 1; j < n; j++ {
				if A[j].acara < A[idxMin].acara {
					idxMin = j
				}
			}
			if idxMin != i {
				temp := A[i]
				A[i] = A[idxMin]
				A[idxMin] = temp
			}
		}
	} 
}

func rekomendasi(A ootd, n int) {
	var targetAcara string
	var ditemukan bool = false

	fmt.Print("Masukkan jenis acara (cth: Formal, Kasual, Pesta): ")
	fmt.Scan(&targetAcara)
	
	fmt.Printf("\n--- Rekomendasi untuk Acara '%s' ---\n", targetAcara)
	fmt.Printf("%-20s %-15s %-15s\n", "Nama Item", "Kategori", "Acara")

	for k := 0; k < 51; k++ {
		fmt.Print("-")
	}
	fmt.Println()


	for i := 0; i < n; i++ {
		if A[i].acara == targetAcara {
			fmt.Printf("%-20s %-15s %-15s\n", A[i].nama, A[i].kategori, A[i].acara)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Printf("Tidak ditemukan item yang cocok untuk acara '%s'.\n", targetAcara)
	}
}

func editData (A *ootd, n *int) {
	var target string
	var pilihan, ubah int
	var found bool = false 
	
	fmt.Printf("\nPilih apa yang ingin dilakukan\n")
	fmt.Printf("1. Edit\n")
	fmt.Printf("2. Hapus\n")
	fmt.Print("Opsi nomor (1/2)? ")
	fmt.Scan(&pilihan)
	
	if pilihan == 1{
		fmt.Printf("Masukkan nama item yang ingin diubah: \n")
		fmt.Scan(&target)
		for i := 0; i < *n; i++ {
			if A[i].nama == target {
				fmt.Printf("Item ditemukan: %s - %s - %s\n", A[i].nama, A[i].kategori, A[i].acara)
				fmt.Printf("\nApa yang ingin anda ubah:\n 1. Nama\n 2. Kategori\n 3. Acara\n")
				fmt.Print("Pilih menu: ")
				fmt.Scan(&ubah)
				if ubah == 1{
					fmt.Print("Masukkan nama baru (gunakan _ jika lebih dari 2 kata): ")
					fmt.Scan(&A[i].nama)
					fmt.Println("Nama item berhasil diubah")
				}else if ubah == 2{
					fmt.Print("Masukkan kategori baru (Atasan/Bawahan/Sepatu): ")
					fmt.Scan(&A[i].kategori)
					fmt.Println("Kategori item berhasil diubah")
				}else if ubah == 3{
					fmt.Print("Masukkan acara baru (Formal/Kasual/Pesta): ")
					fmt.Scan(&A[i].acara)
					fmt.Println("Acara Item berhasil diubah")
				}
			}
		}
	}else if pilihan == 2{
		fmt.Printf("Masukkan nama item yang ingin dihapus: \n")
		fmt.Scan(&target)
		for i := 0 ; i< *n ; i++ {
			if A[i].nama == target {
				for j:= i ; j<*n -1 ; j++ {
				A[j] = A [j+1]
				}
			*n-- 
			found = true 
			i = *n 
			}
		}
		if found {
			fmt.Printf("Item '%s' berhasil dihapus.\n", target)
			fmt.Printf("Total item setelah dihapus : %d\n", *n)
		}
	}
}