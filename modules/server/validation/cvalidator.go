package validation

import (
	"github.com/templatedop/api/module"
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
)

var (
	HOAPattern                                = regexp.MustCompile(`^\d{15}$`)
	PersonnelNamePattern                      = regexp.MustCompile(`^[A-Za-z][A-Za-z\s]{1,48}[A-Za-z]$`)
	AddressPattern                            = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9\s,.-]{1,48}[A-Za-z0-9]$`)
	EmailPattern                              = regexp.MustCompile(`^[a-zA-Z0-9._+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	gValidatePhoneLengthPattern               = regexp.MustCompile(`^\d{10}$`)
	allZerosRegex                             = regexp.MustCompile("^0+$")
	gValidateSOBONamePattern                  = regexp.MustCompile(`^[A-Za-z][A-Za-z\s]{1,48}[A-Za-z]$`)
	gValidatePANNumberPattern                 = regexp.MustCompile(`^[A-Z]{5}[0-9]{4}[A-Z]$`)
	gValidateVehicleRegistrationNumberPattern = regexp.MustCompile(`^[A-Z]{2}\d{2}[A-Z]{1,2}\d{4,7}$ |\d{2}[A-Z]{2}\d{4}[A-Z]{2}$`)
	gValidateBarCodeNumberPattern             = regexp.MustCompile(`^[A-Z]{2}\d{6,12}[A-Z]{2}$`)
	alphanumericRegex                         = regexp.MustCompile(`^[A-Z0-9]+$`)
	trainNoPattern                            = regexp.MustCompile(`^\d{5}$`)
	customValidateGLCodePattern               = regexp.MustCompile(`^GL\d{11}$`)
	timeStampValidatePattern                  = regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-(\d{4}) ([01]\d|2[0-3]):([0-5]\d):([0-5]\d)$`)
	customValidateAnyStringLengthto50Pattern  = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]{0,48}[a-zA-Z]$`)
	dateyyyymmddPattern                       = regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`)
	dateddmmyyyyPattern                       = regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-\d{4}$`)
	validateEmployeeIDPattern                 = regexp.MustCompile(`^\d{8}$`)
	validateGSTINPattern                      = regexp.MustCompile(`^[0-9]{2}[A-Z]{5}[0-9]{4}[A-Z]{1}[A-Z0-9]{1}[Z]{1}[0-9]{1}$`)
	specialCharPattern                        = regexp.MustCompile(`[!@#$%^&*()<>:;"{}[\]\\]`)
	validateBankUserIDPattern                 = regexp.MustCompile(`^[A-Z0-9]{1,50}$`)
	validateOrderNumberPattern                = regexp.MustCompile(`^[A-Z]{2}\d{19}$`)
	validateAWBNumberPattern                  = regexp.MustCompile(`^[A-Z]{4}\d{9}$`)
	validatePNRNoPattern                      = regexp.MustCompile(`^[A-Z]{3}\d{6}$`)
	validatePLIIDPattern                      = regexp.MustCompile(`^[A-Z]{3}\d{10}$`)
	validatePaymentTransIDPattern             = regexp.MustCompile(`^\d{2}[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	validateOfficeCustomerIDPattern           = regexp.MustCompile(`^[a-zA-Z0-9\-]{1,50}$`)
	validateBankIDPattern                     = regexp.MustCompile(`^[A-Z0-9]{1,50}$`)
	validateCSIFacilityIDPattern              = regexp.MustCompile(`^[A-Z]{2}\d{11}$`)
	validatePosBookingOrderNumberPattern      = regexp.MustCompile(`^[A-Z]{2}\d{19}$`)
	validateSOLIDPattern                      = regexp.MustCompile(`^\d{6}\d{2}$`)
	validatePLIOfficeIDPattern                = regexp.MustCompile(`^[A-Z]{3}\d{10}$`)
	validateProductCodePattern                = regexp.MustCompile(`^[A-Z]{3}\d{12}$`)
	validateCustomerIDPattern                 = regexp.MustCompile(`^\d{10}$`)
	validateFacilityIDPattern                 = regexp.MustCompile(`^[A-Z]{2}\d{11}$`)
	validateApplicationIDPattern              = regexp.MustCompile(`^[A-Z]{3}\d{8}-\d{3}$`)
	validateReceiverKYCReferencePattern       = regexp.MustCompile(`^KYCREF[A-Z0-9]{0,44}$`)
	validateOfficeCustomerPattern             = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	validatePRANPattern                       = regexp.MustCompile(`^\d{12}$`)
)

func NewValidateHOAPatternValidator() Rule {
	return NewRule("customHeadOfAccount", ValidateHOAPattern, "the %s ")
}
func NewPersonnelNameValidator() Rule {
	return NewRule("customPersonnelName", ValidatePersonnelNamePattern, "the %s ")
}

func NewEmailValidator() Rule {
	return NewRule("emailValidator", ValidateEmailPattern, "the %s ")
}
func NewAddressPatternValidator() Rule {
	return NewRule("customAddressPattern", ValidateAddressPattern, "the %s ")
}
func NewGValidatePhoneLengthPatternValidator() Rule {
	return NewRule("customValidatePhoneLengthPattern", GValidatePhoneLengthPattern, "the %s ")
}
func NewGValidateSOBONamePatternValidator() Rule {
	return NewRule("gValidateSOBONamePattern", GValidateSOBONamePattern, "the %s ")
}
func NewGValidatePANNumberPatternValidator() Rule {
	return NewRule("gValidatePANNumberPattern", GValidatePANNumberPattern, "the %s ")
}
func NewGValidateVehicleRegistrationNumberPatternValidator() Rule {
	return NewRule("gValidateVehicleRegistrationNumberPattern", GValidateVehicleRegistrationNumberPattern, "the %s ")
}
func NewGValidateBarCodeNumberPatternValidator() Rule {
	return NewRule("gValidateBarCodeNumberPattern", GValidateBarCodeNumberPattern, "the %s ")
}

func NewCustomValidateGLCodePatternValidator() Rule {
	return NewRule("customValidateGLCodePattern", CustomValidateGLCodePattern, "the %s ")
}
func NewTimeStampValidatePatternValidator() Rule {
	return NewRule("timeStampValidatePattern", TimeStampValidatePattern, "the %s ")
}
func NewCustomValidateAnyStringLengthto50PatternValidator() Rule {
	return NewRule("customValidateAnyStringLengthto50Pattern", CustomValidateAnyStringLengthto50Pattern, "the %s ")
}
func NewDateyyyymmddPatternValidator() Rule {
	return NewRule("dateyyyymmddPattern", DateyyyymmddPattern, "the %s ")
}

func NewDateddmmyyyyPatternValidator() Rule {
	return NewRule("dateddmmyyyyPattern", DateddmmyyyyPattern, "the %s ")
}

func NewValidateEmployeeIDPatternValidator() Rule {
	return NewRule("validateEmployeeIDPattern", ValidateEmployeeIDPattern, "the %s ")
}
func NewValidateValidateGSTINPatternValidator() Rule {
	return NewRule("validateGSTINPattern", ValidateGSTINPattern, "the %s ")
}

// ********************************
func NewValidatePRANPatternValidator() Rule {
	return NewRule("validatePRANPattern", ValidatePRANPattern, "the %s ")
}

func NewValidateOfficeCustomerPatternValidator() Rule {
	return NewRule("validateOfficeCustomerPattern", ValidateOfficeCustomerPattern, "the %s ")
}

func NewValidateReceiverKYCReferencePatternValidator() Rule {
	return NewRule("validateReceiverKYCReferencePattern", ValidateReceiverKYCReferencePattern, "the %s ")
}

func NewValidateApplicationIDPatternValidator() Rule {
	return NewRule("validateApplicationIDPattern", ValidateApplicationIDPattern, "the %s ")
}

func NewValidateFacilityIDPatternValidator() Rule {
	return NewRule("validateFacilityIDPattern", ValidateFacilityIDPattern, "the %s ")
}

func NewValidateCustomerIDPatternValidator() Rule {
	return NewRule("validateCustomerIDPattern", ValidateCustomerIDPattern, "the %s ")
}

func NewValidateProductCodePatternValidator() Rule {
	return NewRule("validateProductCodePattern", ValidateProductCodePattern, "the %s ")
}

func NewValidatePLIOfficeIDPatternValidator() Rule {
	return NewRule("validatePLIOfficeIDPattern", ValidatePLIOfficeIDPattern, "the %s ")
}

func NewValidateSOLIDPatternValidator() Rule {
	return NewRule("validateSOLIDPattern", ValidateSOLIDPattern, "the %s ")
}

func NewValidatePosBookingOrderNumberPatternValidator() Rule {
	return NewRule("validatePosBookingOrderNumberPattern", ValidatePosBookingOrderNumberPattern, "the %s ")
}

func NewValidateCSIFacilityIDPatternValidator() Rule {
	return NewRule("validateCSIFacilityIDPattern", ValidateCSIFacilityIDPattern, "the %s ")
}

func NewValidateBankIDPatternValidator() Rule {
	return NewRule("validateBankIDPattern", ValidateBankIDPattern, "the %s ")
}

func NewValidateOfficeCustomerIDPatternValidator() Rule {
	return NewRule("validateOfficeCustomerIDPattern", ValidateOfficeCustomerIDPattern, "the %s ")
}

func NewValidatePaymentTransIDPatternValidator() Rule {
	return NewRule("validatePaymentTransIDPattern", ValidatePaymentTransIDPattern, "the %s ")
}

func NewValidatePLIIDPatternValidator() Rule {
	return NewRule("validatePLIIDPattern", ValidatePLIIDPattern, "the %s ")
}

func NewValidatePNRNoPatternValidator() Rule {
	return NewRule("validatePNRNoPattern", ValidatePNRNoPattern, "the %s ")
}

func NewValidateAWBNumberPatternValidator() Rule {
	return NewRule("validateAWBNumberPattern", ValidateAWBNumberPattern, "the %s ")
}

func NewValidateOrderNumberPatternValidator() Rule {
	return NewRule("validateOrderNumberPattern", ValidateOrderNumberPattern, "the %s ")
}

func NewValidateBankUserIDPatternValidator() Rule {
	return NewRule("validateBankUserIDPattern", ValidateBankUserIDPattern, "the %s ")
}

func NewvalidatePinCodeGlobalValidator() Rule {
	return NewRule("customPincode", validatePinCodeGlobal, "the %s ")
}

// ///////////////////////////////////////////
func ValidateWithGlobalRegex(fl validator.FieldLevel, regex *regexp.Regexp) bool {
	fieldValue := fl.Field().String()
	return regex.MatchString(fieldValue)
}

// ////////////////////////////////////////////////////////////
func ValidateHOAPattern(fl validator.FieldLevel) bool {
	//pattern := `^\d{15}$`
	return ValidateWithGlobalRegex(fl, HOAPattern)
}
func ValidatePersonnelNamePattern(fl validator.FieldLevel) bool {
	return ValidateWithGlobalRegex(fl, PersonnelNamePattern)
}
func ValidateAddressPattern(fl validator.FieldLevel) bool {
	return ValidateWithGlobalRegex(fl, AddressPattern)
}
func ValidateEmailPattern(fl validator.FieldLevel) bool {
	return ValidateWithGlobalRegex(fl, EmailPattern)
}
func GValidatePhoneLengthPattern(fl validator.FieldLevel) bool {
	// Handle the case where the phone number is a string
	if _, ok := fl.Field().Interface().(string); ok {
		// Validate using a regular expression for exactly 10 digits
		// pattern := `^\d{10}$`
		// return ValidateWithRegex(fl, pattern)
		return ValidateWithGlobalRegex(fl, gValidatePhoneLengthPattern)
	}

	// Handle the case where the phone number is a uint64
	if phoneNumber, ok := fl.Field().Interface().(uint64); ok {
		// Check if the phone number has exactly 10 digits
		return phoneNumber >= 1000000000 && phoneNumber <= 9999999999
	}
	//works only for 64 bit system
	// Handle the case where the phone number is an int
	if phoneNumber, ok := fl.Field().Interface().(int); ok {
		// Check if the phone number has exactly 10 digits
		return phoneNumber >= 1000000000 && phoneNumber <= 9999999999
	}

	// If the field is neither a string, uint64, nor int, the validation fails
	return false
}
func GValidateSOBONamePattern(f1 validator.FieldLevel) bool {
	// Define the regex pattern
	// ^[A-Za-z] -> Start with a letter
	// [A-Za-z\s]{1,48} -> 1 to 48 letters or spaces
	// [A-Za-z]$ -> End with a letter
	//pattern := `^[A-Za-z][A-Za-z\s]{1,48}[A-Za-z]$`

	return ValidateWithGlobalRegex(f1, gValidateSOBONamePattern)
}
func GValidatePANNumberPattern(fl validator.FieldLevel) bool {
	// regex pattern for PAN number (5 letters followed by 4 digits followed by 1 letter)
	//pattern := `^[A-Z]{5}[0-9]{4}[A-Z]$`

	return ValidateWithGlobalRegex(fl, gValidatePANNumberPattern)
}
func GValidateVehicleRegistrationNumberPattern(fl validator.FieldLevel) bool {
	// Define the regex pattern for vehicle registration number
	//pattern := `^[A-Z]{2}\d{2}[A-Z]{1,2}\d{4,7}$`
	return ValidateWithGlobalRegex(fl, gValidateVehicleRegistrationNumberPattern)
}
func GValidateBarCodeNumberPattern(fl validator.FieldLevel) bool {

	// Define the regex pattern for vehicle registration number
	//pattern := `^[A-Z]{2}\d{6,12}[A-Z]{2}$`
	return ValidateWithGlobalRegex(fl, gValidateBarCodeNumberPattern)
}
func CustomValidateGLCodePattern(fl validator.FieldLevel) bool {
	//pattern := `^GL\d{11}$`
	return ValidateWithGlobalRegex(fl, customValidateGLCodePattern)
}
func TimeStampValidatePattern(f1 validator.FieldLevel) bool {
	//dateTimeRegex := regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-(\d{4}) ([01]\d|2[0-3]):([0-5]\d):([0-5]\d)$`)
	return ValidateWithGlobalRegex(f1, timeStampValidatePattern)
}
func CustomValidateAnyStringLengthto50Pattern(fl validator.FieldLevel) bool {
	//pattern := `^[a-zA-Z][a-zA-Z0-9]{0,48}[a-zA-Z]$`
	// Check if the string matches the regex pattern
	return ValidateWithGlobalRegex(fl, customValidateAnyStringLengthto50Pattern)
}
func DateyyyymmddPattern(fl validator.FieldLevel) bool {
	//pattern := `^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`
	// Check if the date matches the regex pattern
	return ValidateWithGlobalRegex(fl, dateyyyymmddPattern)

}
func DateddmmyyyyPattern(fl validator.FieldLevel) bool {
	//pattern := `^(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-\d{4}$`

	// Check if the date matches the regex pattern
	return ValidateWithGlobalRegex(fl, dateddmmyyyyPattern)

}
func ValidateEmployeeIDPattern(fl validator.FieldLevel) bool {
	//pattern := `^\d{8}$`
	return ValidateWithGlobalRegex(fl, validateEmployeeIDPattern)
}
func ValidateGSTINPattern(fl validator.FieldLevel) bool {

	// Define the regex pattern for GSTIN validation
	//pattern := `^[0-9]{2}[A-Z]{5}[0-9]{4}[A-Z]{1}[A-Z0-9]{1}[Z]{1}[0-9]{1}$`

	return ValidateWithGlobalRegex(fl, validateGSTINPattern)
}

