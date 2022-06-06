package main

import (
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/network"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"

	"github.com/Thomasparsley/server-stats/data"
)

const (
	PORT = ":9110"
)

var (
	VerificationKey string
)

func main() {
	//s, _ := NewStatistic()

	app := fiber.New(fiber.Config{
		ServerHeader: "vel",
	})

	app.Use(logger.New())
	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(compress.New())
	app.Use(helmet.New())

	app.Get("/stats", func(c *fiber.Ctx) error {
		s, err := NewStatistic()

		if err != nil {
			return err
		}

		return c.JSON(s)
	})

	app.Listen(PORT)
}

func NewStatistic() (*data.Statistic, error) {
	memory, err := memory.Get()

	if err != nil {
		return nil, err
	}

	var networks []data.StatisticNetwork
	rawNetworksBefore, err := network.Get()
	if err != nil {
		return nil, err
	}

	for _, n := range rawNetworksBefore {
		networks = append(networks, data.StatisticNetwork{
			Name:         n.Name,
			RxBytesTotal: n.RxBytes,
			TxBytesTotal: n.TxBytes,
		})
	}

	cpuBefore, err := cpu.Get()
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Duration(1) * time.Second)

	cpuAfter, err := cpu.Get()
	if err != nil {
		return nil, err
	}

	rawNetworksAfter, err := network.Get()
	if err != nil {
		return nil, err
	}

	for i, n := range rawNetworksAfter {
		networks[i].RxBytesPerSec = n.RxBytes - rawNetworksBefore[i].RxBytes
		networks[i].TxBytesPerSec = n.TxBytes - rawNetworksBefore[i].TxBytes
	}

	cpuTotal := float64(cpuAfter.Total - cpuBefore.Total)
	return &data.Statistic{
		Cpu: data.StatisticCpu{
			UserPercent:   (float64(cpuAfter.User-cpuBefore.User) / cpuTotal * 100),
			SystemPercent: (float64(cpuAfter.System-cpuBefore.System) / cpuTotal * 100),
			IdlePercent:   (float64(cpuAfter.Idle-cpuBefore.Idle) / cpuTotal * 100),
			CoreCount:     cpuAfter.CPUCount,
		},
		Memory: data.StatisticMemory{
			Total:       memory.Total,
			Used:        memory.Used,
			UsedPercent: (float64(memory.Used) / float64(memory.Total) * 100),
			Cached:      memory.Cached,
			Free:        memory.Total - memory.Used,
			FreePercent: 100 - (float64(memory.Used) / float64(memory.Total) * 100),
		},
		Networks: networks,
	}, nil
}
