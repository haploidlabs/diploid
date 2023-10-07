<template>
  <div class="container mx-auto prose">
    <h1>Containers</h1>
    <p v-if="error" class="text-red-500">{{ error }}</p>
    <div v-else-if="data" class="space-y-2">
      <div
        v-for="container in data.containers"
        :key="container.id"
        class="card card-bordered card-compact"
      >
        <div class="card-body">
          <h2 class="card-title">{{ container.name }} ({{ container.status }})</h2>
          <div>
            <p>Image: {{ container.image }}({{ container.imageId }})</p>
            <div>
              <p class="font-semibold">Ports</p>
              <ul>
                <li
                  v-for="port in container.ports"
                  :key="`${port.public_port}:${port.private_port}:${port.type}`"
                >
                  - {{ port.public_port }}:{{ port.private_port }}
                </li>
              </ul>
            </div>
          </div>
          <div>
            <p class="font-semibold">Volumes</p>
            <ul>
              <li v-for="volume in container.volumes" :key="`${volume.host}:${volume.container}`">
                - {{ volume.host }}:{{ volume.container }}
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Container } from "~/lib/types/containers";

definePageMeta({
  layout: "landing",
});

const apiStore = useApiStore();

const { data, error } = useQuery({
  queryKey: ["containers"],
  queryFn: () => {
    return apiStore.GET<{ containers: Container[] }>("/containers");
  },
  refetchInterval: 1000,
});
</script>
