package main

func main() {
	mirroredQuery()
}

func mirroredQuery() string {
	responses := make(chan string, 3)

	go func() {
		responses <- request("asia.site.io")
	}()

	go func() {
		responses <- request("europe.site.io")
	}()

	go func() {
		responses <- request("america.site.io")
	}()

}

func request(hostname string) string {

}
