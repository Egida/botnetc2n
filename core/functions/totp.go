package functions

import (
	"crypto/rand"
	"encoding/base32"
	"math/big"
	"strings"
)

// generates our totp code for our system
// this will be used within the process properly
func GenerateSecret() (string, error) { // returns string process

	// creates our custom secret before
	// before we encode properly happens
	secretpre, err := crypto(32, "") //generates
	if err != nil { //err handles
		return "", err 
	}

	// returns our information properly and safely without issues
	return strings.ReplaceAll(base32.StdEncoding.EncodeToString([]byte(secretpre)), "=", "D"), nil
}

// generates our crypto TOTP string properly
// used within the system in the future properly
func crypto(l int, s string) (string, error) {

	// ranges through the amount properly and safely
	// this will ensure its done without errors properly
	for  {

		//if length properly and safely
		//makes sure its only done properly
		if len(s) >= l { //checks len
			break //breaks from loop
		}

		// random number generater properly enforced
		number, err := rand.Int(rand.Reader, big.NewInt(int64(127)))
		if err != nil { //err handles
			return "", err
		}


		// pure 64 int structure
		pure := number.Int64()

		// checks within bounds properly
		// this will ensure its done properly
		if pure > 32 && pure < 127 { // within guidelines
			s += string(pure) // saves
		}
	}

	// returns information
	// majors used properly
	return s, nil
}