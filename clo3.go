package main

import "fmt"

const NMAX int = 10

type snack struct {
	id, nama    string
	harga, stok int
}

type TabSnack [NMAX]snack

func main() {
	var data TabSnack
	var nData int
	var id string
	var harga int

	bacaData(&data, nData)

	cetakData(data, nData)

	fmt.Scan(&id)
	fmt.Print("Nama snack: ", namaSnackSeq(data, nData, id))

	fmt.Println("Omset", omset(data, nData))

	urutIDMenaik(&data, nData)
	fmt.Print("Data id menaik")
	cetakData(data, nData)

	urutHargaMenurun(&data, nData)
	fmt.Print("Data Harga menurun")
	cetakData(data, nData)

	fmt.Scan(&harga)
	fmt.Print("Nama Snack:", namaSnackBin(data, nData, harga))

}

func bacaData(A *TabSnack, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i].id, &A[i].nama, &A[i].harga, &A[i].stok)
	}
}

func cetakData(A TabSnack, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(A[i].id, A[i].nama, A[i].harga, A[i].stok)
	}
}

func namaSnackSeq(A TabSnack, n int, x string) string {
	for i := 0; i < n; i++ {
		if A[i].id == x {
			return A[i].nama
		}
	}
	return "none"
}

func omset(A TabSnack, n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += A[i].stok * A[i].harga
	}
	return total
}

func urutIDMenaik(A *TabSnack, n int) {
	var i, idx, pass int
	var temp snack

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass

		for i < n {
			if A[i].id > temp.id {
				idx = i
			}
			i = i + 1

			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass = pass + 1
		}
	}
}

func urutHargaMenurun(A *TabSnack, n int) {
	var temp snack
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

func namaSnackBin(A TabSnack, n int, bId int) string {
	var left, right, mid, idx int
	idx = -1
	left = 0
	right = n - 1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if A[mid].harga == bId {
			idx = mid
		} else if A[mid].harga < bId {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return "None"
}
