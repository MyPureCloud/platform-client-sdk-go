package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping
type Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping struct { 
	// ContactColumnName
	ContactColumnName *string `json:"contactColumnName,omitempty"`


	// DataActionField
	DataActionField *string `json:"dataActionField,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping

	

	return json.Marshal(&struct { 
		ContactColumnName *string `json:"contactColumnName,omitempty"`
		
		DataActionField *string `json:"dataActionField,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		ContactColumnName: u.ContactColumnName,
		
		DataActionField: u.DataActionField,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
