--------------------------------------------------------------------------------
Go Build Commands: Hands-on Exercises
This set of exercises focuses on the go build command, a fundamental tool for compiling your Go source code into executable binaries. Building static binaries is one of Go's best features, enabling efficient code shipping.
Exercise 1: Basic Build and Execution
Goal: To compile a simple Go program into an executable binary using go build and then run it.
Concepts Covered:
• go build command.
• Default binary naming.
• Executing a Go binary.
Steps:
1. Create a New Directory and Module: Create a fresh directory for this exercise and initialize a Go module inside it.
2. Create a Simple Go Program: Create a file named main.go with the following content:
3. Run go build: Execute the go build command. This will compile your Go program.
    ◦ Observation: This command should produce a binary executable with the name of your module (e.g., myprogram on Linux/macOS, myprogram.exe on Windows) in the current directory.
4. Execute the Binary: Now, simply execute the generated binary.
    ◦ On Linux/macOS: ./<app-name>
    ◦ On Windows:
    ◦ Expected Output: Hello, Go Build!

--------------------------------------------------------------------------------
Exercise 2: Specifying Output Binary Name
Goal: To use the -o flag with go build to specify a custom name for the output executable.
Concepts Covered:
• go build -o flag.
Steps:
1. Ensure you are in the go-build-exercise directory from Exercise 1.
2. Run go build with the -o flag: Compile your program again, but this time specify a different output name, for example, mycustomapp.
    ◦ Observation: A new executable named mycustomapp (or mycustomapp.exe on Windows) should be created in the current directory.
3. Execute the New Binary: Run the binary with the custom name.
    ◦ On Linux/macOS: ./<app-name>
    ◦ On Windows:
    ◦ Expected Output: Hello, Go Build!

--------------------------------------------------------------------------------
Exercise 3: Cross-Compilation with  and 
Goal: To build Go programs for different operating systems and processor architectures. This is a powerful feature of Go.
Concepts Covered:
• GOOS environment variable (Operating System).
• GOARCH environment variable (Processor Architecture).
• Cross-compilation capabilities of Go.
• go tool dist list to list supported architectures.
Steps:
1. Ensure you are in the go-build-exercise directory from Exercise 1 and 2.
2. List Supported Architectures (Optional but Recommended): You can see all the supported combinations of operating systems and architectures by running: go tool dist list
    ◦ Observation: This command will show a comprehensive list like darwin/arm64, linux/amd64, windows/amd64, etc.
3. Build for a Different Operating System (e.g., Windows from macOS/Linux): Set the GOOS and GOARCH environment variables before running go build. For example, to build a Windows executable from a macOS or Linux machine:
    ◦ On Linux/macOS: GOOS=windows GOARCH=amd64 go build -o mywindowsapp.exe
    ◦ On Windows (using PowerShell):
    ◦ Observation: A new executable (mywindowsapp.exe or mylinuxapp) will be created in your current directory, ready to be run on the target system.
4. Build for Another Combination (e.g., Linux from Windows/macOS): Try building for a different target, such as Linux.
    ◦ On Linux/macOS:
    ◦ On Windows (using PowerShell):
    ◦ Observation: You will have a Linux executable (mylinuxapp) that can be deployed to a Linux server.

--------------------------------------------------------------------------------
Exercise 4: Producing Statically Linked Binaries with 
Goal: To understand and use CGO_ENABLED to create statically linked binaries, which are useful for deployment, especially in containerized environments.
Concepts Covered:
• CGO_ENABLED environment variable.
• CGO (calling C code from Go).
• Statically linked binaries and their benefits (no external dependencies).
• Relevance for Docker containers.
Steps:
1. Ensure you are in the go-build-exercise directory.
2. Understand CGO_ENABLED: CGO_ENABLED is an environment variable that configures CGO, which allows Go to call C code. Setting CGO_ENABLED=0 during a build ensures that the resulting binary is statically linked, meaning it includes all necessary libraries and works without any external dependencies. This is very helpful for running Go binaries in environments like Docker containers, where you want minimal external dependencies.
3. Build a Statically Linked Binary: By default, Go often produces statically linked binaries, but explicitly setting CGO_ENABLED=0 ensures it, especially if your project might otherwise link to C libraries.
    ◦ On Linux/macOS: CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp_static
    ◦ On Windows (using PowerShell):
    ◦ Observation: You've now built a self-contained executable. While the size difference might not be immediately obvious for a simple "Hello World" program, for applications with C dependencies, this approach creates a highly portable binary.

--------------------------------------------------------------------------------