package main

import (
    "fmt"
    "pharmanet/internal/config"
    "pharmanet/internal/detector"
    "pharmanet/internal/httpserver"
    "pharmanet/internal/strategy"
)

func main() {
    fmt.Println("PharmaNet Agent Starting...")

    // Load config
    config.LoadConfig("./config.yaml")

    // Step 2: Detect installed apps & prompt user
    apps := detector.DetectInstalledSoftware()
    chosenApp := detector.PromptUserSelection(apps)

    // Step 3: Instantiate Generic Strategy
    strat := strategy.Factory(chosenApp)
    err := strat.Connect()
    if err != nil {
        fmt.Println("Error connecting to app:", err)
        return
    }

    // Example search
    results, _ := strat.SearchInventory("domperidone 10mg")
    fmt.Println("Sample search results:", results)

    // Start HTTP API
    httpserver.StartServer(8081)
}
