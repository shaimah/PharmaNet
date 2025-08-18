// internal/detector/detector.go
package detector

import (
    "bufio"
    "fmt"
    "os"
    "runtime"
    "strings"
)

// DetectInstalledSoftware lists apps installed on the system
func DetectInstalledSoftware() []string {
    var apps []string

    switch runtime.GOOS {
    case "windows":
        // PowerShell Get-StartApps
        out, _ := exec.Command("powershell", "Get-StartApps | Select-Object -ExpandProperty Name").Output()
        apps = strings.Split(string(out), "\n")
    case "darwin":
        out, _ := exec.Command("ls", "/Applications").Output()
        apps = strings.Split(string(out), "\n")
    case "linux":
        out, _ := exec.Command("ls", "/usr/share/applications").Output()
        apps = strings.Split(string(out), "\n")
    }

    var cleanApps []string
    for _, a := range apps {
        trimmed := strings.TrimSpace(a)
        if trimmed != "" {
            cleanApps = append(cleanApps, trimmed)
        }
    }
    return cleanApps
}

// PromptUserSelection prints apps and asks user to choose one
func PromptUserSelection(apps []string) string {
    fmt.Println("Detected installed apps:")
    for i, app := range apps {
        fmt.Printf("%d) %s\n", i+1, app)
    }

    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Enter the number of the software to allow the agent to access: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        choice := 0
        fmt.Sscanf(input, "%d", &choice)

        if choice > 0 && choice <= len(apps) {
            fmt.Println("Selected:", apps[choice-1])
            return apps[choice-1]
        }
        fmt.Println("Invalid choice. Try again.")
    }
}
