package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

var conferenceName string = "Go Conference"

const conferenceTickets uint8 = 50

var remainingTickets uint8 = conferenceTickets
var bookings []string = []string{}

func main() {
	greetUser()

	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		isInvalid := !isValidName || !isValidEmail || !isValidTicket
		if isInvalid {
			if !isValidName {
				fmt.Println("Your first name or last name is too short!")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain '@' sign!")
			}
			if !isValidTicket {
				fmt.Println("Number of tickets you entered is invalid!")
			}

			fmt.Println("Try again!")
			continue
		}

		bookTicket(firstName, lastName, email, userTickets)

		firstNames := getFirstNames()
		fmt.Printf("The first names of the bookings are: %v\n", firstNames)
	}
}

func greetUser() {
	fmt.Printf("conferenceName is %T and conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func getUserInput() (string, string, string, uint8) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint8

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint8) {
	remainingTickets -= userTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining tickets are: %v\n", remainingTickets)
}
