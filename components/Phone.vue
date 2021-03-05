<template>
    <div class="phone-container lg:block sm:hidden mx-8">
        <img v-bind:style="{left: mousePos[0] + 'px', top: mousePos[1] + 50 + 'px'}" class="absolute z-10" src="/floating-phone-icons.png" alt="">
        <img ref="phoneIMG" class="absolute z-0" src="/phone.png" alt="">
    </div> 
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'
@Component
export default class Phone extends Vue{

    mousePos: [number, number] = [0, 0];

    data() {
        return {
            mouseEnter: false,
        }
    }
     
    changePhonePos(x: number, y: number): void {
        this.mousePos = [x/50, y/50];
    }
}
</script>

<style scoped>
.phone-container {
    max-width: 777.75px;
    max-height: 975px;

    @apply relative pt-12;

    @media (max-width: 1460px) {
        top: 2rem;
    }
    @media (max-width: 1100px) {
        top: 7rem;
    }
}

img {
    user-drag: none; 
    user-select: none;
    -moz-user-select: none;
    -webkit-user-drag: none;
    -webkit-user-select: none;
    -ms-user-select: none;
}
</style>
