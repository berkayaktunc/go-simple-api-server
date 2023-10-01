package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:id`
	Title    string `json:"title"`
	Author   string `json:"auther"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
	{ID: "4", Title: "I Know Why The Caged Bird Sings", Author: "Maya Angelou", Quantity: 5},
	{ID: "5", Title: "East of Eden", Author: "John Steinbeck", Quantity: 5},
	{ID: "6", Title: "The Sun Also Rises", Author: "Ernest Hemingway", Quantity: 5},
	{ID: "7", Title: "Do Androids Dream of Electric Sheep?", Author: "Philip K. Dick", Quantity: 5},
	{ID: "8", Title: "The Curious Incident of the Dog in the Night-Time", Author: "Mark Haddon", Quantity: 5},
	{ID: "9", Title: "Cloudy with a Chance of Meatballs", Author: "Judi Barrett", Quantity: 5},
	{ID: "10", Title: "Pride and Prejudice and Zombies", Author: "Seth Grahame-Smith", Quantity: 5},
	{ID: "11", Title: "The House of Mirth", Author: "Edith Wharton", Quantity: 5},
	{ID: "12", Title: "Are You There, Vodka? It's Me, Chelsea", Author: "Chelsea Handler", Quantity: 5},
	{ID: "13", Title: "And Then There Were None", Author: "Agatha Christie", Quantity: 5},
	{ID: "14", Title: "Their Eyes Were Watching God", Author: "Zora Neale Hurston", Quantity: 5},
	{ID: "15", Title: "The Devil Wears Prada", Author: "Lauren Weisberger", Quantity: 5},
	{ID: "16", Title: "Brave New World", Author: "Aldous Huxley", Quantity: 5},
	{ID: "17", Title: "Bury My Heart at Wounded Knee", Author: "Dee Brown", Quantity: 5},
	{ID: "18", Title: "The Man Who Was Thursday", Author: "G.K, Chesterton", Quantity: 5},
}

func handleError(c *gin.Context, err error, ok bool, status int, message string) bool {
	if err != nil || !ok {
		c.IndentedJSON(status, gin.H{"message": message})
		return true
	}
	return false
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if handleError(c, err, true, http.StatusNotFound, "Book not found") {
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if handleError(c, nil, ok, http.StatusBadRequest, "Missing ID in query") {
		return
	}

	book, err := getBookById(id)
	if handleError(c, err, true, http.StatusNotFound, "Book not found") {
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "no more copies left for this title"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if handleError(c, nil, ok, http.StatusBadRequest, "Missing ID in query") {
		return
	}

	book, err := getBookById(id)
	if handleError(c, err, true, http.StatusNotFound, "Book not found") {
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func addBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", addBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)

	router.Run("localhost:8080")
}
