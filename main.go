package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/gomail.v2"
)

// Struct to represent the data received in the HTTP request
type Contact struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Linkedin       string `json:"linkedin"`
	Tech           string `json:"tech"`
	Message        string `json:"message"`
}


// Struct to represent the JSON response
type Response struct {
	Message string `json:"message"`
}

var client *mongo.Client // Global variable for the MongoDB client

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables from .env file:", err)
	}

// Establish connection to MongoDB
	mongoDBURI := os.Getenv("MONGODB_URI")

	// Configure client options
	clientOptions := options.Client().ApplyURI(mongoDBURI)

	// Connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping to confirm connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	r := http.NewServeMux()
	r.HandleFunc("/api/contact", handleContact)

	http.Handle("/", r)
	port := os.Getenv("PORT")
if port == "" {
    port = "8080" // Set a default port if not defined in the .env file
}

fmt.Println("API started on port :" + port)
log.Fatal(http.ListenAndServe(":"+port, nil))

}

func handleContact(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON data from the request
	var contact Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Error decoding JSON:", err)
		return
	}

	// Get the MongoDB collection to store the contact data
	collection := client.Database("franciscoinoque-tech").Collection("contacts")

	// Insert data into MongoDB
	_, err = collection.InsertOne(context.TODO(), contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error inserting data into MongoDB:", err)
		return
	}
	log.Println("Data inserted into MongoDB successfully")

	//Send an email
	err = sendEmail(contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error sending email:", err)
		return
	}
	log.Println("Email successfully sent")

	// Send a JSON response
	response := Response{Message: "Data saved successfully and email sent!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}


func sendEmail(contact Contact) error {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading environment variables from .env file")
        return err
    }

    // Get SMTP server settings from .env file
    smtpHost := os.Getenv("SMTP_HOST")
    smtpPortStr := os.Getenv("SMTP_PORT") // Read as string
    smtpUsername := os.Getenv("SMTP_USERNAME")
    smtpPassword := os.Getenv("SMTP_PASSWORD")

    // Convert port string to int
    smtpPort, err := strconv.Atoi(smtpPortStr)
    if err != nil {
        log.Fatal("Erro ao converter a porta SMTP para int:", err)
        return err
    }

    // Configure the email client
    m := gomail.NewMessage()
    m.SetHeader("From", smtpUsername) // Use the sender's email address
    m.SetHeader("To", contact.Email)  // Recipient's email address
    m.SetHeader("Subject", "I've Received Your Message!")

    // Message body
    body := fmt.Sprintf("Hi,\n\n" +
    "I just wanted to let you know that I've received your message. I'm glad you reached out.\n\n" +
    "I'm taking a look at your message, and I'll get back to you shortly. I love receiving messages and I'm excited to chat more.\n\n" +
    "If you need anything or want to discuss further, feel free to shoot me an email directly at " + smtpUsername + ".\n\n" +
    "Thanks for choosing to write to me. I can't wait to continue our conversation.\n\n" +
    "Cheers,\n\n" +
    "Francisco Inqoue")
    m.SetBody("text/plain", body)

    // Configure the SMTP client with .env variables
    d := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)

    // Send the email
    if err := d.DialAndSend(m); err != nil {
        log.Println("Error sending email:", err)
        return err
    }

    log.Println("Email successfully sent")
    return nil
}
