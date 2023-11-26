package main

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var options = map[string][]string{

	"fruits":    {"Apple", "Blueberry", "Mandarin", "Pineapple", "Pomegranate", "Watermelon"},
	"animals":   {"Hedgehog", "Rhinoceros", "Squirrel", "Panther", "Walrus", "Zebra"},
	"countries": {"India", "Hungary", "Kyrgyzstan", "Switzerland", "Zimbabwe", "Dominica"},
}

var chosenWord string

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

func main() {

	rand.Seed(time.Now().UnixNano())
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler)
	http.HandleFunc("/generate-word", generateWordHandler)

	// Enable CORS
	corsHandler := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			h.ServeHTTP(w, r)
		})
	}

	http.Handle("/", corsHandler(http.HandlerFunc(handler)))
	http.Handle("/generate-word", corsHandler(http.HandlerFunc(generateWordHandler)))

	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {

	err := templates.ExecuteTemplate(w, "html", nil)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}

func generateWordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		// Handle preflight request
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Handle the actual request
	optionValue := r.FormValue("option")
	optionArray, ok := options[optionValue]

	if !ok {
		http.Error(w, "Invalid option", http.StatusBadRequest)
		return
	}

	chosenWord = optionArray[rand.Intn(len(optionArray))]
	chosenWord = strings.ToUpper(chosenWord)

	// Send the chosen word back to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"chosenWord": chosenWord})
}
