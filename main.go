package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Lucifer07/parking-lot/entity"
	"github.com/Lucifer07/parking-lot/parking"
)

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	exit := false
	menu := "Parking Lot\n" +
		"1. Setup\n" +
		"2. Park\n" +
		"3. UnPark\n" +
		"4. Exit"
	attendant := parking.Attendant{}
	attendant.ChangeStyle(parking.StyleDefault())
	var name string
	var (
		login = false
	)
	tickets := []*entity.Ticket{}
	for !exit {
		fmt.Println(menu)
		input := promptInput(scanner, "\nInput menu: ")
		fmt.Println(input)
		switch input {
		case "1":
			if !login {
				dashboard := promptInput(scanner, "Please  Register first, insert y to register :")
				if strings.ToLower(dashboard) != "y" {
					fmt.Println("please create account first to using this app")
					break
				} else {
					names := promptInput(scanner, "Please insert youre name :")
					name = names
					capacity := promptInput(scanner, "Please insert how many youre Parkinglot :")
					size, err := strconv.Atoi(capacity)
					if err != nil || size <= 0 {
						fmt.Println("Invalid capacity")
						break
					}
					attendant.AddParkinglot(size)
					login = true
				}

			} else {
				fmt.Printf("\nhi %s, please select another menu \n", name)
				break
			}

		case "2":
			if !login {
				fmt.Println("please create account first")
				break
			}
			plateNumber := promptInput(scanner, "Input your car plate number: ")
			ticket, err := attendant.Park(entity.AddCar(plateNumber))
			if err != nil {
				fmt.Println(err)
				break
			} else {
				tickets = append(tickets, ticket)
			}
			fmt.Printf("ticket: %v\n", ticket.ID)
		case "3":
			if !login {
				fmt.Println("please create account first")
				break
			}
			fmt.Println("\n Tickets:")
			for i, ticket := range tickets {
				fmt.Printf("%v. %v\n", i+1, ticket.ID)
			}
			ticketNumber := promptInput(scanner, "Select ticket :")
			selectedTicket, err := strconv.Atoi(ticketNumber)
			if err != nil || selectedTicket > len(tickets) || selectedTicket <= 0 {
				fmt.Println("Invalid option")
				break
			}
			car, _ := attendant.UnPark(tickets[selectedTicket-1])
			tickets = append(tickets[:selectedTicket-1], tickets[selectedTicket:]...)
			fmt.Printf("\n this is your car with number plate : %v \n", *car)
			break
		case "4":
			fmt.Printf("Don't forget to smile %s, because your smile is the happiness of the world.", name)
			exit = true
		default:
			fmt.Println("Invalid menu")
		}
		fmt.Println()
	}
}
