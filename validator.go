package main

import (
	"fmt"
	"strings"
)

var messageBag []string

func printErrors(messageBag []string) {
	if len(messageBag) == 0 {
		return
	}
	fmt.Println(Warn(messageBag[0]))
	printErrors(messageBag[1:])
}

func validateUserInput(firstName string, lastName string, email string, tickets uint) bool {
	var validated bool

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := tickets > 0 && tickets <= remainingTickets

	validated = isValidName && isValidEmail && isValidTicketNumber

	if !validated {
		if !isValidName {
			messageBag = append(messageBag, "- first name or last name you entered is too short")
		}
		if !isValidEmail {
			messageBag = append(messageBag, "- email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			messageBag = append(messageBag, "- number of tickets you entered is invalid")
		}

		fmt.Println(Warn("Failed validation:"))
		printErrors(messageBag)
	}

	return validated
}
