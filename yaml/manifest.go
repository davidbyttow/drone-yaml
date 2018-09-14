package yaml

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

type (
	// Manifest is a collection of Drone resources.
	Manifest struct {
		Resources []Resource
	}

	// Resource represents a Drone resource.
	Resource interface {
		// only objects in this package can
		// define themselves as resources.
		resource()
	}

	resource struct {
		Kind string `json:"kind"`
		Type string `json:"type"`
	}
)

// UnmarshalJSON implement the json.Unmarshaler.
func (m *Manifest) UnmarshalJSON(b []byte) error {
	messages := []json.RawMessage{}
	err := json.Unmarshal(b, &messages)
	if err != nil {
		return err
	}
	for _, message := range messages {
		res := new(resource)
		err := json.Unmarshal(message, res)
		if err != nil {
			return err
		}
		var obj Resource
		switch res.Kind {
		case "cron":
			obj = new(Cron)
		case "secret":
			obj = new(Secret)
		case "signature":
			obj = new(Signature)
		case "registry":
			obj = new(Registry)
		default:
			obj = new(Pipeline)
		}
		err = json.Unmarshal(message, obj)
		if err != nil {
			return err
		}
		m.Resources = append(m.Resources, obj)
	}
	return nil
}

// MarshalJSON implement the json.Marshaler.
func (m *Manifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Resources)
}

// MarshalYAML implement the yaml.Marshaler.
func (m *Manifest) MarshalYAML() (interface{}, error) {
	return yaml.Marshal(m.Resources)
}

//
// structures that implement is_resource
//

func (*Cron) resource()      {}
func (*Pipeline) resource()  {}
func (*Registry) resource()  {}
func (*Secret) resource()    {}
func (*Signature) resource() {}
