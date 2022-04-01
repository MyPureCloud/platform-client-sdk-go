package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webmessagingrecipient - Information about the recipient the message is sent to or received from.
type Webmessagingrecipient struct { 
	// FirstName - First name of the recipient.
	FirstName *string `json:"firstName,omitempty"`


	// LastName - Last name of the recipient.
	LastName *string `json:"lastName,omitempty"`


	// Nickname - Nickname or display name of the recipient.
	Nickname *string `json:"nickname,omitempty"`


	// AdditionalIds - List of recipient additional identifiers
	AdditionalIds *[]Recipientadditionalidentifier `json:"additionalIds,omitempty"`

}

func (o *Webmessagingrecipient) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webmessagingrecipient
	
	return json.Marshal(&struct { 
		FirstName *string `json:"firstName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		
		Nickname *string `json:"nickname,omitempty"`
		
		AdditionalIds *[]Recipientadditionalidentifier `json:"additionalIds,omitempty"`
		*Alias
	}{ 
		FirstName: o.FirstName,
		
		LastName: o.LastName,
		
		Nickname: o.Nickname,
		
		AdditionalIds: o.AdditionalIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Webmessagingrecipient) UnmarshalJSON(b []byte) error {
	var WebmessagingrecipientMap map[string]interface{}
	err := json.Unmarshal(b, &WebmessagingrecipientMap)
	if err != nil {
		return err
	}
	
	if FirstName, ok := WebmessagingrecipientMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
	
	if LastName, ok := WebmessagingrecipientMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
	
	if Nickname, ok := WebmessagingrecipientMap["nickname"].(string); ok {
		o.Nickname = &Nickname
	}
	
	if AdditionalIds, ok := WebmessagingrecipientMap["additionalIds"].([]interface{}); ok {
		AdditionalIdsString, _ := json.Marshal(AdditionalIds)
		json.Unmarshal(AdditionalIdsString, &o.AdditionalIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webmessagingrecipient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
