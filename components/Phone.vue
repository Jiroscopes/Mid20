<template>
    <div class="lg:block sm:hidden relative">
        <img v-bind:style="{left: mousePos[0] + 'px', top: mousePos[1] + 'px'}" class="absolute z-10" src="../assets/floating-phone-icons.png" alt="">
        <img ref="phoneIMG" class="absolute z-0" src="../assets/phone.png" alt="">
    </div> 
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'
@Component
export default class Phone extends Vue{

    counter : number = 0;
    mousePos: [number, number] = [0, 0];

    data() {
        return {
            mouseEnter: false,
            //counter: 0
        }
    }
     
    changePhonePos(x: number, y: number): void {
        this.mousePos = [x/50, y/50];
    }
}
</script>

<style scoped>
div {
    max-width: 777.75px;
    max-height: 975px;
}
</style>