//***********************************************

func ValidatePRANPattern(fl validator.FieldLevel) bool {
	// Handle the case where the PRAN is a string
	if _, ok := fl.Field().Interface().(string); ok {
		// Define a regex pattern to match exactly 12 digits
		//pattern := `^\d{12}$`

		return ValidateWithGlobalRegex(fl, validatePRANPattern)
	}

	// Handle the case where the PRAN is an int64
	if pranInt, ok := fl.Field().Interface().(int64); ok {
		// Check if the int64 falls within the 12-digit range
		return pranInt >= 100000000000 && pranInt <= 999999999999
	}

	// If the field is neither a valid string nor a valid integer, the validation fails
	return false
}

func ValidateOfficeCustomerPattern(fl validator.FieldLevel) bool {
	// Regular expression to allow only alphanumeric characters and spaces
	// This will disallow special characters like @, #, $, %, etc.
	//pattern := `^[a-zA-Z0-9\s]+$`

	// Get the field value and convert it to a string
	officeCustomer, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	// Check if the length of the string is within 50 characters
	if len(officeCustomer) > 50 {
		return false
	}

	// Check if the office_customer string matches the allowed pattern
	return ValidateWithGlobalRegex(fl, validateOfficeCustomerPattern)
}

func ValidateReceiverKYCReferencePattern(fl validator.FieldLevel) bool {

	// Define a regex pattern to match the format KYCREF followed by up to 44 alphanumeric characters
	//pattern := `^KYCREF[A-Z0-9]{0,44}$`
	// Check if the string matches the pattern
	return ValidateWithGlobalRegex(fl, validateReceiverKYCReferencePattern)
}

