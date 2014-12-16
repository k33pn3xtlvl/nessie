package nessie

// loginResp is the internal response to login attemps.
type loginResp struct {
	Token string `json:"token"`
}

// ServerProperties is the structure returned by the ServerProperties() method.
type ServerProperties struct {
	Token           string `json:"token"`
	NessusType      string `json:"nessus_type"`
	NessusUIVersion string `json:"nessus_ui_version"`
	ServerVersion   string `json:"server_version"`
	Feed            string `json:"feed"`
	Enterprise      bool   `json:"enterprise"`
	LoadedPluginSet string `json:"loaded_plugin_set"`
	ServerUUID      string `json:"server_uuid"`
	Expiration      int64  `json:"expiration"`
	Notifications   []struct {
		Type string `json:"type"`
		Msg  string `json:"message"`
	} `json:"notifications"`
	ExpirationTime int64 `json:"expiration_time"`
	Capabilities   struct {
		MultiScanner      bool `json:"multi_scanner"`
		ReportEmailConfig bool `json:"report_email_config"`
	} `json:"capabilities"`
	PluginSet       string `json:"plugin_set"`
	IdleTImeout     int64  `json:"idle_timeout"`
	ScannerBoottime int64  `json:"scanner_boottime"`
	LoginBanner     bool   `json:"login_banner"`
}

// ServerStatus is the stucture returned  by the ServerStatus() method.
type ServerStatus struct {
	Status             string `json:"status"`
	Progress           int64  `json:"progress"`
	MustDestroySession bool
}
