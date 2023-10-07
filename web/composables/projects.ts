import { GetProjects } from "~/lib/types/projects";

export const useProjects = () =>
  useQuery({
    queryKey: ["projects"],
    queryFn: async () => {
      const api = useApiStore();
      return await api.GET<GetProjects>("/projects");
    },
    select: (data) => data.projects,
  });
