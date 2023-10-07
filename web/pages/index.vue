<template>
  <div class="container mx-auto prose prose-h1:mb-4 prose-h2:mt-0 flex flex-col gap-8">
    <div class="space-y-4">
      <h1>Projects</h1>
      <form>
        <input class="input input-bordered w-full" type="text" placeholder="Search" />
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

const { data: projects } = useProjects();
</script>
