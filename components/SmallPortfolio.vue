<template>
    <div class="bg-white flex items-center shadow-2xl h-full">
        <div class="px-16 md:px-8">
            <h3 class="font-OpenSans font-bold text-orange text-2xl block pt-8 md:pt-8"> {{ title }} </h3>
            <p class="font-Nunito font-bold text-blue text-lg block mt-4"> {{ type }} </p>
            <a class="font-Nunito font-semibold text-darkBlue underline text-sm md:mb-0 md:mt-8 my-8 block" :href="link"> Read More </a>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'

const SmallProps = Vue.extend({

    props: {
        title: String,
        type: String,
        link: String,

    }

})

export default class SmallPortfolio extends SmallProps {

}

</script>

<style>

</style>
