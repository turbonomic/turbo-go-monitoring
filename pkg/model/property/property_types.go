// The property package contains definitions within the context of metric property.
package property

import "github.com/turbonomic/turbo-go-monitoring/pkg/model"

// List of metric property types
const (
	// The size of the resource
	CAP     model.MetricPropType = "Capacity"
	// The reservation of the resource
	RESERVATION    model.MetricPropType = "Reservation"
	//The amount of consumption in the last polling
	USED    model.MetricPropType = "Used"
	// The peak amount of consumption in the last polling
	PEAK    model.MetricPropType = "Peak"
	//The average amount of consumption in the last polling
	AVERAGE model.MetricPropType = "Average"
	//The maximum amount of consumption during the entire time period
	MAX	model.MetricPropType = "Max"
)
