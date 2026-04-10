package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, Welcome to Wails!", name)
}

// GetInfo returns some information about the app
func (a *App) GetInfo() string {
	return "Wails GUI Application - Built with Go and Web Technologies"
}

// ScriptToolResult represents the result of a script execution
type ScriptToolResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Output  string `json:"output"`
}

// ClearDNS clears the DNS cache
func (a *App) ClearDNS() ScriptToolResult {
	var cmd *exec.Cmd
	
	if runtime.GOOS == "windows" {
		cmd = exec.Command("ipconfig", "/flushdns")
	} else if runtime.GOOS == "darwin" {
		cmd = exec.Command("sudo", "dscacheutil", "-flushcache")
	} else {
		// Linux - different distributions have different services
		cmd = exec.Command("sudo", "systemctl", "restart", "nscd")
	}
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ScriptToolResult{
			Success: false,
			Message: "Failed to clear DNS cache",
			Output:  err.Error(),
		}
	}
	
	return ScriptToolResult{
		Success: true,
		Message: "DNS cache cleared successfully",
		Output:  string(output),
	}
}

// GitPullAll performs git pull on all repositories in a directory
func (a *App) GitPullAll(disk, workPath string) ScriptToolResult {
	if disk == "" || workPath == "" {
		return ScriptToolResult{
			Success: false,
			Message: "Please provide disk and work path",
			Output:  "",
		}
	}
	
	// Change to the specified disk and path
	targetPath := filepath.Join(disk+":", workPath)
	if runtime.GOOS != "windows" {
		targetPath = workPath
	}
	
	entries, err := os.ReadDir(targetPath)
	if err != nil {
		return ScriptToolResult{
			Success: false,
			Message: "Failed to read directory",
			Output:  err.Error(),
		}
	}
	
	var output strings.Builder
	
	for _, entry := range entries {
		if entry.IsDir() {
			dirPath := filepath.Join(targetPath, entry.Name())
			gitPath := filepath.Join(dirPath, ".git")
			
			// Check if .git directory exists
			if _, err := os.Stat(gitPath); err == nil {
				output.WriteString(fmt.Sprintf("[%s]\n", entry.Name()))
				
				cmd := exec.Command("git", "pull")
				cmd.Dir = dirPath
				out, err := cmd.CombinedOutput()
				if err != nil {
					output.WriteString(fmt.Sprintf("Error: %v\n", err))
				}
				output.WriteString(string(out))
				output.WriteString("\n")
			}
		}
	}
	
	return ScriptToolResult{
		Success: true,
		Message: "Git pull completed for all repositories",
		Output:  output.String(),
	}
}

// KillTaskForce forcefully kills a process by PID
func (a *App) KillTaskForce(pid string) ScriptToolResult {
	if pid == "" {
		return ScriptToolResult{
			Success: false,
			Message: "Please provide a PID",
			Output:  "",
		}
	}
	
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("taskkill", "/pid", pid, "/f")
	} else {
		cmd = exec.Command("kill", "-9", pid)
	}
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ScriptToolResult{
			Success: false,
			Message: "Failed to kill process",
			Output:  err.Error(),
		}
	}
	
	return ScriptToolResult{
		Success: true,
		Message: fmt.Sprintf("Process %s killed successfully", pid),
		Output:  string(output),
	}
}

// ViewPIDOccupied shows information about a process by PID
func (a *App) ViewPIDOccupied(pid string) ScriptToolResult {
	if pid == "" {
		return ScriptToolResult{
			Success: false,
			Message: "Please provide a PID",
			Output:  "",
		}
	}
	
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	} else {
		cmd = exec.Command("ps", "-p", pid)
	}
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ScriptToolResult{
			Success: false,
			Message: "Failed to get process information",
			Output:  err.Error(),
		}
	}
	
	// Filter output for the specific PID (on Windows)
	if runtime.GOOS == "windows" {
		lines := strings.Split(string(output), "\n")
		var filtered strings.Builder
		for _, line := range lines {
			if strings.Contains(line, pid) {
				filtered.WriteString(line + "\n")
			}
		}
		if filtered.Len() == 0 {
			return ScriptToolResult{
				Success: false,
				Message: fmt.Sprintf("No process found with PID %s", pid),
				Output:  "",
			}
		}
		output = []byte(filtered.String())
	}
	
	return ScriptToolResult{
		Success: true,
		Message: fmt.Sprintf("Process information for PID %s", pid),
		Output:  string(output),
	}
}

// ViewPortOccupied shows information about processes using a port
func (a *App) ViewPortOccupied(port string) ScriptToolResult {
	if port == "" {
		return ScriptToolResult{
			Success: false,
			Message: "Please provide a port number",
			Output:  "",
		}
	}
	
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("netstat", "-aon")
	} else {
		cmd = exec.Command("lsof", "-i", ":"+port)
	}
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ScriptToolResult{
			Success: false,
			Message: "Failed to get port information",
			Output:  err.Error(),
		}
	}
	
	// Filter output for the specific port
	lines := strings.Split(string(output), "\n")
	var filtered strings.Builder
	for _, line := range lines {
		if strings.Contains(line, ":"+port) || strings.Contains(line, " "+port+" ") {
			filtered.WriteString(line + "\n")
		}
	}
	
	if filtered.Len() == 0 {
		return ScriptToolResult{
			Success: true,
			Message: fmt.Sprintf("No processes found using port %s", port),
			Output:  "",
		}
	}
	
	return ScriptToolResult{
		Success: true,
		Message: fmt.Sprintf("Processes using port %s", port),
		Output:  filtered.String(),
	}
}