func ValidateApplicationIDPattern(fl validator.FieldLevel) bool {
	// Define a regex pattern to match the format <3 uppercase letters><12 digits with hyphen>
	//pattern := `^[A-Z]{3}\d{8}-\d{3}$`
	// Check if the string matches the pattern
	return ValidateWithGlobalRegex(fl, validateApplicationIDPattern)
}

func ValidateFacilityIDPattern(fl validator.FieldLevel) bool {

	// Define a regex pattern to match the format <2 uppercase letters><11 digits>
	//pattern := `^[A-Z]{2}\d{11}$`
	// Check if the string matches the pattern
	return ValidateWithGlobalRegex(fl, validateFacilityIDPattern)
}

func ValidateCustomerIDPattern(fl validator.FieldLevel) bool {
	// Handle the case where the value is a string
	if _, ok := fl.Field().Interface().(string); ok {
		// Define a regex pattern to match a 10-digit number
		//pattern := `^\d{10}$`
		// Check if the string matches the pattern
		return ValidateWithGlobalRegex(fl, validateCustomerIDPattern)
	}

	// Handle the case where the value is an integer
	if customerIDInt, ok := fl.Field().Interface().(int); ok {
		// Convert the integer to a string
		customerIDStr := strconv.Itoa(customerIDInt)

		// Check if the integer has exactly 10 digits
		return len(customerIDStr) == 10
	}

	// If the field is neither a string nor an integer, validation fails
	return false
}

