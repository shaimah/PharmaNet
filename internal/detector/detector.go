package detector

import (
    "fmt"
    "os/exec"
    "runtime"
    "strings"
)

// DetectInstalledSoftware lists apps for Windows, macOS, Linux
func DetectInstalledSoftware() []string {
    var apps []string

    switch runtime.GOOS {
    case "windows":
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

// PromptUserSelection lets the recipient choose which software the agent can access
func PromptUserSelection(apps []string) string {
    fmt.Println("Detected installed apps:")
    for i, app := range apps {
        fmt.Printf("%d) %s\n", i+1, app)
    }

    var choice int
    fmt.Print("Select the software to allow the agent to access: ")
    fmt.Scan(&choice)
    if choice > 0 && choice <= len(apps) {
        return apps[choice-1]
    }
    return ""
}
