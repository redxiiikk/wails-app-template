<script lang="ts" setup>
import { ref } from 'vue';
import echoApi from '../../api/echo';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';

const message = ref<string>('');
const echoMessage = ref<string>('');

const echo = () => {
  echoApi(message.value).then((response) => {
    echoMessage.value = response;
  });
};
</script>

<template>
  <div class="space-y-8">
    <h1 class="text-3xl">
      Echo Page
    </h1>
    <div class="!space-x-4">
      <InputText
        v-model="message"
        type="text"
      />
      <Button
        label="Echo"
        @click="echo"
      />
    </div>
    <div>
      <p>{{ echoMessage }}</p>
    </div>
  </div>
</template>

<style scoped></style>
