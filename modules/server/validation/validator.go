package validation

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	ut "github.com/templatedop/universal-translator-master"
)

type (
	FieldError struct {
		FailedField string `json:"field"`
		//Tag         string      `json:"tag"`
		//Value       interface{} `json:"value"`
		Message string `json:"message"`
	}

	Error struct {
		msg  string
		errs []FieldError
	}
)

// Need to check translator and add additional functionality

var (
	validate           *validator.Validate
	uni                *ut.UniversalTranslator
	trans              ut.Translator
	validationMessages = map[string]func(string, any) string{
		// "required": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s is required", s)
		// },
		// "required_if": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s is a required field", s)
		// },
		// "required_unless": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s is a required field", s)
		// },
		// "required_with": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s is a required field", s)
		// },
		// "required_with_all": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s is a required field", s)
		// },
		// "required_without": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s is a required field", s)
		// },
		// "required_without_all": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s is a required field", s)
		// },
		// "excluded_if": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s is an excluded field", s)
		// },
		// "excluded_unless": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s is an excluded field", s)
		// },
		// "excluded_with": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s is an excluded field", s)
		// },
		// "excluded_with_all": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s is an excluded field", s)
		// },
		// "excluded_without": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s is an excluded field", s)
		// },
		// "excluded_without_all": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s is an excluded field", s)
		// },
		// "isdefault": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be default value", s)
		// },

		// "eq": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must be equal to %v", s, v)
		// },
		// "ne": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must not be equal to %v", s, v)
		// },
		// "gt": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must be greater than %v", s, v)
		// },
		// "gte": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must be greater than or equal to %v", s, v)
		// },
		// "lt": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must be less than %v", s, v)
		// },
		// "lte": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must be less than or equal to %v", s, v)
		// },
		// "len": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must be %v characters long", s, v)
		// },
		// "min": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must be at least %v", s, v)
		// },
		// "max": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must be at most %v", s, v)
		// },
		// "oneof": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must be one of %v", s, v)
		// },
		// "unique": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be unique", s)
		// },
		// "email": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid email address", s)
		// },
		// "uuid": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid UUID", s)
		// },
		// "alpha": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must contain only letters", s)
		// },
		// "alphanum": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must contain only letters and numbers", s)
		// },
		// "numeric": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must contain only numbers", s)
		// },
		// "hexadecimal": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid hexadecimal", s)
		// },
		// "hexcolor": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid hex color", s)
		// },
		// "rgb": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid RGB color", s)
		// },
		// "rgba": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid RGBA color", s)
		// },
		// "hsl": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid HSL color", s)
		// },
		// "hsla": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid HSLA color", s)
		// },
		// "ipv4": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid IPv4 address", s)
		// },
		// "ipv6": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid IPv6 address", s)
		// },
		// "ip": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid IP address", s)
		// },
		// "cidr": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid CIDR address", s)
		// },
		// "cidrv4": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid CIDRv4 address", s)
		// },
		// "cidrv6": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid CIDRv6 address", s)
		// },
		// "tcp_addr": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid TCP address", s)
		// },
		// "udp_addr": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid UDP address", s)
		// },
		// "ip_addr": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid IP address", s)
		// },
		// "unix_addr": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid Unix address", s)
		// },
		// "mac": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid MAC address", s)
		// },
		// "latitude": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid latitude", s)
		// },
		// "longitude": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid longitude", s)
		// },
		// "ssn": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid SSN", s)
		// },
		// "isbn": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid ISBN", s)
		// },
		// "isbn10": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid ISBN10", s)
		// },
		// "isbn13": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid ISBN13", s)
		// },
		// "uuid3": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid UUID3", s)
		// },
		// "uuid4": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid UUID4", s)
		// },
		// "uuid5": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid UUID5", s)
		// },
		// "ascii": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid ASCII", s)
		// },
		// "printableascii": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid printable ASCII", s)
		// },
		// "multibyte": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid multibyte", s)
		// },
		// "datauri": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid data URI", s)
		// },
		// "base64": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid base64", s)
		// },
		// "filepath": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid file path", s)
		// },
		// "uri": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid URI", s)
		// },
		// "uri_rfc3986": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid URI RFC3986", s)
		// },
		// "uri_rfc3987": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid URI RFC3987", s)
		// },
		// "uri_rfc2396": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid URI RFC2396", s)
		// },
		// "uri_rfc2732": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid URI RFC2732", s)
		// },
		// "uri_rfc3986_host": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid URI RFC3986 host", s)
		// },
		// "uri_rfc3986_path": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid URI RFC3986 path", s)
		// },
		// "uri_rfc3986_strict": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid URI RFC3986 strict", s)
		// },
		// "lowercase": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be lowercase", s)
		// },
		// "uppercase": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be uppercase", s)
		// },
		// "datetime": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid datetime", s)
		// },
		// "rfc3339": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid RFC3339", s)
		// },
		// "rfc3339_without_zone": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid RFC3339 without zone", s)
		// },
		// "rfc3339_with_zone": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid RFC3339 with zone", s)
		// },
		// "rfc1123": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid RFC1123", s)
		// },
		// "unix": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid Unix", s)
		// },
		// "alphaunicode": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must contain only letters", s)
		// },
		// "alphanumunicode": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must contain only letters and numbers", s)
		// },
		// "numericunicode": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must contain only numbers", s)
		// },
		// "utfdigit": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must contain only digits", s)
		// },
		// "utfletter": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must contain only letters", s)
		// },
		// "utfletternumeric": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must contain only letters and numbers", s)
		// },
		// "utfnumeric": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must contain only numbers", s)
		// },
		// "utfdigitunicode": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must contain only digits", s)
		// },
		// "utfletterunicode": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must contain only letters", s)
		// },
		// "utfletternumericunicode": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must contain only letters and numbers", s)
		// },
		// "utfnumericunicode": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must contain only numbers", s)
		// },
		// "contains": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must contain %v", s, v)
		// },
		// "containsany": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must contain any of %v", s, v)
		// },
		// "excludes": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must not contain %v", s, v)
		// },
		// "excludesall": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must not contain any of %v", s, v)
		// },
		// "excludesrune": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must not contain %v", s, v)
		// },
		// "excludesallrune": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must not contain any of %v", s, v)
		// },
		// "excludesunicode": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must not contain %v", s, v)
		// },
		// "excludesallunicode": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must not contain any of %v", s, v)
		// },
		// "excludesruneunicode": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must not contain %v", s, v)
		// },
		// "excludesallruneunicode": func(s string, v any) string {
		// 	return fmt.Sprintf("field %s must not contain any of %v", s, v)
		// },
		// "boolean": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid boolean", s)
		// },
		// "number": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid number", s)
		// },
		// "image": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid image", s)
		// },
		// "cve": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid CVE", s)
		// },
		// "postcode_iso3166_alpha2": func(s string, _ any) string {
		// 	return fmt.Sprintf("Does not match the required pattern for field %s", s)
		// },
		// "postcode_iso3166_alpha2_field": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid postcode ISO3166 alpha2", s)
		// },

		// "base64url": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid base64 URL", s)
		// },
		// "hostname": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid hostname", s)
		// },
		// "json": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid JSON", s)
		// },
		// "e164": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid E164", s)
		// },

		// "nefield": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must not be equal to the field", s)
		// },
		// "gtefield": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be greater than or equal to the field", s)
		// },
		// "printascii": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid printable ASCII", s)
		// },
		// "ltecsfield": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be less than or equal to the field", s)
		// },
		// "ltefield": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be less than or equal to the field", s)
		// },
		// "ip6_addr": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid IP6 address", s)
		// },
		// "tcp4_addr": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid TCP4 address", s)
		// },
		// "gtecsfield": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be greater than or equal to the field", s)
		// },
		// "ltcsfield": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be less than the field", s)
		// },
		// "ip4_addr": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid IP4 address", s)
		// },
		// "eqfield": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be equal to the field", s)
		// },
		// "necsfield": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must not be equal to the field", s)
		// },
		// "gtcsfield": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be greater than the field", s)
		// },
		// "ulid": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid ULID", s)
		// },
		// "iscolor": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid color", s)
		// },
		// "fqdn": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid FQDN", s)
		// },
		// "jwt": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid JWT", s)
		// },
		// "udp6_addr": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid UDP6 address", s)
		// },
		// "issn": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid ISSN", s)
		// },
		// "gtfield": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be greater than the field", s)
		// },
		// "url": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid URL", s)
		// },
		// "cron": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid cron", s)
		// },
		// "udp4_addr": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid UDP4 address", s)
		// },
		// "eqcsfield": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be equal to the field", s)
		// },
		// "ltfield": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be less than the field", s)
		// },
		// "tcp6_addr": func(s string, _ any) string {
		// 	return fmt.Sprintf("field %s must be a valid TCP6 address", s)
		// },
	}
)

