package server

import (
	"crud/db"
	"crud/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	idToInt32Error        string = "Erro ao converter id para int32"
	bodyReadError         string = "Erro ao ler o corpo da requisção"
	jsonToStructError     string = "Erro ao converter json para struct"
	structToJsonError     string = "Erro ao converter struct para json"
	dbConnectError        string = "Erro ao conectar com o banco de dados"
	statementPrepareError string = "Erro ao preparar o statement"
	statementExecError    string = "Erro ao executar o statement"
	fetchError            string = "Erro ao buscar usuários"
	scanError             string = "Erro ao escanear usuários"
	idError               string = "Erro ao obter o ID"
)

// CriarUsuario cria um usuário no banco de dados
func CriaUsuario(w http.ResponseWriter, r *http.Request) {
	corpoReq, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(bodyReadError))
		return
	}

	var user dto.User

	if err = json.Unmarshal(corpoReq, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(jsonToStructError))
		return
	}

	db, err := db.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(dbConnectError))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into users (name, email) values (?, ?)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(statementPrepareError))
		return
	}
	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(statementExecError))
		return
	}

	idInserido, err := insert.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(idError))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso. ID: %d", idInserido)))
}

// BuscaUsuarios lista todos os usuários do banco de dados
func BuscaUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := db.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(dbConnectError))
		return
	}
	defer db.Close()

	linhas, err := db.Query("select * from users")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fetchError))
		return
	}
	defer linhas.Close()

	var users []dto.User

	for linhas.Next() {
		var user = dto.User{}
		if err := linhas.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(scanError))
			return
		}
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(structToJsonError))
	}
}

// BuscaUsuarioPorId busca um usuário do banco de dados pelo ID
func BuscaUsuarioPorId(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(idToInt32Error))
		return
	}

	db, err := db.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(dbConnectError))
		return
	}
	defer db.Close()

	linha, err := db.Query("select * from users WHERE id = ?", ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(statementPrepareError))
		return
	}

	var user dto.User

	if linha.Next() {
		if err := linha.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte(scanError))
			return
		}
	}

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte(structToJsonError))
	}
}

// AtualizaUsuario atualiza um usuário do banco de dados
func AtualizaUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(idToInt32Error))
		return
	}

	corpoReq, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(bodyReadError))
		return
	}

	var newUserData dto.User

	if err = json.Unmarshal(corpoReq, &newUserData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(jsonToStructError))
		return
	}

	db, err := db.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(dbConnectError))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update users set name = ?, email = ? where id = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(statementPrepareError))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(newUserData.Name, newUserData.Email, ID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(statementExecError))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletaUsuario deleta um usuário do banco de dados
func DeletaUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(idToInt32Error))
		return
	}

	db, err := db.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(dbConnectError))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from users where id = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(statementPrepareError))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(statementExecError))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
