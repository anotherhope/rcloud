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
		fmt.Println("selfupdate")
		binaryUrl := "https://github.com/anotherhope/rcloud/releases/download/latest/rcloud-" + runtime.GOOS + "-" + runtime.GOARCH
		hashUrl := binaryUrl + ".md5"
		fmt.Println("hash:analyse")
		if hashRemote, err := update.ReadRemote(hashUrl); err == nil {
			binPath, _ := os.Executable()
			file, _ := os.Open(binPath)
			hash := md5.New()
			io.Copy(hash, file)
			hashLocal := fmt.Sprintf("%x", hash.Sum(nil))
			fmt.Println("hash:change")
			if hashRemote != hashLocal {
				fmt.Println("Download in progress")
				update.DownloadFile(
					binPath,
					"https://github.com/anotherhope/rcloud/releases/download/latest/rcloud-"+runtime.GOOS+"-"+runtime.GOARCH,
				)
			}
		}

		fmt.Println("rclone")
		sub := exec.Command("rclone", "selfupdate")
		sub.Stdout = os.Stdout
		sub.Stdin = os.Stdin
		sub.Stderr = os.Stderr

		fmt.Println("rclone:start")
		err := sub.Start()
		if err != nil {
			return err
		}

		fmt.Println("rclone:wait")
		fmt.Println(sub.Wait())

		return nil
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(selfUpdate)
}
