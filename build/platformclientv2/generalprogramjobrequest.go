package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Generalprogramjobrequest
type Generalprogramjobrequest struct { 
	// Dialect - The dialect of the topics to link with the general program, dialect format is {language}-{country} where language follows ISO 639-1 standard and country follows ISO 3166-1 alpha 2 standard
	Dialect *string `json:"dialect,omitempty"`


	// Mode - The mode to use for the general program job, default value is Skip
	Mode *string `json:"mode,omitempty"`

}

// String returns a JSON representation of the model
func (o *Generalprogramjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
