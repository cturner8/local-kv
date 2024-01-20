-- CreateTable
CREATE TABLE "KeyMetadata" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdDate" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedDate" DATETIME NOT NULL,
    "arn" TEXT,
    "awsAccountId" TEXT,
    "cloudHsmClusterId" TEXT,
    "creationDate" INTEGER,
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

-- CreateTable
CREATE TABLE "KeyMetadataAlgorithm" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "keyMetadataId" TEXT NOT NULL,
    "algorithm" TEXT NOT NULL,
    CONSTRAINT "KeyMetadataAlgorithm_keyMetadataId_fkey" FOREIGN KEY ("keyMetadataId") REFERENCES "KeyMetadata" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "KeyMetadataAlgorithm_keyMetadataId_fkey" FOREIGN KEY ("keyMetadataId") REFERENCES "KeyMetadata" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "KeyMetadataAlgorithm_keyMetadataId_fkey" FOREIGN KEY ("keyMetadataId") REFERENCES "KeyMetadata" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "KeyMaterial" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "keyId" TEXT NOT NULL,
    "protectedDataKey" TEXT NOT NULL,
    "protectedKeyMaterial" TEXT NOT NULL,
    "unprotectedKeyMaterial" TEXT NOT NULL,
    CONSTRAINT "KeyMaterial_keyId_fkey" FOREIGN KEY ("keyId") REFERENCES "KeyMetadata" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateIndex
CREATE UNIQUE INDEX "KeyMaterial_keyId_key" ON "KeyMaterial"("keyId");
