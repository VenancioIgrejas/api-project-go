package logging

import (
	"fmt"
	"time"

	"github.com/VenancioIgrejas/api-project-go/pkg/setting"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.Main.RuntimeRootPath, setting.AppSetting.Logger.SavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.Logger.SaveName,
		time.Now().Format(setting.AppSetting.Logger.TimeFormat),
		setting.AppSetting.Logger.FileExt,
	)
}
