package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"

	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(VersionCmd)
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the versions associated with this tool.",
	Long:  `Print the versions associated with this tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("docker-apps:", "v0.0.1")

		docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			log.Panic(err)
		}

		serverVersion, err := docker.ServerVersion(context.TODO())
		if err != nil {
			log.Panic(err)
		}

		fmt.Println("Docker Engine:", serverVersion.Version)
		fmt.Println("Docker API:", docker.ClientVersion())

		execCmd := exec.Command("docker", "compose", "version", "--short")
		if errors.Is(execCmd.Err, exec.ErrDot) {
			execCmd.Err = nil
		}
		output, err := execCmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Docker Compose: %s", output)
	},
}
