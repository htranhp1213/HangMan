package main

// Necessary imports for our code
import (

	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"strings"
	"time"

)

// Creating a map variable with a string key and string slice value to store the category and words.
var options = map[string][]string{

	// For instance, the first key would be fruits and the first values would be the slice of strings
	"fruits":    {"Apple", "Blueberry", "Mandarin", "Pineapple", "Pomegranate", "Watermelon"},
	"animals":   {"Hedgehog", "Rhinoceros", "Squirrel", "Panther", "Walrus", "Zebra"},
	"countries": {"India", "Hungary", "Kyrgyzstan", "Switzerland", "Zimbabwe", "Dominica"},
	
}

// A variable to store the randomly selected word
var chosenWord string

// Creating an HTML template using the included Go package, within it is our HTML code for the hangman interface
var templates = template.Must(template.New("html").Parse(`

<!DOCTYPE html>
<html lang="en">
  <head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Hangman</title>
    <!-- Google Fonts -->
    <link
      href="https://fonts.googleapis.com/css2?family=Poppins:wght@400;600&display=swap"
      rel="stylesheet"
    />
    <!-- Stylesheet -->
    <link rel="stylesheet" href="/static/style.css" />
  </head>
  <body>
    <div class="container">
      <div id="options-container"></div>
      <div id="letter-container" class="letter-container hide"></div>
      <div id="user-input-section"></div>
      <canvas id="canvas"></canvas>
      <div id="new-game-container" class="new-game-popup hide">
        <div id="result-text"></div>
        <button id="new-game-button">New Game</button>
      </div>
    </div>
    <!-- Script -->
    <script src="/static/script.js"></script>
  </body>
</html>

`))

// Here starts our main function
func main() {

	// Seeds the random number generator
	rand.Seed(time.Now().UnixNano())

	// Creates a static file server to handle requests to our "/static/" path with our CSS and JS
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Enable CORS (Cross-Origin Resource Sharing) for all the routes we need
	corsHandler := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// This is what actually sets up the paths to be accepted from any place
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			// This handles preflight requests by returning a successful response
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			// Serves the HTTP request
			h.ServeHTTP(w, r)
		})
	}

	// Handles the root "/" route with the correct handler
	http.Handle("/", corsHandler(http.HandlerFunc(handler)))

	// Handles the "/generate-word" route with the correct handler
	http.Handle("/generate-word", corsHandler(http.HandlerFunc(generateWordHandler)))

	// Starts the HTTP server (localhost:8080)
	http.ListenAndServe(":8080", nil)

}

// Handler function for the main HTML page
func handler(w http.ResponseWriter, r *http.Request) {

	// Try to execute the HTML template
	err := templates.ExecuteTemplate(w, "html", nil)

	// Checks for any possible errors
	if err != nil {

		// If error, returns internal server error along with the error message
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}

// Handler function for the random word generation
func generateWordHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {

		// Handle preflight request
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.WriteHeader(http.StatusOK)
		return

	}

	// Retrieve the selected option from the previous request parameters
	optionValue := r.FormValue("option")

	// Check if said option exists
	optionArray, ok := options[optionValue]

	// IF option is not found, bad request error
	if !ok {

		http.Error(w, "Invalid option", http.StatusBadRequest)
		return

	}
	
	// Randomly chooses a word from the selected option and stores it into the chosenWord variable
	chosenWord = optionArray[rand.Intn(len(optionArray))]

	// Sets string to upper case
	chosenWord = strings.ToUpper(chosenWord)

	// Send the chosen word back to the client in JSON format so it can be used in our script.js file
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"chosenWord": chosenWord}

	// If error, internal server error
	if err := json.NewEncoder(w).Encode(response); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

}
