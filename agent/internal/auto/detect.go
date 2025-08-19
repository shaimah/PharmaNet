package auto

import (
	"pharmanet/agent/internal/connectors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Auto struct {
	detectors []connectors.Detector
}

func New() *Auto {
	// Add any generic detectors here. Each detector discovers one software type.
	return &Auto{detectors: []connectors.Detector{&GenericDetector{}}}
}

func (a *Auto) Discover() ([]connectors.Detection, error) {
	var all []connectors.Detection
	for _, d := range a.detectors {
		if dets, err := d.Detect(); err == nil {
			all = append(all, dets...)
		}
	}
	return all, nil
}

// GenericDetector detects installed software (placeholder)
type GenericDetector struct{}

func (d *GenericDetector) App() string { return "generic_pharmacy_app" }

func (d *GenericDetector) Detect() ([]connectors.Detection, error) {
	var candidates []string
	switch runtime.GOOS {
	case "windows":
		candidates = []string{`C:\Program Files`, `C:\ProgramData`}
	case "darwin":
		candidates = []string{"/Applications", filepath.Join(os.Getenv("HOME"), "Library/Application Support")}
	default:
		candidates = []string{"/opt", "/usr/local/share", "/var/lib"}
	}

	var out []connectors.Detection
	for _, base := range candidates {
		_ = filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
			if err != nil || info == nil || info.IsDir() {
				return nil
			}
			name := strings.ToLower(info.Name())
			if strings.Contains(name, "pharmacy") {
				out = append(out, connectors.Detection{
					App:  d.App(),
					Kind: "generic",
					DSN:  path,
					Meta: map[string]string{"path": path},
				})
			}
			return nil
		})
	}
	return out, nil
}
