package elements

// The TimeZoneDefinition element specifies the periods and transitions that define a time zone.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/timezonedefinition
import "encoding/xml"

type TimeZoneDefinition struct {
	XMLName xml.Name

	// The Period element defines the name, time offset, and unique identifier for a specific stage of the time zone.
	Period *Period `xml:"Period"`
	// The Periods element represents an array of periods that define the time offset at different stages of the time zone.
	Periods *Periods `xml:"Periods"`
	// The Transitions element represents an array of time zone transitions.
	Transitions *Transitions `xml:"Transitions"`
	// The TransitionsGroup element represents an array of time zone transitions.
	TransitionsGroup *TransitionsGroup `xml:"TransitionsGroup"`
	// The TransitionsGroups element represents an array of time zone transition groups.
	TransitionsGroups *TransitionsGroups `xml:"TransitionsGroups"`
	// Represents the unique identifier of the time zone.
	Id *string `xml:"Id,attr"`
	// Represents the descriptive name of the time zone.
	Name *string `xml:"Name,attr"`
}

func (T *TimeZoneDefinition) SetForMarshal() {
	T.XMLName.Local = "t:TimeZoneDefinition"
}

func (T *TimeZoneDefinition) GetSchema() *Schema {
	return &SchemaTypes
}
