package rclone

func GetStatus(s string) string {
	if process, ok := multiton[s]; ok {
		return process.Type
	}

	return "idle"
}
