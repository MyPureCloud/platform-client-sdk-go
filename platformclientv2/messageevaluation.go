package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messageevaluation
type Messageevaluation struct { 
	// ContactColumn - The name of the contact column that was wrapped up
	ContactColumn *string `json:"contactColumn,omitempty"`


	// ContactAddress - The address (phone or email) that was wrapped up
	ContactAddress *string `json:"contactAddress,omitempty"`


	// MessageType - The type of message sent
	MessageType *string `json:"messageType,omitempty"`


	// WrapupCodeId - The id of the wrap-up code
	WrapupCodeId *string `json:"wrapupCodeId,omitempty"`


	// Timestamp - The time that the wrap-up was applied. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`

}

func (u *Messageevaluation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messageevaluation

	
	Timestamp := new(string)
	if u.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(u.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	

	return json.Marshal(&struct { 
		ContactColumn *string `json:"contactColumn,omitempty"`
		
		ContactAddress *string `json:"contactAddress,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		WrapupCodeId *string `json:"wrapupCodeId,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		*Alias
	}{ 
		ContactColumn: u.ContactColumn,
		
		ContactAddress: u.ContactAddress,
		
		MessageType: u.MessageType,
		
		WrapupCodeId: u.WrapupCodeId,
		
		Timestamp: Timestamp,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Messageevaluation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
