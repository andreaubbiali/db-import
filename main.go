package main

import (
	"aubbiali/db-import/database"
	"log"
	"os"
)

func main() {
	log.Println("Start service")

	db := database.SetupDatabase()
	dd, _ := db.DB.DB()
	defer dd.Close()

	file, err := os.ReadFile("treeLst.csv")
	check(err)

	db.Import(file)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// func startContainer() {
// 	cli, err := client.NewEnvClient()
// 	if err != nil {
// 		fmt.Println("Error creating client")
// 		panic(err)
// 	}

// 	hostBinding := nat.PortBinding{
// 		HostIP:   "0.0.0.0",
// 		HostPort: "8000",
// 	}
// 	containerPort, err := nat.NewPort("tcp", "80")
// 	if err != nil {
// 		panic("Unable to get the port")
// 	}

// 	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}

// 	cont, err := cli.ContainerCreate(
// 		context.Background(),
// 		&container.Config{
// 			Image: "ubuntu",
// 		},
// 		&container.HostConfig{
// 			PortBindings: portBinding,
// 		}, nil, nil, "prova")
// 	if err != nil {
// 		panic(err)
// 	}

// 	cli.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
// 	fmt.Printf("Container %s is started", cont.ID)

// }
