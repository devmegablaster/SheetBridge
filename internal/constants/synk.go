package constants

type SynkConstants struct {
	STATUS_INIT    string
	STATUS_HEALTHY string
	STATUS_ERROR   string
	STATUS_PAUSED  string
}

var Synk = SynkConstants{
	STATUS_INIT:    "INIT",
	STATUS_HEALTHY: "HEALTHY",
	STATUS_ERROR:   "ERROR",
	STATUS_PAUSED:  "PAUSED",
}
