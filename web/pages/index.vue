<template>
  <div class="container mx-auto prose prose-h1:mb-4 prose-p:mb-0 prose-h2:mt-0 flex flex-col gap-8">
    <div class="space-y-4">
      <h1>Projects</h1>
      <form>
        <input
          v-model="inputSearch"
          class="input input-bordered w-full"
          type="text"
          placeholder="Search project..."
        />
      </form>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div v-for="project in projects" :key="project.id" class="card card-bordered card-compact">
          <div class="card-body">
            <h2 class="card-title">{{ project.name }}</h2>
            <p>{{ project.description }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
definePageMeta({
  layout: "landing",
});

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
