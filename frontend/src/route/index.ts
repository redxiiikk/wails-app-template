import { createMemoryHistory, createRouter, RouteRecordRaw } from 'vue-router';

import menus, { Menus } from './menus';

function menusConvertToRoute(menu: Menus): RouteRecordRaw[] {
  if (menu.items) {
    return ([] as RouteRecordRaw[]).concat(
      ...menu.items.map(menusConvertToRoute),
    );
  }

  if (!menu.items && menu.component) {
    return [
      {
        path: menu.path,
        component: menu.component,
      },
    ];
  }

  return [];
}

function toRoutes(menus1: Array<Menus>) {
  return ([] as RouteRecordRaw[]).concat(
    ...menus1.map((menu) => menusConvertToRoute(menu)),
  );
}

export default createRouter({
  history: createMemoryHistory(),
  routes: toRoutes(menus),
});
