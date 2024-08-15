package air_parser

import (
	"testing"
)

func TestValidAirValues(t *testing.T) {
	cases := make(map[string]Air)
	cases["urn:air:sd1:model:civitai:2421@43533:layer.safetensors"] = Air{
		Ecosystem: "sd1",
		Type:      "model",
		Source:    "civitai",
		Id:        "2421",
		Version:   "43533",
		Layer:     "layer",
		Format:    "safetensors",
	}

	cases["urn:air:sd1:model:civitai:2421@43533:layer"] = Air{
		Ecosystem: "sd1",
		Type:      "model",
		Source:    "civitai",
		Id:        "2421",
		Version:   "43533",
		Layer:     "layer",
	}
	cases["urn:air:sd1:model:civitai:2421@43533"] = Air{
		Ecosystem: "sd1",
		Type:      "model",
		Source:    "civitai",
		Id:        "2421",
		Version:   "43533",
	}

	cases["urn:air:sd1:model:civitai:2421"] = Air{
		Ecosystem: "sd1",
		Type:      "model",
		Source:    "civitai",
		Id:        "2421",
	}

	cases["urn:air:sd1:model:civitai:2421:layer"] = Air{
		Ecosystem: "sd1",
		Type:      "model",
		Source:    "civitai",
		Id:        "2421",
		Layer:     "layer",
	}

	cases["urn:air:sd1:model:civitai:2421:layer.safetensors"] = Air{
		Ecosystem: "sd1",
		Type:      "model",
		Source:    "civitai",
		Id:        "2421",
		Layer:     "layer",
		Format:    "safetensors",
	}

	cases["urn:air:sd1:model:civitai:2421.safetensors"] = Air{
		Ecosystem: "sd1",
		Type:      "model",
		Source:    "civitai",
		Id:        "2421",
		Format:    "safetensors",
	}

	for urn, expectedResult := range cases {
		parsedResult, err := NewAirFromString(urn)
		if err != nil {
			panic(err)
		}

		if *parsedResult != expectedResult {
			t.Fatalf("expected: %v got: %v", expectedResult, parsedResult)
		}

		if parsedResult.Urn() != urn {
			t.Fatalf("urn mismatch. expect: '%s' got '%s'", urn, parsedResult.Urn())
		}

	}
}
