--------------------------------------------------------------------------------
Go Workspaces: Hands-on Exercises
This set of exercises will guide you through setting up and utilizing Go workspaces, a powerful feature introduced in Go 1.18. Workspaces allow you to work with multiple modules simultaneously without needing to manually edit go.mod files for replace directives. Each module within a workspace is treated as a root module when resolving dependencies.
Exercise 1: Setting Up Your First Workspace and Main Module
Goal: To initialize a Go workspace and create a basic executable module within it.
Concepts Covered:
• What is a Go Workspace.
• The go.work file.
• Initializing a workspace with go work init.
• Adding modules to a workspace with go work use.
Steps:
1. Create a Root Directory for Your Workspace: First, create a new directory that will serve as the root for your workspace. This directory will contain the go.work file and your individual modules.
2. Initialize the Go Workspace: Inside my-go-workspace, initialize the workspace. This will create a go.work file, which tracks the modules included in your workspace.
    ◦ Verify the go.work file was created: ls
3. Create Your First Module (hello-app): Now, create a new Go module that will be our main executable application.
    ◦ This command initializes a new Go module named example.com/hello-app and creates a go.mod file within hello-app.
4. Add hello-app to the Workspace: Go back to your workspace root directory and add hello-app to the workspace. This step is crucial; it tells Go that hello-app is part of your multi-module setup.
    ◦ Inspect the go.work file now: cat go.work. It should contain a line like go 1.x and use ./hello-app.
5. Write a Simple Go Program: Inside the hello-app directory, create a file named main.go with a basic "Hello, World!" program.
    ◦ Recall that fmt is part of the Go standard library.
6. Run Your Program from the Workspace Root: From the my-go-workspace directory, run your hello-app module. Go's module system will understand to look within the workspace for the hello-app module.
    ◦ Expected Output: Hello from hello-app!

Exercise 2: Managing Local Dependencies with Workspaces
Goal: To demonstrate how Go workspaces allow you to develop and use local dependencies (other modules) without relying on replace directives in go.mod.
Concepts Covered:
• Resolving dependencies within a workspace.
• Eliminating replace directives for local development.
• The benefit of instant reflection of changes in local modules.
Steps:
1. Create a New Local Dependency Module (stringutil): Still in your my-go-workspace directory, create a new module that will serve as a utility library for our hello-app. Let's call it stringutil.
    ◦ This creates stringutil/go.mod.
2. Add a Simple Function to stringutil: Inside the stringutil directory, create a file named reverse.go and define an exported function. Remember, exported values (like functions) are defined with an uppercase identifier to be visible from other packages.
3. Add stringutil to the Workspace: Go back to the my-go-workspace root and add stringutil to your workspace.
    ◦ Verify the go.work file again. It should now list both hello-app and stringutil.
4. Modify hello-app to Use stringutil: Edit hello-app/main.go to import and use the stringutil module you just created. Crucially, you DO NOT need to add a replace directive in hello-app/go.mod. The workspace handles this resolution.
    ◦ Note: example.com/stringutil is the import path for the module, not just the package name.
5. Run hello-app and Verify Local Dependency Usage: From the my-go-workspace directory, run hello-app again.
    ◦ Expected Output:
    ◦ This confirms that hello-app is successfully using the stringutil module from your local workspace.
6. Demonstrate Live Changes (No replace Needed): Now, let's make a small change to the Reverse function in your stringutil module and see how it immediately affects hello-app without any extra steps. Edit stringutil/reverse.go:
    ◦ Go back to the my-go-workspace root and run hello-app again:
    ◦ Expected Output:
    ◦ Notice how the change in stringutil was immediately reflected in hello-app. This is a very underrated feature for local development with multiple modules.

--------------------------------------------------------------------------------