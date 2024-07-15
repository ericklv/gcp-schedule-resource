package gcp

import (
	"log"
	"os/exec"
	"scheduler/utils"
	"strings"
)

const cmd_sql = "sql instances"
const start_stop = "patch"
const restart = "restart"
const start = "--activation-policy=ALWAYS"
const stop = "--activation-policy=NEVER"
const tier = "--tier="

func getTier(p utils.Params) string{
	tier_ :=[]string{tier,utils.GetRezise(p.Instance,p.Resize)}
	return strings.Join(tier_,"")
}

func Action(p utils.Params) []string {
	switch p.Action {
	case "start":
		return []string{start_stop, start}
	case "stop":
		return []string{start_stop, stop}
	case "resize":
		return []string{start_stop, getTier(p)}
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
		log.Println("Something is wrong", cmd_l)
	} else {
		log.Println("Change applied: ", res, cmd_l)
	}
}