func ValidateProductCodePattern(fl validator.FieldLevel) bool {
	// Assume the fl value is always a string

	// Define a regex pattern to match the format <3 uppercase letters><12 digits>
	//pattern := `^[A-Z]{3}\d{12}$`

	// Check if the string matches the pattern
	return ValidateWithGlobalRegex(fl, validateProductCodePattern)
}

func ValidatePLIOfficeIDPattern(fl validator.FieldLevel) bool {
	// Assume the fl value is always a string

	// Define a regex pattern to match the format <3 uppercase letters><10 digits>
	//pattern := `^[A-Z]{3}\d{10}$`
	// Check if the string matches the pattern
	return ValidateWithGlobalRegex(fl, validatePLIOfficeIDPattern)
}

func ValidateSOLIDPattern(fl validator.FieldLevel) bool {
	// Assume the fl value is always a string

	// Define a regex pattern to match the format <6 digits pincode><2 digits office type number>
	//pattern := `^\d{6}\d{2}$`
	// Check if the string matches the pattern
	return ValidateWithGlobalRegex(fl, validateSOLIDPattern)
}

func ValidatePosBookingOrderNumberPattern(fl validator.FieldLevel) bool {
	// Assume the fl value is always a string

	// Define a regex pattern to match the format <2 uppercase letters><19 digits>
	//pattern := `^[A-Z]{2}\d{19}$`
	// Check if the string matches the pattern
	return ValidateWithGlobalRegex(fl, validatePosBookingOrderNumberPattern)
}

