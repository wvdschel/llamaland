package xdg

import "os"

// $XDG_DATA_HOME defines the base directory relative to which user-specific data files should be stored. If $XDG_DATA_HOME is either not set or empty, a default equal to $HOME/.local/share should be used.
func DataHome() string {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome != "" {
		return dataHome
	}

	return os.Getenv("HOME") + "/.local/share"
}

// $XDG_CONFIG_HOME defines the base directory relative to which user-specific configuration files should be stored. If $XDG_CONFIG_HOME is either not set or empty, a default equal to $HOME/.config should be used.
func ConfigHome() string {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome != "" {
		return configHome
	}
	return os.Getenv("HOME") + "/.config"
}

// $XDG_STATE_HOME defines the base directory relative to which user-specific state files should be stored. If $XDG_STATE_HOME is either not set or empty, a default equal to $HOME/.local/state should be used.
func StateHome() string {
	stateHome := os.Getenv("XDG_STATE_HOME")
	if stateHome != "" {
		return stateHome
	}
	return os.Getenv("HOME") + "/.local/state"
}
