/*
  Warnings:

  - You are about to drop the column `creationDate` on the `KeyMetadata` table. All the data in the column will be lost.

*/
-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_KeyMetadata" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdDate" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedDate" DATETIME NOT NULL,
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
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
