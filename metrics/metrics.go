package metrics

import (
	mp "github.com/mackerelio/go-mackerel-plugin"
	"github.com/nasa9084/go-switchbot/v3"
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
	AmbientBrightness = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "ambient_brightness",
			Label: "SwitchBot (Ambient Brightness)",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			value, _ := status.Brightness.AmbientBrightness()
			switch value {
			case switchbot.AmbientBrightnessDim:
				return 0
			case switchbot.AmbientBrightnessBright:
				return 1
			default:
				return 0
			}
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
	Weight = &MetricSource{
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
	Calibrate = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "calibrate",
			Label: "SwitchBot (Calibrate)",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			if status.IsCalibrated {
				return 1
			}

			return 0
		},
	}
	Group = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "group",
			Label: "SwitchBot (Group)",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			if status.IsGrouped {
				return 1
			}

			return 0
		},
	}
	Moving = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "moving",
			Label: "SwitchBot (Moving)",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			if status.IsMoving {
				return 1
			}

			return 0
		},
	}
	MoveDetected = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "move_detected",
			Label: "SwitchBot (Move Detected)",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			if status.IsMoveDetected {
				return 1
			}

			return 0
		},
	}
	OnlineStatus = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "online_status",
			Label: "SwitchBot (Online Status)",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			switch status.OnlineStatus {
			case switchbot.CleanerOffline:
				return 0
			case switchbot.CleanerOnline:
				return 1
			default:
				return 0
			}
		},
	}
	Auto = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "auto",
			Label: "SwitchBot (Auto)",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			if status.IsAuto {
				return 1
			}

			return 0
		},
	}
	ChildLock = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "child_lock",
			Label: "SwitchBot (Child Lock)",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			if status.IsChildLock {
				return 1
			}

			return 0
		},
	}
	Sound = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "sound",
			Label: "SwitchBot (Sound)",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			if status.IsSound {
				return 1
			}

			return 0
		},
	}
	LackWater = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "lack_water",
			Label: "SwitchBot (Lack Water)",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			if status.IsLackWater {
				return 1
			}

			return 0
		},
	}
	LightLevel = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "light_level",
			Label: "SwitchBot (Light Level)",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.LightLevel)
		},
	}
	FanSpeed = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "fan_speed",
			Label: "SwitchBot (Fan Speed)",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.FanSpeed)
		},
	}
	CO2 = &MetricSource{
		Metrics: &mp.Metrics{
			Name:  "co2",
			Label: "SwitchBot (CO2)",
		},
		Unit: mp.UnitInteger,
		Value: func(status *switchbot.DeviceStatus) float64 {
			return float64(status.CO2)
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
	Weight,
	ElectricityOfDay,
	ElectricCurrent,
	NebulizationEfficiency,
	SlidePosition,
	Calibrate,
	Group,
	Moving,
	MoveDetected,
	AmbientBrightness,
	OnlineStatus,
	Auto,
	ChildLock,
	Sound,
	LackWater,
	LightLevel,
	FanSpeed,
	CO2,
}

var SupportedMetrics = map[switchbot.PhysicalDeviceType][]*MetricSource{
	// https://github.com/OpenWonderLabs/SwitchBotAPI/blob/main/README.md#responses-1
	switchbot.Bot:                      {Battery},
	switchbot.Curtain:                  {Calibrate, Group, Moving, Battery, SlidePosition},
	"Curtain3":                         {Calibrate, Group, Moving, Battery, SlidePosition},
	switchbot.Meter:                    {Temperature, Battery, Humidity},
	switchbot.MeterPlus:                {Battery, Temperature, Humidity},
	"MeterPro(CO2)":                    {Battery, Temperature, Humidity, CO2},
	switchbot.WoIOSensor:               {Battery, Temperature, Humidity},
	switchbot.Lock:                     {Battery /* lockState, doorState */, Calibrate},
	"Smart Lock Pro":                   {Battery /* lockState, doorState */, Calibrate},
	switchbot.KeyPad:                   {},
	switchbot.KeyPadTouch:              {},
	switchbot.MotionSensor:             {Battery, MoveDetected, AmbientBrightness},
	switchbot.ContactSensor:            {Battery, MoveDetected /* openState */, AmbientBrightness},
	"Water Detector":                   {Battery /* status */},
	switchbot.CeilingLight:             {Brightness, ColorTemperature},
	switchbot.CeilingLightPro:          {Brightness, ColorTemperature},
	switchbot.PlugMiniUS:               {Voltage, Weight, ElectricityOfDay, ElectricCurrent},
	switchbot.PlugMiniJP:               {Voltage, Weight, ElectricityOfDay, ElectricCurrent},
	switchbot.Plug:                     {},
	switchbot.StripLight:               {Brightness},
	switchbot.ColorBulb:                {Brightness, ColorTemperature},
	switchbot.RobotVacuumCleanerS1:     { /* workingStatus */ OnlineStatus, Battery},
	switchbot.RobotVacuumCleanerS1Plus: { /* workingStatus */ OnlineStatus, Battery},
	"K10+":                             { /* workingStatus */ OnlineStatus, Battery},
	"K10+ Pro":                         { /* workingStatus */ OnlineStatus, Battery},
	"Robot Vacuum Cleaner S10":         { /* workingStatus */ OnlineStatus, Battery /* waterBaseBatterym, taskType */},
	switchbot.Humidifier:               {Humidity, Temperature, NebulizationEfficiency, Auto, ChildLock, Sound, LackWater},
	switchbot.BlindTilt:                {Calibrate, Group, Moving, SlidePosition},
	switchbot.Hub2:                     {Temperature, LightLevel, Humidity},
	"Battery Circulator Fan":           {Battery /* nightStatus */, FanSpeed},
}
