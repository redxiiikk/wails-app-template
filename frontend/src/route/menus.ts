import { Component } from 'vue';

export interface Menus {
  name: string;
  path: string;
  component?: () => Promise<Component>;
  icon?: string;
  items?: Menus[];
}

export default [
  {
    name: 'Home',
    path: '/',
    component: () => import('../pages/home/index.vue'),
    icon: 'pi pi-home',
  },
  {
    name: 'Echo',
    path: '/echo',
    component: () => import('../pages/echo/index.vue'),
    icon: 'pi pi-wrench',
  },
  {
    name: 'Setting',
    path: '/setting',
    icon: 'pi pi-cog',
    items: [
      {
        name: 'Profile',
        path: '/profile',
        icon: 'pi pi-cog',
        component: () => import('../pages/profile/index.vue'),
      },
    ],
  },
] as Array<Menus>;
