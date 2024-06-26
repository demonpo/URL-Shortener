package main

import (
	"ariga.io/atlas-provider-gorm/gormschema"
	"fmt"
	analyticsEntitiesInfra "goHexBoilerplate/src/modules/analytics/infra/entities"
	shortenerEntitiesInfra "goHexBoilerplate/src/modules/shortener/infra/entities"
	userEntitiesInfra "goHexBoilerplate/src/modules/user/infra/entities"
	"io"
	"os"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(&userEntitiesInfra.User{}, shortenerEntitiesInfra.Shortener{}, analyticsEntitiesInfra.Click{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
