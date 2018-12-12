package yaml

import "errors"

type (
	// Predicate declares a custom step in the pipeline to be run if triggered by
	// a conditional step
	Predicate struct {
		Version string `json:"version,omitempty"`
		Kind    string `json:"kind,omitempty"`
		Type    string `json:"type,omitempty"`
		Name    string `json:"name,omitempty"`

		Spec Spec `json:"spec,omitempty"`
	}

	// Spec defines the parameters for the predicate to run
	Spec struct {
		Image string `json:"image,omitempty"`
	}
)

// GetVersion returns the resource version.
func (p *Predicate) GetVersion() string { return p.Version }

// GetKind returns the resource k ind.
func (p *Predicate) GetKind() string { return p.Kind }

// Validate returns an error if the predicate is invalid.
func (p *Predicate) Validate() error {
	if p.Name == "" || p.Spec.Image == "" || p.Type == "" {
		return errors.New("yaml: invalid predicate resource")
	}
	return nil
}
