package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentactivitychangedtopicoutofoffice
type Agentactivitychangedtopicoutofoffice struct { 
	// Active
	Active *bool `json:"active,omitempty"`


	// ModifiedDate
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

func (o *Agentactivitychangedtopicoutofoffice) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentactivitychangedtopicoutofoffice
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Active *bool `json:"active,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		*Alias
	}{ 
		Active: o.Active,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Agentactivitychangedtopicoutofoffice) UnmarshalJSON(b []byte) error {
	var AgentactivitychangedtopicoutofofficeMap map[string]interface{}
	err := json.Unmarshal(b, &AgentactivitychangedtopicoutofofficeMap)
	if err != nil {
		return err
	}
	
	if Active, ok := AgentactivitychangedtopicoutofofficeMap["active"].(bool); ok {
		o.Active = &Active
	}
	
	if modifiedDateString, ok := AgentactivitychangedtopicoutofofficeMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agentactivitychangedtopicoutofoffice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
