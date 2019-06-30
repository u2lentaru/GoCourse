package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

//1. В методичку из статьи добавлять ничего не надо, так как в методичке имеется информация по
//  главной особенности работы с временем в Golang - форматированию. Для остального достаточно
// имеющегося в методичке краткого изложения и ссылки на документацию.

func main() {
	fileRead()
	rwCSV()
}

//2. Добавлены выбор открываемого файла, печать ошибок, проверка IsDir().
func fileRead() {
	fr := ""
	fmt.Println("Enter file name for reading:")
	fmt.Scanln(&fr)

	file, err := os.Open(fr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//getting size of file
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	if stat.IsDir() {
		fmt.Println(stat.Name(), "is a directory.")
		return
	}

	//reading file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}

func rwCSV() {
	records := [][]string{
		{"name", "phone"},
		{"Bob", "89167243814"},
		{"Alex", "89155243629"},
		{"Mike", "89167245643"},
	}

	wfile, err := os.Create("somecsvfile.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wfile.Close()

	w := csv.NewWriter(wfile)
	w.WriteAll(records)

	if err := w.Error(); err != nil {
		fmt.Println("error writing csv:", err)
		return
	}
	rfile, err := os.Open("somecsvfile.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer rfile.Close()

	reader := csv.NewReader(rfile)
	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sanity check, display to standard output
	for _, each := range rawCSVdata {
		fmt.Printf("%s and %s\n", each[0], each[1])
	}

}
