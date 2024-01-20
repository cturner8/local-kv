CREATE TABLE
    IF NOT EXISTS KeyMetadata (
        KeyId TEXT PRIMARY KEY NOT NULL,
        Arn TEXT,
        AWSAccountId TEXT,
        CloudHsmClusterId TEXT,
        CreationDate INTEGER DEFAULT CURRENT_TIMESTAMP,
        CustomerMasterKeySpec TEXT,
        CustomKeyStoreId TEXT,
        DeletionDate INTEGER,
        Description TEXT,
        Enabled INTEGER,
        EncryptionAlgorithms TEXT,
        ExpirationModel TEXT,
        KeyManager TEXT,
        KeySpec TEXT,
        KeyState TEXT,
        KeyUsage TEXT,
        MacAlgorithms TEXT,
        MultiRegion INTEGER,
        Origin TEXT,
        PendingDeletionWindowInDays INTEGER,
        SigningAlgorithms TEXT,
        ValidTo INTEGER,
    );

CREATE TABLE
    IF NOT EXISTS KeyMaterial (
        KeyId TEXT PRIMARY KEY NOT NULL,
        ProtectedDataKey TEXT NOT NULL,
        ProtectedKeyMaterial TEXT NOT NULL,
        UnprotectedKeyMaterial TEXT,
    )