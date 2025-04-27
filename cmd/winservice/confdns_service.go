package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/kardianos/service"
)

type program struct {
	cmd *exec.Cmd
}

// Start is called when the service is started
func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

// run contains the service logic
func (p *program) run() {
	// Define the command to run
	cmdPath := "E:\\programs\\bingo-hi\\ConfDNS\\dnsclient.exe"

	p.cmd = exec.Command(cmdPath)
	p.cmd.Stdout = os.Stdout // Redirect standard output to console
	p.cmd.Stderr = os.Stderr // Redirect standard error to console

	// Start the command
	err := p.cmd.Start()
	if err != nil {
		log.Fatalf("Failed to start command: %v", err)
	}

	// Log the successful start of the command
	log.Println("dnsclient.exe has started successfully.")
	err = p.cmd.Wait()
	if err != nil {
		log.Printf("Command exited with error: %v", err)
	}
}

// Stop is called when the service is stopped
func (p *program) Stop(s service.Service) error {
	// Attempt to stop the command if it's running
	if p.cmd != nil && p.cmd.Process != nil {
		log.Println("Stopping dnsclient.exe...")
		err := p.cmd.Process.Kill()
		if err != nil {
			log.Printf("Failed to stop dnsclient.exe: %v", err)
			return err
		}
		log.Println("dnsclient.exe stopped successfully.")
	}
	return nil
}

func main() {
	// Define service configuration
	svcConfig := &service.Config{
		Name:        "ConfDNS",                 // Name of the service
		DisplayName: "My DNS Client",           // Display name of the service
		Description: "This is ConfDNS Client.", // Description of the service
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Parse command line arguments
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "/install":
			err = s.Install()
			if err != nil {
				log.Fatalf("Failed to install service: %v", err)
			}
			log.Println("Service installed successfully.")
			return
		case "/uninstall":
			err = s.Uninstall()
			if err != nil {
				log.Fatalf("Failed to uninstall service: %v", err)
			}
			log.Println("Service uninstalled successfully.")
			return
		}
	}

	// Run the service
	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
