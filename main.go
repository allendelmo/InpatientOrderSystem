// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"

// 	_ "github.com/mattn/go-sqlite3" // SQLite driver
// )

// // TODO: struct for Todo
// type Todo struct {
// 	Id          int64
// 	Description string
// 	IsCompleted bool
// }

// // TODO: initialize DB
// var DB *sql.DB

// tmpl := template.Must(template.New("index").Parse(`
// 	<!DOCTYPE html>
// <html lang="en">
// <head>
//     <meta charset="UTF-8">
//     <meta name="viewport" content="width=device-width, initial-scale=1.0">
//     <title>TODO LIST</title>
//     <style>
//         /* General Styling */
//         body {
//             font-family: 'Arial', sans-serif;
//             background-color: #f4f4f9;
//             color: #333;
//             display: flex;
//             justify-content: center;
//             align-items: center;
//             height: 100vh;
//             margin: 0;
//         }

//         /* Container */
//         .container {
//             background-color: #fff;
//             padding: 20px;
//             border-radius: 10px;
//             box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
//             max-width: 500px;
//             width: 100%;
//         }

//         /* Header */
//         h1 {
//             text-align: center;
//             color: #333;
//             margin-bottom: 20px;
//         }

//         /* Form styling */
//         form {
//             display: flex;
//             justify-content: space-between;
//             margin-bottom: 20px;
//         }

//         input[type="text"] {
//             width: 70%;
//             padding: 10px;
//             font-size: 1rem;
//             border: 1px solid #ddd;
//             border-radius: 5px;
//             outline: none;
//         }

//         button {
//             padding: 10px 20px;
//             background-color: #28a745;
//             color: #fff;
//             border: none;
//             border-radius: 5px;
//             cursor: pointer;
//             transition: background-color 0.3s;
//         }

//         button:hover {
//             background-color: #218838;
//         }

//         /* Todo List Styling */
//         ul {
//             list-style-type: none;
//             padding: 0;
//         }

//         li {
//             display: flex;
//             justify-content: space-between;
//             align-items: center;
//             background-color: #f9f9f9;
//             padding: 10px;
//             border: 1px solid #ddd;
//             border-radius: 5px;
//             margin-bottom: 10px;
//         }

//         li label {
//             flex-grow: 1;
//             margin-left: 10px;
//         }

//         /* Completed task style */
//         .completed {

//             color: #888;
//         }

//         /* Delete link */
//         a {
//             color: #dc3545;
//             text-decoration: none;
//             font-weight: bold;
//         }

//         a:hover {
//             text-decoration: underline;
//         }

//         /* Checkbox Styling */
//         input[type="checkbox"] {
//             transform: scale(1.5);
//             cursor: pointer;
//         }
//     </style>
// </head>
// <body>

//     <div class="container">
//         <h1>Todo List</h1>

//         <form action="/create" method="POST">
//             <input type="text" name="DESCRIPTION" placeholder="New Todo" required>
//             <button type="submit">Add</button>
//         </form>

//         <ul>
//             {{range .}}
//             <li>
//                 <input type="checkbox" id="isCompleted_{{.Id}}" name="isCompleted" {{if .IsCompleted}}checked{{end}}>
//                 <label for="isCompleted_{{.Id}}" class="{{if .IsCompleted}}completed{{end}}">
//                     {{.Description}}
//                 </label>
//                 <a href="/delete?ID={{.Id}}" onclick="return confirm('Are you sure you want to delete this item?')">Delete</a>
//             </li>
//             {{end}}
//         </ul>
//     </div>

// </body>
// </html>