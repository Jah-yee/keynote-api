package constants

import "github.com/google/uuid"

const (
	KindUsers     = "users"
	KindTenant    = "tenant"
	KindWorkspace = "workspace"
	KindProject   = "project"

	StateActive   = "active"
	StateDeleted  = "deleted"
	StateArchived = "archived"
)

var StateIDs = map[string]map[string]uuid.UUID{
	KindUsers: {
		StateActive:   uuid.MustParse("664b2db7-b6fa-4bdc-8be7-4b47272d5f9e"),
		StateDeleted:  uuid.MustParse("44ea4a85-e250-462c-bfb4-dc14efd75713"),
		StateArchived: uuid.MustParse("537a902b-6b8b-4c11-8ec8-805ffcaed75c"),
	},
	KindTenant: {
		StateActive:   uuid.MustParse("7f7e76af-6359-4199-a2f3-1a3302f80112"),
		StateDeleted:  uuid.MustParse("4a60dc6f-5f6d-44d7-bb6a-e5e9cbd68ce7"),
		StateArchived: uuid.MustParse("f4686af6-0523-4007-a295-bada3c3aaafd"),
	},
	KindWorkspace: {
		StateActive:   uuid.MustParse("7ccd00b8-e708-45f6-b3f5-955ecc597694"),
		StateDeleted:  uuid.MustParse("8af31ae3-453a-4c1c-af74-7c486a6e8702"),
		StateArchived: uuid.MustParse("1ba4db36-6270-4c1c-b807-02394403bfbc"),
	},
	KindProject: {
		StateActive:   uuid.MustParse("b7d042e8-4739-46a3-bb6f-61b9000188b9"),
		StateDeleted:  uuid.MustParse("c339575c-2c6f-4565-ba44-fdbfda91a1bb"),
		StateArchived: uuid.MustParse("7cd9c2f4-0453-4d98-af58-9a7653236ecf"),
	},
}