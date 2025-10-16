package main

import "fmt"

const NMAX int = 20

type barang struct {
	id, nama    string
	stok, harga int
}

type tabBarang [NMAX]barang

func main() {
	var data tabBarang
	var nData int
	var id, nama string
	var pilih int

	var running bool = true
	for running {

		menu()
		fmt.Print("Pilih (1/2/3/4/5/6/7): ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			fmt.Print("Masukkan banyak barang: ")
			fmt.Scan(&nData)
			tambah(&data, nData)
		} else if pilih == 2 {
			fmt.Println("Masukkan ID dan Nama: ")
			fmt.Scan(&id, &nama)
			editD(&data, nData, id, nama)

		} else if pilih == 3 {
			fmt.Println("Masukkan ID barang: ")
			fmt.Scan(&id)
			hapusD(&data, &nData, id)

		} else if pilih == 4 {
			carikategori(&data, nData)

		} else if pilih == 5 {
			fmt.Println("Data diurutkan berdasarkan stok: ")
			urutStok(&data, nData)
			cetak(data, nData)

			fmt.Println("Data diurutkan berdasarkan harga: ")
			urutHarga(&data, nData)
			cetak(data, nData)

		} else if pilih == 6 {
			cetak(data, nData)

		} else if pilih == 7 {
			running = false

		} else {
			fmt.Print("Input ulang, pilihan tidak ada")
		}
	}
}

func menu() {
	fmt.Println("                          ")
	fmt.Println("--------------------------")
	fmt.Println("         M E N U          ")
	fmt.Println("--------------------------")
	fmt.Println("1. Input Barang")
	fmt.Println("2. Edit Barang")
	fmt.Println("3. Hapus Barang")
	fmt.Println("4. Cari Barang")
	fmt.Println("5. Urutkan Barang")
	fmt.Println("6. Cetak Barang")
	fmt.Println("7. Exit")
	fmt.Println("-------------------------")
	fmt.Println("-------------------------")
}

func tambah(A *tabBarang, n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Masukkan barang ke-", i+1)

		fmt.Print("Masukkan id: ")
		fmt.Scan(&A[i].id)

		fmt.Print("Masukkan nama: ")
		fmt.Scan(&A[i].nama)

		fmt.Print("Masu1kkan stok: ")
		fmt.Scan(&A[i].stok)

		fmt.Print("Masukkan harg: ")
		fmt.Scan(&A[i].harga)
	}
}

func cetak(A tabBarang, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(A[i].id, A[i].nama, A[i].stok, A[i].harga)
	}
}

// Sequential search
func finddata(A tabBarang, n int, bld string) int {
	var i, idx int
	idx = -1
	i = 0
	for i < n && idx == -1 {
		if A[i].id == bld {
			idx = i
		}
		i++
	}
	return idx
}

func hapusD(A *tabBarang, n *int, bld string) {
	var idx, i int

	idx = finddata(*A, *n, bld)
	if idx != -1 {
		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
	} else {
		fmt.Println("Data tidak ada")
	}

}

func editD(A *tabBarang, n int, bld string, x string) {
	var idx int

	idx = finddata(*A, n, bld)
	if idx != -1 {
		A[idx].nama = x
		fmt.Printf("Data berhasil di edit")
	} else {
		fmt.Printf("Tidak ada data")
	}
}

func carikategori(A *tabBarang, n int) {
	var idx string

	fmt.Print("Masukan nama yg ingin dicari: ")
	fmt.Scan(&idx)

	ditemukan := false
	for i := 0; i < n; i++ {
		if A[i].nama == idx {
			tampilkategori(A[i])
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada barang.")
	}
}

func tampilkategori(x barang) {
	fmt.Printf("ID: %s | Nama: %s | Stok: %d | Harga: %d", x.id, x.nama, x.stok, x.harga)
}

func urutHarga(A *tabBarang, n int) {
	var temp barang
	var i, pass int

	for pass = 1; pass <= n-1; pass++ {
		temp = A[pass]
		i = pass

		for (i > 0 && A[i-1].harga < temp.harga) || (i > 0 && A[i-1].harga == temp.harga && A[i-1].harga > temp.harga) {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
	}
}

func urutStok(A *tabBarang, n int) {
	var i, idx, pass int
	var temp barang

	for pass = 1; pass <= n-1; pass++ {
		idx = pass - 1
		for i = pass; i <= n-1; i++ {
			if A[i].stok > A[idx].stok {
				idx = i
			}
		}
		temp = A[idx]
		A[idx] = A[pass-1]
		A[pass-1] = temp
	}
}
