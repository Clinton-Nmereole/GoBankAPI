package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string
	store         Storage
}

func NewAPIServer(listenAddress string, store Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (a *APIServer) Start() {
	router := mux.NewRouter()
	router.HandleFunc("/login", makeHTTPHandleFunc(a.handleLogin))
	router.HandleFunc("/accounts", makeHTTPHandleFunc(a.handleAccounts))
	router.HandleFunc(
		"/accounts/{id}",
		withJWTAuth(makeHTTPHandleFunc(a.handleGetAccountByID),
			a.store,
		))
	router.HandleFunc("/accounts", makeHTTPHandleFunc(a.handleCreateAccount))
	router.HandleFunc("/transactions", makeHTTPHandleFunc(a.handleTransactions))
	log.Println("JSON API server started. Listening on port", a.listenAddress)
	http.ListenAndServe(a.listenAddress, router)
}

func (a *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("Unsupported method: %s", r.Method)
	}
	loginRequest := new(LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(loginRequest); err != nil {
		return err
	}

	// Validate login request

	return WriteJSON(w, http.StatusOK, loginRequest)
}

func (a *APIServer) handleAccounts(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return a.handleGetAccounts(w, r)
	}
	if r.Method == "POST" {
		return a.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return a.handleDeleteAccountByID(w, r)
	}
	return fmt.Errorf("Unsupported method: %s", r.Method)
}

func (a *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {

		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return fmt.Errorf("Invalid ID: %s", idStr)
		}
		// account := NewAccount("Clinton", "Sensei")

		// database will go here (something like db get id)
		account, err := a.store.GetAccountByID(id)
		if err != nil {
			return err
		}
		return WriteJSON(w, http.StatusOK, account)
	}
	if r.Method == "DELETE" {
		return a.handleDeleteAccountByID(w, r)
	}
	return fmt.Errorf("Unsupported method: %s", r.Method)
}

func (a *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := a.store.GetAllAccounts()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, accounts)
}

func (a *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountRequest := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(createAccountRequest); err != nil {
		return err
	}

	account, err := NewAccount(
		createAccountRequest.FirstName,
		createAccountRequest.LastName,
		createAccountRequest.Password,
	)
	if err := a.store.CreateAccount(account); err != nil {
		return err
	}

	tokenString, err := createJWT(account)
	if err != nil {
		return err
	}
	fmt.Println("JWT token: ", tokenString)

	return WriteJSON(w, http.StatusOK, account)
}

func (a *APIServer) handleDeleteAccountByID(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("Invalid ID: %s", idStr)
	}
	if err := a.store.DeleteAccount(id); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, map[string]int{"Account deleted": id})
}

func (a *APIServer) handleTransactions(w http.ResponseWriter, r *http.Request) error {
	transactionRequest := new(Transaction)
	if err := json.NewDecoder(r.Body).Decode(transactionRequest); err != nil {
		return err
	}
	defer r.Body.Close()
	return WriteJSON(w, http.StatusOK, transactionRequest)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func createJWT(account *Account) (string, error) {
	claims := &jwt.MapClaims{
		"accountNumber": account.Number,
		"exp":           time.Now().Add(time.Hour * 24).Unix(),
	}
	secret := os.Getenv("JWT_SECRET")
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}

func permissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusForbidden, APIError{Error: "Permission denied"})
}

// Middleware
func withJWTAuth(handlerFunc http.HandlerFunc, s Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("withJWTAuth")
		tokenString := r.Header.Get("x-jwt-token")
		token, err := validateJWT(tokenString)
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
		if err != nil {
			permissionDenied(w)
			return
		}
		if !token.Valid {
			permissionDenied(w)
			return
		}
		userID, err := getID(r)
		if err != nil {
			permissionDenied(w)
			return
		}
		account, err := s.GetAccountByID(userID)

		// panic(reflect.TypeOf(claims["accountNumber"]))
		if account.Number != int32(claims["accountNumber"].(float64)) {
			permissionDenied(w)
			return
		}
		if err != nil {
			WriteJSON(w, http.StatusForbidden, APIError{Error: "account not found"})
			return
		}

		handlerFunc(w, r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}

type APIHandler func(w http.ResponseWriter, r *http.Request) error

type APIError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f APIHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("Invalid ID: %s", idStr)
	}
	return id, nil
}
