package controllers

import "net/http"

// CreateUser create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
}

// FetchUsers fetch all users
func FetchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Fetching Users"))
}

// FetchUserById fetch user by id
func FetchUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Fetching by id"))
}

// UpdateUser update user by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

// DeleteUser delete user by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
