package common

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

var newdag_Version = "0.0.1"

func DefaultConfig() *Config {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	storeType := "badger"
	storePath, _ := ioutil.TempDir("", "badger")
	return &Config{
		HeartbeatTimeout: 1000 * time.Millisecond,
		TCPTimeout:       1000 * time.Millisecond,
		CacheSize:        500,
		SyncLimit:        100,
		StoreType:        storeType,
		StorePath:        storePath,
		Logger:           logger,
	}
}

func DefaultBadgerDir() string {
	dataDir := DefaultDataDir()
	if dataDir != "" {
		return filepath.Join(dataDir, "badger_db")
	}
	return ""
}

func DefaultDataDir() string {
	// Try to place the data folder in the user's home dir
	home := HomeDir()
	if home != "" {
		if runtime.GOOS == "darwin" {
			return filepath.Join(home, ".babble")
		} else if runtime.GOOS == "windows" {
			return filepath.Join(home, "babble")
		} else {
			return filepath.Join(home, ".babble")
		}
	}
	// As we cannot guess a stable location, return empty and handle later
	return ""
}

func HomeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}

func GetVersion() string {
	return newdag_Version
}

type Config struct {
	HeartbeatTimeout time.Duration
	TCPTimeout       time.Duration
	CacheSize        int
	SyncLimit        int
	StoreType        string
	StorePath        string
	Logger           *logrus.Logger
}

func NewConfig(heartbeat time.Duration,
	timeout time.Duration,
	cacheSize int,
	syncLimit int,
	storeType string,
	storePath string,
	logger *logrus.Logger) *Config {
	return &Config{
		HeartbeatTimeout: heartbeat,
		TCPTimeout:       timeout,
		CacheSize:        cacheSize,
		SyncLimit:        syncLimit,
		StoreType:        storeType,
		StorePath:        storePath,
		Logger:           logger,
	}
}
