import { Output, array, object, string } from "valibot";

export const schemaEnvironment = object({
  id: string(),
  name: string(),
  description: string(),
});

export type Environment = Output<typeof schemaEnvironment>;

export const schemaGetEnvironments = object({
  environments: array(schemaEnvironment),
});

export type GetEnvironments = Output<typeof schemaGetEnvironments>;
