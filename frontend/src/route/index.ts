import { createMemoryHistory, createRouter, RouteRecordRaw } from 'vue-router';

import menus from './menus';

const routes: RouteRecordRaw[] = menus.map((menu) => {
  return {
    path: menu.path,
    component: menu.component,
  };
});

export default createRouter({
  history: createMemoryHistory(),
  routes,
});
