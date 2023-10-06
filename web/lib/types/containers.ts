import { Output, array, enumType, object, string } from "valibot";

export const schemaContainer = object({
  id: string(),
  name: string(),
  image: string(),
  imageId: string(),
  status: enumType(["running", "stopped"]),
  ports: array(
    object({
      public_port: string(),
      private_port: string(),
      type: enumType(["tcp", "udp"]),
    }),
  ),
  volumes: array(
    object({
      host: string(),
      container: string(),
    }),
  ),
});

export type Container = Output<typeof schemaContainer>;
