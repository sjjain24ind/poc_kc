package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Nerzal/gocloak/v13"
)

var (
	keycloakClient *gocloak.GoCloak
	realm          = "master"
	clientID       = "your-client-id"     // Replace with your Keycloak client ID
	clientSecret   = "your-client-secret" // Replace with your Keycloak client secret
)

func main() {
	// Initialize Keycloak client
	keycloakClient = gocloak.NewClient("http://keycloak:8080")

	// Define routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/welcome", welcomeHandler)

	// Start the server
	log.Println("App is running on :9081...")
	if err := http.ListenAndServe(":9081", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// Home handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Render the login page
	tmpl := template.Must(template.ParseFiles("login.html"))
	tmpl.Execute(w, nil)
}

// Login handler
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	userType := r.FormValue("userType") // "admin" or "user"

	// Authenticate with Keycloak
	token, err := keycloakClient.GetToken(r.Context(), realm, gocloak.TokenOptions{
		ClientID:     &clientID,
		ClientSecret: &clientSecret,
		Username:     &username,
		Password:     &password,
		GrantType:    gocloak.StringP("password"), // Use StringP for pointer to string
	})
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Save the token in a cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token.AccessToken,
		Path:  "/",
	})

	// Redirect to the welcome page
	http.Redirect(w, r, "/welcome?userType="+userType, http.StatusFound)
}

// Welcome handler
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user type from the query parameter
	userType := r.URL.Query().Get("userType")

	// Get the token from the cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Validate the token
	_, err = keycloakClient.RetrospectToken(r.Context(), cookie.Value, clientID, clientSecret, realm)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Show the welcome message based on the user type
	if userType == "admin" {
		w.Write([]byte("Welcome Admin!"))
	} else {
		w.Write([]byte("Welcome User!"))
	}
}
