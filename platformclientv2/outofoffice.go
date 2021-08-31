package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outofoffice
type Outofoffice struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// User
	User **User `json:"user,omitempty"`


	// StartDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`


	// Active
	Active *bool `json:"active,omitempty"`


	// Indefinite
	Indefinite *bool `json:"indefinite,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Outofoffice) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outofoffice
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		User **User `json:"user,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		Indefinite *bool `json:"indefinite,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		User: o.User,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		Active: o.Active,
		
		Indefinite: o.Indefinite,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Outofoffice) UnmarshalJSON(b []byte) error {
	var OutofofficeMap map[string]interface{}
	err := json.Unmarshal(b, &OutofofficeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OutofofficeMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := OutofofficeMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if User, ok := OutofofficeMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if startDateString, ok := OutofofficeMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := OutofofficeMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if Active, ok := OutofofficeMap["active"].(bool); ok {
		o.Active = &Active
	}
	
	if Indefinite, ok := OutofofficeMap["indefinite"].(bool); ok {
		o.Indefinite = &Indefinite
	}
	
	if SelfUri, ok := OutofofficeMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outofoffice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
