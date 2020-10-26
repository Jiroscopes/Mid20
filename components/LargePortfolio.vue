<template>
    <div class="bg-white grid grid-cols-4 shadow-2xl">
        <!-- Image -->
        <div class="col-span-2 h-portImg overflow-hidden bg-darkBlue">
            <img class="" :src="image" alt="Large portfolio image" >
        </div>
        <!-- Right Side -->
        <div class="col-span-2 flex items-center">
            <div class="px-16">
                <h3 class="font-OpenSans font-bold text-orange text-4xl block mt-16">{{title}}</h3>
                <p class="font-Nunito font-bold text-blue text-xl block mt-4">{{type}}</p>
                <a class="font-Nunito font-semibold text-darkBlue underline text-md mt-16 block" :href="link">Read More</a>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'

const LargePortProps = Vue.extend({

    props: {
        image: String,
        title: String,
        type: String,
        link: String,
    }

})

@Component
export default class LargePortfolio extends LargePortProps {

}
</script>

<style>

</style>
