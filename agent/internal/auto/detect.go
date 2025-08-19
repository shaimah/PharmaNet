package auto

import (
	"pharmanet/agent/internal/connectors"
)

// Auto orchestrates all detectors
type Auto struct {
	detectors []connectors.Detector
}

// New returns an Auto orchestrator with generic detectors
func New() *Auto {
	return &Auto{
		detectors: []connectors.Detector{
			&GenericDetector{},
		},
	}
}

// Discover returns all detections from user-selected software
func (a *Auto) Discover(selectedApps []string) ([]connectors.Detection, error) {
	var all []connectors.Detection
	for _, d := range a.detectors {
		dets, err := d.Detect()
		if err != nil {
			continue
		}
		for _, det := range dets {
			// Only include if the user selected this app
			for _, sel := range selectedApps {
				if det.App == sel {
					all = append(all, det)
				}
			}
		}
	}
	return all, nil
}

// GenericDetector is a placeholder detector for any installed software
type GenericDetector struct{}

// App returns the generic name
func (d *GenericDetector) App() string { return "generic_pharmacy_app" }

// Detect finds installed apps (scan paths)
func (d *GenericDetector) Detect() ([]connectors.Detection, error) {
	// Placeholder: in production, scan installed programs, AppData, /Applications, etc.
	// Return dummy detection for demonstration
	return []connectors.Detection{
		{App: "generic_pharmacy_app", Kind: "generic", DSN: "path/to/software", Meta: map[string]string{"info": "dummy"}},
	}, nil
}
