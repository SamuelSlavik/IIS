<script setup lang="ts">

</script>

<template>
  <div class="dashboard">
    <div class="dashboard-content">
      <div>
        <h2>Arthur Dent</h2>
        <p>arthur@protonmail.com</p>
        <br/>
        <p>Admin</p>
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
        <button class="small-button">Edit profile</button>
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