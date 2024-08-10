import { createMemoryHistory, createRouter, RouteRecordRaw } from 'vue-router';

import menus from './menus';

const routes = menus.map((menu) => {
  return {
    path: menu.path,
    component: menu.component,
  } as RouteRecordRaw;
});

export default createRouter({
  history: createMemoryHistory(),
  routes,
});
