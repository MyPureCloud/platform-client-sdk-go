package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Defaultgreetinglist
type Defaultgreetinglist struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Owner
	Owner *Greetingowner `json:"owner,omitempty"`


	// OwnerType
	OwnerType *string `json:"ownerType,omitempty"`


	// Greetings
	Greetings *map[string]Greeting `json:"greetings,omitempty"`


	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// CreatedBy
	CreatedBy *string `json:"createdBy,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// ModifiedBy
	ModifiedBy *string `json:"modifiedBy,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Defaultgreetinglist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Defaultgreetinglist
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Owner *Greetingowner `json:"owner,omitempty"`
		
		OwnerType *string `json:"ownerType,omitempty"`
		
		Greetings *map[string]Greeting `json:"greetings,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Owner: o.Owner,
		
		OwnerType: o.OwnerType,
		
		Greetings: o.Greetings,
		
		CreatedDate: CreatedDate,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedDate: ModifiedDate,
		
		ModifiedBy: o.ModifiedBy,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Defaultgreetinglist) UnmarshalJSON(b []byte) error {
	var DefaultgreetinglistMap map[string]interface{}
	err := json.Unmarshal(b, &DefaultgreetinglistMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DefaultgreetinglistMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DefaultgreetinglistMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Owner, ok := DefaultgreetinglistMap["owner"].(map[string]interface{}); ok {
		OwnerString, _ := json.Marshal(Owner)
		json.Unmarshal(OwnerString, &o.Owner)
	}
	
	if OwnerType, ok := DefaultgreetinglistMap["ownerType"].(string); ok {
		o.OwnerType = &OwnerType
	}
    
	if Greetings, ok := DefaultgreetinglistMap["greetings"].(map[string]interface{}); ok {
		GreetingsString, _ := json.Marshal(Greetings)
		json.Unmarshal(GreetingsString, &o.Greetings)
	}
	
	if createdDateString, ok := DefaultgreetinglistMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if CreatedBy, ok := DefaultgreetinglistMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
    
	if modifiedDateString, ok := DefaultgreetinglistMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if ModifiedBy, ok := DefaultgreetinglistMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
    
	if SelfUri, ok := DefaultgreetinglistMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Defaultgreetinglist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
