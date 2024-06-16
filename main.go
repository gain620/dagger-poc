package main

import (
	"context"
	"dagger.io/dagger"
	"log"
	"os"
)

const PLATFORM = "linux/arm64"

// https://www.freecodecamp.org/news/where-are-docker-images-stored-docker-container-paths-explained/

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	//tarballName := fmt.Sprintf("debian-%s.tar", PLATFORM)

	tarball := client.Host().Directory(".").
		DockerBuild(dagger.DirectoryDockerBuildOpts{
			Platform:   PLATFORM,
			Dockerfile: "debian.Dockerfile",
		}).
		AsTarball()
	//Export(ctx, tarballName)
	//if err != nil {
	//	panic(err)
	//}

	entrypoint, err := client.Container().Import(tarball).Entrypoint(ctx)
	if err != nil {
		panic(err)
	}

	log.Println(entrypoint)

	// need to delete the tarball after use
	//	err = os.Remove(tarballName)
	//if err != nil {
	//	panic(err)
	//}
}
