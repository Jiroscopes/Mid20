<template>
    <div>
        <img class="lg:block hidden h-16 w-16 bounce relative" src="../assets/scroll_icon.png" alt="Scroll Icon" >
    </div>
</template>

<script lang="ts">

import Vue from "vue"
import Component from "vue-class-component"

@Component
export default class ScrollIcon extends Vue {

}

</script>

<style scoped>

.bounce {
  -moz-animation: bounce 1s infinite ease-in;
  -o-animation: bounce 1s infinite ease-in;
  -webkit-animation: bounce 1s infinite ease-in;
  animation: bounce 1s infinite ease-in;
}
@-moz-keyframes bounce {
    0% { top: 0; }
    50% {top: -0.3em; }
    70% {top: -0.4em; }
    100% {top: 0; }
}
@-o-keyframes bounce {
    0% { top: 0; }
    50% { top: -0.3em; }
    70% { top: -0.4em; }
    100% { top: 0; }
}
@-ms-keyframes bounce {
    0% { top: 0; }
    50% { top: -0.3em; }
    70% { top: -0.4em; }
    100% { top: 0; }
}
@keyframes bounce {
    0% { top: 0; }
    50% { top: -0.3em; }
    70% { top: -0.4em; }
    100% { top: 0; }
}
@-webkit-keyframes bounce {
    0% { top: 0; }
    50% { top: -0.3em; }
    70% { top: -0.4em; }
    100% { top: 0; }
}
</style>
