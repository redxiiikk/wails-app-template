<script setup lang="ts">
import Tag from 'primevue/tag';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import { onMounted, ref } from 'vue';
import healthcheckApi, { HealthStatus } from '../../api/healthStatus';

const healthStatus = ref();

onMounted(() => {
  healthcheckApi().then((data: HealthStatus[]) => {
    healthStatus.value = data;
  });
});
</script>

<template>
  <DataTable :value="healthStatus" tableStyle="min-width: 50rem">
    <Column field="systemName" header="Name"></Column>
    <Column field="status" header="Status">
      <template #body="slotProps">
        <Tag
          :value="slotProps.data.status"
          :severity="slotProps.data.status === 'UP' ? 'success' : 'danger'"
        />
      </template>
    </Column>
    <Column field="message" header="message"></Column>
  </DataTable>
</template>

<style scoped></style>
