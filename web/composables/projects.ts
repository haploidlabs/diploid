import { GetProjects } from "~/lib/types/projects";

export const useProjects = () => {
  const api = useApiStore();
  return useQuery({
    queryKey: ["projects"],
    queryFn: async () => {
      return await api.GET<GetProjects>("/projects");
    },
    select: (data) => data.projects,
  });
};

export const useProject = (id: number) => {
  const { data } = useProjects();
  return computed(() => {
    if (!data.value) return undefined;
    return data.value.find((project) => project.id === id);
  });
};
