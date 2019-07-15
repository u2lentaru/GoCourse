package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	serial()
	parallel()
}

func serial() {
	fmt.Println("Serial")

	start := time.Now()
	for _, url := range os.Args[1:] {
		serialFetch(url)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func serialFetch(url string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	//bytes,err:=ioutil.ReadAll(resp.Body)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}

func parallel() {
	fmt.Println("Parallel")
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go parallelFetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Print(<-ch)
	}

	fmt.Printf("%.2fs elapse\n", time.Since(start).Seconds())
}

func parallelFetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch <- fmt.Sprintf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}

//1. Исследуйте работу последовательного и параллельного сканера веб-сайтов, задав большое
//(не менее 10) количество URL. Какие выводы можно сделать?
//
//PS C:\Golang_work\src\GoCourse\homework-8> go run .\homework-8.go http://ya.ru http://yandex.ru http://r0.ru http://rambler.ru
//http://mail.ru http://geekbrains.ru http://google.com http://vk.com http://ok.ru http://facebook.com
//Serial
//0.49s   13615 http://ya.ru
//0.66s  161501 http://yandex.ru
//0.35s   49101 http://r0.ru
//0.41s  407108 http://rambler.ru
//0.32s  288490 http://mail.ru
//0.92s   47931 http://geekbrains.ru
//1.86s   14320 http://google.com
//0.55s    9314 http://vk.com
//0.33s  181963 http://ok.ru
//1.23s  128508 http://facebook.com
//7.13s elapsed
//Parallel
//0.10s   47931 http://geekbrains.ru
//0.13s   14319 http://google.com
//0.19s  181955 http://ok.ru
//0.19s   13619 http://ya.ru
//0.19s    9315 http://vk.com
//0.20s  288415 http://mail.ru
//0.20s   48842 http://r0.ru
//0.25s  407108 http://rambler.ru
//0.30s  141474 http://yandex.ru
//0.61s  128104 http://facebook.com
//0.61s elapse

//Выводы:
//
//Serial 7.13s elapsed
//Parallel 0.61s elapse
//Параллельный сканер значительно превосходит последовательный в скорости
//для одинакового количества сканируемых сайтов.

//2. Какие практические варианты применения сканера веб-сайтов вы можете предложить?
//
//1. Индексация сайтов в поисковых системах.
//2. Поиск уязвимостей в целях проверки безопасности.
//3. Тестирование работоспособности и производительности сайтов.

//3. Сформулируйте предложения по улучшению сетевого чата.
//
//1. Убрать дублирование сообщения для отправителя.
//2. Добавить возможность передачи файлов.
//3. Интегрированная VoIP-связь
//4. Видеоконференцсвязь.
//5. Интерактивная доска.
//6. Совместный доступ к экрану или отдельным приложениям.
//7. Запись хода веб-конференции

//4. * Реализуйте одно из своих предложений по улучшению сетевого чата
//
//Пытался убрать дублирование сообщения отправителя - пока не работает.
//Изменения в chatserver.go закомментированы.
