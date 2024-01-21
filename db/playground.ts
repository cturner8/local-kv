import { PrismaClient } from "@prisma/client";

const prisma = new PrismaClient();

const main = async () => {
  return Promise.all([
    prisma.keyMetadata.findMany({
      take: 10,
      select: {
        id: true,
        // createdDate: true,
        // updatedDate: true,
        arn: true,
        awsAccountId: true,
        cloudHsmClusterId: true,
        customerMasterKeySpec: true,
        customKeyStoreId: true,
        deletionDate: true,
        description: true,
        enabled: true,
        expirationModel: true,
        keyManager: true,
        keySpec: true,
        keyState: true,
        keyUsage: true,
        multiRegion: true,
        origin: true,
        pendingDeletionWindowInDays: true,
        validTo: true,
      },
    }),
    prisma.keyMaterial.findMany({
      take: 10,
      select: {
        id: true,
        keyId: true,
        // keyMetadata: true,
        protectedDataKey: true,
        protectedKeyMaterial: true,
        unprotectedKeyMaterial: true,
        // createdDate: true,
        // updatedDate: true,
      },
    }),
    prisma.keyMetadataAlgorithm.findMany({
      take: 10,
      select: {
        id: true,
        keyMetadataId: true,
        algorithm: true,
      },
    }),
  ]);
};

main()
  .then((result) => console.info(result))
  .catch((error) => console.error(error));
