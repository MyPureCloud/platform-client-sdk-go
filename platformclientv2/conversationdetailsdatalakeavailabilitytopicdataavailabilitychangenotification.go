package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification
type Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification struct { 
	// DataAvailabilityDate - Date and time before which data is guaranteed to be available in the datalake
	DataAvailabilityDate *time.Time `json:"dataAvailabilityDate,omitempty"`

}

func (o *Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification
	
	DataAvailabilityDate := new(string)
	if o.DataAvailabilityDate != nil {
		
		*DataAvailabilityDate = timeutil.Strftime(o.DataAvailabilityDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DataAvailabilityDate = nil
	}
	
	return json.Marshal(&struct { 
		DataAvailabilityDate *string `json:"dataAvailabilityDate,omitempty"`
		*Alias
	}{ 
		DataAvailabilityDate: DataAvailabilityDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification) UnmarshalJSON(b []byte) error {
	var ConversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotificationMap)
	if err != nil {
		return err
	}
	
	if dataAvailabilityDateString, ok := ConversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotificationMap["dataAvailabilityDate"].(string); ok {
		DataAvailabilityDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", dataAvailabilityDateString)
		o.DataAvailabilityDate = &DataAvailabilityDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
