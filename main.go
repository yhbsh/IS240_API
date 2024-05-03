package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {
	http.HandleFunc("/signup", signUpHandler)
	http.HandleFunc("/signin", signInHandler)
	http.HandleFunc("/points", pointsHandler)

	fmt.Println("Server Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

var (
	users = []User{}
	lock  = sync.Mutex{}
)

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var req SignUpRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Error parsing JSON body", http.StatusBadRequest)
		return
	}

	newUser := User{
		ID:     req.ID,
		Email:  req.Email,
		Points: 500, // Default points
	}

	// Add user to the global slice
	lock.Lock()
	users = append(users, newUser)
	lock.Unlock()

	res := SignUpResponse{
		Message: fmt.Sprintf("Signed up user: %s, Points: 500", newUser.ID),
	}

	responseData, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var req SignInRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Error parsing JSON body", http.StatusBadRequest)
		return
	}

	// Check if user exists
	lock.Lock()
	defer lock.Unlock()
	for _, user := range users {
		if user.ID == req.ID {
			res := SignInResponse{
				Message: fmt.Sprintf("Signed in user: %s", user.ID),
				Email:   user.Email,
				Points:  user.Points,
			}
			responseData, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "Error creating response", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseData)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

func pointsHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "User ID required", http.StatusBadRequest)
		return
	}

	lock.Lock()
	defer lock.Unlock()

	switch r.Method {
	case "GET":
		for _, user := range users {
			if user.ID == userID {
				res := PointsResponse{
					Points: user.Points,
				}
				responseData, err := json.Marshal(res)
				if err != nil {
					http.Error(w, "Error creating response", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(responseData)
				return
			}
		}
	case "POST":
		var update struct {
			Points int `json:"points"`
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		if err := json.Unmarshal(body, &update); err != nil {
			http.Error(w, "Error parsing JSON body", http.StatusBadRequest)
			return
		}

		for i, user := range users {
			if user.ID == userID {
				users[i].Points = update.Points
				fmt.Fprintf(w, "{\"message\":\"Updated points to %d\"}", update.Points)
				return
			}
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	http.Error(w, "User not found", http.StatusNotFound)
}
