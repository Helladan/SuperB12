package App

import "github.com/Helladan/guakaTrack/Controllers"

const DATABASE_HOST = "localhost"
const DATABASE_NAME = "GuakaTrack"

var Database = Controllers.NewDatabase(
	DATABASE_HOST,
	DATABASE_NAME,
)