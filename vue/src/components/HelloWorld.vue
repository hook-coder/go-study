<script setup>
import { ref } from "vue";
import { useWebSocket } from "@vueuse/core";
const { status, data, send, open, close } = useWebSocket(
  "ws://localhost:1234",
  {
    autoReconnect: true,
    delay: 1000,
    onFailed() {
      alert("Failed to connect WebSocket after 3 retries");
    },
    // heartbeat: {
    //   message: "jump",
    //   interval: 1000,
    // },
  }
);

// send handle
const inputValue = ref(null)
const handleSend = ()=>{
  if(inputValue.value){
    send(inputValue.value)
    inputValue.value = null
  }
}
defineProps({
  msg: String,
});

const count = ref(0);
</script>

<template>
  <h1>WebSocket</h1>
  <input
    class="
      py-2
      px-4
      font-semibold
      rounded-lg
      shadow-md
      text-white
      bg-green-500
      hover:bg-green-400
      border-none
      cursor-pointer
    "
    v-model="inputValue"
    style="outline: none"
    type="text"
    placeholder="enter the message to send"
    @keyup.enter="handleSend"
  />

  <button
    class="
      py-2
      px-4
      font-semibold
      rounded-lg
      shadow-md
      text-white
      bg-green-500
      hover:bg-green-400
      border-none
      cursor-pointer
    "
    style="margin-left:5px;"
    @click="handleSend"
  >
    send
  </button>
  <div>
    <h1>content</h1>
    <div v-if="data">
      <p>{{ data }}</p>
    </div>
  </div>
</template>

<style scoped>
a {
  color: #42b983;
}
</style>
