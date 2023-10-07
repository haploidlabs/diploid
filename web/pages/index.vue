<template>
  <div class="container mx-auto max-w-3xl flex flex-col gap-8">
    <div class="space-y-4">
      <div class="flex justify-between items-end">
        <h1 class="text-4xl font-bold">Projects</h1>
        <NuxtLink class="btn btn-neutral btn-sm" to="/projects/new">New Project</NuxtLink>
      </div>
      <form class="form-control">
        <input
          v-model="inputSearch"
          class="input input-bordered w-full"
          type="text"
          placeholder="Search project..."
        />
      </form>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <NuxtLink
          v-for="project in projects"
          :key="project.id"
          :to="`/projects/${project.id}`"
          class="card card-bordered card-compact hover:bg-base-200 cursor-pointer transition ease-out duration-100 hover:ease-in hover:duration-75 active:bg-base-300"
        >
          <div class="card-body">
            <h2 class="card-title">{{ project.name }}</h2>
            <p>{{ project.description }}</p>
          </div>
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
const inputSearch = ref("");

const { data } = useProjects();

const projects = computed(() => {
  if (!data.value) return [];
  if (!inputSearch.value) return data.value;
  return data.value.filter((project) => {
    return project.name.toLowerCase().includes(inputSearch.value.toLowerCase());
  });
});
</script>
