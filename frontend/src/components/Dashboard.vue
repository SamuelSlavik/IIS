<script setup lang="ts">
import {useUserStore} from "@/stores/user-store";
import {useRouter} from "vue-router";

const user = useUserStore()
const router = useRouter();

const logOut = async () => {
  user.logOut()

  await router.push('/')


  // TODO endpoint for forgetting a cookie
}

</script>

<template>
  <div class="dashboard">
    <div class="dashboard-content">
      <div>
        <h2>{{ user.firstName + " " + user.lastName }}</h2>
        <p>{{user.email}}</p>
        <br/>
        <p>{{user.role}}</p>
      </div>
      <div class="hr"></div>
      <div>
        <p><a>Action 1</a></p>
        <p><a>Action 2</a></p>
        <p><a>Action 3</a></p>
        <br/>
        <p><a>Action 4</a></p>
        <p><a>Action 5</a></p>
      </div>
      <div class="hr"></div>
      <div>
        <button :onclick="logOut" class="small-button">Log Out</button>
      </div>
    </div>
  </div>
</template>

<style>
  .dashboard {
    background-color: rgb(246, 246, 247);
    width: calc((100% - (1280px - 64px)) / 2 + 350px - 32px);
    min-width: 350px;
    height: auto;
    min-height: 100vh;
  }
  .dashboard-content {
    width: 350px;
    margin-right: 0;
    margin-left: auto;
    padding: 2rem 2rem 2rem 1rem;
  }
</style>