func ValidateCSIFacilityIDPattern(fl validator.FieldLevel) bool {
	// Handle the case where the csi_facility_id is a string
	if _, ok := fl.Field().Interface().(string); ok {
		// Define a regex pattern that matches the format <2 uppercase letters><11 digit numeric>
		//pattern := `^[A-Z]{2}\d{11}$`
		// Check if the csi_facility_id matches the pattern
		return ValidateWithGlobalRegex(fl, validateCSIFacilityIDPattern)
	}

	// If the field is not a string, the validation fails
	return false
}

func ValidateBankIDPattern(fl validator.FieldLevel) bool {
	// Handle the case where the value is a string
	if _, ok := fl.Field().Interface().(string); ok {
		// Define a regex pattern to match a string with up to 50 characters consisting of uppercase letters and digits
		//pattern := `^[A-Z0-9]{1,50}$`
		// Check if the string matches the pattern
		return ValidateWithGlobalRegex(fl, validateBankIDPattern)
	}

	// If the field is not a string, validation fails
	return false
}

func ValidateOfficeCustomerIDPattern(fl validator.FieldLevel) bool {
	// Handle the case where the value is a string
	if _, ok := fl.Field().Interface().(string); ok {
		// Define a regex pattern to match any string with up to 50 characters

		//pattern := `^[a-zA-Z0-9\-]{1,50}$`
		// Check if the string matches the pattern
		return ValidateWithGlobalRegex(fl, validateOfficeCustomerIDPattern)
	}

	// If the field is not a string, validation fails
	return false
}

