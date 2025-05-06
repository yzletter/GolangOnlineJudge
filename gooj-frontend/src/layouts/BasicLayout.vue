<template>
  <div id="basicLayout">
    <a-layout style="min-height: 100vh">
      <!--   首部   -->
      <a-layout-header class="header">
        <GlobalHeader></GlobalHeader>
      </a-layout-header>
      <!--   中部   -->
      <a-layout-content class="content">
        <router-view />
      </a-layout-content>
      <!--   尾部   -->
      <a-layout-footer class="footer">
        <a href="https://github.com/yzletter" target="_blank" a>
          Designed by yzletter
        </a>
      </a-layout-footer>
    </a-layout>
  </div>
</template>

<style scoped>
#basicLayout {
}

#basicLayout .header {
  margin-bottom: 16px;
  box-shadow: #eee 1px 1px 5px;
}

#basicLayout .content {
  background: linear-gradient(to right, #bbb, #fff);
  margin-bottom: 16px;
}

#basicLayout .footer {
  background: #efefef;
  padding: 16px;
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  text-align: center;
}
</style>

<script setup lang="ts">
import GlobalHeader from "@/components/GlobalHeader.vue";
</script>
