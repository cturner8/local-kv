-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_KeyMaterial" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "keyId" TEXT NOT NULL,
    "protectedDataKey" TEXT NOT NULL,
    "protectedKeyMaterial" TEXT NOT NULL,
    "unprotectedKeyMaterial" TEXT,
    "createdDate" TEXT NOT NULL,
    "updatedDate" TEXT NOT NULL,
    CONSTRAINT "KeyMaterial_keyId_fkey" FOREIGN KEY ("keyId") REFERENCES "KeyMetadata" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
INSERT INTO "new_KeyMaterial" ("createdDate", "id", "keyId", "protectedDataKey", "protectedKeyMaterial", "unprotectedKeyMaterial", "updatedDate") SELECT "createdDate", "id", "keyId", "protectedDataKey", "protectedKeyMaterial", "unprotectedKeyMaterial", "updatedDate" FROM "KeyMaterial";
DROP TABLE "KeyMaterial";
ALTER TABLE "new_KeyMaterial" RENAME TO "KeyMaterial";
CREATE UNIQUE INDEX "KeyMaterial_keyId_key" ON "KeyMaterial"("keyId");
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
