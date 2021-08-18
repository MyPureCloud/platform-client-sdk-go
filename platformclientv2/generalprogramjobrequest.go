package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Generalprogramjobrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Generalprogramjobrequest

	

	return json.Marshal(&struct { 
		Dialect *string `json:"dialect,omitempty"`
		
		Mode *string `json:"mode,omitempty"`
		*Alias
	}{ 
		Dialect: u.Dialect,
		
		Mode: u.Mode,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Generalprogramjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
