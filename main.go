package main

import (
	"fmt"
	"time"

	"github.com/giantswarm/microerror"
	"sigs.k8s.io/kind/pkg/cluster"
	"sigs.k8s.io/kind/pkg/cluster/config"
	"sigs.k8s.io/kind/pkg/cluster/create"
)

func main() {
	err := mainWithError()
	if err != nil {
		panic(fmt.Sprintf("%#v\n", err))
	}
}

func mainWithError() (err error) {
	ctx := cluster.NewContext("kind")
	cfg := &config.Cluster{}

	err = ctx.Create(cfg, create.Retain(true), create.WaitForReady(5*time.Minute))
	if err != nil {
		return microerror.Mask(err)
	}

	err = ctx.CollectLogs("/home/circleci/logs")
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
