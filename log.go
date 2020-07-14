package main

import (
	"fmt"
	"github.com/chanyipiaomiao/hlog"
	"time"
)

func InitLogger(logPath string) error {
	_, err := hlog.New(&hlog.Option{
		LogPath:                logPath,
		LogType:                hlog.JSON,
		LogLevel:               hlog.DebugLevel,
		MaxAge:                 15 * 24 * time.Hour,
		RotationTime:           24 * time.Hour,
		JSONPrettyPrint:        true,
		IsEnableRecordFileInfo: true,
	})

	if err != nil {
		return fmt.Errorf("init logger error: %s", err)
	}

	return nil
}
