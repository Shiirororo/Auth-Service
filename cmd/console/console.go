package console

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func Console() {
	// Fixed log directory
	logDir := `C:\Users\Admin\go\src\github.com\auth_service\test\log`

	// Ensure directory exists
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	// Create log file with timestamp
	logFileName := fmt.Sprintf(
		"app_log_%s.txt",
		time.Now().Format("20060102_150405"),
	)
	logPath := logDir + `\` + logFileName

	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	defer logFile.Close()

	fmt.Printf("Logging to: %s\n", logPath)

	// Open new PowerShell window and tail the log
	cmdStr := fmt.Sprintf(
		`
		start powershell -NoExit -NoProfile -Command "
		Clear-Host; 
		Set-Location C:\; 
		Write-Host '📄 LOG VIEWER (READ-ONLY)' -ForegroundColor Cyan;
		Write-Host 'Press Ctrl+C to close' -ForegroundColor DarkGray;

	
		[Console]::TreatControlCAsInput = $true;
		
		$null = Register-EngineEvent PowerShell.Exiting -Action {};
		Get-Content -Path '%s' -Wait"
		
		`,
		logPath,
	)

	spawnCmd := exec.Command("cmd", "/c", cmdStr)
	if err := spawnCmd.Run(); err != nil {
		log.Fatalf("Failed to spawn terminal: %v", err)
	}

	fmt.Println("New terminal opened. Streaming logs...")
	fmt.Println("Press Ctrl+C to stop.")

	// Handle Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Simulate log writing
	go func() {
		counter := 1
		for {
			logMsg := fmt.Sprintf(
				"[%s] INFO: Log message #%d\n",
				time.Now().Format("15:04:05"),
				counter,
			)
			if _, err := logFile.WriteString(logMsg); err != nil {
				fmt.Printf("Error writing to log: %v\n", err)
				return
			}
			logFile.Sync()
			counter++
			time.Sleep(1 * time.Second)
		}
	}()

	<-sigChan
	fmt.Println("\nStopping logger...")
}
