import {Component} from "vue";

interface Menus {
    name: string;
    path: string;
    component: () => Promise<Component>;
    icon?: string;
}

export default [
    {
        name: "Home",
        path: "/",
        component: () => import("../pages/home/index.vue"),
        icon: "pi pi-home",
    },
    {
        name: "Echo",
        path: "/echo",
        component: () => import("../pages/echo/index.vue"),
        icon: "pi pi-wrench",
    },
] as Array<Menus>
