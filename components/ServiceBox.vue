<template>
    <div class="serviceBox">
        <div class="flex py-8 font-Nunito font-semibold text-2xl mx-8">
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

<style scoped>

.serviceBox {
    @apply grid grid-flow-row auto-rows-max bg-blue rounded-3xl wwd-shadow transform  transition duration-500 ease-in-out;
}

.serviceBox-title {

}


.wwd-shadow:hover {
    box-shadow: 6px 6px 17px 3px rgba(7, 13, 89, 0.4);
}
</style>
