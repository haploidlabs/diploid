import { Project } from "~/lib/types/projects";

export const useProjectsStore = defineStore("projects", () => {
  const projects = ref([] as Project[]);

  const fetch = async () => {
    const api = useApiStore();
    projects.value = await api.GET<Project[]>("/projects");
  };

  return { projects, fetch };
});
