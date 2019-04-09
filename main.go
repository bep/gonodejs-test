package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	jsSource = `[1, 2, 3].map((n) => n + 1)`
	nodeHTTP = "http://localhost:8182"
)

func main() {
	log.Println("Starting gonodejs test …")

	nodeCmd := startNode()

	time.Sleep(300 * time.Millisecond)

	res, err := transpileViaHTTP()
	log.Println("HTTP:", res, err)

	res, err = transpileViaExec()
	log.Println("EXEC:", res, err)

	nodeCmd.Process.Signal(os.Kill)

}

func transpileViaHTTP() (string, error) {
	resp, err := http.DefaultClient.Post(nodeHTTP, "text/plain", strings.NewReader(jsSource))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func transpileViaExec() (string, error) {
	bin := filepath.FromSlash("node_modules/@babel/cli/bin/babel.js")
	var out bytes.Buffer

	cmd := exec.Command(bin, "--no-babelrc")

	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	go func() {
		defer stdin.Close()
		io.Copy(stdin, strings.NewReader(jsSource))
	}()

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
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
