package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Billingusagereport
type Billingusagereport struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// StartDate - The period start date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The period end date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`


	// Status - Generation status of report
	Status *string `json:"status,omitempty"`


	// Usages - The usages for the given period.
	Usages *[]Billingusage `json:"usages,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Billingusagereport) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Billingusagereport
	
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
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Usages *[]Billingusage `json:"usages,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		Status: o.Status,
		
		Usages: o.Usages,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Billingusagereport) UnmarshalJSON(b []byte) error {
	var BillingusagereportMap map[string]interface{}
	err := json.Unmarshal(b, &BillingusagereportMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BillingusagereportMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := BillingusagereportMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if startDateString, ok := BillingusagereportMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := BillingusagereportMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if Status, ok := BillingusagereportMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Usages, ok := BillingusagereportMap["usages"].([]interface{}); ok {
		UsagesString, _ := json.Marshal(Usages)
		json.Unmarshal(UsagesString, &o.Usages)
	}
	
	if SelfUri, ok := BillingusagereportMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Billingusagereport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
