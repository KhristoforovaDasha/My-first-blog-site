package main

import (
	"context"
	"flag"
	"fmt"
	"hristoforovada-project/backend/internal/entity"
	handler "hristoforovada-project/backend/internal/handler/http"
	"hristoforovada-project/backend/internal/repository"
	repo_sqlite "hristoforovada-project/backend/internal/repository/sqlite"
	"hristoforovada-project/backend/internal/service"
	"hristoforovada-project/backend/pkg/auth"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func checkRepository(repo *repository.Repository) {
	//--User check
	repo.User.Create(&entity.User{UserRegister: entity.UserRegister{Login: "login1", Password: "hahaha"}})
	repo.User.Create(&entity.User{UserRegister: entity.UserRegister{Login: "login2", Password: "hahaha"}})
	repo.User.Create(&entity.User{UserRegister: entity.UserRegister{Login: "login3", Password: "hahaha"}})
	repo.User.Create(&entity.User{UserRegister: entity.UserRegister{Login: "login1", Password: "hahaha"}})

	users, _ := repo.User.GetAll()
	fmt.Printf("%+v\n", users)
	fmt.Println()

	user_to_update := (*users)[0]
	user_to_update.Login = "login_update"
	repo.User.Update(&user_to_update)
	users, _ = repo.User.GetAll()
	fmt.Printf("%+v\n", (*users)[0])

	user1, _ := repo.User.Get(1)
	fmt.Printf("%+v\n", user1)
	fmt.Println()
}

func checkRepositoryComment(repo *repository.Repository) {
	//--Comment check
	repo.Comment.Create(&entity.Comment{PostId: 1, UserId: 1, CommentText: "i am user 1"})
	repo.Comment.Create(&entity.Comment{PostId: 2, UserId: 2, CommentText: "i am user 2"})
	repo.Comment.Create(&entity.Comment{PostId: 3, UserId: 3, CommentText: "i am user 3"})

	comments, _ := repo.Comment.GetAll()
	fmt.Printf("%+v\n", comments)
	fmt.Println()

	comment_to_update := (*comments)[0]
	comment_to_update.CommentText = "text_update"
	repo.Comment.Update(&comment_to_update)
	comments, _ = repo.Comment.GetAll()
	fmt.Printf("%+v\n", (*comments)[0])

	comment2, _ := repo.Comment.Get(2)
	fmt.Printf("%+v\n", comment2)
	fmt.Println()
}

func checkRepositoryPost(repo *repository.Repository) {
	//--Post check
	repo.Post.Create(&entity.Post{PostText: "i am post 1"})
	repo.Post.Create(&entity.Post{PostText: "i am post 2"})
	repo.Post.Create(&entity.Post{PostText: "i am post 3"})

	posts, _ := repo.Post.GetAll()
	fmt.Printf("%+v\n", posts)
	fmt.Println()

	post_to_update := (*posts)[0]
	post_to_update.PostText = "text_update"
	repo.Post.Update(&post_to_update)
	posts, _ = repo.Post.GetAll()
	fmt.Printf("%+v\n", (*posts)[0])

	post2, _ := repo.Post.Get(2)
	fmt.Printf("%+v\n", post2)
	fmt.Println()
}

func checkService(service *service.Service) {
	//-------START------------------------
	fmt.Println("Try to register")
	userReg := entity.UserRegister{
		Login:    "abacaba",
		Password: "b",
	}
	err := service.User.Register(&userReg)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("oki")
	}

	fmt.Println("Try to register")
	userReg1 := entity.UserRegister{
		Login:    "beme",
		Password: "b",
	}
	err = service.User.Register(&userReg1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("oki")
	}
	//-----------END--------------------------

	//------------START--------------------------------
	fmt.Println("Try to login with incorrect password")
	userLog := entity.UserRegister{
		Login:    "abacaba",
		Password: "e",
	}
	_, err = service.User.Login(&userLog)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("ne oki")
	}

	fmt.Println("Try to login with correct password")
	userLog = entity.UserRegister{
		Login:    "abacaba",
		Password: "b",
	}
	_, err = service.User.Login(&userLog)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("oki")
	}
	//--------------END-----------------------------

	fmt.Println("Try to get user")
	user, err := service.User.Get(1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%+v\n", *user)
	}
}

func checkPostService(service *service.Service) {
	//--------START-------------------
	fmt.Println("Try to create a post")
	post := entity.Post{
		PostText: "Мой первый постец",
	}
	err := service.Post.Create(&post)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("normik")
	}

	fmt.Println("Try to create a post")
	post = entity.Post{
		PostText: "Мой second постец",
	}
	err = service.Post.Create(&post)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("normik")
	}
	//--------END--------------------

	//---START---------------------
	fmt.Println("Try to update post")
	post1, err := service.Post.Get(1)
	post1.PostText = "update text"
	err = service.Post.Update(post1)
	post1, err = service.Post.Get(1)
	fmt.Printf("%+v", post1)
	//--------END--------------------
}

func main() {
	//checkEntities()
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	logFile, err := os.OpenFile("log/backend.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	dbUri, ok := os.LookupEnv("DB_URI")
	if !ok {
		log.Println("cannot get DB_URI from ENV")
		dbUri = "test.db"
	}

	db, err := repo_sqlite.NewSQLiteDB(dbUri)
	if err != nil {
		log.Panicf("Failed to initialize database")
	} else {
		log.Println("database initialized")
	}

	repo := repository.NewRepository(db)
	myService := service.NewService(repo)
	//checkPostService(s_ervice)

	signingKey, ok := os.LookupEnv("AUTH_SIGNING_KEY")
	if !ok {
		log.Println("cannot get AUTH_SIGNING_KEY from ENV")
		signingKey = "siuefui4nfweu"
	}

	authManager := auth.NewAuthManager([]byte(signingKey))

	h := handler.NewHandler(myService, authManager)

	srv := &http.Server{
		Addr: ":3001",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.NewRouter(), // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	//repo.DropAll(db)
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

}
