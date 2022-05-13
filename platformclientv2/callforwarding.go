package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Callforwarding
type Callforwarding struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// User
	User *User `json:"user,omitempty"`


	// Enabled - Whether or not CallForwarding is enabled
	Enabled *bool `json:"enabled,omitempty"`


	// PhoneNumber - This property is deprecated. Please use the calls property
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// Calls - An ordered list of CallRoutes to be executed when CallForwarding is enabled
	Calls *[]Callroute `json:"calls,omitempty"`


	// Voicemail - The type of voicemail to use with the callForwarding configuration
	Voicemail *string `json:"voicemail,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Callforwarding) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callforwarding
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		Calls *[]Callroute `json:"calls,omitempty"`
		
		Voicemail *string `json:"voicemail,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		User: o.User,
		
		Enabled: o.Enabled,
		
		PhoneNumber: o.PhoneNumber,
		
		Calls: o.Calls,
		
		Voicemail: o.Voicemail,
		
		ModifiedDate: ModifiedDate,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Callforwarding) UnmarshalJSON(b []byte) error {
	var CallforwardingMap map[string]interface{}
	err := json.Unmarshal(b, &CallforwardingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CallforwardingMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CallforwardingMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if User, ok := CallforwardingMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Enabled, ok := CallforwardingMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if PhoneNumber, ok := CallforwardingMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
    
	if Calls, ok := CallforwardingMap["calls"].([]interface{}); ok {
		CallsString, _ := json.Marshal(Calls)
		json.Unmarshal(CallsString, &o.Calls)
	}
	
	if Voicemail, ok := CallforwardingMap["voicemail"].(string); ok {
		o.Voicemail = &Voicemail
	}
    
	if modifiedDateString, ok := CallforwardingMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if SelfUri, ok := CallforwardingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Callforwarding) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
