package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=123456 -p 3306:3306 -d mysql:latest

func main() {
	stringConexao := "root:123456@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil { // só pega erros inesperados que impeçam de fazer a conexão, não pega erros de autenticação / banco inexistente
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil { // aqui pega erros de autenticação / banco inexistente
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta")

	linhas, erro := db.Query("select * from usuarios")

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	defer linhas.Close()

	fmt.Println(linhas)

}
