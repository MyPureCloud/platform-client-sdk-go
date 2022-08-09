package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Availablemediatype
type Availablemediatype struct { 
	// MediaType - Name of the available media type
	MediaType *string `json:"mediaType,omitempty"`


	// AvailableSubTypes - List of available subtypes for this media type
	AvailableSubTypes *[]string `json:"availableSubTypes,omitempty"`

}

func (o *Availablemediatype) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Availablemediatype
	
	return json.Marshal(&struct { 
		MediaType *string `json:"mediaType,omitempty"`
		
		AvailableSubTypes *[]string `json:"availableSubTypes,omitempty"`
		*Alias
	}{ 
		MediaType: o.MediaType,
		
		AvailableSubTypes: o.AvailableSubTypes,
		Alias:    (*Alias)(o),
	})
}

func (o *Availablemediatype) UnmarshalJSON(b []byte) error {
	var AvailablemediatypeMap map[string]interface{}
	err := json.Unmarshal(b, &AvailablemediatypeMap)
	if err != nil {
		return err
	}
	
	if MediaType, ok := AvailablemediatypeMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if AvailableSubTypes, ok := AvailablemediatypeMap["availableSubTypes"].([]interface{}); ok {
		AvailableSubTypesString, _ := json.Marshal(AvailableSubTypes)
		json.Unmarshal(AvailableSubTypesString, &o.AvailableSubTypes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Availablemediatype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
