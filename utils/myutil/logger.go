package myutil

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func NewExecutionLogger() (*log.Logger, func(), error) {
	rootDir, err := findProjectRoot()
	if err != nil {
		return nil, nil, err
	}

	fileName := fmt.Sprintf("execution-log-%s.txt", time.Now().Format("2006-01-02"))
	logFilePath := filepath.Join(rootDir, fileName)

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, nil, err
	}

	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	closeFn := func() {
		_ = file.Close()
	}

	return logger, closeFn, nil
}

func RunStep(logger *log.Logger, stepName string, step func()) {
	stepStart := time.Now()
	logger.Printf("step started: %s", stepName)
	step()
	logger.Printf("step completed: %s | duration: %s", stepName, FormatDuration(time.Since(stepStart)))
}

func FormatDuration(duration time.Duration) string {
	return fmt.Sprintf("%.3f ms", float64(duration.Microseconds())/1000)
}

func findProjectRoot() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return "", fmt.Errorf("go.mod not found from current directory")
		}

		currentDir = parentDir
	}
}
