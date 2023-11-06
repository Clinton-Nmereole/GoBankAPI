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
	router.HandleFunc("/logout", makeHTTPHandleFunc(a.handleLogout))
	router.HandleFunc("/deposit", makeHTTPHandleFunc(a.handleDeposit))
	// router.HandleFunc("/me", makeHTTPHandleFunc(a.handleMe))
	log.Println("JSON API server started. Listening on port", a.listenAddress)
	err := http.ListenAndServe(a.listenAddress, router)
	fmt.Printf("%+v\n", err)
}

var (
	loginid     string
	loggedin    bool
	logintoken  string
	loginnumber int32
)

var loginResponse LoginResponse

func (a *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("Unsupported method: %s", r.Method)
	}
	loginRequest := new(LoginRequest) // LoginRequest
	if err := json.NewDecoder(r.Body).Decode(loginRequest); err != nil {
		return err
	}

	// Validate login request
	// fmt.Println(loginRequest.AccountNumber)
	account, err := a.store.GetAccountByNumber(loginRequest.AccountNumber)
	if err != nil {
		fmt.Println("Account not found")
		return err
	}

	if account.PasswordMatches(loginRequest.Password) != nil {
		return fmt.Errorf("Not Authenticated")
	}

	token, err := createJWT(account)
	if err != nil {
		return err
	}

	response := LoginResponse{
		Token:  token,
		Number: account.Number,
	}

	loginResponse = response

	// fmt.Printf("account: %+v\n", account)
	// fmt.Println("your account id is:", account.ID)

	// http.Redirect(w, r, "/accounts/"+strconv.Itoa(account.ID), http.StatusOK)
	w.Header().Add("x-jwt-token", token)
	w.Header().Add("user-id", strconv.Itoa(account.ID))
	// fmt.Println("getting user-id token: ", w.Header().Get("user-id"))
	loginid = strconv.Itoa(account.ID)
	loginnumber = response.Number
	loggedin = true
	logintoken = loginResponse.Token

	// fmt.Println("your header is:", r.Header)

	return WriteJSON(w, http.StatusOK, response)
}

func (a *APIServer) handleLogout(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("x-jwt-token", "")
	w.Header().Add("user-id", "")
	loginid = ""
	loggedin = false
	logintoken = ""
	return WriteJSON(w, http.StatusOK, map[string]string{"message": "Logged out"})
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
		if !loggedin {
			return fmt.Errorf("Not logged in")
		} else {
			fmt.Println("you are logged in")
		}
		// account := NewAccount("Clinton", "Sensei")
		// fmt.Println("you are logged in")
		// w.Header().Add("x-jwt-token", loginResponse.Token)
		// fmt.Println("this is your x-jwt-token: ", r.Header.Get("x-jwt-token"))

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
	if r.Method == "POST" {
		createAccountRequest := new(CreateAccountRequest)
		fmt.Println(json.NewDecoder(r.Body))
		if err := json.NewDecoder(r.Body).Decode(createAccountRequest); err != nil {
			return fmt.Errorf("Failed to decode request body: %w", err)
		}

		account, err := NewAccount(
			createAccountRequest.FirstName,
			createAccountRequest.LastName,
			createAccountRequest.Password,
		)
		if err := a.store.CreateAccount(account); err != nil {
			return fmt.Errorf("Failed to create account: %w", err)
		}

		tokenString, err := createJWT(account)
		if err != nil {
			return fmt.Errorf("Failed to create JWT token: %w", err)
		}
		fmt.Println("JWT token: ", tokenString)

		// return WriteJSON(w, http.StatusOK, account)
		http.Redirect(w, r, "http://localhost:5173/", http.StatusOK)
		r.Header.Add("x-jwt-token", tokenString)
		return WriteJSON(w, http.StatusOK, account)
	}
	return fmt.Errorf("Unsupported method: %s", r.Method)
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

func (a *APIServer) handleDeposit(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		depositRequest := new(DepositRequest)
		if err := json.NewDecoder(r.Body).Decode(depositRequest); err != nil {
			return fmt.Errorf("Failed to decode request body: %w", err)
		}
		// Perform the handleDeposit
		if err := a.store.Deposit(depositRequest.Amount, loginnumber); err != nil {
			return err
		}
		// Create a token for the handleDeposit
		// To do this we first have to get the fromAccount so we can get it by account number
		account, err := a.store.GetAccountByNumber(loginnumber)
		if err != nil {
			return err
		}
		token, err := createJWT(account)
		if err != nil {
			return err
		}
		response := DepositResponse{
			Balance: account.Balance,
			Token:   token,
		}
		return WriteJSON(w, http.StatusOK, response)
	}
	return fmt.Errorf("Unsupported method: %s", r.Method)
}

func (a *APIServer) handleTransactions(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		transactionRequest := new(Transaction)
		if err := json.NewDecoder(r.Body).Decode(transactionRequest); err != nil {
			return fmt.Errorf("Failed to decode request body: %w", err)
		}
		// Perform the handleTransactions
		if err := a.store.Transfer(transactionRequest.Amount, transactionRequest.ToAccount, loginnumber); err != nil {
			return err
		}
		// Create a token for the handleTransactions
		// To do this we first have to get the fromAccount so we can get it by account number
		account, err := a.store.GetAccountByNumber(loginnumber)
		if err != nil {
			return err
		}
		token, err := createJWT(account)
		if err != nil {
			return err
		}
		response := TransactionResponse{
			Transaction: *transactionRequest,
			Balance:     account.Balance,
			Token:       token,
		}
		// fromAccount, err := a.store.GetAccountByID(r.)
		defer r.Body.Close()
		return WriteJSON(w, http.StatusOK, response)
	}
	return fmt.Errorf("Unsupported method: %s", r.Method)
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
		fmt.Println("calling withJWTAuth")
		tokenString := logintoken
		// fmt.Println("logintoken: ", loginResponse.Token)
		// fmt.Println("getting user-id token from JWT Auth: ", w.Header().Get("user-id"))
		// fmt.Println("your x-jwt-token: ", tokenString)
		token, err := validateJWT(tokenString)
		if err != nil {
			permissionDenied(w)
			return
		}
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
