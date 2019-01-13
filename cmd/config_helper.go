package cmd

import "os"

var defaultConfigPath string = "/etc/kaz.cfg"

// Default config data
// If not specified and config is not found. generate file by it.
var defaultConfigData string = `
########################################
# KAZ cofigs                           #
########################################

# ------------------
# Base directory settings
# ------------------
bin_dir = "/opt/kaz/bin"
log_dir = "/var/log/kaz"
app_dir = "/var/opt/kaz"
`

// Create config file from data
func createConfigFile(dest string) error {
	if cfg, err := os.Create(dest); err != nil {
		return err
	} else if err := cfg.Chmod(0644); err != nil {
		return err
	} else if _, err := cfg.Write([]byte(defaultConfigData)); err != nil {
		return err
	}
	return nil
}
