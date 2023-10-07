import { Output, date, number, object, string, array } from "valibot";

export const schemaProject = object({
  id: number(),
  name: string(),
  description: string(),
  createdBy: number(),
  createdAt: date(),
});

export type Project = Output<typeof schemaProject>;

export const schemaGetProjects = object({
  projects: array(schemaProject),
});

export type GetProjects = Output<typeof schemaGetProjects>;
