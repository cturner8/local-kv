-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_KeyMetadata" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdDate" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedDate" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "arn" TEXT,
    "awsAccountId" TEXT,
    "cloudHsmClusterId" TEXT,
    "customerMasterKeySpec" TEXT,
    "customKeyStoreId" TEXT,
    "deletionDate" INTEGER,
    "description" TEXT,
    "enabled" BOOLEAN NOT NULL,
    "expirationModel" TEXT,
    "keyManager" TEXT NOT NULL,
    "keySpec" TEXT,
    "keyState" TEXT,
    "keyUsage" TEXT,
    "multiRegion" BOOLEAN NOT NULL,
    "origin" TEXT,
    "pendingDeletionWindowInDays" INTEGER,
    "validTo" INTEGER
);
INSERT INTO "new_KeyMetadata" ("arn", "awsAccountId", "cloudHsmClusterId", "createdDate", "customKeyStoreId", "customerMasterKeySpec", "deletionDate", "description", "enabled", "expirationModel", "id", "keyManager", "keySpec", "keyState", "keyUsage", "multiRegion", "origin", "pendingDeletionWindowInDays", "updatedDate", "validTo") SELECT "arn", "awsAccountId", "cloudHsmClusterId", "createdDate", "customKeyStoreId", "customerMasterKeySpec", "deletionDate", "description", "enabled", "expirationModel", "id", "keyManager", "keySpec", "keyState", "keyUsage", "multiRegion", "origin", "pendingDeletionWindowInDays", "updatedDate", "validTo" FROM "KeyMetadata";
DROP TABLE "KeyMetadata";
ALTER TABLE "new_KeyMetadata" RENAME TO "KeyMetadata";
CREATE TABLE "new_KeyMaterial" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "keyId" TEXT NOT NULL,
    "protectedDataKey" TEXT NOT NULL,
    "protectedKeyMaterial" TEXT NOT NULL,
    "unprotectedKeyMaterial" TEXT,
    "createdDate" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedDate" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "KeyMaterial_keyId_fkey" FOREIGN KEY ("keyId") REFERENCES "KeyMetadata" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
INSERT INTO "new_KeyMaterial" ("id", "keyId", "protectedDataKey", "protectedKeyMaterial", "unprotectedKeyMaterial") SELECT "id", "keyId", "protectedDataKey", "protectedKeyMaterial", "unprotectedKeyMaterial" FROM "KeyMaterial";
DROP TABLE "KeyMaterial";
ALTER TABLE "new_KeyMaterial" RENAME TO "KeyMaterial";
CREATE UNIQUE INDEX "KeyMaterial_keyId_key" ON "KeyMaterial"("keyId");
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
