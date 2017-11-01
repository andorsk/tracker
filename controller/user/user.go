package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	umodel "tracker/model/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type UserController struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *UserController) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.InitializeRoutesNoRouter()
}

func (a *UserController) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *UserController) createUser(w http.ResponseWriter, r *http.Request) {
	var u umodel.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := u.CreateUser(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, u)
}

func getUser(db *sql.DB, id int) (umodel.User, error) {
	statement := fmt.Sprintf("SELECT is FROM user where id = %d", id)
	row, err := db.Query(statement)

	if err != nil {
		return umodel.User{}, err
	}

	defer row.Close()
	user := umodel.User{}

	for row.Next() {
		var u umodel.User
		if err := row.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return umodel.User{}, err
		}
		user = u
	}

	return user, nil

}

func (a *UserController) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	u := umodel.User{ID: id}
	if err := u.GetUser(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, u)
}

func (a *UserController) getUsers(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))
	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}
	users, err := umodel.GetUsers(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *UserController) InitializeRoutes(r *mux.Router) {
	r.HandleFunc("/users", a.getUsers).Methods("GET")
	r.HandleFunc("/create", a.createUser).Methods("POST")
	r.HandleFunc("/user/{id:[0-9]+}", a.getUser).Methods("GET")
}

func (a *UserController) InitializeRoutesNoRouter() {
	fmt.Println("Initialized the routes")
	a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
	a.Router.HandleFunc("/create", a.createUser).Methods("POST")
	a.Router.HandleFunc("/user/{id:[0-9]+}", a.getUser).Methods("GET")
}
