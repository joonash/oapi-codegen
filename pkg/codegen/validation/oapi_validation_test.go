package validation

import (
	"encoding/base64"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type SubStruct struct {
	Name string `validate:"omitempty,pattern='XmhlbGxvJHxed29ybGRbMC05XSsk'"` // "^hello$|^world[0-9]+$"
}

type PatternStruct struct {
	Name string `validate:"omitempty,pattern='XmhlbGxvJHxed29ybGRbMC05XSsk'"` // "^hello$|^world[0-9]+$"
	Relatives []SubStruct `validate:"omitempty,dive"`
}

func TestGenerateBase64EncodedString(t *testing.T) {
	encodedPattern := base64.StdEncoding.EncodeToString([]byte("^hello$|^world[0-9]+$"))
	log.Printf("Pattern: '%s'", encodedPattern)
}

func generatePatternObject(names... string) *PatternStruct {
	pStruct := PatternStruct{
		Name: names[0],
	}

	var subStructs []SubStruct

	for _, s := range names[1:] {
		subStructs = append(subStructs, SubStruct{Name: s})
	}

	pStruct.Relatives = subStructs
	return &pStruct
}

func TestPatternValidation(t *testing.T) {
	validatable := generatePatternObject("hello", "", "hello", "world2221")

	testValidator := validator.New()
	RegisterValidations(testValidator)
	err := testValidator.Struct(validatable)
	assert.NoError(t, err, "Pattern validation failed ")
}

func TestPatternValidationFailMainLevel(t *testing.T) {
	validatable := generatePatternObject("hellox", "", "hello", "world2221")
	testValidator := validator.New()
	RegisterValidations(testValidator)
	err := testValidator.Struct(validatable)
	assert.Error(t, err, "Pattern validation failed ")
	validationErrors, _ := err.(validator.ValidationErrors)
	assert.Len(t, validationErrors, 1, "Invalid number of validation errors")
	assert.Equal(t, validationErrors[0].Namespace(), "PatternStruct.Name")
	assert.Equal(t, validationErrors[0].Tag(), "pattern")
}

func TestPatternValidationSubStructLevel(t *testing.T) {
	validatable := generatePatternObject("hello", "", "hellox", "World2221")
	testValidator := validator.New()
	RegisterValidations(testValidator)
	err := testValidator.Struct(validatable)
	assert.Error(t, err, "Pattern validation failed ")
	validationErrors, _ := err.(validator.ValidationErrors)
	assert.Len(t, validationErrors, 2, "Invalid number of validation errors")
	assert.Equal(t, validationErrors[0].Namespace(), "PatternStruct.Relatives[1].Name")
	assert.Equal(t, validationErrors[0].Tag(), "pattern")
	assert.Equal(t, validationErrors[1].Namespace(), "PatternStruct.Relatives[2].Name")
	assert.Equal(t, validationErrors[1].Tag(), "pattern")
}
