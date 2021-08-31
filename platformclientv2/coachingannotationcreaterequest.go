package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingannotationcreaterequest
type Coachingannotationcreaterequest struct { 
	// Text - The text of the annotation.
	Text *string `json:"text,omitempty"`


	// AccessType - Determines the permissions required to view this item.
	AccessType *string `json:"accessType,omitempty"`

}

func (o *Coachingannotationcreaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Coachingannotationcreaterequest
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		AccessType *string `json:"accessType,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		
		AccessType: o.AccessType,
		Alias:    (*Alias)(o),
	})
}

func (o *Coachingannotationcreaterequest) UnmarshalJSON(b []byte) error {
	var CoachingannotationcreaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &CoachingannotationcreaterequestMap)
	if err != nil {
		return err
	}
	
	if Text, ok := CoachingannotationcreaterequestMap["text"].(string); ok {
		o.Text = &Text
	}
	
	if AccessType, ok := CoachingannotationcreaterequestMap["accessType"].(string); ok {
		o.AccessType = &AccessType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Coachingannotationcreaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
