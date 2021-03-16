package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Metricdefinition) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
