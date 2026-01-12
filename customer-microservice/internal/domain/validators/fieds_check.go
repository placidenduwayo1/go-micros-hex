package validators

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"github.com/trng-tr/customer-microservice/internal/domain"
)

func CheckInputFields(fileds map[string]string) error {
	for key, value := range fileds {
		value = strings.TrimSpace(value)
		if value == "" {
			return fmt.Errorf("%w %s", errEmptyFields, key)
		} else if len(value) > 255 {
			return fmt.Errorf("%w %s", errTooLong, key)
		}

	}

	return nil
}

func CheckEmailValid(email string) bool {
	if _, err := mail.ParseAddress(email); err != nil {
		return false
	}
	return true
}

func CheckInputGenda(genda domain.Genda) error {
	switch genda {
	case domain.Female, domain.Male:
		return nil
	default:
		{
			return fmt.Errorf("%w %s", errInvalidGenda, genda)
		}
	}
}

func CheckInputId(id int64) error {
	if id > 0 {
		return nil
	}
	return fmt.Errorf("%w %d", errInvalidId, id)
}

func CheckPhoneValid(phone string) error {
	var regex = regexp.MustCompile(`^\+?[0-9]{8,15}$`)
	if regex.MatchString(phone) {
		return nil
	}

	return fmt.Errorf("%w %s", errInvalidPhone, phone)
}

func GenerateDate() time.Time {
	return time.Now()
}
