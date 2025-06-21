package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/rpowelson12/Easypass/internal/config"
	"github.com/rpowelson12/Easypass/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	cfg.KEY = ensureKeyFromEnv(".env")

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]commandEntry),
	}
	cmds.register("login", "Log into a registered user account", handlerLogin)
	cmds.register("register", "Register a new user", handlerRegister)
	cmds.register("users", "List all registered users", handlerListUsers)
	cmds.register("generate", "Generate a new password for the given platform", handlerGenerate)
	cmds.register("get", "Get a password for the given platform", handlerGetPassword)
	cmds.register("platforms", "List all platforms for current user", handlerGetPlatforms)
	cmds.register("delete", "Delete given platform", handlerDeletePlatform)
	cmds.register("deactivate", "Deactivate the given user", handlerDeleteUser)
	cmds.register("help", "List all commands and descriptions", func(s *state, c command) error {
		return handlerHelp(&cmds, s, c)
	})
	cmds.register("new", "Update password for given platform", handlerUpdatePassword)
	cmds.register("update", "Update to the newest version of Easypass", handlerUpdate)
	cmds.register("version", "Print Easypass version info", handlerVersion)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

}
