package validation

import (
	"encoding/base64"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"strings"
)

func base64Regex(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()
	parameter := fl.Param()

	if strings.Index(parameter, "'") == 0 && strings.LastIndex(parameter, "'") == len(parameter) - 1 {
		parameter = parameter[1:]
		parameter = parameter[0:len(parameter) - 1]
	}

	decoded, err := base64.StdEncoding.DecodeString(parameter)

	if err != nil {
		panic(err.Error())
	}

	matched, err := regexp.MatchString(string(decoded), fieldValue)

	if err != nil {
		panic(err.Error())
	}

	return matched
}

func multipleOf(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().Int()
	parameter := fl.Param()
	nominator, err := strconv.Atoi(parameter)

	if err != nil {
		panic(err.Error())
	}

	return fieldValue % int64(nominator) == 0
}

func base64JsonEnum(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()
	parameter := fl.Param()

	if strings.Index(parameter, "'") == 0 && strings.LastIndex(parameter, "'") == len(parameter) - 1 {
		parameter = parameter[1:]
		parameter = parameter[0:len(parameter) - 1]
	}

	decoded, err := base64.StdEncoding.DecodeString(parameter)

	if err != nil {
		panic(err.Error())
	}

	var enumValues []string
	err = json.Unmarshal(decoded, &enumValues)

	if err != nil {
		panic(err.Error())
	}

	for _, value := range enumValues {
		if value == fieldValue {
			return true
		}
	}

	return false
}

func RegisterValidations(validate *validator.Validate) {
	validate.RegisterValidation("pattern", base64Regex)
	validate.RegisterValidation("multipleOf", multipleOf)
	validate.RegisterValidation("enum", base64JsonEnum)
}
