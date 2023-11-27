<script setup lang="ts">
// @ts-ignore
import MenuIcon from 'vue-material-design-icons/Menu.vue';
// @ts-ignore
import Close from 'vue-material-design-icons/Close.vue';
import {ref} from "vue";
import {useUserStore} from "@/stores/user-store";

const user = useUserStore()

const displayMenu = ref<boolean>(false)

const toggleMenu = () => {
  displayMenu.value = !displayMenu.value
}
</script>

<template>
  <div class="navigation-border">
    <div class="navigation-wrapper">
      <div class="navigation">
        <router-link to="/" class="logo"><h3>MCBP.php</h3></router-link>
        <!--<img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />-->
        <a class="navigation__icon" :onclick="toggleMenu" v-if="!displayMenu"><MenuIcon :size="32"/></a>
        <a class="navigation__icon" :onclick="toggleMenu" v-if="displayMenu"><Close :size="32"/></a>
      </div>
      <div v-if="displayMenu" class="menu">
        <router-link to="/" :onclick="toggleMenu">Homepage</router-link>
        <router-link v-if="!user.id" to="/login" :onclick="toggleMenu">Login</router-link>
        <router-link v-if="user.id" to="/profile" :onclick="toggleMenu">Profile</router-link>
      </div>
    </div>
  </div>
</template>

<style>
.navigation-border {
  width: 98%;
  border-bottom: 1px solid rgba(60, 60, 67, .12);
  margin: auto
}
.navigation-wrapper {
  width: 100%;
  max-width: 1280px;
  margin: auto;
  position: relative;
}
.navigation {
  width: 100%;
  position: relative;
  top: 0;
  left: 0;
  height: auto;
  padding: 1rem;
  display: flex;
  justify-content: space-between;
  background-color: transparent;
}
.logo {
  margin: auto 0;
  color: #00bd7e;
}
.navigation__icon {
  margin: auto 0;
}
.navigation__icon:hover {
  cursor: pointer;
}

.menu {
  width: auto;
  position: absolute;
  right: 1rem;
  padding: 1rem;
  display: flex;
  flex-direction: column;
  text-align: right;
  gap: 1rem;
  border-radius: 10px;
  background-color: rgb(255, 255, 255);
  border: 1px solid rgba(60, 60, 67, 0.12);
  box-shadow: rgba(0, 0, 0, 0.1) 0px 12px 32px 0px, rgba(0, 0, 0, 0.08) 0px 2px 6px 0px;
}
.menu a {
  width: auto;
}
</style>