// RunLaravelQueue runs Laravel queue worker (Windows batch simulation)
func (a *App) RunLaravelQueue(disk, workPath string) ScriptToolResult {
	if disk == "" || workPath == "" {
		return ScriptToolResult{
			Success: false,
			Message: "Please provide disk and work path",
			Output:  "",
		}
	}
	
	targetPath := filepath.Join(disk+":", workPath, "artisan")
	if runtime.GOOS != "windows" {
		targetPath = filepath.Join(workPath, "artisan")
	}
	
	// Check if artisan file exists
	if _, err := os.Stat(targetPath); err != nil {
		return ScriptToolResult{
			Success: false,
			Message: "Laravel artisan not found at specified path",
			Output:  err.Error(),
		}
	}
	
	// Note: This is a simplified version. The original batch runs in a loop.
	cmd := exec.Command("php", targetPath, "schedule:run")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ScriptToolResult{
			Success: false,
			Message: "Failed to run Laravel schedule",
			Output:  err.Error(),
		}
	}
	
	return ScriptToolResult{
		Success: true,
		Message: "Laravel schedule executed successfully",
		Output:  string(output),
	}
}

// RunSkeletonDocker runs the Hyperf skeleton Docker container
func (a *App) RunSkeletonDocker() ScriptToolResult {
	// This is an informational function as running Docker interactively
	// from a GUI app is complex. We provide the command instead.
	dockerCommand := "docker run -v d:/code:/skeleton -p 9501:9501 -it --entrypoint /bin/sh hyperf/hyperf:7.4-alpine-v3.9-cli"
	
	return ScriptToolResult{
		Success: true,
		Message: "Docker command ready to execute",
		Output:  dockerCommand,
	}
}

// GetSupervisorCommands returns supervisor control commands
func (a *App) GetSupervisorCommands() ScriptToolResult {
	commands := `sudo supervisord -c /etc/supervisord.d/supervisord.conf

# Start hyperf application
supervisorctl start hyperf
# Restart hyperf application
supervisorctl restart hyperf
# Stop hyperf application
supervisorctl stop hyperf
# View status of all managed projects
supervisorctl status
# Reload configuration files
supervisorctl update
# Restart all programs
supervisorctl reload`
	
	return ScriptToolResult{
		Success: true,
		Message: "Supervisor commands",
		Output:  commands,
	}
}

// GetFiddlerScript returns the Fiddler image download script
func (a *App) GetFiddlerScript() ScriptToolResult {
	script := `static function OnDone(oSession: Session) {

    // Check Content-Type
    if (oSession.ResponseHeaders["Content-Type"] != null || oSession.ResponseHeaders["content-type"] != null) {
        // Avoid non-standard headers
        var contentType = oSession.ResponseHeaders["Content-Type"];
        if (String.IsNullOrEmpty(contentType))
            contentType = oSession.ResponseHeaders["content-type"];

        // Check if request is an image
        if (contentType.Contains("image")) {
            // Determine filename (for saving)
            var fileName = "";
            var fileIndex = oSession.RequestHeaders.RequestPath.LastIndexOf("/");
            if (fileIndex > 0)
                fileName = oSession.RequestHeaders.RequestPath.Substring(fileIndex + 1);

            // If filename is invalid (contains invalid characters)
            if (fileName.IndexOf('?') > 0 || fileName.IndexOf('&'))
                fileName = String.Empty;

            // If filename is Null, create a new filename (Guid)
            if (String.IsNullOrEmpty(fileName)) {
                fileName = Guid.NewGuid().ToString();
                var extName = contentType.Replace("image/", "");
                fileName = fileName + "." + extName;
            }

            // Don't save images that are too small, such as placeholder images (adjust as needed)
            if (oSession.ResponseBody.Length > 100) {
                // Specify save directory
                var saveDir = "e:\\\\Temp\\\\";
                // Create directory if it doesn't exist
                if (!System.IO.Directory.Exists(saveDir))
                    System.IO.Directory.CreateDirectory(saveDir);

                // Save response body
                oSession.SaveResponseBody(saveDir + fileName);
                // Write log
                // FiddlerObject.log("[File Saved]:" + fileName)
            }
        }
    }
}`
	
	return ScriptToolResult{
		Success: true,
		Message: "Fiddler script retrieved",
		Output:  script,
	}
}

// GetSystemInfo returns system information
func (a *App) GetSystemInfo() ScriptToolResult {
	info := fmt.Sprintf(`Operating System: %s
Architecture: %s
Go Version: %s
Number of CPUs: %d`,
		runtime.GOOS,
		runtime.GOARCH,
		runtime.Version(),
		runtime.NumCPU(),
	)
	
	return ScriptToolResult{
		Success: true,
		Message: "System information",
		Output:  info,
	}
}
