package main

import (
	"context"
	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
  	"strings"
)

func setEnvironment(ip string) {
	os.Setenv("DOCKER_HOST", ip)
}

func ListAllContainers() []Data {
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All:true})
	if err != nil {
		panic(err)
	}
	list := make([]Data, 0)
	for _, container := range containers {
		container_name := container.Names
		cname := fmt.Sprint(container_name)
		temp := Data{container.ID[:10], container.Image, container.Command, container.Status, cname[2 : len(cname)-1]}
		list = append(list, temp)
	}
	return list
}

type Data struct {
	Imgid          string
	Imgname        string
	Command        string
	Status         string
	Container_Name string
}

func main() {
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		commandtorun := r.FormValue("commandtorun")
		imagename := r.FormValue("imagename")
		if imagename != "" {
      commands := strings.Split(commandtorun, " ")
      if len(commands)==1 {
        output, err := exec.Command("docker","run",  "-d", imagename, commandtorun).CombinedOutput()
        if err != nil {
  				fmt.Println(err.Error())
  			}
  			fmt.Println("\nContainter Created: "+string(output))
      } else if len(commands)==2 {
        output, err := exec.Command("docker","run",  "-d", imagename, commands[0], commands[1]).CombinedOutput()
        if err != nil {
  				fmt.Println(err.Error())
  			}
  			fmt.Println("\nContainter Created: "+string(output))
      } else if len(commands)==3 {
        output, err := exec.Command("docker","run",  "-d", imagename, commands[0], commands[1], commands[3]).CombinedOutput()
        if err != nil {
  				fmt.Println(err.Error())
  			}
  			fmt.Println("\nContainter Created: "+string(output))
      } else {
        output, err := exec.Command("docker","run",  "-d", imagename).CombinedOutput()
        if err != nil {
  				fmt.Println(err.Error())
  			}
  			fmt.Println("\nContainter Created: "+string(output))
      }
		}
		List := ListAllContainers()
		if err := templates.ExecuteTemplate(w, "welcome-template.html", List); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/removecontainer", func(w http.ResponseWriter, r *http.Request) {
		container_name := r.FormValue("container_name")
		output, err := exec.Command("docker","rm",  "-f", container_name).CombinedOutput()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(output)+" Removed")
		List := ListAllContainers()
		if err := templates.ExecuteTemplate(w, "welcome-template.html", List); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/startcontainer", func(w http.ResponseWriter, r *http.Request) {
		container_name := r.FormValue("container_name")
		output, err := exec.Command("docker","start", container_name).CombinedOutput()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(output)+" Started")
		List := ListAllContainers()
		if err := templates.ExecuteTemplate(w, "welcome-template.html", List); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/stopcontainer", func(w http.ResponseWriter, r *http.Request) {
		container_name := r.FormValue("container_name")
		output, err := exec.Command("docker","stop", container_name).CombinedOutput()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(output)+" Stopped")
		List := ListAllContainers()
		if err := templates.ExecuteTemplate(w, "welcome-template.html", List); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":80", nil))
	fmt.Printf("DONE")
}
