package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Presenceeventuserpresence
type Presenceeventuserpresence struct { 
	// Source
	Source *string `json:"source,omitempty"`


	// PresenceDefinition
	PresenceDefinition *Presenceeventorganizationpresence `json:"presenceDefinition,omitempty"`


	// Primary
	Primary *bool `json:"primary,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// ModifiedDate
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

func (u *Presenceeventuserpresence) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Presenceeventuserpresence

	
	ModifiedDate := new(string)
	if u.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(u.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	

	return json.Marshal(&struct { 
		Source *string `json:"source,omitempty"`
		
		PresenceDefinition *Presenceeventorganizationpresence `json:"presenceDefinition,omitempty"`
		
		Primary *bool `json:"primary,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		*Alias
	}{ 
		Source: u.Source,
		
		PresenceDefinition: u.PresenceDefinition,
		
		Primary: u.Primary,
		
		Message: u.Message,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Presenceeventuserpresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
