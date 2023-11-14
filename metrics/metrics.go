package metrics

import (
	mp "github.com/mackerelio/go-mackerel-plugin"
	"github.com/nasa9084/go-switchbot/v3"

	switchbot2 "github.com/SlashNephy/mackerel-plugin-switchbot/switchbot"
)

type MetricSource struct {
	*mp.Metrics
	Unit  string
	Value func(status *switchbot.DeviceStatus) float64
}

var (
	Battery = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "battery",
			Label: "SwitchBot (Battery)",
		},
		Unit: mp.UnitPercentage,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.Battery)
		},
	}
	Temperature = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "temperature",
			Label: "SwitchBot (Temperature)",
		},
		Unit: mp.UnitFloat,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return status.Temperature
		},
	}
	Humidity = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "humidity",
			Label: "SwitchBot (Humidity)",
		},
		Unit: mp.UnitPercentage,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.Humidity)
		},
	}
	Brightness = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "brightness",
			Label: "SwitchBot (Brightness)",
		},
		Unit: mp.UnitPercentage,
		Value: func(status *switchbot.DeviceStatus) float64 {
			value, _ := status.Brightness.Int()
			return float64(value)
		},
	}
	ColorTemperature = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "color_temperature",
			Label: "SwitchBot (Color Temperature)",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.ColorTemperature)
		},
	}
	Voltage = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "voltage",
			Label: "SwitchBot (Voltage)",
		},
		Unit: mp.UnitFloat,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return status.Voltage
		},
	}
	weight = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "weight",
			Label: "SwitchBot (Weight)",
		},
		Unit: mp.UnitFloat,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return status.Weight
		},
	}
	ElectricityOfDay = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "electricity_of_day",
			Label: "SwitchBot (Electricity of Day)",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.ElectricityOfDay)
		},
	}
	ElectricCurrent = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "electric_current",
			Label: "SwitchBot (Electric Current)",
		},
		Unit: mp.UnitFloat,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return status.ElectricCurrent
		},
	}
	NebulizationEfficiency = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "nebulization_efficiency",
			Label: "SwitchBot (Nebulization Efficiency)",
		},
		Unit: mp.UnitPercentage,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.NebulizationEfficiency)
		},
	}
	SlidePosition = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "slide_position",
			Label: "SwitchBot (Slide Position)",
		},
		Unit: mp.UnitPercentage,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.SlidePosition)
		},
	}
)

var AllMetrics = []*MetricSource{
	Battery,
	Temperature,
	Humidity,
	Brightness,
	ColorTemperature,
	Voltage,
	weight,
	ElectricityOfDay,
	ElectricCurrent,
	NebulizationEfficiency,
	SlidePosition,
}

var SupportedMetrics = map[switchbot.PhysicalDeviceType][]*MetricSource{
	// https://github.com/OpenWonderLabs/SwitchBotAPI/blob/main/README.md#get-device-status
	switchbot.Bot:                      {Battery},
	switchbot.Curtain:                  {Battery},
	switchbot.Meter:                    {Temperature, Battery, Humidity},
	switchbot.MeterPlus:                {Battery, Temperature, Humidity},
	switchbot2.OutdoorMeter:            {Battery, Temperature, Humidity},
	switchbot2.SmartLock:               {Battery},
	switchbot.MotionSensor:             {Battery},
	switchbot.ContactSensor:            {Battery},
	switchbot.CeilingLight:             {Brightness, ColorTemperature},
	switchbot.CeilingLightPro:          {Brightness, ColorTemperature},
	switchbot.PlugMiniUS:               {Voltage, weight, ElectricityOfDay, ElectricCurrent},
	switchbot.PlugMiniJP:               {Voltage, weight, ElectricityOfDay, ElectricCurrent},
	switchbot.StripLight:               {Brightness},
	switchbot.ColorBulb:                {Brightness, ColorTemperature},
	switchbot.RobotVacuumCleanerS1:     {Battery},
	switchbot.RobotVacuumCleanerS1Plus: {Battery},
	switchbot.Humidifier:               {Humidity, Temperature, NebulizationEfficiency},
	switchbot.BlindTilt:                {SlidePosition},
	switchbot2.Hub2:                    {Temperature, Humidity}, // missing lightLevel
}
