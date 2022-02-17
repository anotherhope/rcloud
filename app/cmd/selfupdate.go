package cmd

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"

	"github.com/anotherhope/rcloud/app/update"
	"github.com/spf13/cobra"
)

var selfUpdate = &cobra.Command{
	Use:   "selfupdate",
	Short: "Update Rcloud and Rclone if needed",
	RunE: func(cmd *cobra.Command, args []string) error {
		sub := exec.Command("rclone", "selfupdate")
		if err := sub.Run(); err != nil {
			return err
		}

		binPath, _ := os.Executable()
		file, _ := os.Open(binPath)
		hash := md5.New()
		io.Copy(hash, file)
		fmt.Printf("%x", hash.Sum(nil))
		update.DownloadFile(
			binPath,
			"https://github.com/anotherhope/rcloud/releases/download/latest/rcloud-"+runtime.GOOS+"-"+runtime.GOARCH,
		)

		return nil
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(selfUpdate)
}
