package rclone

// Kill running process and delete socket before exit
func Kill() {
	for _, process := range multiton {
		if process.Command.Process != nil {
			process.Command.Process.Kill()
		}
	}
}
