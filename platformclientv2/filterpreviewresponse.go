package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Filterpreviewresponse
type Filterpreviewresponse struct { 
	// FilteredContacts
	FilteredContacts *int `json:"filteredContacts,omitempty"`


	// TotalContacts
	TotalContacts *int `json:"totalContacts,omitempty"`


	// Preview
	Preview *[]Dialercontact `json:"preview,omitempty"`

}

func (o *Filterpreviewresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Filterpreviewresponse
	
	return json.Marshal(&struct { 
		FilteredContacts *int `json:"filteredContacts,omitempty"`
		
		TotalContacts *int `json:"totalContacts,omitempty"`
		
		Preview *[]Dialercontact `json:"preview,omitempty"`
		*Alias
	}{ 
		FilteredContacts: o.FilteredContacts,
		
		TotalContacts: o.TotalContacts,
		
		Preview: o.Preview,
		Alias:    (*Alias)(o),
	})
}

func (o *Filterpreviewresponse) UnmarshalJSON(b []byte) error {
	var FilterpreviewresponseMap map[string]interface{}
	err := json.Unmarshal(b, &FilterpreviewresponseMap)
	if err != nil {
		return err
	}
	
	if FilteredContacts, ok := FilterpreviewresponseMap["filteredContacts"].(float64); ok {
		FilteredContactsInt := int(FilteredContacts)
		o.FilteredContacts = &FilteredContactsInt
	}
	
	if TotalContacts, ok := FilterpreviewresponseMap["totalContacts"].(float64); ok {
		TotalContactsInt := int(TotalContacts)
		o.TotalContacts = &TotalContactsInt
	}
	
	if Preview, ok := FilterpreviewresponseMap["preview"].([]interface{}); ok {
		PreviewString, _ := json.Marshal(Preview)
		json.Unmarshal(PreviewString, &o.Preview)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Filterpreviewresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
