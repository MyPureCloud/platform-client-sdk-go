package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactsunresolvedcontactchangedtopiclineuserid
type Externalcontactsunresolvedcontactchangedtopiclineuserid struct { 
	// UserId
	UserId *string `json:"userId,omitempty"`

}

func (o *Externalcontactsunresolvedcontactchangedtopiclineuserid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactsunresolvedcontactchangedtopiclineuserid
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		*Alias
	}{ 
		UserId: o.UserId,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalcontactsunresolvedcontactchangedtopiclineuserid) UnmarshalJSON(b []byte) error {
	var ExternalcontactsunresolvedcontactchangedtopiclineuseridMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactsunresolvedcontactchangedtopiclineuseridMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := ExternalcontactsunresolvedcontactchangedtopiclineuseridMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactsunresolvedcontactchangedtopiclineuserid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
