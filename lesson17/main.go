package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	/* Задача 17.1
	Необходимо код примера 1 изменить так, чтобы tcp-сервер
	обрабатывал подключения параллельно
	*/
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
	log.Println("завершение работы")

	/* Задача 17.2
	Измените код примера 4 так, чтобы логировались абсолютно все
	запросы, даже которые не прошли авторизацию

	Задача 17.3
	Измените код примера 4 так, чтобы лог попадал не в stdout, а в файл log
	*/
	file, err := os.OpenFile("log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Не удалось открыть файл:", err)
	}
	defer file.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	l := log.New(file, "", log.LstdFlags)
	logHandler := logMiddleware(l)
	httpServer := &http.Server{Addr: ":8080", Handler: logHandler(authHandler(mux))}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln(fmt.Errorf("Неудалось запустить сервер: %w", err))
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	var i int
	for {
		_, err := io.WriteString(c, fmt.Sprintf("%d\n", i))
		if err != nil {
			log.Println(err)
			return
		}
		i++
		time.Sleep(time.Second)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	msg := "Hello, Go!"
	log.Println("resp:", msg)
	fmt.Fprint(res, msg)
}
func closeAll(res http.ResponseWriter, req *http.Request) {
	log.Fatalln()
}
func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xId := r.Header.Get("x-my-app-id")
		if xId != "my_secret" {
			http.Error(w, "пользователь не авторизован", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}
func logMiddleware(l *log.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Println("url:", r.URL)
			h.ServeHTTP(w, r)
		})
	}
}
