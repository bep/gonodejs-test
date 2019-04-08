package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	log.Println("Starting gonodejs test …")

	nodeCmd := startNode()

	time.Sleep(300 * time.Millisecond)

	transpileViaHTTP()

	time.Sleep(3 * time.Second)

	nodeCmd.Process.Signal(os.Kill)

}

const (
	jsSource = `[1, 2, 3].map((n) => n + 1)`
	nodeHTTP = "http://localhost:8182"
)

func transpileViaHTTP() {
	resp, err := http.DefaultClient.Post(nodeHTTP, "text/plain", strings.NewReader(jsSource))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Got", string(b))
}

func startNode() *exec.Cmd {
	cmd := exec.Command("node", "index.js")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	return cmd
}
