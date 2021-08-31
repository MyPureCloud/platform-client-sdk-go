package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicalimportdeletejobresponse
type Historicalimportdeletejobresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Status - Property denoting the status of the delete.
	Status *string `json:"status,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Historicalimportdeletejobresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historicalimportdeletejobresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Status: o.Status,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Historicalimportdeletejobresponse) UnmarshalJSON(b []byte) error {
	var HistoricalimportdeletejobresponseMap map[string]interface{}
	err := json.Unmarshal(b, &HistoricalimportdeletejobresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := HistoricalimportdeletejobresponseMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := HistoricalimportdeletejobresponseMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Status, ok := HistoricalimportdeletejobresponseMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if SelfUri, ok := HistoricalimportdeletejobresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Historicalimportdeletejobresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
