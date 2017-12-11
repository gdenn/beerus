package manager

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lib/pq"

	"golang.org/x/crypto/bcrypt"
)

// IUserTable interfaces functions handling db
// operations for user domain object
type IUserTable interface {
	AllUsers() (*[]User, error)
	UserByID(uuid string) error
	DeleteUserByID(uuid int64) error
	UserByName(name string) (*User, error)
	Save(user *User) error
}

// UserTable handles db operations for user domain
// object
type UserTable struct {
	db *sql.DB
}

// User stands for a manager user
type User struct {
	ID           int64
	Name         string
	Password     []byte
	BrokerAccess []int64
	AgentAccess  []int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Save user to database and replace Password
// by hash
func (u *UserTable) Save(user *User) error {
	s := "INSERT INTO users(name, password, brokerAccess, agentAcccess, createdAt, updatedAt) VALUES($1, $2, $3, $4, $5, $6)"
	stmt, err := u.db.Prepare(s)
	if err != nil {
		panic(err)
	}
	stmt.Exec(user.Name, user.Password, pq.Array(user.BrokerAccess), pq.Array(user.AgentAccess), time.Now(), time.Now())
	return nil
}

// DeleteUserByID deletes user given by id from
// db
func (u *UserTable) DeleteUserByID(id int64) error {
	stmt, err := u.db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		panic(err)
	}
	stmt.Exec(id)
	return nil
}

// UserByID returns user with id from db
func (u *UserTable) UserByID(id int64) (*User, error) {
	var ur User
	err := u.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(
		&ur.ID,
		&ur.Name,
		&ur.Password,
		&ur.BrokerAccess,
		&ur.AgentAccess,
		&ur.CreatedAt,
		&ur.UpdatedAt,
	)
	if err != nil {
		panic(err)
	}

	return &ur, nil
}

// UserByName returns user with name from db
func (u *UserTable) UserByName(name string) (*User, error) {
	var ur User
	err := u.db.QueryRow("SELECT * FROM users WHERE Name = $1", name).Scan(
		&ur.ID,
		&ur.Name,
		&ur.Password,
		&ur.BrokerAccess,
		&ur.AgentAccess,
		&ur.CreatedAt,
		&ur.UpdatedAt,
	)
	if err != nil {
		panic(err)
	}

	return &ur, nil
}

// AllUsers returns all users from db
func (u *UserTable) AllUsers() (*[]User, error) {
	var usrs []User
	rows, err := u.db.Query("SELECT * FROM users")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var ur User

		err := rows.Scan(
			&ur.ID,
			&ur.Name,
			&ur.Password,
			&ur.BrokerAccess,
			&ur.AgentAccess,
			&ur.CreatedAt,
			&ur.UpdatedAt,
		)

		if err != nil {
			panic(err)
		}

		usrs = append(usrs, ur)
	}

	return &usrs, nil
}

// UserRequest holding user information for new
// user in db
type UserRequest struct {
	Name         string
	Password     string
	BrokerAccess []int64
	AgentAccess  []int64
}

// DeleteUserRequest holding ID of user that shall be
// deleted from db
type DeleteUserRequest struct {
	ID int64
}

// GetUserByIDRequest holding ID of user that shall be
// returned from db
type GetUserByIDRequest struct {
	ID int64
}

// GetUserByNameRequest holding name of user that shall
// be returned from db
type GetUserByNameRequest struct {
	Name string
}

// CreateUser creates user in db
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	c := r.Context()
	ut, ok := c.Value(contextKey("user-table")).(UserTable)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic("failed to receive UserTable from context")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var ur UserRequest
	err := decoder.Decode(&ur)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panic("failed to parse request body")
		return
	}
	defer r.Body.Close()

	var u User
	u.Name = ur.Name
	u.AgentAccess = ur.AgentAccess
	u.BrokerAccess = ur.BrokerAccess
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	// hash password for db
	password := []byte(os.Getenv("BCRYPT_PASS"))
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic("failed to generate hashedPassword")
		return
	}
	u.Password = hashedPassword

	err = ut.Save(&u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic("failed to create user in database")
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteUserByID removes user by ID from db
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	c := r.Context()
	ut, ok := c.Value(contextKey("user-table")).(UserTable)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic("failed to receive UserTable from context")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var du DeleteUserRequest
	err := decoder.Decode(&du)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panic("failed to parse request body")
	}

	err = ut.DeleteUserByID(du.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panic("failed to delete user with id")
	}

	w.WriteHeader(http.StatusOK)
}

// GetUser returns user by ID from db
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	c := r.Context()
	ut, ok := c.Value(contextKey("user-table")).(UserTable)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic("failed to receive UserTable from context")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var gu GetUserByIDRequest
	err := decoder.Decode(&gu)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panic("failed to parse request body")
	}

	u, err := ut.UserByID(gu.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panic("failed to find user with ID")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

// GetUserByName returns user from db with name
func GetUserByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	c := r.Context()
	ut, ok := c.Value(contextKey("user-table")).(UserTable)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic("failed to receive UserTable from context")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var gn GetUserByNameRequest
	err := decoder.Decode(&gn)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panic("failed to parse request body")
	}

	u, err := ut.UserByName(gn.Name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panic("failed to find user with ID")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

// ListUser returns all users from db
func ListUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	c := r.Context()
	ut, ok := c.Value(contextKey("user-table")).(UserTable)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic("failed to receive UserTable from context")
		return
	}

	usrs, err := ut.AllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic("failed to load users from database")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usrs)
}