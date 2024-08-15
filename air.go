// airparser is a parser for Civitai AIR URNs
// https://developer.civitai.com/docs/getting-started/ai-resource-identifier
package air_parser

import (
	"fmt"
	"strings"
)

// This struct represents an AIR resource URN
type Air struct {
	Ecosystem string
	Type      string
	Source    string
	Id        string
	Version   string
	Layer     string
	Format    string
	urn       string
}

// Get the AIR URN that this struct represents.
func (a *Air) Urn() string {
	if a.urn != "" {
		return a.urn
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("urn:air:%s:%s:%s:%s", a.Ecosystem, a.Type, a.Source, a.Id))

	if a.Version != "" {
		builder.WriteByte('@')
		builder.WriteString(a.Version)
	}

	if a.Layer != "" {
		builder.WriteByte(':')
		builder.WriteString(a.Layer)
	}

	if a.Format != "" {
		builder.WriteByte('.')
		builder.WriteString(a.Format)
	}

	a.urn = builder.String()
	return a.urn
}

// Create a new AirResource from a string e.g. urn:air:sd1:model:civitai:2421@43533
func NewAirFromString(airStr string) (*Air, error) {

	if len(airStr) < 7 {
		return nil, fmt.Errorf("value too short to be an AIR urn")
	}

	if !strings.HasPrefix(strings.ToLower(airStr), "urn:air") {
		return nil, fmt.Errorf("missing urn:air prefix")
	}

	var resourse Air
	var elements []string

	// if a format is specified strip and store it
	delimterIndex := strings.LastIndex(airStr, ".")
	if delimterIndex != -1 {
		elements = strings.Split(airStr[:delimterIndex], ":")
		resourse.Format = airStr[delimterIndex+1:]
	} else {
		elements = strings.Split(airStr, ":")
	}

	if len(elements) < 6 {
		return nil, fmt.Errorf("invalid AIR urn expected 6 ':' delimited elements got: %d", len(elements))
	} else if len(elements) == 7 {
		resourse.Layer = elements[6]
	}

	resourse.Ecosystem = elements[2]
	resourse.Type = elements[3]
	resourse.Source = elements[4]

	delimiterIndex := strings.LastIndex(elements[5], "@")

	if delimiterIndex == -1 {
		resourse.Id = elements[5]
	} else {
		resourse.Id = elements[5][:delimiterIndex]
		resourse.Version = elements[5][delimiterIndex+1:]
	}

	return &resourse, nil
}
