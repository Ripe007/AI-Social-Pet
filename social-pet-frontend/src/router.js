import * as VueRouter from "vue-router";
import Home from "@/page/Home";
import Collection from "@/page/Collection";
import Comminity from "@/page/Community";
import Foundation from "@/page/Foundation";
import Detail from "@/page/Detail";

const routes = [
  { path: "/", component: Home, name: "Home" },
  { path: "/collection", component: Collection, name: "collection" },
  { path: "/community", component: Comminity, name: "community" },
  { path: "/foundation", component: Foundation, name: "foundation" },
  { path: "/detail", component: Detail, name: "detail" },
];

export const router = VueRouter.createRouter({
  history: VueRouter.createWebHistory(),
  routes,
});
