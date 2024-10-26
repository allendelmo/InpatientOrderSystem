package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write HTML content with text boxes to the response
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Medical Request Form</title>
				<style>
					body { font-family: Arial, sans-serif; text-align: center; margin-top: 50px; }
					h1 { color: #333; }
					p { color: #666; }
					form { max-width: 300px; margin: 0 auto; }
					input[type="text"], input[type="submit"] {
						width: 100%;
						padding: 10px;
						margin: 5px 0;
						font-size: 1em;
					}
				</style>
			</head>
			<body>
				<h1>Medical Request Form</h1>
				<p>Fill out the form below:</p>
				
				<!-- Form with text boxes -->
				<form method="POST" action="/submit">
					<label for="name">File Number</label>
					<input type="text" id="name" name="name" placeholder="Enter your name" required>
					<br></br>
					<label for="email">Nurse Name</label>
					<input type="text" id="email" name="email" placeholder="Enter Nurse Name" required>
					<br></br>
					<label for="email">Ward</label>
					<input type="text" id="email" name="email" placeholder="Enter Ward Number" required>
					<br></br>
					<label for="email">Bed No.</label>
					<input type="text" id="email" name="email" placeholder="Enter Bed Number" required>
					<br></br>
					<label for="email">Medication</label>
					<input type="text" id="email" name="email" placeholder="Enter Medication" required>
					<br></br>
					<label for="email">UOM</label>
					<input type="text" id="email" name="email" placeholder="Enter UOM" required>
					<br></br>
					<label for="email">Remarks</label>
					<input type="text" id="email" name="email" placeholder="Enter Remarks" required>


					<input type="submit" value="Submit">
				</form>
			</body>
			</html>
		`)
	})

	// Handler for form submission
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			email := r.FormValue("email")
			fmt.Fprintf(w, `
				<!DOCTYPE html>
				<html lang="en">
				<head>
					<meta charset="UTF-8">
					<meta name="viewport" content="width=device-width, initial-scale=1.0">
					<title>Form Submission</title>
				</head>
				<body>
					<h1>Thank You!</h1>
					<p>Your submission has been received.</p>
					<p><strong>Name:</strong> %s</p>
					<p><strong>Email:</strong> %s</p>
					<a href="/">Go back</a>
				</body>
				</html>
			`, name, email)
		}
	})

	// Start the server on port 8080
	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
