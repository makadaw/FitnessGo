package conf

import (
	"encoding/json"
	"fmt"
	"os"
    "strconv"
    "log"
)

type config struct {
    Debug bool

	WebHost string "web address"
	WebPort int    "web port"

	SessionSecret             string

	DbHost string
	DbPort int
	DbName string
    DbUrl string
}

var Path = "./config.json"
var Config = new(config)

func (c *config) HostString() string {
	return fmt.Sprintf("%s:%d", c.WebHost, c.WebPort)
}

func (c *config) DbHostString() string {
    if len(c.DbUrl) > 0 {
        return c.DbUrl
    }
	if c.DbPort > 0 {
		return fmt.Sprintf("mongodb://%s:%d", c.DbHost, c.DbPort)
	}
	return fmt.Sprintf("mongodb://%s", c.DbHost)

}

func (c *config) String() string {
	s := "Config:"
	s += fmt.Sprintf("   Host: %s,\n", c.HostString())
	s += fmt.Sprintf("   DB: %s,\n", c.DbHostString())
	s += fmt.Sprintf("   Debug: %v\n", c.Debug)
	return s
}

func init() {
    Config.WebHost = "0.0.0.0"
	Config.WebPort = 4000
	Config.DbHost = "127.0.0.1"
	Config.DbPort = 0
	Config.DbName = "fitness"
	Config.SessionSecret = "SECRET-KEY-SET-IN-CONFIG"
	Config.Debug = false
	
    mode := os.Getenv("MARTINI_ENV")
    log.Println("mode", mode)
    if mode == "production" {
        Path = "./conf/config.json"
    } else {
        Path = "./conf/config.dev.json"
    }
    
	file, err := os.Open(Path)
	if err != nil {
		if len(Path) > 1 {
			fmt.Printf("Error: could not read config file %s.\n", Path)
		}
		return
	}
	decoder := json.NewDecoder(file)
	// overwrite in-mem config with new values
	err = decoder.Decode(Config)
    if len(os.Getenv("PORT")) > 0 {
        port, _ := strconv.ParseInt(os.Getenv("PORT"), 10, 0)
        log.Println("heroku port", port)
        Config.WebPort = int(port)
    }
    if len(os.Getenv("MONGOLAB_URI")) > 0 {
        Config.DbUrl = os.Getenv("MONGOLAB_URI")
    }
	if err != nil {
		fmt.Printf("Error decoding file %s\n%s\n", Path, err)
	}
}