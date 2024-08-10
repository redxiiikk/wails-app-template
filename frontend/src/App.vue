<script lang="ts" setup>
import {RouterView, useRouter} from 'vue-router'
import {ref} from "vue";
import Menus from 'primevue/menu';
import menus from './route/menus';

const router = useRouter();

const navMenus = ref(
    menus.map(menu => {
      return {
        icon: menu.icon,
        label: menu.name,
        command: () => {
          router.push(menu.path);
        }
      }
    })
);


</script>

<template>
  <main class="grid grid-cols-8 h-screen">
    <nav>
      <Menus :model="navMenus" class="h-screen"/>
    </nav>
    <main class="col-span-7 pl-2">
      <section>
        <router-view/>
      </section>
    </main>
  </main>
</template>
