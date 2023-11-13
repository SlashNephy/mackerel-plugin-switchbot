package main

import (
	mp "github.com/mackerelio/go-mackerel-plugin"
	"github.com/nasa9084/go-switchbot"

	switchbot2 "github.com/SlashNephy/mackerel-plugin-switchbot/switchbot"
)

type metricSource struct {
	mp.Metrics
	Unit  string
	Value func(status *switchbot.DeviceStatus) float64
}

var (
	battery = &metricSource{
		Metrics: mp.Metrics{
			Name:  "battery",
			Label: "Battery",
		},
		Unit: mp.UnitPercentage,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.Battery)
		},
	}
	temperature = &metricSource{
		Metrics: mp.Metrics{
			Name:  "temperature",
			Label: "Temperature",
		},
		Unit: mp.UnitFloat,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return status.Temperature
		},
	}
	humidity = &metricSource{
		Metrics: mp.Metrics{
			Name:  "humidity",
			Label: "Humidity",
		},
		Unit: mp.UnitPercentage,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.Humidity)
		},
	}
	brightness = &metricSource{
		Metrics: mp.Metrics{
			Name:  "brightness",
			Label: "Brightness",
		},
		Unit: mp.UnitPercentage,
		Value: func(status *switchbot.DeviceStatus) float64 {
			value, _ := status.Brightness.Int()
			return float64(value)
		},
	}
	colorTemperature = &metricSource{
		Metrics: mp.Metrics{
			Name:  "color_temperature",
			Label: "Color Temperature",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.ColorTemperature)
		},
	}
	voltage = &metricSource{
		Metrics: mp.Metrics{
			Name:  "voltage",
			Label: "Voltage",
		},
		Unit: mp.UnitFloat,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.Voltage)
		},
	}
	weight = &metricSource{
		Metrics: mp.Metrics{
			Name:  "weight",
			Label: "Weight",
		},
		Unit: mp.UnitFloat,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.Weight)
		},
	}
	electricityOfDay = &metricSource{
		Metrics: mp.Metrics{
			Name:  "electricity_of_day",
			Label: "Electricity of Day",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.ElectricityOfDay)
		},
	}
	electricCurrent = &metricSource{
		Metrics: mp.Metrics{
			Name:  "electric_current",
			Label: "Electric Current",
		},
		Unit: mp.UnitFloat,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return status.ElectricCurrent
		},
	}
	nebulizationEfficiency = &metricSource{
		Metrics: mp.Metrics{
			Name:  "nebulization_efficiency",
			Label: "Nebulization Efficiency",
		},
		Unit: mp.UnitPercentage,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.NebulizationEfficiency)
		},
	}
	slidePosition = &metricSource{
		Metrics: mp.Metrics{
			Name:  "slide_position",
			Label: "Slide Position",
		},
		Unit: mp.UnitPercentage,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.SlidePosition)
		},
	}
)

var allMetricSources = []*metricSource{
	battery,
	temperature,
	humidity,
	brightness,
	colorTemperature,
	voltage,
	weight,
	electricityOfDay,
	electricCurrent,
	nebulizationEfficiency,
	slidePosition,
}

// https://github.com/OpenWonderLabs/SwitchBotAPI/blob/main/README.md#get-device-status
var supportedMetrics = map[switchbot.PhysicalDeviceType][]*metricSource{
	switchbot.Bot:                      {battery},
	switchbot.Curtain:                  {battery},
	switchbot.Meter:                    {temperature, battery, humidity},
	switchbot.MeterPlus:                {battery, temperature, humidity},
	switchbot2.OutdoorMeter:            {battery, temperature, humidity},
	switchbot2.SmartLock:               {battery},
	switchbot.MotionSensor:             {battery},
	switchbot.ContactSensor:            {battery},
	switchbot.CeilingLight:             {brightness, colorTemperature},
	switchbot.CeilingLightPro:          {brightness, colorTemperature},
	switchbot.PlugMiniUS:               {voltage, weight, electricityOfDay, electricCurrent},
	switchbot.PlugMiniJP:               {voltage, weight, electricityOfDay, electricCurrent},
	switchbot.StripLight:               {brightness},
	switchbot.ColorBulb:                {brightness, colorTemperature},
	switchbot.RobotVacuumCleanerS1:     {battery},
	switchbot.RobotVacuumCleanerS1Plus: {battery},
	switchbot.Humidifier:               {humidity, temperature, nebulizationEfficiency},
	switchbot2.BlindTilt:               {slidePosition},
	switchbot2.Hub2:                    {temperature, humidity}, // missing lightLevel
}
