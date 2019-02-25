package file_util

import (
	"os"
)

func DeleteFile(downloadPath string) error {
	if err := os.Remove(downloadPath); err != nil {
		return err
	}
	return nil
}
