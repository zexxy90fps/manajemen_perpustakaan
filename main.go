package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Buku struct {
	judul     string
	pengarang string
	penerbit  string
	tahun     int
	halaman   int
	kode      string
}

var listBuku []Buku

func TambahBuku() {
	judulBuku := bufio.NewReader(os.Stdin)
	pengarangBuku := bufio.NewReader(os.Stdin)
	penerbitBuku := bufio.NewReader(os.Stdin)
	tahunBuku := 0
	halamanBuku := 0
	kodeBuku := bufio.NewReader(os.Stdin)

	fmt.Println("\n<=======================>")
	fmt.Println("Tambah Buku")
	fmt.Println("<=======================>")

	// masukan kode buku
	fmt.Print("masukan kode buku: ")
	kodeBukuBaru, err := kodeBuku.ReadString('\n')
	if err != nil {
		fmt.Println("terjadi error: ", err)
		return
	}
	kodeBukuBaru = strings.TrimSpace(kodeBukuBaru)

	// input judul buku
	fmt.Print("Masukkan Judul Buku: ")
	judulBukuBaru, err := judulBuku.ReadString('\n')
	if err != nil {
		fmt.Println("terjadi error: ", err)
		return
	}
	judulBukuBaru = strings.TrimSpace(judulBukuBaru)

	// input pengarang
	fmt.Print("Masukkan Pengarang Buku: ")
	pengarangBukuBaru, err := pengarangBuku.ReadString('\n')
	if err != nil {
		fmt.Println("terjadi error: ", err)
		return
	}
	pengarangBukuBaru = strings.TrimSpace(pengarangBukuBaru)

	// penerbit buku baru
	fmt.Print("Masukan Penerbit Buku: ")
	penerbitBukuBaru, err := penerbitBuku.ReadString('\n')
	if err != nil {
		fmt.Println("terjadi error: ", err)
		return
	}
	penerbitBukuBaru = strings.TrimSpace(penerbitBukuBaru)

	// input tahun buku
	fmt.Print("masukkan Tahun Buku: ")
	_, err = fmt.Scanln(&tahunBuku)
	if err != nil {
		fmt.Println("terjadi error: ", err)
		return
	}

	// halaman buku
	fmt.Print("masukan halaman buku: ")
	_, err = fmt.Scanln(&halamanBuku)
	if err != nil {
		fmt.Println("terjadi error: ", err)
		return
	}

	listBuku = append(listBuku, Buku{
		judul:     judulBukuBaru,
		pengarang: pengarangBukuBaru,
		penerbit:  penerbitBukuBaru,
		tahun:     tahunBuku,
		halaman:   halamanBuku,
		kode:      kodeBukuBaru,
	})

	fmt.Println("berhasil menambahkan buku!")
}

func editBuku() {
	var kode string
	fmt.Println("\n<=================>")
	fmt.Println("edit data buku")
	fmt.Println("<=================>")
	lihatbuku()
	fmt.Print("\nmasukan kode buku yang ingin di edit: ")
	fmt.Scanln(&kode)
	kode = strings.TrimSpace(kode)

	for i, buku := range listBuku {
		if strings.TrimSpace(buku.kode) == kode {
			inputanUser := bufio.NewReader(os.Stdin)

			fmt.Println("masukan data baru: ")

			fmt.Print("judul buku: ")
			listBuku[i].judul, _ = inputanUser.ReadString('\n')
			listBuku[i].judul = strings.TrimSpace(listBuku[i].judul)

			fmt.Print("pengarang buku: ")
			listBuku[i].pengarang, _ = inputanUser.ReadString('\n')
			listBuku[i].pengarang = strings.TrimSpace(listBuku[i].pengarang)

			fmt.Print("penerbit buku: ")
			listBuku[i].penerbit, _ = inputanUser.ReadString('\n')
			listBuku[i].penerbit = strings.TrimSpace(listBuku[i].penerbit)

			fmt.Print("jumlah halaman buku: ")
			fmt.Scanln(&listBuku[i].halaman)

			fmt.Print("tahun terbit buku: ")
			fmt.Scanln(&listBuku[i].tahun)

			fmt.Println("data berhasil diperbarui")
			fmt.Println("tekan 'enter' untuk melanjutkan.......")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			return
		}
	}
	fmt.Println("terjadi error pada pemilihan kode buku atau kode buku tidak ada")
	fmt.Println("tekan 'enter' untuk melanjutkan.......")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func lihatbuku() {
	fmt.Println("\n\n lihat buku")
	for urutan, buku := range listBuku {
		fmt.Printf("%d. kode buku: %s, judul buku: %s, pengarang: %s, penerbit: %s, tahun: %d, halaman: %d\n",
			urutan+1,
			buku.kode,
			buku.judul,
			buku.pengarang,
			buku.penerbit,
			buku.tahun,
			buku.halaman,
		)
	}
}

func HapusBuku() {
	var kodeInput string

	fmt.Println("<=======================>")
	fmt.Println("Hapus Buku")
	fmt.Println("<=======================>")
	lihatbuku()

	fmt.Println("\n<=======================>")
	fmt.Print("Masukkan kode buku yang ingin dihapus: ")
	fmt.Scanln(&kodeInput)
	kodeInput = strings.TrimSpace(kodeInput)

	for i, v := range listBuku {
		if strings.TrimSpace(v.kode) == kodeInput {
			listBuku = append(
				listBuku[:i],
				listBuku[i+1:]...,
			)
			fmt.Println("Buku telah berhasil dihapus")
			fmt.Println("tekan 'enter' untuk melanjutkan.......")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			return
		}
	}
	fmt.Println("terjadi error pada pemilihan kode buku atau kode buku tidak ada")
	fmt.Println("tekan 'enter' untuk melanjutkan.......")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	for {
		var pilihBuku int

		fmt.Println("<=====================>")
		fmt.Println("Pilih opsi anda")
		fmt.Println("<=====================>")
		fmt.Println("Silahkan Pilih :")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Lihat Buku")
		fmt.Println("3. Hapus Buku")
		fmt.Println("4. Edit Data Buku")
		fmt.Println("5. Selesai")
		fmt.Print("masukkan opsi anda: ")
		fmt.Scanln(&pilihBuku)

		switch pilihBuku {
		case 1:
			TambahBuku()
		case 2:
			lihatbuku()
		case 3:
			HapusBuku()
		case 4:
			editBuku()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("pilihan anda tidak tersedia")
		}
	}
}
