package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Metricdefinition
type Metricdefinition struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// UnitType - The type of associated metric unit
	UnitType *string `json:"unitType,omitempty"`


	// ShortName - An alternate name for this metric definition, often abbreviation
	ShortName *string `json:"shortName,omitempty"`


	// DividendMetrics - Metric names used as dividend
	DividendMetrics *[]string `json:"dividendMetrics,omitempty"`


	// DivisorMetrics - Metric names used as divisor
	DivisorMetrics *[]string `json:"divisorMetrics,omitempty"`


	// DefaultObjective - A predefined default objective for this metric
	DefaultObjective *Defaultobjective `json:"defaultObjective,omitempty"`


	// LockTemplateId - An optional field to specify if this metric definition is locked to certain template. e.g. punctuality
	LockTemplateId *string `json:"lockTemplateId,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Metricdefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Metricdefinition

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		UnitType *string `json:"unitType,omitempty"`
		
		ShortName *string `json:"shortName,omitempty"`
		
		DividendMetrics *[]string `json:"dividendMetrics,omitempty"`
		
		DivisorMetrics *[]string `json:"divisorMetrics,omitempty"`
		
		DefaultObjective *Defaultobjective `json:"defaultObjective,omitempty"`
		
		LockTemplateId *string `json:"lockTemplateId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		UnitType: u.UnitType,
		
		ShortName: u.ShortName,
		
		DividendMetrics: u.DividendMetrics,
		
		DivisorMetrics: u.DivisorMetrics,
		
		DefaultObjective: u.DefaultObjective,
		
		LockTemplateId: u.LockTemplateId,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Metricdefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
