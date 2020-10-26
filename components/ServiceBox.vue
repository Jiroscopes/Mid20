<template>
    <div class="grid grid-flow-row auto-rows-max bg-darkBlue rounded-3xl wwd-shadow py-8 h-wwd transform hover:translate-y-1 hover:scale-110 transition duration-500 ease-in-out">
        <div class="flex justify-center h-24 " >
            <img :src='image' class="h-auto w-auto object-contain" alt="Website Icon" >
        </div>
        <div class="flex py-8 font-Nunito font-semibold text-2xl mx-8 justify-center">
            <h4 class="text-white">{{ title }} </h4>
        </div>
        <div class="flex mx-8 font-Nunito text-md">
            <p class="text-white" >{{ text }}</p>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'

// In order to access props you have to extend Vue here
// We will extend this so that we can still use our class component
// https://class-component.vuejs.org/guide/props-definition.html
const ServiceProps = Vue.extend({
     props: {
        title: String,
        image: String,
        text: String,
    }
})

@Component
export default class ServiceBox extends ServiceProps {


}
</script>

<style>
.wwd-shadow {
    box-shadow: 6px 6px 17px 3px rgba(7, 13, 89, 0.4);
}
.wwd-shadow:hover {
    
}
</style>
