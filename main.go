package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/smtp"
	"os"
)

const conferenceName = "Go Conference"
const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(Fata("Error loading .env file"))
	}

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			fmt.Printf("The first names of bookings are: %v\n", getFirstNames())

			if remainingTickets == 0 {
				// end program
				fmt.Println(Info("Our conference is booked out. Come back next year."))
				break
			}
		} else {
			fmt.Println(Warn("Failed validation:"))
			if !isValidName {
				fmt.Println(Warn("- first name or last name you entered is too short"))
			}
			if !isValidEmail {
				fmt.Println(Warn("- email address you entered doesn't contain @ sign"))
			}
			if !isValidTicketNumber {
				fmt.Println(Warn("- number of tickets you entered is invalid"))
			}
			fmt.Println("Try again...")
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their input
	fmt.Println("Enter your first name: ")
	_, _ = fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	_, _ = fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	_, _ = fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	_, _ = fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	auth := smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST"))

	to := []string{email}
	from := os.Getenv("MAIL_FROM_ADDRESS")
	msg := []byte(
		fmt.Sprintf("From: %v\n", from) +
			fmt.Sprintf("To: %v\n", email) +
			fmt.Sprintf("Subject: '%v' tickets\n\n", conferenceName) +
			fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName))

	err := smtp.SendMail(fmt.Sprintf("%v:%v", os.Getenv("MAIL_HOST"), os.Getenv("MAIL_PORT")), auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