func ValidatePaymentTransIDPattern(fl validator.FieldLevel) bool {
	// Handle the case where the payment_trans_id is a string
	if _, ok := fl.Field().Interface().(string); ok {
		// Define a regex pattern that matches the format <2digit><uuid v4>
		//pattern := `^\d{2}[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`

		// Check if the payment_trans_id matches the pattern
		return ValidateWithGlobalRegex(fl, validatePaymentTransIDPattern)
	}

	// If the field is not a string, the validation fails
	return false
}

func ValidatePLIIDPattern(fl validator.FieldLevel) bool {
	// Handle the case where the pli_id is a string
	if _, ok := fl.Field().Interface().(string); ok {
		// Define a regex pattern that matches the format <3 uppercase letters><10 digit numeric>
		//pattern := `^[A-Z]{3}\d{10}$`
		// Check if the awbnumber matches the pattern
		return ValidateWithGlobalRegex(fl, validatePLIIDPattern)
	}

	// If the field is not a string, the validation fails
	return false
}

func ValidatePNRNoPattern(fl validator.FieldLevel) bool {

	// Handle the case where the pnr_no is a string
	if _, ok := fl.Field().Interface().(string); ok {
		// Define a regex pattern that matches the format
		//pattern := `^[A-Z]{3}\d{6}$`
		// Check if the pnr_no matches the pattern
		return ValidateWithGlobalRegex(fl, validatePNRNoPattern)
	}

	// If the field is not a string, the validation fails
	return false
}

func ValidateAWBNumberPattern(fl validator.FieldLevel) bool {
	// Handle the case where the awbnumber is a string
	if _, ok := fl.Field().Interface().(string); ok {
		// Define a regex pattern that matches the format <4 uppercase letters><9 digit numeric>
		//pattern := `^[A-Z]{4}\d{9}$`

		// Check if the awbnumber matches the pattern
		return ValidateWithGlobalRegex(fl, validateAWBNumberPattern)
	}

	// If the field is not a string, the validation fails
	return false
}

func ValidateOrderNumberPattern(fl validator.FieldLevel) bool {
	// Handle the case where the order_number is a string
	if _, ok := fl.Field().Interface().(string); ok {
		// Define a regex pattern that matches the format <2 uppercase letters><19 digit numeric>
		//pattern := `^[A-Z]{2}\d{19}$`
		// Check if the order_number matches the pattern
		return ValidateWithGlobalRegex(fl, validateOrderNumberPattern)
	}

	// If the field is not a string, the validation fails
	return false
}

