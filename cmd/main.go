package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	cachePkg "github.com/SkinGad/rmd/internal/cache"
	"github.com/SkinGad/rmd/internal/config"
	"github.com/SkinGad/rmd/internal/ros"
	"github.com/SkinGad/rmd/pkg/lookup"
)

var cache *cachePkg.Cache
var cfg config.Config
var miks []*ros.Mikrotik
var err error

func main() {
	initLogger()
	cfg, err = config.ParsingConfig("config.yml")
	if err != nil {
		slog.Error("Error read config", "err", err)
	}
	slog.Info("Config read")

	cache, err = cachePkg.NewCache()
	if err != nil {
		slog.Error("Error read cache", "err", err)
	}
	slog.Info("Cache read")

	look()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigChan
		slog.Info("Stop service", "sig", sig)
		closeMikrotiksConnection()
		cache.Save()
		os.Exit(0)
	}()

	for _, router := range cfg.Routers {
		go connectToMikrotik(router)
	}

	go func() {
		for {
			look()
			time.Sleep(cfg.RMD.LookUp)
		}
	}()

	go syncMikrotikRoute()
	for {
		cleaner()
		time.Sleep(time.Second)
	}
}

func initLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		// Level:     slog.LevelDebug,
	}))
	slog.SetDefault(logger)
	return logger
}

func look() {
	for _, v := range cfg.Domains {
		slog.Debug("start lookup fo domain", "domain", v)
		IPs := lookup.LookUp(v)
		cache.UpdateAddresses(cfg.RMD.TTL, v, IPs...)
	}
}

func closeMikrotiksConnection() {
	for _, v := range miks {
		v.Con.Close()
		slog.Info("Close connected to mikrotik", "address", v.Address, "port", v.Port)
	}
}

func syncMikrotikRoute() {
	var wg sync.WaitGroup
	for {
		for _, mik := range miks {
			wg.Go(func() {
				var ro []ros.Routes
				for _, v := range cache.Addresses {
					ro = append(ro, ros.Routes{
						Address: v.Address,
						Comment: v.Domain,
					})
				}
				routes := mik.GetRouts()
				for _, v := range routes {
					cache.AddAddresses(cfg.RMD.TTL, v.Comment, v.Address)
				}
				mik.SyncRoute(ro)
			})
		}
		wg.Wait()
		time.Sleep(time.Second * 10)
	}
}

func cleaner() {
	addresses := cache.FindClear()
	for _, v := range addresses {
		for _, m := range miks {
			m.DisableRoute(v.Address)
		}
		cache.Clear(v.Address)
	}
}

func connectToMikrotik(router config.Router) {
	for {
		mik, err := ros.Connect(
			router.Address,
			router.Login,
			router.Password,
			router.RoutingTable,
			router.Gateway,
			router.Port,
		)
		if err != nil {
			time.Sleep(time.Second * 5)
			continue
		}
		routes := mik.GetRouts()
		for _, v := range routes {
			cache.AddAddresses(cfg.RMD.TTL, v.Comment, v.Address)
		}
		slog.Info("Add routes to cache from mikrotik", "routes", routes)
		miks = append(miks, mik)
		helthcheck(*mik)
		break
	}
}

func helthcheck(m ros.Mikrotik) {
	for {
		if m.HethCheck() {
		} else {
			for i, v := range miks {
				if v.Address == m.Address && v.Port == m.Port {
					miks = append(miks[:i], miks[i+1:]...)
					break
				}
			}
			m.Con.Close()
			slog.Info("Close connected to mikrotik", "address", m.Address, "port", m.Port)
			for _, v := range cfg.Routers {
				if v.Address == m.Address && v.Port == m.Port {
					go connectToMikrotik(v)
					break
				}
			}
		}
		time.Sleep(time.Second * 15)
	}
}
