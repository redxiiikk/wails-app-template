<script setup lang="ts">
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import { onMounted, ref } from 'vue';
import queryAllMigrateHistoryApi, { MigrateHistory } from '../../api/migrate';

const migrateHistory = ref();

onMounted(() => {
  queryAllMigrateHistoryApi().then((data: MigrateHistory[]) => {
    migrateHistory.value = data;
  });
});
</script>

<template>
  <DataTable :value="migrateHistory" tableStyle="min-width: 50rem">
    <Column field="id" header="Id"></Column>
    <Column field="key" header="Key"> </Column>
    <Column field="hash" header="Hash"></Column>
    <Column field="createdAt" header="CreatedAt"></Column>
  </DataTable>
</template>

<style scoped></style>
