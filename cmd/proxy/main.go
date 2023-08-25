package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"/restock-monitor/cmd/pystrings"
)

// Credentials are http proxy credentials
type Credentials struct {
	username string
	password string
}

// ProxyString creates proxy string from credentials and given uri
func (c Credentials) ProxyString(uri, protocol string) string {
	r := strings.NewReplacer("{protocol}", protocol, "{username}", c.username, "{password}", c.password, "{uri}", uri)
	return r.Replace("{protocol}://{username}:{password}@{uri}")
}

// SetProxyEnv sets proxy environtment variable
func (c Credentials) SetProxyEnv(uri, protocol string) error {
	var envKey string
	switch protocol {
	case "http":
		envKey = "http_proxy"
	case "https":
		envKey = "https_proxy"
	default:
		return fmt.Errorf("unkown protocol %s", protocol)

	}

	return os.Setenv(envKey, c.ProxyString(uri, protocol))
}

var (
	delay      time.Duration
	proxieFile string
)

func init() {
	flag.DurationVar(&delay, "delay", time.Second*30, "delay between proxy change")
	flag.StringVar(&proxieFile, "proxieFile", "./proxies.txt", "filepath to proxie list")

	flag.Parse()
}

func main() {
	f, err := os.Open(proxieFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	protocol := pystrings.String(lines[0]).Split(":")[1].Trim(" ").Lower()
	username := pystrings.String(lines[1]).Split(":")[1].Trim(" ")
	password := pystrings.String(lines[2]).Split(":")[1].Trim(" ")

	credentials := Credentials{
		username: string(username),
		password: string(password),
	}

	_ = credentials
	_ = protocol

	proxies := lines[7:]
	for range time.NewTicker(delay).C {
		uri := proxies[rand.Intn(len(proxies))]
		credentials.SetProxyEnv(uri, string(protocol))
	}
}
