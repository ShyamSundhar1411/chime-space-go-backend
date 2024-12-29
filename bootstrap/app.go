package bootstrap

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDBConnection(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
