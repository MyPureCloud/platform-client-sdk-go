package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactscontactchangedtopiclineuserid
type Externalcontactscontactchangedtopiclineuserid struct { 
	// UserId
	UserId *string `json:"userId,omitempty"`

}

func (o *Externalcontactscontactchangedtopiclineuserid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactscontactchangedtopiclineuserid
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		*Alias
	}{ 
		UserId: o.UserId,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalcontactscontactchangedtopiclineuserid) UnmarshalJSON(b []byte) error {
	var ExternalcontactscontactchangedtopiclineuseridMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactscontactchangedtopiclineuseridMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := ExternalcontactscontactchangedtopiclineuseridMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactscontactchangedtopiclineuserid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
