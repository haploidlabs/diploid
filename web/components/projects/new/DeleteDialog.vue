<template>
  <dialog id="my_modal_1" ref="dialogRef" class="modal" aria-label="Delete Project">
    <div class="modal-box">
      <h3 class="font-bold text-lg">
        Delete <span>{{ project?.name }}</span>
      </h3>
      <p class="py-4">
        Type <span class="font-bold">{{ project?.name }}</span> to delete the project.
      </p>
      <input
        v-model="inputCheck"
        class="input input-bordered w-full"
        :placeholder="project?.name"
      />
      <div class="modal-action w-full">
        <form method="dialog" class="flex gap-4 w-full">
          <button class="btn btn-error w-full" :disabled="!canDelete" @click="handleDelete()">
            Delete
          </button>
        </form>
      </div>
    </div>
  </dialog>
</template>

<script lang="ts" setup>
import { ref, onMounted, defineEmits } from "vue";

const props = defineProps<{
  projectId: number;
  handleDelete: () => void;
}>();
const emit = defineEmits(["update:dialogRef"]);

const dialogRef = ref<HTMLDialogElement | null>(null);

const inputCheck = ref("");

const canDelete = computed(() => {
  return inputCheck.value === project.value?.name;
});

onMounted(() => {
  emit("update:dialogRef", dialogRef);
});

const project = useProject(props.projectId);
</script>
