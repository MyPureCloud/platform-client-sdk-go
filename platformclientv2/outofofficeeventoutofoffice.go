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

func (u *Outofofficeeventoutofoffice) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outofofficeeventoutofoffice

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if u.EndDate != nil {
		
		*EndDate = timeutil.Strftime(u.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		User: u.User,
		
		Active: u.Active,
		
		Indefinite: u.Indefinite,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Outofofficeeventoutofoffice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
