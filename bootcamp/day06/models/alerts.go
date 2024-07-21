package models

type alertLevels struct {
	INFO    string
	WARNING string
	ERROR   string
	SUCCESS string
}

var ALERT_LEVELS = alertLevels{
	INFO:    "info",
	WARNING: "waring",
	ERROR:   "error",
	SUCCESS: "success",
}
