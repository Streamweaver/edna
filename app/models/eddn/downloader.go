package eddn

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// DownloadSources from eddn if they are more than 24 hours old.
func DownloadSources() {
	names := []string{"systems", "stations", "commodities"}
	for _, val := range names {
		url := "http://eddb.io/archive/" + val + ".json"
		dst := "data/" + val + ".json"
		if isStale(dst) != true {
			fmt.Println("Skipping", dst, "- file is up to date.")
			continue
		}
		err := downloadData(url, dst)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// isStale checks if the named file is over 24 hours old.
func isStale(name string) bool {
	fi, err := os.Stat(name)
	if err == nil {
		age := time.Since(fi.ModTime())
		if age.Hours() < 23 {
			return false
		}
	}

	return true
}

func downloadData(url string, dst string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	n, err := io.Copy(f, res.Body)
	if err != nil {
		return err
	}

	fmt.Println(n, "bytes downloaded to", dst)
	return nil
}
