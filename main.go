package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var trashDir string

func main() {
	rootCmd := &cobra.Command{
		Use:   "trash",
		Short: "Move files to the trash directory",
		Run:   trashFile,
	}

	rootCmd.Flags().StringVarP(&trashDir, "trash-dir", "t", filepath.Join(os.Getenv("HOME"), ".trash"), "Directory to store trashed files")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func trashFile(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Please provide a file or directory to delete")
		os.Exit(1)
	}

	for _, arg := range args {
		if err := moveToTrash(arg); err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting %s: %v\n", arg, err)
		} else {
			fmt.Printf("%s moved to trash\n", arg)
		}
	}
}

func moveToTrash(path string) error {
	// Create the trash directory if it doesn't exist
	if _, err := os.Stat(trashDir); os.IsNotExist(err) {
		if err := os.MkdirAll(trashDir, 0755); err != nil {
			return err
		}
	}

	// Move the file or directory to the trash directory
	trashPath := filepath.Join(trashDir, filepath.Base(path))
	return os.Rename(path, trashPath)
}
