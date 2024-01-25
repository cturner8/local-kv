package testing

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func setupKeyMetadataTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE 
		KeyMetadata (
			id TEXT NOT NULL PRIMARY KEY,
			createdDate TEXT NOT NULL,
			updatedDate TEXT NOT NULL,
			arn TEXT,
			awsAccountId TEXT,
			cloudHsmClusterId TEXT,
			customerMasterKeySpec TEXT,
			customKeyStoreId TEXT,
			deletionDate INTEGER,
			description TEXT,
			enabled BOOLEAN NOT NULL,
			expirationModel TEXT,
			keyManager TEXT NOT NULL,
			keySpec TEXT,
			keyState TEXT,
			keyUsage TEXT,
			multiRegion BOOLEAN NOT NULL,
			origin TEXT,
			pendingDeletionWindowInDays INTEGER,
			validTo INTEGER
		);
	`)
	if err != nil {
		panic(err)
	}
}

func setupKeyMaterialTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE 
		KeyMaterial (
			id TEXT NOT NULL PRIMARY KEY,
			keyId TEXT NOT NULL,
			protectedDataKey TEXT NOT NULL,
			protectedKeyMaterial TEXT NOT NULL,
			unprotectedKeyMaterial TEXT,
			createdDate TEXT NOT NULL,
			updatedDate TEXT NOT NULL,
			CONSTRAINT KeyMaterial_keyId_fkey FOREIGN KEY (keyId) REFERENCES KeyMetadata (id) ON DELETE RESTRICT ON UPDATE CASCADE
		);

	CREATE UNIQUE INDEX KeyMaterial_keyId_key ON KeyMaterial(keyId);
	`)
	if err != nil {
		panic(err)
	}
}

func setupTables(db *sql.DB) {
	setupKeyMetadataTable(db)
	setupKeyMaterialTable(db)
}

func SetupDatabase(tempDirName string) *sql.DB {
	dir, err := os.MkdirTemp("", tempDirName)
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	db, err := sql.Open("sqlite3", dir+"/vault.db")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	setupTables(db)

	return db
}
