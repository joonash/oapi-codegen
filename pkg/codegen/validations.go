package codegen

import (
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pkg/errors"
)

func CompilePropertyValidations(propertyName string, propertySchema *openapi3.Schema) (*[]string, error) {
	var validations []string

	if propertySchema.Max != nil {
		if propertySchema.ExclusiveMax {
			validations = append(validations, fmt.Sprintf("lt=%f", *propertySchema.Max))
		} else {
			validations = append(validations, fmt.Sprintf("max=%f", *propertySchema.Max))
		}
	}

	if propertySchema.Min != nil {
		if propertySchema.ExclusiveMin {
			validations = append(validations, fmt.Sprintf("gt=%f", *propertySchema.Min))
		} else {
			validations = append(validations, fmt.Sprintf("min=%f", *propertySchema.Min))
		}
	}

	if propertySchema.MaxLength != nil {
		validations = append(validations, fmt.Sprintf("max=%f", *propertySchema.Max))
	}

	if propertySchema.MinLength > 0 {
		validations = append(validations, fmt.Sprintf("min=%f", *propertySchema.Min))
	}

	if len(propertySchema.Pattern) > 0 {
		// Regex not supported by validation
	}

	if len(propertySchema.Enum) > 0 {

	}

	if val, ok := propertySchema.Extensions["x-oapi-codegen-validate"]; ok {
		msg := val.(json.RawMessage)
		if err := json.Unmarshal(msg, &pValidations); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error generating validations for property '%s'.", propertyName))
		}
	}

}
