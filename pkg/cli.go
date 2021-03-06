package pkg

import (
	"flag"
	"os"
)

type Config struct {
	Directory   string
	Bucket      string
	PoolInteval int
	Profile     string
}

// ParseConfigFromCli gets the config from cli and returns the config
func ParseConfigFromCli() *Config {
	dir := flag.String("dir", "", "directory to watch")
	interval := flag.Int("interval", 30, "The time interval between sync runs in seconds")
	bucket := flag.String("bucket", "", "destination bucket for the files")
	profile := flag.String("profile", "default", "The aws credential profile to use")

	flag.Parse()

	if *dir == "" || *bucket == "" {
		flag.Usage()
		os.Exit(1)
	}

	return &Config{
		Bucket:      *bucket,
		Directory:   *dir,
		PoolInteval: *interval,
		Profile:     *profile,
	}
}
