## win-rce-checker

A simple and benign Windows RCE checker written in Go. Red teamers can drop and execute it to validate remote code execution on a Windows host.    

I created this tool to assist in RCE validation during a security assessment at work. While there are many ways to verify RCE, this is built mainly for having a consistent and purpose-built binary that allows for simplicity, easy retrieval and standardisation.


### What is does (by default)

* Writes an indicator file to `C:\Users\Public\exploit_success.txt` with the contents `Exploited`.
* Launches `calc.exe` as a visible proof of execution.
* Prints minimal status to stdout. Ignores all CLI args.

```go
var (
    path    string = "C:\\Users\\Public\\exploit_success.txt"
    content string = "Exploited"
    cmd     string = "calc.exe"
)
```

* No network, no persistence, no privilege escalation, no lateral movement. Purely a local execution signal.

### Quick start

* Option 1: Use the prebuilt binary in `bin/win-rce-checker.exe` (if you trust me).
* Option 2: Build from source.

Then run on the target (locally or via your RCE vector).

### To-do

* Add a simple dropper to this repo for remote retrieval and execution.
