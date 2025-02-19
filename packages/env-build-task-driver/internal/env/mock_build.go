package env

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/docker/docker/client"
	docker "github.com/fsouza/go-dockerclient"
	"go.opentelemetry.io/otel"
)

func MockBuild(envID, buildID string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	tracer := otel.Tracer("test")

	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	legacyClient, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	apiSecret := "SUPER_SECR3T_4PI_K3Y"

	writer := NewWriter(envID, buildID, apiSecret)

	e := Env{
		BuildID:               buildID,
		EnvID:                 envID,
		EnvsDiskPath:          "/mnt/disks/fc-envs/v1",
		VCpuCount:             2,
		MemoryMB:              512,
		DockerContextsPath:    "/mnt/disks/docker-contexts/v1",
		DockerRegistry:        "us-central1-docker.pkg.dev/e2b-prod/custom-environments",
		StartCmd:              "",
		KernelsDir:            "/fc-kernels",
		KernelMountDir:        "/fc-vm",
		KernelName:            "vmlinux.bin",
		KernelVersion:         "vmlinux-5.10.186",
		DiskSizeMB:            512,
		FirecrackerBinaryPath: "/usr/bin/firecracker",
		EnvdPath:              "/fc-vm/envd",
		ContextFileName:       "context.tar.gz",
		BuildLogsWriter:       writer,
	}

	err = e.Build(ctx, tracer, client, legacyClient)
	if err != nil {
		panic(err)
	}

	err = writer.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error closing writer: %v\n", err)
	}

	<-writer.Done
}
