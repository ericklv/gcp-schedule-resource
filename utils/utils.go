package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Params struct {
	Action   string `param:"action"`
	Instance string `param:"inst_name"`
	Resize   string `param:"resize"`
}

type Machine struct {
	Name string `json:"name"`
	Up   string `json:"up"`
	Down string `json:"down"`
}

type Machines struct {
	Machines []Machine `json:"machines"`
}

func S200(msg string) Response {
	return Response{200, msg}
}

func S4xx(msg string) Response {
	return Response{400, msg}
}

func S5xx(msg string) Response {
	return Response{500, msg}
}

func ReadMachines(path string) (Machines, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Println("Load JSON failed", err)
		return Machines{}, err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Println("Load JSON failed", err)
		return Machines{}, err
	}

	data := Machines{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Println("failed to unmarshal json file, error:", err)
		return Machines{}, err
	}
	return data, nil
}

func GetRezise(instance string, resize string) string {

	data, err := ReadMachines("./resize.json")

	if err != nil {
		return ""
	}

	for _, mch := range data.Machines {
		if mch.Name == instance {
			if resize == "up" {
				return mch.Up
			}
			if resize == "down" {
				return mch.Down
			}
		}
	}
	return ""
}
