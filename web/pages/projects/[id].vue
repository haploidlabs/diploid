<template>
  <div v-if="project" class="container mx-auto max-w-3xl flex flex-col gap-8">
    <div class="space-y-6">
      <div class="space-y-4">
        <div class="flex justify-between items-end">
          <h1 class="text-4xl font-bold">Project</h1>
          <div class="space-x-2">
            <NuxtLink class="btn btn-neutral btn-sm gap-1" to="/projects/new">
              <Icon name="heroicons:pencil" class="w-4 h-4" />
              Edit
            </NuxtLink>
            <button class="btn btn-error btn-sm gap-1" @click="deleteDialogRef?.showModal()">
              <Icon name="heroicons:trash" class="w-4 h-4" />
              Delete
            </button>
            <ProjectsNewDeleteDialog
              :project-id="project?.id"
              :handle-delete="deleteProject"
              @update:dialog-ref="(ref) => (deleteDialogRef = ref)"
            />
          </div>
        </div>
        <p class="text-sm text-black/80">{{ project?.description }}</p>
      </div>
      <div class="space-y-4">
        <h1 class="text-2xl font-bold">Environments</h1>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <NuxtLink
            v-for="environment in environments"
            :key="environment.id"
            :to="`/projects/${environment.id}`"
            class="card card-bordered card-compact hover:bg-base-200 cursor-pointer transition ease-out duration-100 hover:ease-in hover:duration-75 active:bg-base-300"
          >
            <div class="card-body">
              <h2 class="card-title">{{ environment.name }}</h2>
              <p>{{ environment.description }}</p>
            </div>
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { parse } from "valibot";
import { schemaId } from "~/lib/types/util";

const route = useRoute();
const router = useRouter();

const id = parse(schemaId, route.params.id);

const project = useProject(id);
const { data: environments } = useEnvironments(id);

const deleteDialogRef = ref<HTMLDialogElement | null>(null);

const { mutate: deleteProject } = useMutation({
  mutationKey: ["projects.delete"],
  mutationFn: async () => {
    const api = useApiStore();
    return await api.DELETE(`/projects/${id}`);
  },
  onSuccess: () => {
    router.push("/");
  },
});
</script>
