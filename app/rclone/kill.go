package rclone

// Kill running process and delete socket before exit
func Kill() {
	mu.Lock()
	for _, process := range multiton {
		if process != nil && process.Command != nil {
			if process.Command.Process != nil {
				process.Command.Process.Kill()
			}
		}
	}
	mu.Unlock()
}
