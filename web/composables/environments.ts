import { GetEnvironments } from "~/lib/types/environments";

export const useEnvironments = (projectId: number) =>
  useQuery({
    queryKey: ["environments", projectId],
    queryFn: async () => {
      const apiStore = useApiStore();
      return await apiStore.GET<GetEnvironments>(`/projects/${projectId}/environments`);
    },
    select: (data) => data.environments,
  });
