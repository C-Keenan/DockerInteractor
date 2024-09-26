package controllers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func ListAllContainers(w http.ResponseWriter, r *http.Request) {
	tmplt, err := template.ParseGlob("views/index.html")
	if err != nil {
		panic(err)
	}

	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	CtrData := map[string]interface{}{
		"CtrID":   "ID",
		"CtrImg":  "Img",
		"CtrStat": "Stat",
		"CtrName": "Name",
	}

	for _, ctr := range containers {
		CtrData["CtrID"] = ctr.ID
		CtrData["CtrImg"] = ctr.Image
		CtrData["CtrStat"] = ctr.Status
		CtrData["CtrName"] = ctr.Names

	}

	fmt.Printf("%+v\n", CtrData)
	tmplt.Execute(w, CtrData)
}
