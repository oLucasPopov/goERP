package migrations

import (
	"embed"
	"fmt"
	migrationsRepository2 "goERP/repositories/migration"
	"goERP/types/migration"
	"io/fs"
	"log"
	"strings"
	"time"
)

//go:embed queries/*.sql
var efs embed.FS
var migrationsRepository = migrationsRepository2.Migration{}

func Migrate() {
	initMigration()
	saveMigrations()
	persistMigrations()
}

func initMigration() {
	if err := migrationsRepository.Init(); err != nil {
		log.Fatal("Erro ao inicializar migration!", err.Error())
	}
}

func persistMigrations() {
	migrations := loadMigrations()

	err := migrationsRepository.Persist(migrations)
	if err != nil {
		log.Fatalln("Erro ao persistir as queries em banco de dados!", err.Error())
	}
}

func loadMigrations() migration.Migrations {
	if migrations, err := migrationsRepository.List(); err != nil {
		log.Fatalln("Erro ao carregar queries:", err.Error())
	} else {
		return migrations
	}
	return migration.Migrations{}
}

func saveMigrations() {
	scripts := getMigrations()

	if err := migrationsRepository.Add(scripts); err != nil {
		log.Fatalln("Erro ao salvar queries:", err.Error())
	}
}

func getMigrations() migration.Migrations {
	const MigrationDir = "queries"
	var migrationsList migration.Migrations

	files, err := efs.ReadDir(MigrationDir)
	if err != nil {
		log.Fatalln("Erro ao ler diretório das queries:", err.Error())
	}

	for _, file := range files {
		content, err := fs.ReadFile(efs, fmt.Sprintf("%s/%s", MigrationDir, file.Name()))

		if err != nil {
			log.Fatalln("Erro ao ler a migration", err.Error())
		}

		fileFullName := file.Name()

		fileDateString := fileFullName[0:10]
		fileTimeString := strings.ReplaceAll(fileFullName[11:19], "-", ":")
		fileTime, err := time.Parse("2006-01-02 15:04:05", fileDateString+" "+fileTimeString)
		if err != nil {
			log.Fatalln("Erro ao converter data", err.Error())
		}

		migrationName := fileFullName[20 : len(fileFullName)-4]

		migrationsList = append(migrationsList,
			migration.Migration{
				Timestamp: fileTime,
				Name:      migrationName,
				Content:   string(content),
			})
	}

	return migrationsList
}
