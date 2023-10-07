import { Output, date, number, object, string, array, minLength, maxLength } from "valibot";

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

export const schemaGetProject = object({
  project: schemaProject,
});

export type GetProject = Output<typeof schemaGetProject>;

export const schemaCreateProject = object({
  name: string([
    minLength(3, "The name must be at least 3 characters long."),
    maxLength(30, "The name must be at most 30 characters long."),
  ]),
  description: string([maxLength(256, "The description must be at most 256 characters long.")]),
});

export type CreateProject = Output<typeof schemaCreateProject>;
