package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func getCakeHandler(w http.ResponseWriter, r *http.Request, u User) {
	if u.IsBan {
		handleAccessError(errors.New("you were banned"), w)
		return
	}
	w.Write([]byte(u.FavoriteCake))
}

func getHistoryHandler(w http.ResponseWriter, r *http.Request, u User) {
	var history string
	history += "User " + u.Email + " history\n"
	for i := 0; i < len(u.BanHistory); i++ {
		log := u.BanHistory[i]
		if log.IsBan {
			history += "BANNED\n"
		} else {
			history += "UNBANNED\n"
		}
		history += "	by " + log.Who + "\n"
		history += "	at " + log.When.Format("2006-01-02 15:04:05") + "\n"
		if log.IsBan {
			history += "	reason: " + log.Why + "\n"
		}
	}
	w.Write([]byte(history))
}

func main() {
	r := mux.NewRouter()
	users := NewInMemoryUserStorage()
	userService := UserService{repository: users}

	jwtService, err := NewJWTService("pubkey.rsa", "privkey.rsa")
	if err != nil {
		panic(err)
	}
	initEnv()

	r.HandleFunc("/cake", logRequest(jwtService.jwtAuth(users, getCakeHandler))).Methods(http.MethodGet)
	r.HandleFunc("/user/register", logRequest(userService.Register)).Methods(http.MethodPost)
	r.HandleFunc("/user/jwt", logRequest(wrapJwt(jwtService, userService.JWT))).Methods(http.MethodPost)

	r.HandleFunc("/user/me", logRequest(userService.ShowMyCake)).Methods(http.MethodGet)
	r.HandleFunc("/user/favorite_cake", logRequest(userService.ChangeCake)).Methods(http.MethodPost)
	r.HandleFunc("/user/email", logRequest(userService.ChangeEmail)).Methods(http.MethodPost)
	r.HandleFunc("/user/password", logRequest(userService.ChangePassword)).Methods(http.MethodPost)

	r.HandleFunc("/admin/promote", logRequest(userService.AdminPromote)).Methods(http.MethodGet)
	r.HandleFunc("/admin/fire", logRequest(userService.AdminFire)).Methods(http.MethodGet)
	r.HandleFunc("/admin/ban", logRequest(userService.UserBan)).Methods(http.MethodGet)
	r.HandleFunc("/admin/unban", logRequest(userService.UserUnban)).Methods(http.MethodGet)

	srv := http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {

		<-interrupt
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}()

	log.Println("Server started, hit Ctrl+C to stop")
	err = srv.ListenAndServe()
	if err != nil {
		log.Println("Server exited with error:", err)
	}
	log.Println("Good bye :)")
}

func wrapJwt(jwt *JWTService, f func(http.ResponseWriter, *http.Request, *JWTService)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		f(rw, r, jwt)
	}
}

func initEnv() {
	os.Setenv("CAKE_ADMIN_EMAIL", "admin@mail.com")
	os.Setenv("CAKE_ADMIN_PASSWORD", "admin1111")
}
