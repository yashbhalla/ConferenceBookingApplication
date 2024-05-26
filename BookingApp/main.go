package main

import (
	"bookingapp/helper"
	"fmt"
	"sync"
	"time"
)

// package level variables aka global variables
var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0) //Map replaced by Struct - make([]map[string]string, 0)

// Structure
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	//Welcome Page
	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {

		//User Information
		firstName, lastName, email, userTickets := getUserInput()

		//Validate User Input
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		//Conditional Stmts
		if isValidName && isValidEmail && isValidTicketNumber {
			//User Booking
			bookTickets(userTickets, firstName, lastName, email)

			wg.Add(1)                                              // Sets the number of goroutines to wait for, it increases the counter by the provided number "wg.Add(int)""
			go sendTicket(userTickets, firstName, lastName, email) //"go" - starts a new goroutine which is a lightweight thread managed by the Go routine

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			//noTicketsRemaining is bool type
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

		/* Switch Case Stmts
		city := "London"

		switch city{
			case "New York":
				//Code for NYC
			case "Berlin":
				//Code for Berlin
			case "London":
				//Code for London
			case "Hong Kong", "Singapore":
				//Code for Hong Kong or Singapore - same logic is applied to both
			case "Mexico City":
				//Code for Mexico City
		}
		*/
	}
	wg.Wait() //Blocks until WaitGroup counter is 0
}

// WaitGroup waits for the launched goroutine to finish else terminal ends before the end of all code execution - especially the sleep timer that we have put
var wg = sync.WaitGroup{}

func greetUsers() {
	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available! \n", conferenceName, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
	fmt.Println()
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//for-each loop
		//the underscore is considered as a blank identifier "_", it comes in place of index - meaning we have a variable there but we don't need to use it right now so we skipping it
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets required: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	//creating map - collection of key value pair
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is \n %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets for the conference with us. You will receive a confirmation email on %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("\n%v are the number of tickets remaining for the %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) //saving print line into a variable
	fmt.Println("\n#######")
	fmt.Printf("Sending %v to email address %v\n", ticket, email)
	fmt.Println("######")
	wg.Done() //Decrements the WaitGroup counter by 1, so it is called by the goroutine to indicate that it is finished executing
}
