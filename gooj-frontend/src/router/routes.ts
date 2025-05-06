import { RouteRecordRaw } from "vue-router";
import HomeView from "@/views/HomeView.vue";
import AdminView from "@/views/AdminView.vue";
import NoAuthView from "@/views/NoAuthView.vue";

export const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "题目列表",
    component: HomeView,
  },
  {
    path: "/admin",
    name: "后台管理",
    component: AdminView,
    meta: {
      access: "canAdmin",
    },
  },
  {
    path: "/noauth",
    name: "没有权限",
    component: NoAuthView,
  },
  {
    path: "/about",
    name: "关于我们",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/AboutView.vue"),
  },
];
