package procfinder

import (
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func FindProcesses(pathToBinary string) ([]int32, error) {
	rootPath := "/"
	rootFs := os.DirFS(rootPath)
	matches, err := fs.Glob(rootFs, "proc/*/exe")
	if err != nil {
		return nil, err
	}
	pids := make([]int32, 0)
	for _, match := range matches {
		pathMatch := filepath.Join(rootPath, match)
		linkTarget, err := filepath.EvalSymlinks(pathMatch)
		if err != nil {
			// probably didn't have perm?
			//fmt.Printf("Encountered error while resolving %q: %v\n", pathMatch, err)
			continue
		}
		pidStr := strings.Split(match, "/")[1]
		if linkTarget == pathToBinary {
			// Ignoring errors, TODO - confirm pids fit into an int32
			intPid, _ := strconv.ParseInt(pidStr, 10, 32)
			pids = append(pids, int32(intPid))
		}
	}

	return pids, nil
}
