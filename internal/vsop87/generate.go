package vsop87

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

// UnmarshalYAML implements custom unmarshaling to support both sequence
// style ([A, B, C]) and mapping style (A:..., B:..., C:...).
func (c *Coeff) UnmarshalYAML(node *yaml.Node) error {
	switch node.Kind {
	case yaml.SequenceNode:
		if len(node.Content) != 3 {
			return fmt.Errorf("expected 3 elements in sequence for Coeff, got %d", len(node.Content))
		}
		var err error
		if c.A, err = strconv.ParseFloat(node.Content[0].Value, 64); err != nil {
			return fmt.Errorf("invalid A value %q: %w", node.Content[0].Value, err)
		}
		if c.B, err = strconv.ParseFloat(node.Content[1].Value, 64); err != nil {
			return fmt.Errorf("invalid B value %q: %w", node.Content[1].Value, err)
		}
		if c.C, err = strconv.ParseFloat(node.Content[2].Value, 64); err != nil {
			return fmt.Errorf("invalid C value %q: %w", node.Content[2].Value, err)
		}
		return nil
	case yaml.MappingNode:
		type plain Coeff
		var p plain
		if err := node.Decode(&p); err != nil {
			return err
		}
		*c = Coeff(p)
		return nil
	default:
		return fmt.Errorf("unexpected YAML node kind %v for Coeff", node.Kind)
	}
}

// VSOPData maps planet names to series names (e.g., "L0","B1") to nested coefficient groups.
// Each [][]Coeff corresponds to coefficient groups for increasing powers of t.
type VSOPData map[string]map[string][][]Coeff

// GenerateSources reads the VSOP87 YAML at inFile and writes generated Go files in outDir.
// modulePath is the module import path, e.g., "github.com/yourusername/vsop87-go".
func GenerateSources(modulePath, inFile, outDir string) error {
	dataBytes, err := os.ReadFile(inFile)
	if err != nil {
		return fmt.Errorf("read input file: %w", err)
	}

	var data VSOPData
	if err := yaml.Unmarshal(dataBytes, &data); err != nil {
		return fmt.Errorf("unmarshal yaml: %w", err)
	}

	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}

	for planet, seriesMap := range data {
		fname := strings.ToLower(planet) + ".go"
		outFile := filepath.Join(outDir, fname)
		f, err := os.Create(outFile)
		if err != nil {
			return fmt.Errorf("create file %s: %w", outFile, err)
		}
		defer f.Close()

		// Write file header
		fmt.Fprintf(f, "package generated\n\n")
		fmt.Fprintf(f, "// Auto-generated VSOP87 data for %s\n", planet)
		fmt.Fprintf(f, "import \"%s/internal/vsop87\"\n\n", modulePath)
		fmt.Fprintf(f, "var (\n")

		// Generate each series
		for seriesName, nested := range seriesMap {
			varName := fmt.Sprintf("%s_%s", planet, seriesName)
			fmt.Fprintf(f, "    // %s holds coefficient groups for series %s (power of t)\n", varName, seriesName)
			fmt.Fprintf(f, "    %s = [][]vsop87.Coeff{\n", varName)
			for _, group := range nested {
				fmt.Fprintf(f, "        {\n")
				for _, c := range group {
					fmt.Fprintf(f, "            {A: %g, B: %g, C: %g},\n", c.A, c.B, c.C)
				}
				fmt.Fprintf(f, "        },\n")
			}
			fmt.Fprintf(f, "    }\n")
		}
		fmt.Fprintf(f, ")\n")
	}
	return nil
}
