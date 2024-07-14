package gcp

import (
	"log"
	"os/exec"
	"strings"
)

const cmd_sql = "sql instances"
const start_stop = "patch"
const restart = "restart"
const start = "--activation-policy=ALWAYS"
const stop = "--activation-policy=NEVER"
const tier = "--tier="

func Action(path string) []string {
	switch path {
	case "start":
		return []string{start_stop, start}
	case "stop":
		return []string{start_stop, stop}
	case "resize":
		return []string{start_stop, tier}
	case "restart":
		return []string{restart, ""}
	}
	return nil
}

func CallGCP(values []string) {
	cmd_ := []string{"gcloud", cmd_sql, values[0], values[2], values[1]}
	cmd_l := strings.Join(cmd_, " ")

	log.Println(cmd_l)

	cmd := exec.Command("/bin/sh", "-c", cmd_l)

	res, err := cmd.Output()

	if err != nil {
		log.Println("Something is wrong")
	} else {
		log.Println("Change applied: ", res)
	}
}
