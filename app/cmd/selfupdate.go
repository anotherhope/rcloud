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
		binaryUrl := "https://github.com/anotherhope/rcloud/releases/download/latest/rcloud-" + runtime.GOOS + "-" + runtime.GOARCH
		hashUrl := binaryUrl + ".md5"
		if hashRemote, err := update.ReadRemote(hashUrl); err == nil {
			binPath, _ := os.Executable()
			file, _ := os.Open(binPath)
			hash := md5.New()
			io.Copy(hash, file)
			hashLocal := fmt.Sprintf("%x", hash.Sum(nil))
			if hashRemote != hashLocal {
				fmt.Println("Download in progress")
				update.DownloadFile(
					binPath,
					"https://github.com/anotherhope/rcloud/releases/download/latest/rcloud-"+runtime.GOOS+"-"+runtime.GOARCH,
				)
			}
		}

		sub := exec.Command("rclone", "selfupdate")
		sub.Stdout = os.Stdout
		sub.Stdin = os.Stdin
		sub.Stderr = os.Stderr
		return nil
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(selfUpdate)
}
