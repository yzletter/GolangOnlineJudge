<template>
  <div>
    <a-row
      id="globalHeader"
      class="grid-demo"
      style="margin-bottom: 16px"
      align="center"
    >
      <!--      <a-col flex="100px"> </a-col>-->
      <a-col flex="auto">
        <a-menu
          mode="horizontal"
          :selected-keys="selectedKeys"
          @menu-item-click="doMenuClick"
        >
          <a-menu-item
            key="0"
            :style="{ padding: 0, marginRight: '38px' }"
            disabled
          >
            <div class="title-bar">
              <img class="logo" src="../assets/ojlogo.jpg" />
              <div class="title">GoOJ</div>
            </div>
          </a-menu-item>
          <a-menu-item v-for="item in routes" :key="item.path">
            {{ item.name }}
          </a-menu-item>
        </a-menu>
      </a-col>
      <a-col flex="100px">
        <div>
          {{ store.state.user?.loginUser?.userName ?? "未登录" }}
        </div>
      </a-col>
    </a-row>
  </div>
</template>

<script setup lang="ts">
import { routes } from "../router/routes";
import { useRoute, useRouter } from "vue-router";
import { ref } from "vue";
import { useStore } from "vuex";

const router = useRouter();
const store = useStore();
// 默认主页
const selectedKeys = ref(["/"]);

// 跳转路由更新菜单
router.afterEach((to, from, failure) => {
  selectedKeys.value = [to.path];
});

setTimeout(() => {
  store.dispatch("user/getLoginUser", {
    userName: "yzletter",
    role: "admin",
  });
}, 3000);

const doMenuClick = (key: string) => {
  router.push({
    path: key,
  });
};
</script>

<style scoped>
.title-bar {
  display: flex;

  align-items: center;
}

.title {
  color: #444;
  margin-left: 16px;
}

.logo {
  height: 50px;
}
</style>
