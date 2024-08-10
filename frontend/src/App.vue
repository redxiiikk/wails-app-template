<script lang="ts" setup>
import { RouterView } from 'vue-router';
import Navigation from './components/navigation/index.vue';
</script>

<template>
  <main class="grid grid-cols-8 h-screen">
    <nav>
      <Navigation />
    </nav>
    <main class="col-span-7 pl-2">
      <section>
        <router-view />
      </section>
    </main>
  </main>
</template>
