package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	var message string;
    
    message =  fmt.Sprintf("Welcome to my party, %s!", name );
    return (message);
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	var message string;

    message = fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age);
    return (message);
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var message, line1, line2, line3 string;
	
    line1 = Welcome(name);
    line2 = fmt.Sprintf("You have been assigned to table %003d. Your table is %s, exactly %.1f meters from here.", table, direction, distance);
    line3 = fmt.Sprintf("You will be sitting next to %s.", neighbor);
    message = fmt.Sprintf("%s\n%s\n%s", line1, line2, line3);
    return (message);
}

