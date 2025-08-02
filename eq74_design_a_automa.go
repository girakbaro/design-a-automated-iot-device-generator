package main

import (
    "encoding/json"
    "fmt"
    "time"
)

type IoTDevice struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Type        string    `json:"type"` // sensor, actuator, etc.
    IP          string    `json:"ip"`
    Port        int       `json:"port"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type GeneratorConfig struct {
    DeviceType string `json:"device_type"` // sensor, actuator, etc.
    NumDevices int    `json:"num_devices"`
    IPRange    struct {
        Start string `json:"start"`
        End   string `json:"end"`
    } `json:"ip_range"`
    PortRange struct {
        Start int `json:"start"`
        End   int `json:"end"`
    } `json:"port_range"`
}

type AutomatedGenerator struct {
    Config   GeneratorConfig `json:"config"`
    Devices  []IoTDevice    `json:"devices"`
    Generate func(config GeneratorConfig) ([]IoTDevice, error)
}

func (g *AutomatedGenerator) GenerateDevices() ([]IoTDevice, error) {
    return g.Generate(g.Config)
}

func NewAutomatedGenerator(config GeneratorConfig) *AutomatedGenerator {
    return &AutomatedGenerator{
        Config: config,
        Generate: func(config GeneratorConfig) ([]IoTDevice, error) {
            devices := make([]IoTDevice, config.NumDevices)
            for i := 0; i < config.NumDevices; i++ {
                device := IoTDevice{
                    ID:          fmt.Sprintf("device-%d", i),
                    Name:        fmt.Sprintf("Device %d", i),
                    Description: "Automatically generated device",
                    Type:        config.DeviceType,
                    IP:          fmt.Sprintf("%s.%d", config.IPRange.Start, i),
                    Port:        config.PortRange.Start + i,
                    CreatedAt:   time.Now(),
                    UpdatedAt:   time.Now(),
                }
                devices[i] = device
            }
            return devices, nil
        },
    }
}

func main() {
    config := GeneratorConfig{
        DeviceType: "sensor",
        NumDevices: 5,
        IPRange: struct {
            Start string
            End   string
        }{
            Start: "192.168.1.",
            End:   "192.168.1.",
        },
        PortRange: struct {
            Start int
            End   int
        }{
            Start: 8080,
            End:   8089,
        },
    }

    generator := NewAutomatedGenerator(config)
    devices, err := generator.GenerateDevices()
    if err != nil {
        fmt.Println(err)
        return
    }

    jsonDevices, err := json.MarshalIndent(devices, "", "  ")
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(jsonDevices))
}