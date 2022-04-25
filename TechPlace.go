package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	var message string;

    message = "Welcome to the Tech Palace, " + strings.ToUpper(customer);
    return (message);
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    var border string;
	var borderedMessage string;
    
    border = strings.Repeat("*", numStarsPerLine);
    borderedMessage = border + "\n" + welcomeMsg + "\n" + border;
    return (borderedMessage);
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    var finalMessage string;

    finalMessage = strings.ReplaceAll(oldMsg, "*", " ");
    finalMessage = strings.TrimSpace(finalMessage);
    return (finalMessage);
}

