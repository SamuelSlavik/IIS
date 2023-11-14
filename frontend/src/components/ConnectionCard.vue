<script setup lang="ts">
import Tram from 'vue-material-design-icons/Tram.vue';
import Bus from 'vue-material-design-icons/Bus.vue';
import Settings from 'vue-material-design-icons/Cog.vue';
import type {ConnectionStop} from "@/lib/models";

interface Props {
  id: string
  lineName: string
  type: string
  stops: ConnectionStop[]
}
const props = defineProps<Props>()

</script>

<template>
  <div class="card">
    <div class="connection-title">
      <Tram v-if="props.type === 'tram'" class="connection-icon"/>
      <Bus v-if="props.type === 'bus'" class="connection-icon"/>
      <h2>{{ props.lineName }}</h2>
    </div>
    <div class="connection-route">
      <ul>
        <li><div><span>{{props.stops[0].DepartureTime}}</span><span>{{props.stops[0].StopName}}</span></div></li>
        <li><div><span>{{props.stops[stops.length -1].DepartureTime}}</span><span>{{props.stops[stops.length -1].StopName}}</span></div></li>
      </ul>
    </div>
    <button class="small-button">Maybe a button</button>
    <div class="connection-context-icon">
      <Settings class="connection-icon"/>
    </div>
  </div>
</template>

<style>
  .card {
    width: 100%;
    background-color: rgb(246, 246, 247);
    border-radius: 5px;
    margin: 2rem 0;
    padding: 1rem;
    position: relative;
  }
  .card:hover .connection-context-icon {
    opacity: 1;
  }

  .connection-title {
    display: flex;
    justify-content: left;
    gap: 1rem;
  }
  .connection-title * {
    margin: auto 0;
  }
  .connection-icon {
    height: 24px;
    width: 24px;
  }

  .connection-route ul {
    padding: 1rem;
    list-style-type: disc;
  }
  .connection-route ul li div {
    display: flex;
    justify-content: left;
    gap: 1rem;
    font-weight: bold;
  }
  .connection-context-icon {
    position: absolute;
    top: 1rem;
    right: 1rem;
    opacity: 0;
    color: rgb(142, 142, 147);
  }
  .connection-context-icon:hover {
    cursor: pointer;
    color: rgb(60, 60, 67);
  }
</style>