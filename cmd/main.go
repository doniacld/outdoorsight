package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/doniacld/outdoorsight/routers"
)

func main() {
	fmt.Println("  ____       __     __              _____      __   __ \n / __ \\__ __/ /____/ /__  ___  ____/ __(_)__ _/ /  / /_\n/ /_/ / // / __/ _  / _ \\/ _ \\/ __/\\ \\/ / _ `/ _ \\/ __/\n\\____/\\_,_/\\__/\\_,_/\\___/\\___/_/ /___/_/\\_, /_//_/\\__/ \n                                       /___/           ")
	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
