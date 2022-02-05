package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/anotherhope/rcloud/app/repositories"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "daemon",
		Short: "Daemon management",
	}

	install := &cobra.Command{
		Args:  cobra.ExactArgs(0),
		Use:   "install",
		Short: "Install daemon",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO
			fmt.Println("install")
			return nil
		},
	}

	remove := &cobra.Command{
		Args:  cobra.ExactArgs(0),
		Use:   "remove",
		Short: "Remove daemon",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO
			fmt.Println("remove")
			return nil
		},
	}

	standalone := &cobra.Command{
		Args:  cobra.ExactArgs(0),
		Use:   "standalone",
		Short: "Run daemon in standalone mode",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO
			watcher, err := fsnotify.NewWatcher()
			done := make(chan os.Signal, 1)
			signal.Notify(done, os.Interrupt, syscall.SIGTERM)

			if err != nil {
				return err
			}
			defer watcher.Close()

			go func() {
				for {
					select {
					case event, ok := <-watcher.Events:
						if !ok {
							return
						}
						log.Println("event:", event)
						if event.Op&fsnotify.Write == fsnotify.Write {
							log.Println("modified file:", event.Name)
						}
					case err, ok := <-watcher.Errors:
						if !ok {
							return
						}
						log.Println("error:", err)
					}
				}
			}()

			for _, dir := range repositories.List() {
				fmt.Println(dir.Source)
				watcher.Add(dir.Source)
			}

			<-done

			fmt.Println("standalone")
			return nil
		},
	}

	rootCmd.AddCommand(cmd)
	cmd.AddCommand(install)
	cmd.AddCommand(remove)
	cmd.AddCommand(standalone)
}
