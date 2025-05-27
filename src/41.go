package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Generate a random number between 0 and 1
    var randomNumber float64 = rand.Float64()

    // Convert the random number to minutes since midnight
    var secondsSinceMidnight int = time.Since(time.Now()).Seconds()

    // Calculate the hours of the current day
    var hourOfDay int = time.Hour

    // Calculate the days remaining in the month
    var daysRemainingInMonth int = 31 - hoursSinceMidnight

    if daysRemainingInMonth <= 0 {
        // If there are no more days left, set the remaining days to 7 (all days)
        daysRemainingInMonth = 7
    }

    for i := 1; i < daysRemainingInMonth + 1; i++ {
        hourOfDay += randomNumber * float64(i)

        if hoursSinceMidnight > 0 {
            // If there are more than 23 minutes left until midnight, add 60 to the current time
            hourOfDay = int(hourOfDay) % 24
        }

        fmt.Printf("%d:%02d\n", hourOfDay/60, hourOfDay%60)
    }
}
