package god

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	SignatureRegex      = `^([a-zA-Z_][a-zA-Z0-9_]*)\((.*)\)$`
	SignaturePartsCount = 3
)

type abiInput struct {
	Indexed      bool        `json:"indexed"`
	InternalType string      `json:"internalType"`
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Components   []*abiInput `json:"components,omitempty"`
}
type abiEvent struct {
	Anonymous bool        `json:"anonymous"`
	Inputs    []*abiInput `json:"inputs"`
	Name      string      `json:"name"`
	Type      string      `json:"type"`
}

func splitArguments(argsString string) ([]string, error) {
	var args []string

	var currentArg strings.Builder

	level := 0

	for i, r := range argsString {
		char := string(r)

		if char == "," && level == 0 {
			trimmedArg := strings.TrimSpace(currentArg.String())
			if trimmedArg == "" {
				return nil, fmt.Errorf("argument list contains an empty element: '%s' (position %d)", argsString, i)
			}

			args = append(args, trimmedArg)

			currentArg.Reset()
		} else {
			currentArg.WriteRune(r)

			if char == "(" {
				level++
			} else if char == ")" {
				if level == 0 {
					return nil, fmt.Errorf("unmatched closing parenthesis: '%s' (position %d)", argsString, i)
				}

				level--
			}
		}
	}

	trimmedArg := strings.TrimSpace(currentArg.String())
	if len(args) == 0 && trimmedArg == "" {
		if len(strings.TrimSpace(argsString)) == 0 {
			return []string{}, nil
		}
	}

	if trimmedArg != "" {
		args = append(args, trimmedArg)
	} else if len(argsString) > 0 && strings.HasSuffix(strings.TrimSpace(argsString), ",") {
		return nil, fmt.Errorf("argument list ends with a comma: '%s'", argsString)
	}

	if level != 0 {
		return nil, fmt.Errorf("unmatched opening parenthesis: '%s'", argsString)
	}

	return args, nil
}

func parseArgType(argTypeString string, index int) (*abiInput, error) {
	trimmedType := strings.TrimSpace(argTypeString)
	name := fmt.Sprintf("arg%d", index)

	if strings.HasPrefix(trimmedType, "(") && strings.HasSuffix(trimmedType, ")") {
		innerArgsString := trimmedType[1 : len(trimmedType)-1]

		innerArgTypes, err := splitArguments(innerArgsString)
		if err != nil {
			return nil, fmt.Errorf("tuple '%s' parse error: %w", trimmedType, err)
		}

		var components []*abiInput

		for i, innerArgType := range innerArgTypes {
			component, err := parseArgType(innerArgType, i)
			if err != nil {
				return nil, fmt.Errorf("tuple '%s' component parse error: %w", trimmedType, err)
			}

			components = append(components, component)
		}

		return &abiInput{Indexed: false, InternalType: "tuple", Name: name, Type: "tuple", Components: components}, nil
	} else {
		return &abiInput{Indexed: false, InternalType: trimmedType, Name: name, Type: trimmedType, Components: nil}, nil
	}
}

func generateABIJSONFromSignature(signature string) (string, string, error) {
	re := regexp.MustCompile(SignatureRegex)
	matches := re.FindStringSubmatch(signature)

	if len(matches) != SignaturePartsCount {
		return "", "", fmt.Errorf("invalid signature format: '%s'", signature)
	}

	eventName := matches[1]
	argsString := matches[2]

	var inputs []*abiInput

	if len(strings.TrimSpace(argsString)) > 0 {
		argTypeStrings, err := splitArguments(argsString)
		if err != nil {
			return "", "", fmt.Errorf("event '%s' argument list parse error: %w", eventName, err)
		}

		for i, argTypeStr := range argTypeStrings {
			input, err := parseArgType(argTypeStr, i)
			if err != nil {
				return "", "", fmt.Errorf("event '%s' argument parse error: %w", eventName, err)
			}

			inputs = append(inputs, input)
		}
	}

	event := abiEvent{Anonymous: false, Inputs: inputs, Name: eventName, Type: "event"}
	abiArray := []abiEvent{event}

	abiBytes, err := json.MarshalIndent(abiArray, "", "  ")
	if err != nil {
		return "", "", fmt.Errorf("error marshaling ABI JSON: %w", err)
	}

	return string(abiBytes), eventName, nil
}

func isStruct(x any) bool {
	t := reflect.TypeOf(x)
	if t == nil {
		return false
	}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.PkgPath() == "math/big" {
		return false
	}

	return t.Kind() == reflect.Struct
}

func DecodeSignature(signature string, data []byte) ([]any, error) {
	generatedABIJson, eventName, err := generateABIJSONFromSignature(signature)
	if err != nil {
		return nil, err
	}

	contractAbi, err := abi.JSON(strings.NewReader(generatedABIJson))
	if err != nil {
		return nil, err
	}

	eventDefinition, ok := contractAbi.Events[eventName]
	if !ok {
		return nil, fmt.Errorf("event definition not found: %s", eventName)
	}

	decodedArgsDirect, errDirect := eventDefinition.Inputs.Unpack(data)
	if errDirect != nil {
		return nil, errDirect
	}

	for i, args := range decodedArgsDirect {
		if !isStruct(args) {
			continue
		}

		var nestedArgs []any

		typeOf := reflect.TypeOf(args)
		for j := range typeOf.NumField() {
			val := reflect.ValueOf(args).Field(j).Interface()
			nestedArgs = append(nestedArgs, val)
		}

		decodedArgsDirect[i] = nestedArgs
	}

	return decodedArgsDirect, nil
}
