package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

//1. В методичку из статьи добавлять ничего не надо, так как в методичке имеется информация по
//  главной особенности работы с временем в Golang - форматированию. Для остального достаточно
// имеющегося в методичке краткого изложения и ссылки на документацию.

func main() {
	fileRead()
	rwCSV()
}

//2. Выбор открываемого файла, IsDir(). Выводить количество прочитанных байт, проверить кириллический текст.
func fileRead() {
	fr := ""
	fmt.Scanln(&fr)

	file, err := os.Open(fr)
	if err != nil {
		return
	}
	defer file.Close()

	//getting size of file
	stat, err := file.Stat()
	if err != nil {
		return
	}

	if stat.IsDir() {
		fmt.Println(stat.Name(), "является директорией.")
		return
	}

	//reading file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	fmt.Println(string(bs))
}

func rwCSV() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	wfile, err := os.Create("somecsvfile.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wfile.Close()

	w := csv.NewWriter(wfile)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)

		rfile, err := os.Open("somecsvfile.csv")

		if err != nil {
			fmt.Println(err)
			return
		}

		defer rfile.Close()

		reader := csv.NewReader(rfile)
		reader.FieldsPerRecord = -1 // see the Reader struct information below
		rawCSVdata, err := reader.ReadAll()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// sanity check, display to standard output
		for _, each := range rawCSVdata {
			fmt.Printf("email : %s and timestamp : %s\n", each[0], each[1])
		}
	}
}
