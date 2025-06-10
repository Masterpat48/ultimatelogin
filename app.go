package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type App struct {
	ctx    context.Context
	dbPath string
}

type User struct {
	Name       string `json:"name"`
	Pass       string `json:"pass"`
	Permission string `json:"permission"`
}

type LoginResult struct {
	Success    bool   `json:"success"`
	Name       string `json:"name"`
	Permission string `json:"permission"`
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// Database init
	if err := a.initDatabase(); err != nil {
		fmt.Println("Error in database init:", err.Error())
	}
}

func (a *App) Login(username string, password string) LoginResult {
	data, err := os.ReadFile(a.dbPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return LoginResult{Success: false}
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return LoginResult{Success: false}
	}

	for _, user := range users {
		if strings.EqualFold(user.Name, username) && user.Pass == password {
			return LoginResult{
				Success:    true,
				Name:       user.Name,
				Permission: user.Permission,
			}
		}
	}

	return LoginResult{Success: false}
}

// Optional: delete a user (admin only)
func (a *App) DeleteUser(username string) bool {
	data, err := os.ReadFile(a.dbPath)
	if err != nil {
		return false
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		return false
	}

	newUsers := []User{}
	found := false
	for _, user := range users {
		if strings.EqualFold(user.Name, username) {
			found = true
			continue
		}
		newUsers = append(newUsers, user)
	}

	if !found {
		return false
	}

	newData, err := json.MarshalIndent(newUsers, "", "  ")
	if err != nil {
		return false
	}

	err = os.WriteFile(a.dbPath, newData, 0644)
	return err == nil
}

func (a *App) RegisterUser(username string, password string) bool {
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		return false
	}

	data, err := os.ReadFile(a.dbPath)
	if err != nil {
		return false
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		return false
	}

	// Controlla se l'utente esiste già
	for _, user := range users {
		if strings.EqualFold(user.Name, username) {
			return false // Username already exist
		}
	}

	// Add a new user
	newUser := User{
		Name:       username,
		Pass:       password,
		Permission: "user",
	}
	users = append(users, newUser)

	// Saving the new user
	newData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return false
	}
	err = os.WriteFile(a.dbPath, newData, 0644)
	return err == nil
}