var structFieldTags = []string{"json", "param", "query"}

func getStructFieldName(fld reflect.StructField) string {
	for _, st := range structFieldTags {
		name := strings.SplitN(fld.Tag.Get(st), ",", 2)[0]

		if name == "" {
			continue
		}

		if name == "-" {
			return ""
		}

		return name
	}

	return fld.Name
}

func Init(rules []Rule) {

	validate = validator.New()
	eng := en.New()
	uni = ut.New(eng, eng)
	trans, _ = uni.GetTranslator("en")

	// Register default translations for the validator
	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		panic(fmt.Sprintf("Failed to register translations: %v", err))
		//log.Fatalf("Failed to register translations: %v", err)
	}
	validate.RegisterTagNameFunc(getStructFieldName)

	for _, r := range rules {
		if err := validate.RegisterValidation(r.Name(), r.Apply); err != nil {
			panic(err.Error())
		}
		validationMessages[r.Name()] = r.Message
	}
}

func ValidateStruct(s interface{}) error {

	//fmt.Println("calling validate struct")
	if validate == nil {
		panic("validator not initialized")
	}

	if trans == nil {
		panic("translator not initialized")
	}
	//var fieldErrors validation.FieldErrors
	err := validate.Struct(s)
	if err != nil {
		var fieldErrors []FieldError
		var validatorErrors validator.ValidationErrors

		errors.As(err, &validatorErrors)

		// for _, err := range validatorErrors {
		// 	fieldErrors = append(fieldErrors, FieldError{
		// 		FailedField: err.Field(),
		// 		Tag:         err.Tag(),
		// 		Value:       err.Value(),
		// 		Message:     getTagMessage(err),
		// 	})
		// }

		for _, e := range validatorErrors {

			tag := e.Tag()
			if _, ok := validationMessages[tag]; ok {
				fmt.Println("coming inside validation messages")
				fieldErrors = append(fieldErrors, FieldError{
					FailedField: e.Field(),
					//Tag:         e.Tag(),
					//Value:       e.Value(),
					Message: getTagMessage(e),
				})

			} else {
				fieldErrors = append(fieldErrors, FieldError{
					FailedField: e.Field(),
					//Tag:         e.Tag(),
					//Value:       e.Value(),
					Message: e.Translate(trans),
				})

			}

		}

		return &Error{
			msg:  "validation error",
			errs: fieldErrors,
		}

		//err := response.Error("validation error", fieldErrors...)

		//return err

	}
	return nil
}

// func Validate(s any) error {
// 	if validate == nil {
// 		panic("validator not initialized")
// 	}

// 	errs := validate.Struct(s)
// 	if errs != nil {
// 		var fieldErrors []FieldError

// 		var validatorErrors validator.ValidationErrors
// 		errors.As(errs, &validatorErrors)

// 		for _, err := range validatorErrors {
// 			fieldErrors = append(fieldErrors, FieldError{
// 				FailedField: err.Field(),
// 				Tag:         err.Tag(),
// 				Value:       err.Value(),
// 				Message:     getTagMessage(err),
// 			})
// 		}

// 		return &Error{
// 			msg:  "validation error",
// 			errs: fieldErrors,
// 		}
// 	}

// 	return nil
// }

func (ve *Error) Unwrap() error {
	return fiber.ErrBadRequest
}

func (ve *Error) Error() string {
	return ve.msg
}

func (ve *Error) FieldErrors() []FieldError {
	return ve.errs
}

func (fe FieldError) Error() string {
	return fe.Message
}

func getTagMessage(err validator.FieldError) string {
	if mr, ok := validationMessages[err.Tag()]; ok {
		return mr(err.Field(), err.Value())
	}

	return err.Error()
}
