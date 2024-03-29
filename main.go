package main

import (
    "fmt"
    "log"
    "strings"

    "github.com/gocolly/colly/v2"
)

// Match represents a football match.
type Match struct {
    FirstTeam  string 
    Score      string 
    SecondTeam string 
    Timing     string 
    LeagueName string 
}

// Constants for color formatting.
const (
    ColorRed    = "\033[31m"
    ColorGreen  = "\033[32m"
    ColorYellow = "\033[33m"
    ColorReset  = "\033[0m"
)

func main() {
    // Create a new collector instance.
    c := colly.NewCollector()

    // Slice to hold match data.
    matches := []Match{}

    // Callback function to extract match data from HTML.
    c.OnHTML(".match-container", func(h *colly.HTMLElement) {
        match := Match{
            FirstTeam:  h.ChildText("div.left-team .team-name"),
            Score:      h.ChildText("div#result-now"),
            SecondTeam: h.ChildText(".right-team .team-name"),
            Timing:     h.ChildText(".match-timing .date"),
            LeagueName: h.ChildTexts(".match-info span")[2],
        }

        // Append match to matches slice.
        matches = append(matches, match)
    })

    // Visit the URL to start scraping.
    if err := c.Visit("https://koora-live.tv/livestreaming/"); err != nil {
        log.Fatalf("Failed to visit website: %v", err)
    }

    // Print match information.
    for _, match := range matches {
        // Color code timing information.
        var timingColor string
        switch strings.ToLower(match.Timing) {
        case "live":
            timingColor = ColorGreen
        case "finished":
            timingColor = ColorRed
        default:
            timingColor = ColorYellow
        }

        // Print formatted match information.
        fmt.Printf("( %-24s )\t %-25s  %-13s  %-25s  %-30s \n", 
            timingColor+match.Timing+ColorReset, 
            match.FirstTeam, 
            match.Score, 
            match.SecondTeam, 
            match.LeagueName,
        )
    }
}
