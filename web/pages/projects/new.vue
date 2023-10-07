<template>
  <div class="container mx-auto flex flex-col gap-8">
    <div class="space-y-8">
      <h1 class="text-4xl font-bold text-center">Create new project</h1>
      <form
        class="form-control max-w-lg mx-auto flex flex-col gap-4"
        @submit.prevent="handleSubmit"
      >
        <div class="w-full">
          <label for="name" class="label font-medium">Project Name</label>
          <input
            v-model="inputName"
            type="text"
            name="name"
            class="input input-bordered w-full"
            placeholder="Project name..."
          />
          <p v-if="errors.name" class="text-sm text-red-500">
            {{ errors.name }}
          </p>
        </div>
        <div class="w-full">
          <label for="description" class="label font-medium">Project Description</label>
          <textarea
            v-model="inputDescription"
            type="text"
            name="description"
            class="textarea textarea-bordered w-full"
            placeholder="Describe your project..."
          />
          <p v-if="errors.description" class="text-sm text-red-500">
            {{ errors.description }}
          </p>
        </div>
        <div>
          <p v-if="errors.general" class="text-sm text-red-500">
            {{ errors.general }}
          </p>
        </div>
        <button type="submit" class="btn w-full sm:btn-wide btn-neutral">
          <span v-if="isLoading">
            <Icon name="tdesign:loading" class="w-6 h-6 animate-spin" />
          </span>
          <span v-else>Create Project</span>
        </button>
      </form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { FetchError } from "ofetch";
import { safeParse } from "valibot";
import { GetProject, CreateProject, schemaCreateProject } from "~/lib/types/projects";

const router = useRouter();
const inputName = ref("");
const inputDescription = ref("");

const errors = ref<{
  name?: string;
  description?: string;
  general?: string;
  [key: string]: string | undefined;
}>({});

const { refetch: refetchProjects } = useProjects();

const validate = (project: CreateProject) => {
  errors.value = {};
  const validation = safeParse(schemaCreateProject, project);
  if (!validation.success) {
    validation.issues.forEach((issue) => {
      const pathKey = issue.path ? issue.path.map((path) => path.key).join(".") : "general";
      errors.value[pathKey] = issue.message;
    });
    return false;
  }
  return true;
};

const { mutate, isLoading } = useMutation({
  mutationKey: ["projects.create"],
  mutationFn: async (project: CreateProject) => {
    const api = useApiStore();
    return await api.POST<GetProject>("/projects", project);
  },
  onSuccess: (data) => {
    refetchProjects();
    router.push(`/projects/${data.project.id}`);
  },
  onError: (err) => {
    if (err instanceof FetchError) {
      errors.value.general = err.message;
    }
  },
});

const handleSubmit = () => {
  const project = { name: inputName.value, description: inputDescription.value };
  if (!validate(project)) return;
  mutate(project);
};
</script>
