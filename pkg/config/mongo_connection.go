package config

type MongoConnection struct {
	Uri string `long:"uri" required:"true" description:"URI to MongoDB"`
	Db  string `long:"db" required:"true" description:"Database name"`
	Col string `long:"col" required:"true" description:"Collection name"`
}
