package rclone

func Kill() {
	for _, process := range multiton {
		process.Command.Process.Kill()
	}
}
