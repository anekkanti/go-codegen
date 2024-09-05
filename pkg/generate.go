package pkg

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

// Input is the input for the generate function
type Input struct {
	DirectoryPath string `validate:"required"`
	PluginPath    string `validate:"required"`
	OutputPath    string `validate:"required"`

	NameFilters []string
}

func Generate(input *Input) error {

	if err := validate.Struct(input); err != nil {
		return fmt.Errorf("invalid input: %w", err)
	}

	files, err := ParseDir(input.DirectoryPath, nil)
	if err != nil {
		return fmt.Errorf("failed to parse: %w", err)
	}

	if err := Render(
		input.PluginPath,
		input.OutputPath,
		input.NameFilters,
		files,
	); err != nil {
		return fmt.Errorf("failed to render: %w", err)
	}
	return nil
}