func ValidateBankUserIDPattern(fl validator.FieldLevel) bool {
	// Handle the case where the bank_user_id is a string
	if _, ok := fl.Field().Interface().(string); ok {
		// Define a regex pattern that ensures the bank_user_id is alphanumeric and between 1 to 50 characters
		//pattern := `^[A-Z0-9]{1,50}$`

		// Check if the bank_user_id matches the pattern
		return ValidateWithGlobalRegex(fl, validateBankUserIDPattern)
	}

	// If the field is not a string, the validation fails
	return false
}

func ValidateHONamePattern(fl validator.FieldLevel) bool {
	// Handle the case where the ho_name is a string
	if hoName, ok := fl.Field().Interface().(string); ok {
		// Check if the ho_name is not empty and has a maximum length of 50 characters
		if len(hoName) == 0 || len(hoName) > 50 {
			return false
		}

		// Define a regex pattern that disallows special characters @,#/$%!^&*()<>:;"{}[]
		// specialCharPattern := `[!@#$%^&*()<>:;"{}[\]\\]`
		// regex := regexp.MustCompile(specialCharPattern)

		// Check if the ho_name contains any special characters
		if specialCharPattern.MatchString(hoName) {
			return false
		}

		// If all checks pass, return true
		return true
	}

	// If the field is not a string, the validation fails
	return false
}
func validatePinCodeGlobal(fl validator.FieldLevel) bool {
	zipCode := fl.Field().String()

	// Check if the length is 6
	if len(zipCode) != 6 {
		return false
	}
	// Check if the pin code contains only digits
	if _, err := strconv.Atoi(zipCode); err != nil {
		return false
	}

	// Check if the first digit is in the range 1 to 9
	firstDigit, err := strconv.Atoi(string(zipCode[0]))
	if err != nil || firstDigit < 1 || firstDigit > 9 {
		return false
	}

	// Check if the last five digits are not all zeros
	lastFiveDigits := zipCode[1:6]
	//allZerosRegex := regexp.MustCompile("^0+$")
	if allZerosRegex.MatchString(lastFiveDigits) {
		return false
	}
	// Check if the last three digits are not all zeros
	lastThreeDigits := zipCode[3:6]
	if allZerosRegex.MatchString(lastThreeDigits) {
		return false
	}
	return true

}

// //////////////////////////////////////////////////////
var InternalValidatorModule = module.New("internalcustomvalidator").Provide(
	NewValidateHOAPatternValidator,
	NewPersonnelNameValidator,
	NewAddressPatternValidator,
	NewEmailValidator,
	NewGValidatePhoneLengthPatternValidator,
	NewGValidateSOBONamePatternValidator,
	NewGValidatePANNumberPatternValidator,
	NewGValidateVehicleRegistrationNumberPatternValidator,
	NewGValidateBarCodeNumberPatternValidator,
	NewCustomValidateGLCodePatternValidator,
	NewTimeStampValidatePatternValidator,
	NewCustomValidateAnyStringLengthto50PatternValidator,
	NewDateyyyymmddPatternValidator,
	NewDateddmmyyyyPatternValidator,
	NewValidateEmployeeIDPatternValidator,
	NewValidateValidateGSTINPatternValidator,
	//***
	NewValidateOfficeCustomerPatternValidator,
	NewValidatePRANPatternValidator,
	NewValidateReceiverKYCReferencePatternValidator,
	NewValidateApplicationIDPatternValidator,
	NewValidateFacilityIDPatternValidator,
	NewValidateCustomerIDPatternValidator,
	NewValidateProductCodePatternValidator,
	NewValidatePLIOfficeIDPatternValidator,
	NewValidateSOLIDPatternValidator,
	NewValidatePosBookingOrderNumberPatternValidator,
	NewValidateCSIFacilityIDPatternValidator,
	NewValidateBankIDPatternValidator,
	NewValidateOfficeCustomerIDPatternValidator,
	NewValidatePaymentTransIDPatternValidator,
	NewValidatePLIIDPatternValidator,
	NewValidatePNRNoPatternValidator,
	NewValidateAWBNumberPatternValidator,
	NewValidateOrderNumberPatternValidator,
	NewValidateBankUserIDPatternValidator,
	NewvalidatePinCodeGlobalValidator,
)
