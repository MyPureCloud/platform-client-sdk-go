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

func (o *Generalprogramjobrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Generalprogramjobrequest
	
	return json.Marshal(&struct { 
		Dialect *string `json:"dialect,omitempty"`
		
		Mode *string `json:"mode,omitempty"`
		*Alias
	}{ 
		Dialect: o.Dialect,
		
		Mode: o.Mode,
		Alias:    (*Alias)(o),
	})
}

func (o *Generalprogramjobrequest) UnmarshalJSON(b []byte) error {
	var GeneralprogramjobrequestMap map[string]interface{}
	err := json.Unmarshal(b, &GeneralprogramjobrequestMap)
	if err != nil {
		return err
	}
	
	if Dialect, ok := GeneralprogramjobrequestMap["dialect"].(string); ok {
		o.Dialect = &Dialect
	}
	
	if Mode, ok := GeneralprogramjobrequestMap["mode"].(string); ok {
		o.Mode = &Mode
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Generalprogramjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
