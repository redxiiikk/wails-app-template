<script lang="ts" setup>
import { useRouter } from 'vue-router';
import { ref } from 'vue';
import type { MenuItem } from 'primevue/menuitem';
import PanelMenu from 'primevue/panelmenu';
import type { Menus as MenuTypes } from '../../route/menus';
import menus from '../../route/menus';

const router = useRouter();

const menusConvert = (menu: MenuTypes): MenuItem => ({
  icon: menu.icon || '',
  label: menu.name,
  command: () => {
    if (!menu.items) {
      router.push(menu.path);
    }
  },
  items: menu.items?.map((item) => menusConvert(item)),
});

const navMenus = ref(menus.map((menu) => menusConvert(menu)));
</script>

<template>
  <PanelMenu :model="navMenus" class="h-screen border-solid border-2" />
</template>
<style scoped>
.p-panelmenu {
    gap: 0;
    border-color: var(--p-panelmenu-panel-border-color);
}

:deep .p-panelmenu-panel {
    border-left: none;
    border-right: none;
    border-radius: 0 !important;
}


:deep .p-panelmenu-panel:not(:last-child) {
    border-bottom-width: 0;
}
</style>
