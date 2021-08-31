package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trustuserdetails
type Trustuserdetails struct { 
	// DateCreated - Date Trust User was added. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// CreatedBy - User that added trusted user.
	CreatedBy *Orguser `json:"createdBy,omitempty"`

}

func (o *Trustuserdetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trustuserdetails
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		DateCreated *string `json:"dateCreated,omitempty"`
		
		CreatedBy *Orguser `json:"createdBy,omitempty"`
		*Alias
	}{ 
		DateCreated: DateCreated,
		
		CreatedBy: o.CreatedBy,
		Alias:    (*Alias)(o),
	})
}

func (o *Trustuserdetails) UnmarshalJSON(b []byte) error {
	var TrustuserdetailsMap map[string]interface{}
	err := json.Unmarshal(b, &TrustuserdetailsMap)
	if err != nil {
		return err
	}
	
	if dateCreatedString, ok := TrustuserdetailsMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if CreatedBy, ok := TrustuserdetailsMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trustuserdetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
