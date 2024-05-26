package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string //Defining an array with size is an array and w/o a size is a slice

	//Welcome Page
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available! \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
	fmt.Println()

	//User Information
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for remainingTickets > 0 && len(bookings) < 50 {

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets required: ")
		fmt.Scan(&userTickets)

		//Conditional Stmts
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets for the conference with us. You will receive a confirmation email on %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("\n%v are the number of tickets remaining for the %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings { //for-each loop
				//the underscore is considered as a blank identifier "_", it comes in place of index - meaning we have a variable there but we don't need to use it right now so we skipping it
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("Bookings done by: %v\n", bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				fmt.Println("Our conference is booked out. Come back next year!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The first name or last name entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("The email address doesn't contain the @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets is invalid.")
			}
		}
	}
}
