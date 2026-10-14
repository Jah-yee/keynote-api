package project

import (
	"github.com/caneel7/keynote-api/internal/ent/gen"
	"github.com/caneel7/keynote-api/internal/features/project/create"
	"github.com/danielgtaylor/huma/v2"
)

func Mount(api huma.API,db *gen.Client) {

	/*
	* @path: "/project"
	*/
	create.CreateProjectRoute(api,create.NewService(db))
	
}
