package controllers

import (
	"context"
	"html/template"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func ListAllContainers(w http.ResponseWriter, r *http.Request) {
	tmplt, _ := template.ParseGlob("views/index.html")

	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	ctrdata := map[string]interface{}{
		"CtrID":   "",
		"CtrImg":  "",
		"CtrStat": "",
		"CtrName": "",
	}

	for _, ctr := range containers {
		ctrdata["CtrID"] = ctr.ID
		ctrdata["CtrImg"] = ctr.Image
		ctrdata["CtrStat"] = ctr.Status
		ctrdata["CtrName"] = ctr.Names

	}

	tmplt.Execute(w, ctrdata)
}
