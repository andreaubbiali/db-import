package main

import (
	"aubbiali/db-import/config"
	"log"
)

func main() {
	log.Println("Start service")

	config.ReadConfig()

	log.Println("CONFIG LETTO")
	/*
		db := database.SetupDatabase()
		dd, _ := db.DB.DB()
		defer dd.Close()

		file, err := os.ReadFile("example.csv")
		utility.CheckError(err)

		db.Import(file)
	*/
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
