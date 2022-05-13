package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outofofficeeventoutofoffice
type Outofofficeeventoutofoffice struct { 
	// User
	User *Outofofficeeventuser `json:"user,omitempty"`


	// Active
	Active *bool `json:"active,omitempty"`


	// Indefinite
	Indefinite *bool `json:"indefinite,omitempty"`


	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate
	EndDate *time.Time `json:"endDate,omitempty"`

}

func (o *Outofofficeeventoutofoffice) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outofofficeeventoutofoffice
	
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
		User *Outofofficeeventuser `json:"user,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		Indefinite *bool `json:"indefinite,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		*Alias
	}{ 
		User: o.User,
		
		Active: o.Active,
		
		Indefinite: o.Indefinite,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Outofofficeeventoutofoffice) UnmarshalJSON(b []byte) error {
	var OutofofficeeventoutofofficeMap map[string]interface{}
	err := json.Unmarshal(b, &OutofofficeeventoutofofficeMap)
	if err != nil {
		return err
	}
	
	if User, ok := OutofofficeeventoutofofficeMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Active, ok := OutofofficeeventoutofofficeMap["active"].(bool); ok {
		o.Active = &Active
	}
    
	if Indefinite, ok := OutofofficeeventoutofofficeMap["indefinite"].(bool); ok {
		o.Indefinite = &Indefinite
	}
    
	if startDateString, ok := OutofofficeeventoutofofficeMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := OutofofficeeventoutofofficeMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outofofficeeventoutofoffice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
