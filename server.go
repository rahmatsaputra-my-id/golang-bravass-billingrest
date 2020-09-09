package main

import (
	"belajargolang/billingrest/handler"
	"belajargolang/billingrest/lib"
	"belajargolang/billingrest/model"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

var (
	dbUser    = "postgres"
	dbPass    = "postgres"
	dbDefault = "postgres"
	dbUse     = "billing"
	notsecure = false
	port      = 8088
)

func main() {
	// cmd := exec.Command("polymer", "build")
	// cmd.Dir = fmt.Sprintf("frontend")
	// err := cmd.Start()
	// if err != nil {
	// 	log.Printf("command Finished with error: %v", err.Error())
	// }
	// log.Printf("waiting for command to finish..")
	// err = cmd.Wait()
	// if err != nil {
	// 	log.Printf("Command finished with error : %v", err.Error())
	// }

	// db, err := initDatabase()
	// if err != nil {
	// 	return
	// }

	db, err := lib.Connect(dbUser, dbPass, dbDefault)
	if err != nil {
		_, err := initDatabase()
		if err != nil {
			return
		}
	}

	defer db.Close()
	handler.RegisDB(db)

	http.HandleFunc("/v1/billingrest/", handler.SS)
	polymer := http.FileServer(http.Dir("frontend/build/es6-bundled"))
	http.Handle("/", http.StripPrefix("/", polymer))

	log.Println("localhost:8088")
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func initDatabase() (*sql.DB, error) {
	dbInit, err := lib.Connect(dbUser, dbPass, dbDefault)
	if err != nil {
		fmt.Println("error 1 dbInitPass")
		return nil, err
	}
	if err = lib.CreateDB(dbInit, dbUse); err != nil {
		fmt.Println("error 2 dbUse")
		return nil, err
	}
	db, err := lib.Connect(dbUser, dbPass, dbUse)
	if err != nil {
		fmt.Println("error 3 dbPass")
		return nil, err
	}
	if err = lib.CreateTable(db, model.TbUser); err != nil {
		fmt.Println("error 4 TbUser")
		fmt.Println(err)
		return nil, err
	}
	// user := model.User{Nama: "Fahtul", Username: "Fahtul"}
	// if err := user.Insert(db); err != nil {
	// 	return nil, err
	// }
	if err = lib.CreateTable(db, model.TbPeriode); err != nil {
		fmt.Println("error 5 TbPeriode")
		return nil, err
	}
	if err = lib.CreateTable(db, model.TbInvoice); err != nil {
		fmt.Println("error 6 TbInvoice")
		return nil, err
	}
	if err = lib.CreateTable(db, model.TbMember); err != nil {
		fmt.Println("error 7 TbMember")
		return nil, err
	}
	if err = lib.CreateTable(db, model.TbRoom); err != nil {
		fmt.Println("error 8 TbRoom")
		return nil, err
	}
	return db, nil
}

// bang akbar

// package main

// import (
// 	"database/sql"
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	// "github.com/f6/webserver/webhandler"
// 	// "github.com/fabarj4/example/simpleapi/handler"
// 	// "github.com/fabarj4/example/simpleapi/model"
// 	// "github.com/lib/pq"
// )

// var (
// 	db        *sql.DB
// 	dbgolang  = "simple_rest"
// 	defaultdb = "postgres"
// 	dbuser    = "postgres"
// 	dbpass    = "postgres"

// 	notsecure = false
// 	port      = 8088
// )

// func init() {
// 	flag.BoolVar(&notsecure, "notsecure", false, "by default web server run on https, if notsecure true run in http")
// 	flag.IntVar(&port, "port", 8088, "port used in web server")
// 	flag.StringVar(&dbuser, "dbuser", "postgres", "User for db postgres")
// 	flag.StringVar(&dbpass, "dbpass", "postgres", "password for db postgres")
// }

// func main() {
// 	flag.Parse()
// 	var err error

// 	db, err = connectDB(dbgolang, dbuser, dbpass)
// 	if err != nil {
// 		if !isErrDBNotExist(err) {
// 			log.Fatalf("Gagal konek database %s", err)
// 		}
// 		db, err = prepareDB()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	webhandler.RegisterDB(db)
// 	webhandler.DebugOn()
// 	apiUrl := ""
// 	http.Handle(apiUrl, handler.webhandler(apiUrl))
// 	fmt.Printf("Port used: %d\n", port)
// 	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
// }
